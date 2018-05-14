#!/bin/bash

source .env
export AWS_ACCESS_KEY_ID=AKIAIHTXMBGZCTLMU7FA
export AWS_SECRET_ACCESS_KEY=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

set -eux

eval `aws ecr get-login --no-include-email --region us-west-2`
docker push ${IMAGE}:${TAG}

