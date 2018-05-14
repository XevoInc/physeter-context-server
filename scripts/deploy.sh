#!/bin/bash

source .env

set -eux

eval `aws ecr get-login --no-include-email --region us-west-2`
docker push ${IMAGE}:${TAG}

