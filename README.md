# Physeter Context Server


# Development

## Prepare

### Install apps from brew

```sh
brew install go dep
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
./build_proto.sh  # if proto file updated
go build main.go
```

## Usage

### launch server

```sh
./main
```

### show API document

```sh
open https://localhost:11000/openapi-ui/
```

### examples of REST API request

```sh
# gas is low
curl -v -k -G -d "car_state.current_location.latitude=35.650645" -d "car_state.current_location.longitude=139.710048" -d "car_state.fuel_level_percentage=10" --data-urlencode "time=2018-05-01T12:30:00+09:00" "https://localhost:11000/api/v1/recommends"
# morning cafe
curl -v -k -G -d "car_state.current_location.latitude=35.650645" -d "car_state.current_location.longitude=139.710048" -d "car_state.fuel_level_percentage=50" --data-urlencode "time=2018-05-01T07:30:00+09:00" "https://localhost:11000/api/v1/recommends"
# lunch time
curl -v -k -G -d "car_state.current_location.latitude=35.650645" -d "car_state.current_location.longitude=139.710048" -d "car_state.fuel_level_percentage=50" --data-urlencode "time=2018-05-01T12:30:00+09:00" "https://localhost:11000/api/v1/recommends"
# evening to grosary
curl -v -k -G -d "car_state.current_location.latitude=35.650645" -d "car_state.current_location.longitude=139.710048" -d "car_state.fuel_level_percentage=50" --data-urlencode "time=2018-05-01T19:00:00+09:00" "https://localhost:11000/api/v1/recommends"
# take child to school
curl -v -k -G -d "car_state.current_location.latitude=35.650645" -d "car_state.current_location.longitude=139.710048" -d "car_state.fuel_level_percentage=50" -d "car_state.number_of_passengers=2" --data-urlencode "time=2018-05-01T07:30:00+09:00" "https://localhost:11000/api/v1/recommends"
```
