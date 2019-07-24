.PHONY: docker-build
docker-build:
	docker build --build-arg=GOPROXY=${GOPROXY} -t playground/compose-cpu-test --rm .