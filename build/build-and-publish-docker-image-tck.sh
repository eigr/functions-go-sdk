#!/usr/bin/env bash

set -o nounset
set -o errexit
set -o pipefail

DOCKER_BUILDKIT=1 docker build -t gcr.io/eigr-io/eigr-go-tck:latest -f ./build/TCK.Dockerfile . || exit $?
docker push gcr.io/eigr-io/eigr-go-tck:latest
