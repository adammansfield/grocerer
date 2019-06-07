APP_NAME := ourgroceries-rest-api
CONTAINER_PORT := 8080
DOCKERFILE_DIRECTORY := ./internal
HOST_PORT := 1200

ifeq ($(OS),Windows_NT)
	GO_FILES:=
else
	GO_FILES:=$(shell find internal/ -type f -name '*.go')
endif


.PHONY: build
build: gen bin/openapi ## Build the container

.PHONY: build-nc
build-nc: gen ## Build the container without caching
	docker build --no-cache -t $(APP_NAME) $(DOCKERFILE_DIRECTORY)
	./scripts/extract-output-from-docker.sh $(APP_NAME)

.PHONY: clean
clean: ## Clean the output and generated files
	rm -fr ./bin/
	rm -fr ./gen/
	rm -fr ./tools/openapi-generator/
	rm -f ./internal/go/logger.go
	rm -f ./internal/go/model_version.go
	rm -f ./internal/go/routers.go
	rm -f ./internal/Dockerfile
	rm -f ./internal/main.go
	if [ -d tools/ ]; then rmdir tools/; fi

.PHONY: help
help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
.DEFAULT_GOAL := help

.PHONY: run
run: ## Run the container
	docker run -i -t --rm -p=$(HOST_PORT):$(CONTAINER_PORT) --name="$(APP_NAME)" $(APP_NAME)

.PHONY: stop
stop: ## Stop and remove a running container
	docker stop $(APP_NAME); docker rm $(APP_NAME)

.PHONY: up
up: build run ## Build and run the container


bin/openapi: $(GO_FILES)
	docker build -t $(APP_NAME) $(DOCKERFILE_DIRECTORY)
	./scripts/extract-output-from-docker.sh $(APP_NAME)

gen: api/openapi.yaml
	./scripts/generate-client-and-server.sh
