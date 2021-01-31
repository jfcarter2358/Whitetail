.Phony: dependencies run build

dependencies:
	go get -u github.com/gin-gonic/gin
	go get k8s.io/client-go

build-darwin:
	rm -rf dist || true
	mkdir dist
	cd whitetail; env GOOS=darwin GOARCH=amd64 go build -v -o whitetail
	mv whitetail/whitetail dist/whitetail
	cp -r resources/* dist

build-linux:
	rm -rf dist || true
	mkdir dist
	cd whitetail; env GOOS=linux GOARCH=amd64 go build -v -o whitetail
	mv whitetail/whitetail dist/whitetail
	cp -r resources/templates dist/templates

build-windows:
	rm -rf dist || true
	mkdir dist
	cd whitetail; env GOOS=windows GOARCH=amd64 go build -v -o whitetail
	mv whitetail/whitetail dist/whitetail
	cp -r resources/templates dist/templates

build-docker:
	make build-linux
	docker build -t modelop/whitetail:dev-bp .
	docker push modelop/whitetail:dev-bp

run: 
	cd dist; ./whitetail