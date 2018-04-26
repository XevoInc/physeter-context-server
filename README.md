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

```sh
./main
open https://localhost:11000/openapi-ui/
```
