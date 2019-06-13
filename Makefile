APP_NAME := ourgroceries-rest-api
CONTAINER_PORT := 8080
HOST_PORT ?= 1200
OUTPUT := bin/openapi

ifeq ($(OS),Windows_NT)
	COMMENT := @REM
	GO_FILES :=
	PYTHON3 := python
	SEP := \\# TODO: Remove when integrate_server_stub.py is done
else
	COMMENT := @\#
	GO_FILES := $(shell find internal/ -type f -name '*.go')
	PYTHON3 := /usr/bin/env python3
	SEP := /# TODO: Remove when integrate_server_stub.py is done
endif
EXTRACT := $(PYTHON3) scripts/extract.py

define build_image
	docker build $(1) -t $(APP_NAME) internal
	$(EXTRACT) $(APP_NAME) openapi $(OUTPUT)
	$(COMMENT) Update the target's timestamp so it is newer than its prerequisites.
	$(COMMENT) This will ensure that make will not unnecessarily rebuild.
	$(PYTHON3) scripts/touch.py $(OUTPUT)
endef

.PHONY: build
build: gen $(OUTPUT) ## Build the container

.PHONY: build-nc
build-nc: gen bin ## Build the container without caching
	$(call build_image,--no-cache)

.PHONY: clean
clean: ## Clean the output and generated files
	$(PYTHON3) scripts/clean.py

.PHONY: help
help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
.DEFAULT_GOAL := help

.PHONY: run
run: ## Run the container
	docker run -i -t --rm -p=$(HOST_PORT):$(CONTAINER_PORT) --name="$(APP_NAME)" $(APP_NAME)

.PHONY: stop
stop: ## Stop and remove a running container
	docker stop $(APP_NAME)
	docker rm $(APP_NAME)

.PHONY: up
up: build run ## Build and run the container

$(OUTPUT): $(GO_FILES) bin
	$(call build_image)

bin:
	mkdir bin

gen: api/openapi.yaml
	docker build -t $(APP_NAME)-generate -f build/package/Dockerfile.generate .
	$(COMMENT) # TODO: Replace mkdir with mkdir.py
	mkdir gen
	mkdir gen$(SEP)clients
	mkdir gen$(SEP)servers
	$(EXTRACT) $(APP_NAME)-generate gen/clients/go gen/clients/go
	$(EXTRACT) $(APP_NAME)-generate gen/servers/go gen/servers/go
	$(COMMENT) # TODO: Replace extract with $(PYTHON3) scripts/integrate_server_stub.py gen/servers/go internal
	$(EXTRACT) $(APP_NAME)-generate gen/servers/go internal
