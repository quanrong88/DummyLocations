package main

import (
  "flag"
  "fmt"
  "log"
  "net/http"
  "os"
  "os/signal"
  "strings"
  "time"

  "github.com/grpc-ecosystem/grpc-gateway/runtime"
  "github.com/spf13/pflag"
  "github.com/spf13/viper"
  "golang.org/x/net/context"
  "google.golang.org/grpc"

  gw "gen/pb-go"
)

type proxyConfig struct {
  backend string
  swagger string
  corsAllowOrigin      string
  corsAllowCredentials string
}

func allowCors(cfg proxyConfig, handler http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", cfg.corsAllowOrigin)
    w.Header().Set("Access-Control-Allow-Credentials", cfg.corsAllowCredentials)
    handler.ServeHTTP(w, req)
  })
}

func SetupMux(ctx context.Context, cfg proxyConfig) *http.ServeMux {
  mux := http.NewServeMux()

  mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, cfg.swagger)
  })

  opts := []grpc.DialOption{grpc.WithInsecure()}
  gwmux := runtime.NewServeMux()
  fmt.Printf("Proxying requests to gRPC service at '%s'\n", cfg.backend)
  
  err := gw.RegisterLocationsHandlerFromEndpoint(ctx, gwmux, cfg.backend, opts)
  if err != nil {
    log.Fatalf("Could not register gateway: %v", err)
  }
  mux.Handle("/", allowCors(cfg, gwmux))
  return mux
}

// SetupViper returns a viper configuration object
func SetupViper() *viper.Viper {
  viper.SetConfigName("config")
  viper.AddConfigPath(".")
  viper.SetEnvPrefix("Locations")
  viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
  viper.AutomaticEnv()

  flag.String("backend", "", "The gRPC backend service to proxy.")

  pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
  pflag.Parse()
  viper.BindPFlags(pflag.CommandLine)

  err := viper.ReadInConfig()
  if err != nil {
    log.Fatalf("Could not read config: %v", err)
  }

  return viper.GetViper()
}

// SignalRunner runs a runner function until an interrupt signal is received, at which point it
// will call stopper.
func SignalRunner(runner, stopper func()) {
  signals := make(chan os.Signal, 1)
  signal.Notify(signals, os.Interrupt, os.Kill)

  go func() {
    runner()
  }()

  fmt.Println("hit Ctrl-C to shutdown")
  select {
  case <-signals:
    stopper()
  }
}

func main() {

  cfg := SetupViper()
  ctx := context.Background()
  ctx, cancel := context.WithCancel(ctx)
  defer cancel()

  mux := SetupMux(ctx, proxyConfig{
    backend: cfg.GetString("backend"),
    swagger: cfg.GetString("swagger.file"),
    corsAllowOrigin:      cfg.GetString("cors.allow-origin"),
    corsAllowCredentials: cfg.GetString("cors.allow-credentials"),
  })

  addr := fmt.Sprintf(":%v", cfg.GetInt("proxy.port"))
  server := &http.Server{Addr: addr, Handler: mux}

  SignalRunner(
    func() {
      fmt.Printf("launching http server on %v\n", server.Addr)
      if err := server.ListenAndServe(); err != nil {
        log.Fatalf("Could not start http server: %v", err)
      }
    },
    func() {
      shutdown, _ := context.WithTimeout(ctx, 10*time.Second)
      server.Shutdown(shutdown)
    })
}
