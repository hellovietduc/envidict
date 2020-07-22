#!/usr/bin/env bash

VERSION="$(date +"%y.%m%d.%H%M")"

docker --context default build -t docker.pkg.github.com/vietduc01100001/envidict/envidict:"$VERSION" .

docker --context default push docker.pkg.github.com/vietduc01100001/envidict/envidict:"$VERSION"

echo "docker.pkg.github.com/vietduc01100001/envidict/envidict:$VERSION"
