.PHONY: run-docker build-docker

build-docker:
	wsc compile
	docker build -t whitetail .

publish-docker:
	wsc compile
	docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 -t jfcarter2358/whitetail:$$(cat whitetail/VERSION) --push .

run-docker:
	docker-compose rm -f
	docker-compose up
