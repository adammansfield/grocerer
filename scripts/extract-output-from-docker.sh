#!/usr/bin/env bash

outputDir="./bin"
outputFilename="openapi"
outputFilepath="$outputDir/$outputFilename"
if [[ ! -d $outputDir ]]; then
  mkdir $outputDir
fi

id=$(docker create $1)
docker cp $id:$outputFilename $outputFilepath
docker rm -v $id &> /dev/null

# Update the timestamp so that the output is newer than the source.
# This will ensure that make will not unnecessarily rebuild.
touch $outputFilepath
