#!/usr/bin/env bash
set -o nounset

function rnd() {
  cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w ${1:-32} | head -n 1
}

FUNC_IMAGE=${1:-gcr.io/eigr-io/eigr-go-tck:latest}
FUNC="cloudstate-function-$(rnd)"
PROXY_IMAGE=${2:-cloudstateio/functions-proxy-core:latest}
PROXY="cloudstate-proxy-$(rnd)"
TCK_IMAGE=${3:-cloudstateio/functions-tck:latest}
TCK="cloudstate-tck-$(rnd)"

finally() {
  docker rm -f "$PROXY"
  docker rm -f "$FUNC"
}
trap finally EXIT
set -x

docker pull "$PROXY_IMAGE"
docker pull "$TCK_IMAGE"

# run the function and the proxy
docker run -d --name "$FUNC" --net=host "${FUNC_IMAGE}" || exit $?
docker run -d --name "$PROXY" --net=host -e USER_FUNCTION_PORT=8090 "${PROXY_IMAGE}" -Dconfig.resource=dev-mode.conf || exit $?

# run the tck
docker run --rm --name $TCK --net=host "${TCK_IMAGE}"
tck_status=$?
if [ "$tck_status" -ne "0" ]; then
  docker logs functions-function
fi
exit $tck_status
