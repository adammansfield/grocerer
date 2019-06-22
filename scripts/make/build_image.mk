# The output extracted from the docker image might have an older timestamp.
# So update the output's timestamp to ensure that it is newer than its prerequisites.
# Then make will not unnecessarily rebuild.
define build_image
	docker build $(1) -t $(app) internal
	$(EXTRACT) $(app) openapi $(output)
	$(TOUCH) $(output)
endef
