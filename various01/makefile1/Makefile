# https://github.com/Vungle/docker-golang/blob/master/Makefile

ifndef BUILD_SCOPE
BUILD_SCOPE=dev
endif

PROJECT_IMAGE=gomakefile1:$(BUILD_SCOPE)

.PHONY: all test clean

build:
	@echo "Build options: $(BUILD_OPTS)"
	@docker build \
	$(BUILD_OPTS) \
	-t $(PROJECT_IMAGE) \
	.
	@docker build \
	$(BUILD_OPTS) \
	-t $(PROJECT_IMAGE)-fromscratch \
	-f Dockerfile.builder.scratch \
	.

#test:
#	@echo "Testing standard image..."
#	docker run --rm \
#	-v `pwd`/test:/var/test \
#	$(PROJECT_IMAGE) \
#	go run /var/test/mustcompile/mustcompile.go -version="$$(cat .version)"
#	@echo "Testing alpine image..."
#	docker run --rm \
#	-v `pwd`/test:/var/test \
#	$(PROJECT_IMAGE)-alpine \
#	go run /var/test/mustcompile/mustcompile.go -version="$$(cat .version)"
#
#publish:
#	@docker push $(PROJECT_IMAGE)
#	@docker push $(PROJECT_IMAGE)-alpine
