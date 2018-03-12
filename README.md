# DummyLocations

My test grpc with micro service project

### Installing
* Install [Docker](https://docs.docker.com)
* Change directory to the project root forder.
* Create network for this project 
```
$ docker network create -d bridge qNet
```
*  Run this command in command line to run project
```
$ docker-compose up
```
* Create grpc container client using following commands:
```
$ docker build -t location-client-module ./locations-clients
$ docker run -d -it --name location-client-module --net qNet -v $(pwd)/locations-clients/:/code location-client-module
```
* Run test
```
$ docker exec -it location-client-module bash
# python client.py
```
* Create API gateway (Experimental)
```
$ docker build -t locations-gateway ./grpc-gateway
$ docker run --name locations-gateway -p 8080:80 --net qNet  locations-gateway --backend=locations-module:22222
```
* Run test gateway
```
$ curl http://localhost:8080/locations
$ curl http://localhost:8080/locations/1
```
