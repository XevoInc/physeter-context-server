#!/bin/bash

set -eux

`dirname $0`/build_image.sh
`dirname $0`/push_image.sh
`dirname $0`/upgrade_grpc.sh
`dirname $0`/upgrade_gateway.sh

