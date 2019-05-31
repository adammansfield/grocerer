#!/usr/bin/env bash

dependencies="golang-go jq maven python-minimal"

if [ $(id -u) = 0 ]; then
  aptCommand="apt"
else
  aptCommand="sudo apt"
fi

echo "Installing dependencies: $dependencies"
$aptCommand update -qq
$aptCommand install -y $dependencies
