#!/bin/sh

echo "Building protobuf"

DEPENDENCIES="autoconf automake libtool curl make g++ unzip git"

apk add --no-cache ${DEPENDENCIES}
git clone https://github.com/google/protobuf --depth 1 -b ${PROTOBUF_BRANCH}
cd protobuf
./autogen.sh
./configure
make -j 3
make check
make install
cd ..
rm -rf protobuf
apk del ${DEPENDENCIES}
apk add --update libstdc++
rm -rf /var/cache/apk/*

