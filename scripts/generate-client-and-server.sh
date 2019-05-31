#!/usr/bin/env bash

main() {
  local -r projectDir=..
  local -r openApiSpec=$projectDir/api/openapi.yaml
  local -r outputDirectory=$projectDir/gen

  install_missing_dependencies
  get_openapi_generator $projectDir
  generate_client $openApiSpec $outputDirectory
  generate_server $openApiSpec $outputDirectory
}

generate_client() {
  local -r openApiSpec=$1
  local -r outputDirectory=$2
  openapi-generator-cli generate -i $openApiSpec -o $outputDirectory/clients/go -g go
}

generate_server() {
  local -r openApiSpec=$1
  local -r outputDirectory=$2
  openapi-generator-cli generate -i $openApiSpec -o $outputDirectory/servers/go -g go-server
  cp -r $outputDirectory/servers/go $projectDir/internal
}

get_openapi_generator() {
  local -r projectDir=$1
  if ! command -v openapi-generator-cli > /dev/null; then
    openApiGeneratorDir=$projectDir/gen/openapi-generator
    if [[ ! -d $openApiGeneratorDir ]]; then
      mkdir -p $openApiGeneratorDir
    fi
    curl -s https://raw.githubusercontent.com/OpenAPITools/openapi-generator/master/bin/utils/openapi-generator-cli.sh > $openApiGeneratorDir/openapi-generator-cli
    chmod u+x $openApiGeneratorDir/openapi-generator-cli
    export PATH=$PATH:$openApiGeneratorDir
  fi
}

install_missing_dependencies() {
  missingDependencies=
  if ! command -v jq > /dev/null; then
    missingDependencies="$missingDependencies jq"
  fi

  if ! command -v mvn > /dev/null; then
    missingDependencies="$missingDependencies maven"
  fi

  if ! command -v python > /dev/null; then
    missingDependencies="$missingDependencies python-minimal"
  fi

  if [ ! -z "$missingDependencies" ] ; then
    echo "Running: sudo apt install -y $missingDependencies"
    sudo apt update -qq
    sudo apt install $missingDependencies
  fi
}

main "$@"
