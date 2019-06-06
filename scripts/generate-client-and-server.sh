#!/usr/bin/env bash

main() {
  local -r scriptDir=$(dirname "$(readlink -f "$0")")
  local -r projectDir=$(readlink -f "$scriptDir/..")
  local -r openApiSpec="$projectDir/api/openapi.yaml"
  local -r outputDir="$projectDir/gen"

  echo "Installing dependencies..."
  install_dependencies

  echo "Getting openapi-generator-cli..."
  get_openapi_generator "$projectDir/gen/openapi-generator"

  echo "Generating client..."
  generate_client "$openApiSpec" "$outputDir"

  echo "Generating server stub..."
  generate_server_stub "$openApiSpec" "$outputDir"

  echo "Integrating server stub with existing code..."
  integrate_server_stub "$outputDir/servers/go" "$projectDir/internal"
}

integrate_server_stub() {
  local -r serverStubDir=$1
  local -r internalDir=$2
  cp -r $serverStubDir $internalDir
}

generate_client() {
  local -r openApiSpec=$1
  local -r outputDir=$2
  openapi-generator-cli generate -i $openApiSpec -o $outputDir/clients/go -g go
}

generate_server_stub() {
  local -r openApiSpec=$1
  local -r outputDir=$2
  openapi-generator-cli generate -i $openApiSpec -o $outputDir/servers/go -g go-server
}

get_openapi_generator() {
  local -r directory="$projectDir/gen/openapi-generator"
  local -r filename="openapi-generator-cli"
  if ! command -v $filename > /dev/null; then
    if [[ ! -d "$directory" ]]; then
      mkdir -p "$directory"
    fi
    if [[ ! -f "$directory/$filename" ]]; then
      echo "Downloading $filename"
      curl -s https://raw.githubusercontent.com/OpenAPITools/openapi-generator/master/bin/utils/openapi-generator-cli.sh > "$directory/openapi-generator-cli"
      chmod u+x "$directory/openapi-generator-cli"
    fi
    export PATH=$PATH:$directory
  fi
}

install_dependencies() {
  dependencies=""
  if ! command -v jq > /dev/null; then
    dependencies+="jq "
  fi
  if ! command -v mvn > /dev/null; then
    dependencies+="maven "
  fi
  if ! command -v python > /dev/null; then
    dependencies+="python-minimal "
  fi

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
