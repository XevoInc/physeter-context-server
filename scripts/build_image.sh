#!/bin/bash

source .env

docker build . -t ${IMAGE}:${TAG}

