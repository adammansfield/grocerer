# grocerer
[![Build Status](https://api.travis-ci.org/adammansfield/grocerer.svg?branch=master)](https://travis-ci.org/adammansfield/grocerer)
[![Go Report Card](https://goreportcard.com/badge/github.com/adammansfield/grocerer)](https://goreportcard.com/report/github.com/adammansfield/grocerer)
[![License](https://img.shields.io/github/license/adammansfield/grocerer.svg?style=flat-square)](https://github.com/adammansfield/grocerer/blob/master/LICENSE)

A RESTful API for <a href="https://www.ourgroceries.com"><img src="https://www.ourgroceries.com/static/images/header.png" width="auto" height="40px"/></a>

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
