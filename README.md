# Physeter Context Server

## Build

```sh
go build .
```

## Usage

```
$ ./physeter-context-server
NAME:
   context-server - start GRPC server or GRPC gateway

USAGE:
   physeter-context-server [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
     gateway-server, gateway  start gRPC-Gateway server
     grpc-server, grpc        start gRPC server
     help, h                  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --gateway-host host, -H host  The gRPC-Gateway server host (default: "localhost") [$GATEWAY_HOST]
   --gateway-port port, -P port  The gRPC-Gateway server port (default: 11000) [$GATEWAY_PORT]
   --grpc-host host, -h host     The gRPC server host (default: "localhost") [$GRPC_HOST]
   --grpc-port port, -p port     The gRPC server port (default: 10000) [$GRPC_PORT]
   --help                        show help
   --version, -v                 print the version
```

## Start Server by Docker

```sh
./scripts/build.sh
docker-compose -f scripts/docker-compose.yml up
open https://localhost:11000/
```

# Development

## Prepare

### Install apps from brew

```sh
brew install go dep protobuf
```

### Clone src codes

```sh
cd $GOPATH
mkdir -p src/xevo
cd src/xevo
git clone git@github.com:XevoInc/physeter-context-server.git
cd physeter-context-server
```

### Get vendor libraries

```sh
dep ensure
go install \
        ./vendor/github.com/gogo/protobuf/protoc-gen-gogoslick \
        ./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
        ./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
        ./vendor/github.com/mwitkow/go-proto-validators/protoc-gen-govalidators \
        ./vendor/github.com/rakyll/statik
```

## Build

```sh
./scripts/update_proto.sh  # if proto file updated
go build .
```

## Usage

### launch server

```sh
./physeter-context-server grpc &
./physeter-context-server gateway &
open https://localhost:11000/  # and ignore self sign certification error
```

### show API document

```sh
open https://localhost:11000/openapi-ui/
```

### examples of REST API request

```sh
# gas is low
curl -v -k -G -d "car_state.current_location.latitude=35.650645" -d "car_state.current_location.longitude=139.710048" -d "car_state.fuel_level_percentage=10" --data-urlencode "time=2018-05-01T12:30:00+09:00" "https://localhost:11000/api/v1/recommends"
# in the morning
curl -v -k -G -d "car_state.current_location.latitude=35.650645" -d "car_state.current_location.longitude=139.710048" -d "car_state.fuel_level_percentage=50" --data-urlencode "time=2018-05-01T07:30:00+09:00" "https://localhost:11000/api/v1/recommends"
# at lunch time
curl -v -k -G -d "car_state.current_location.latitude=35.650645" -d "car_state.current_location.longitude=139.710048" -d "car_state.fuel_level_percentage=50" --data-urlencode "time=2018-05-01T12:30:00+09:00" "https://localhost:11000/api/v1/recommends"
# in the evening
curl -v -k -G -d "car_state.current_location.latitude=35.650645" -d "car_state.current_location.longitude=139.710048" -d "car_state.fuel_level_percentage=50" --data-urlencode "time=2018-05-01T19:00:00+09:00" "https://localhost:11000/api/v1/recommends"
# drive a child to school
curl -v -k -G -d "car_state.current_location.latitude=35.650645" -d "car_state.current_location.longitude=139.710048" -d "car_state.fuel_level_percentage=50" -d "car_state.number_of_passengers=2" --data-urlencode "time=2018-05-01T07:30:00+09:00" "https://localhost:11000/api/v1/recommends"
```

### Conditions of available recommendation

1. gas is low (gas station)
    - when the fuel level <= 20.0
1. in the morning (cafe)
    - when hour is in [6, 9]
1. at lunch time (restaurant)
    - when hour is in [12, 13]
1. in the evening (grocery)
    - when hour is in [18, 21]
1. drive a child (school)
    - when hour is in [6, 9] and passengers > 1
