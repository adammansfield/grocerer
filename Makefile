include scripts/make/build_image.mk
include scripts/make/crossplatform.mk
include scripts/make/run_tests.mk
include scripts/make/silent.mk

app := $(shell $(BASENAME) $(CURDIR))
port ?= 1200

output := bin/openapi
src := $(shell $(FIND) internal *.go version.go)
version_file := internal/go/version.go

.PHONY: build
build: gen $(output) ## Build the container

.PHONY: build-nc
build-nc: gen $(version_file) bin ## Build the container without caching
	$(call build_image,--no-cache)

.PHONY: clean
clean: ## Clean the project
	$(CLEAN)

.PHONY: help
help:
	$(HELP)
.DEFAULT_GOAL := help

.PHONY: run
run: ## Run the container
	docker run -i -t --rm -p=$(port):8080 --name="$(app)" $(app)

.PHONY: stop
stop: ## Stop and remove the running container
	docker stop $(app)
	docker rm $(app)

.PHONY: test
test: gen $(src) $(version_file) ## Run the small (unit) tests
	$(call run_tests,small_test)

.PHONY: test-large
test-large: gen $(src) $(version_file) ## Run the large (end-to-end) tests
	$(call run_tests,large_test)

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
