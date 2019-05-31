#!/usr/bin/env bash

dependencies="golang-go gom jq maven python-minimal"

if [ $(id -u) = 0 ]; then
  aptCommand="apt"
else
  aptCommand="sudo apt"
fi

set -x
$aptCommand update -qq
$aptCommand install -y $dependencies
go get github.com/gorilla/mux
