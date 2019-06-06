APP_NAME=ourgroceries-rest-api
CONTAINER_PORT=8080
DOCKERFILE_DIRECTORY=./internal
HOST_PORT=1200

.PHONY: help
help: ## Print the list of commands
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
.DEFAULT_GOAL := help

.PHONY: build
build: ## Build the container
	docker build -t $(APP_NAME) $(DOCKERFILE_DIRECTORY)

.PHONY: build-nc
build-nc: ## Build the container without caching
	docker build --no-cache -t $(APP_NAME) $(DOCKERFILE_DIRECTORY)

.PHONY: generate
generate: ## Generate the client and server stub
	./scripts/generate-client-and-server.sh

.PHONY: run
run: ## Run container
	docker run -i -t --rm -p=$(HOST_PORT):$(CONTAINER_PORT) --name="$(APP_NAME)" $(APP_NAME)

.PHONY: stop
stop: ## Stop and remove a running container
	docker stop $(APP_NAME); docker rm $(APP_NAME)
