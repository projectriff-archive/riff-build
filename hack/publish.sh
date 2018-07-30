#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

version=`cat VERSION`

echo "Publishing riff-build"

# build and publish containers, generate yaml
ko resolve -f config/ > riff-build.yaml

# publish yaml
gsutil cp -a public-read "riff-build.yaml" "gs://riff-releases/latest/riff-build.yaml"
gsutil cp -a public-read "riff-build.yaml" "gs://riff-releases/previous/riff-build/riff-build-${version}.yaml"
gsutil cp -a public-read "riff-build.yaml" "gs://riff-releases/previous/riff-build/riff-build-${version}-ci-${TRAVIS_COMMIT}.yaml"
