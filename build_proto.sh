#!/bin/bash

protoc \
  -I proto \
  -I vendor/github.com/grpc-ecosystem/grpc-gateway/ \
  -I vendor/github.com/gogo/googleapis/ \
  -I vendor/ \
  --gogoslick_out=plugins=grpc,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:\
$GOPATH/src/ \
  --grpc-gateway_out=\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:\
$GOPATH/src/ \
  --swagger_out=logtostderr=true:./third_party/OpenAPI \
  proto/context.proto

