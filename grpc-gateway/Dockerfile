# Template for grpc-gateway 

FROM namely/protoc-all AS build
WORKDIR /app
ENV GOPATH=/app

# Copy all of the staged files (protos plus go source)
COPY . /app/

# Download the go dependencies.
RUN go get ./...

WORKDIR /app

# Build the gateway
RUN go build -o grpc_gateway src/pkg/main/main.go

FROM alpine:3.6
WORKDIR /app
COPY --from=build /app/grpc_gateway /app/
COPY --from=build /app/config.yaml /app/
COPY --from=build /app/src/gen/pb-go/locations.swagger.json /app/

EXPOSE 80
ENTRYPOINT ["/app/grpc_gateway"]
