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
```

## Build

```sh
./build_proto.sh  # if proto file updated
go build main.go
```
