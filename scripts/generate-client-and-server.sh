#!/usr/bin/env bash

main() {
  local -r scriptDir=$(dirname "$(readlink -f "$0")")
  local -r projectDir=$(readlink -f "$scriptDir/..")
  local -r openApiSpec="$projectDir/api/openapi.yaml"
  local -r outputDirectory="$projectDir/gen"

  echo "Installing dependencies..."
  install_dependencies

  echo "Getting openapi-generator-cli..."
  get_openapi_generator "$projectDir/gen/openapi-generator"

  echo "Generating client..."
  generate_client $openApiSpec $outputDirectory

  echo "Generating server..."
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
  local -r openApiGeneratorDir="$projectDir/gen/openapi-generator"
  if ! command -v openapi-generator-cli > /dev/null; then
    if [[ ! -d $openApiGeneratorDir ]]; then
      mkdir -p $openApiGeneratorDir
    fi
    curl -s https://raw.githubusercontent.com/OpenAPITools/openapi-generator/master/bin/utils/openapi-generator-cli.sh > $openApiGeneratorDir/openapi-generator-cli
    chmod u+x $openApiGeneratorDir/openapi-generator-cli
    export PATH=$PATH:$openApiGeneratorDir
  fi
}

install_dependencies() {
  dependencies=""
  #if ! command -v jq > /dev/null; then
    dependencies+="jq "
  #fi
  #if ! command -v mvn > /dev/null; then
    dependencies+="maven "
  #fi
  #if ! command -v python > /dev/null; then
    dependencies+="python-minimal "
  #fi

  echo "dependencies=$dependencies"

  if [ -z "$dependencies" ]; then
    return
  fi

  if [ $(id -u) = 0 ]; then
    aptCommand="apt"
  else
    aptCommand="sudo apt"
  fi
  $aptCommand update -qq
  $aptCommand install -y $dependencies
}

main "$@"
