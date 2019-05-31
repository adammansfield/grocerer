#!/usr/bin/env bash
declare -r scriptDir=$(dirname "$(readlink -f "$0")")
declare -r projectDir=$(readlink -f "$scriptDir/..")
go get -d -v "$projectDir/internal/..."
go build -a -installsuffix cgo -o "$projectDir/bin/openapi" "$projectDir/internal/."
