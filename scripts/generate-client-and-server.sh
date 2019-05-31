#!/usr/bin/env bash

main() {
  local -r scriptDir=$(dirname "$(readlink -f "$0")")
  local -r projectDir=$(readlink -f "$scriptDir/..")
  local -r openApiSpec="$projectDir/api/openapi.yaml"
  local -r outputDirectory="$projectDir/gen"

  get_openapi_generator "$projectDir/gen/openapi-generator"
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

main "$@"
