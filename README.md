# OurGroceries REST API
[![Build Status](https://api.travis-ci.org/adammansfield/ourgroceries-rest-api.svg?branch=master)](https://travis-ci.org/adammansfield/ourgroceries-rest-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/adammansfield/ourgroceries-rest-api)](https://goreportcard.com/report/github.com/adammansfield/ourgroceries-rest-api)

## Prerequisites
- docker
- make
- python3

## Make commands
```
make build       Build the container
make build-nc    Build the container without caching
make clean       Clean the output and generated files
make run         Run the container
make stop        Stop and remove a running container
make test        Run the small (unit) tests
make test-large  Run the large (end-to-end) tests
make up          Build, test, and run the container
```
