#!/usr/bin/env bash
projectDir=..

if ! command -v mvn > /dev/null; then
  sudo apt update
  sudo apt install maven
fi

if ! command -v python > /dev/null; then
  sudo apt update
  sudo apt install python-minimal
fi

if ! command -v openapi-generator-cli > /dev/null; then
  openApiGeneratorDir=$projectDir/gen/openapi-generator
  if [[ ! -d $openApiGeneratorDir ]]; then
    mkdir -p $openApiGeneratorDir
  fi
  curl -s https://raw.githubusercontent.com/OpenAPITools/openapi-generator/master/bin/utils/openapi-generator-cli.sh > $openApiGeneratorDir/openapi-generator-cli
  chmod u+x $openApiGeneratorDir/openapi-generator-cli
  export PATH=$PATH:$openApiGeneratorDir
fi

openApiSpec=$projectDir/api/openapi.yaml
outputDirectory=$projectDir/gen
openapi-generator-cli generate -i $openApiSpec -o $outputDirectory/clients/go -l go
openapi-generator-cli generate -i $openApiSpec -o $outputDirectory/servers/go -g go-server

cp -r $outputDirectory/servers/go $projectDir/internal
