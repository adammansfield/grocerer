# OurGroceries API [![OurGroceries Icon](https://www.ourgroceries.com/favicon.ico)](https://www.ourgroceries.com/overview)
[![Build Status](https://api.travis-ci.org/adammansfield/ourgroceries-rest-api.svg?branch=master)](https://travis-ci.org/adammansfield/ourgroceries-rest-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/adammansfield/ourgroceries-rest-api)](https://goreportcard.com/report/github.com/adammansfield/ourgroceries-rest-api)

A RESTful API for [OurGroceries](https://www.ourgroceries.com/overview)

## Prerequisites
For building and running:
- docker
- make
- python3

For development:
- go
- golint
- ./scripts/git/install-git-hooks.py

## Make commands
```
make             # Print help
make build       # Build the container
make build-nc    # Build the container without caching
make clean       # Clean the project
make lint        # Run gofmt, golint, and go vet
make run         # Run the container
make stop        # Stop and remove the running container
make test        # Run the small (unit) tests
make test-large  # Run the large (end-to-end) tests
```
