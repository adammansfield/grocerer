APP_NAME := ourgroceries-rest-api
CONTAINER_PORT := 8080
HOST_PORT ?= 1200
OUTPUT := bin/openapi

ifeq ($(OS),Windows_NT)
	COMMENT := @REM
	PYTHON3 := python

	# TODO: Remove when integrate_server_stub.py is done
	CP := xcopy /s
	RM := del
	SEP := \\
else
	COMMENT := @\#
	PYTHON3 := /usr/bin/env python3

	# TODO: Remove when integrate_server_stub.py is done
	CP := cp -r --no-target-directory
	RM := rm -f
	SEP := /
endif
EXTRACT := $(PYTHON3) scripts/extract.py
GO_FILES := $(shell $(PYTHON3) scripts/find.py internal '*.go')

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

.PHONY: test
test: gen ## Run the tests
	docker build -t $(APP_NAME)-test -f build/package/Dockerfile.test .

.PHONY: up
up: build test run ## Build, test, and run the container

$(OUTPUT): $(GO_FILES) bin
	$(call build_image)

bin:
	mkdir bin

gen: api/openapi.yaml
	docker build -t $(APP_NAME)-generate -f build/package/Dockerfile.generate .
	$(EXTRACT) $(APP_NAME)-generate /gen gen
	$(COMMENT) # TODO: Replace below with $(PYTHON3) scripts/integrate_server_stub.py gen/servers/go internal
	$(CP) gen$(SEP)servers$(SEP)go internal
	$(RM) internal$(SEP)go$(SEP)api_default.go
