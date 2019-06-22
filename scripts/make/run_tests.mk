define run_tests
	docker build -t $(app)-test -f build/package/Dockerfile.test --build-arg tag=$(1) .
endef
