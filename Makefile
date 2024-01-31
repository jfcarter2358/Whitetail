.PHONY: run-docker build-docker publish-docker clean

clean:  ## Remove build and test artifacts
	rm -rf dist || true
	docker-compose rm -f

build-docker:  ## Build a Whitetail docker image
	wsc compile
	docker build -t whitetail .

publish-docker:  ## Build and publish the Whitetail docker image
	wsc compile
	docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 -t jfcarter2358/whitetail:$$(cat whitetail/VERSION) --push .

run-docker:  ## Run local docker-compose
	docker-compose down
	docker-compose rm -f
	docker-compose up


