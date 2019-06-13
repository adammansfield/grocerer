#!/usr/bin/env bash

main() {
  local -r openApiSpec="openapi.yaml"
  local -r outputDir="gen"

  echo "Getting openapi-generator-cli..."
  get_openapi_generator "openapi-generator"

  echo "Generating client..."
  generate_client "$openApiSpec" "$outputDir"

  echo "Generating server stub..."
  generate_server_stub "$openApiSpec" "$outputDir"
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
  local -r directory=$1
  local -r filename="openapi-generator-cli"
  if ! command -v $filename > /dev/null; then
    if [[ ! -d "$directory" ]]; then
      mkdir -p "$directory"
    fi
    if [[ ! -f "$directory/$filename" ]]; then
      echo "Downloading $filename..."
      curl -s https://raw.githubusercontent.com/OpenAPITools/openapi-generator/master/bin/utils/openapi-generator-cli.sh > "$directory/openapi-generator-cli"
      chmod u+x "$directory/openapi-generator-cli"
    fi
    export PATH=$PATH:$directory
  fi
}

main "$@"
