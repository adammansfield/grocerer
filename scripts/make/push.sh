#!/usr/bin/env bash

if [ -z $DOCKER_ID ]; then
  echo "DOCKER_ID environment variable is unset"
  exit 1
elif [ -z $DOCKER_PASSWORD ]; then
  echo "DOCKER_PASSWORD environment variable is unset"
  exit 1
fi

tag="$1"

echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_ID" --password-stdin
docker tag $tag $DOCKER_ID/$tag
docker push $DOCKER_ID/$tag
