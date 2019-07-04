# $(1): image name and tag (docker build -t parameter)
# $(2): output filepath
# $(3): additional docker build args (e.g. --no-cache)
# The output extracted from the docker image might have an older timestamp.
# So update the output's timestamp to ensure that it is newer than its prerequisites.
# Then make will not unnecessarily rebuild.
define build_image
	docker build \
		$(3) \
		-t $(1) \
		-f build/package/Dockerfile \
		.
	$(EXTRACT) $(1) openapi $(2)
	$(TOUCH) $(2)
endef
