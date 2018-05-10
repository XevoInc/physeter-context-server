FROM golang:alpine AS builder
# install protoc
ENV PROTOBUF_BRANCH 3.5.x
ADD scripts/build_protobuf.sh /tmp/build_protobuf.sh
RUN /tmp/build_protobuf.sh
# Go dep! and install tools
RUN mkdir /app
WORKDIR /go/src/xevo/physeter-context-server
RUN apk add --update --no-cache git && \
  rm -rf /tmp/* /var/cache/apk/*
RUN go get -u github.com/golang/dep/...
COPY . .
RUN dep ensure && go install \
        ./vendor/github.com/gogo/protobuf/protoc-gen-gogoslick \
        ./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
        ./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
        ./vendor/github.com/mwitkow/go-proto-validators/protoc-gen-govalidators \
        ./vendor/github.com/rakyll/statik
# build server
RUN ./scripts/update_proto.sh && go build -o /app/server .


FROM alpine
RUN apk add --no-cache ca-certificates
ADD certificates /app/certificates
COPY --from=builder /app/server /app/
ENV PRODUCTION 1
EXPOSE 10000 11000
WORKDIR /app
CMD ["/app/server"]
