#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

version=`cat VERSION`

docker login -u "$DOCKER_USERNAME" -p "$DOCKER_PASSWORD"

make dockerize

echo "Publishing riff-init"

docker tag "projectriff/riff-init:${version}" "projectriff/riff-init:latest"
docker tag "projectriff/riff-init:${version}" "projectriff/riff-init:${version}-ci-${TRAVIS_COMMIT}"

docker push "projectriff/riff-init"
