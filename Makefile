include scripts/crossplatform.mk

app := $(shell $(BASENAME) $(CURDIR))
port ?= 1200

output := bin/openapi
src := $(shell $(FIND) internal *.go version.go)
version_file := internal/go/version.go

# The output extracted from the docker image might have an older timestamp.
# So update the output's timestamp to ensure that it is newer than its prerequisites.
# Then make will not unnecessarily rebuild.
define build_image
	docker build $(1) -t $(app) internal
	$(EXTRACT) $(app) openapi $(output)
	$(TOUCH) $(output)
endef

define run_tests
	docker build -t $(app)-test -f build/package/Dockerfile.test --build-arg tag=$(1) .
endef

.PHONY: build
build: gen $(output) ## Build the container

.PHONY: build-nc
build-nc: gen $(version_file) bin ## Build the container without caching
	$(call build_image,--no-cache)

.PHONY: clean
clean: ## Clean the output and generated files
	$(CLEAN)

.PHONY: help
help:
	$(HELP)
.DEFAULT_GOAL := help

.PHONY: run
run: ## Run the container
	docker run -i -t --rm -p=$(port):8080 --name="$(app)" $(app)

.PHONY: stop
stop: ## Stop and remove a running container
	docker stop $(app)
	docker rm $(app)

.PHONY: test
test: gen $(version_file) ## Run the small (unit) tests
	$(call run_tests,small_test)

.PHONY: test-large
test-large: gen $(version_file) ## Run the large (end-to-end) tests
	$(call run_tests,large_test)

.PHONY: up
up: build test run ## Build, test, and run the container

# Run `make <target> verbose=1` to echo every command
ifndef verbose
MAKEFLAGS += --silent
endif

$(output): $(src) $(version_file) bin
	$(call build_image)

bin:
	mkdir bin

# TODO: Remove $(CP) and $(RM) commands when openapi-generator is removed
gen: api/openapi.yaml
	docker build -t $(app)-generate -f build/package/Dockerfile.generate .
	$(EXTRACT) $(app)-generate /gen gen
	$(CP) gen$(SEP)servers$(SEP)go internal
	$(RM) internal$(SEP)go$(SEP)api_default.go

$(version_file): gen $(src)
	$(VERSION)
