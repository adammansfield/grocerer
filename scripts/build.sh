#!/usr/bin/env bash
echo $PWD
go get -d -v internal/...
go build -a -installsuffix cgo -o ./bin/openapi ./internal/.
