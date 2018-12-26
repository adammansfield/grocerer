#!/usr/bin/env bash

if ! command -v mvn > /dev/null; then
  sudo apt update
  sudo apt install maven
fi

if ! command -v python > /dev/null; then
  sudo apt update
  sudo apt install python-minimal
fi

if ! command -v openapi-generator > /dev/null; then
  openApiGeneratorDir=./openapi-generator
  if [[ ! -d $openApiGeneratorDir ]]; then
    mkdir $openApiGeneratorDir
  fi
  curl -s https://raw.githubusercontent.com/OpenAPITools/openapi-generator/master/bin/utils/openapi-generator-cli.sh > $openApiGeneratorDir/openapi-generator-cli
  chmod u+x $openApiGeneratorDir/openapi-generator-cli
  export PATH=$PATH:$openApiGeneratorDir
fi

openapi-generator-cli version
