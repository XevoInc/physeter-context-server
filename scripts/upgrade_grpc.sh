#!/bin/sh

SERVICE_NAME=grpc-server \
FAMILY_NAME=physeter-context-server-grpc \
`dirname $0`/upgrade_ecs.sh


