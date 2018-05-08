FROM golang:alpine AS builder
# install protoc
WORKDIR /tmp
RUN apk --no-cache add --virtual .builddeps autoconf automake libtool curl make g++ unzip
RUN apk --no-cache add libstdc++ git && \
  git clone https://github.com/google/protobuf.git --branch v3.5.1.1 --depth 1
RUN cd protobuf && ./autogen.sh && ./configure && make -j 3 && make install
# build directories
RUN mkdir /app
WORKDIR /go/src/xevo/physeter-context-server
COPY . .
# Go dep! and install tools
RUN apk add --update --no-cache git && \
  rm -rf /tmp/* /var/cache/apk/*
RUN go get -u github.com/golang/dep/...
RUN dep ensure && go get \
        ./vendor/github.com/gogo/protobuf/protoc-gen-gogoslick \
        ./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
        ./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
        ./vendor/github.com/mwitkow/go-proto-validators/protoc-gen-govalidators \
        ./vendor/github.com/rakyll/statik
# build server
RUN ./build_proto.sh && go build -o /app/server .


FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/server /app/
ENV PRODUCTION 1
EXPOSE 10000 11000
CMD ["/app/server"]
