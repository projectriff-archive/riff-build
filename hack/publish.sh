#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

version=`cat VERSION`

echo "Publishing riff-init"

# build and publish containers, generate yaml
ko resolve -f config/ > riff-init.yaml

# publish yaml
gsutil cp -a public-read "riff-init.yaml" "gs://riff-releases/latest/riff-init.yaml"
gsutil cp -a public-read "riff-init.yaml" "gs://riff-releases/previous/riff-init/riff-init-${version}.yaml"
gsutil cp -a public-read "riff-init.yaml" "gs://riff-releases/previous/riff-init/riff-init-${version}-ci-${TRAVIS_COMMIT}.yaml"
