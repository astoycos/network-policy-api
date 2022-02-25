#!/usr/bin/env bash

set -xv
set -euo pipefail


IMAGE=mfenwick100/sonobuoy-cyclonus:latest

docker build -t $IMAGE .
docker push $IMAGE
