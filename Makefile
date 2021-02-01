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
	mkdir dist/data

build-linux:
	rm -rf dist || true
	mkdir dist
	cd whitetail; env GOOS=linux GOARCH=amd64 go build -v -o whitetail
	mv whitetail/whitetail dist/whitetail
	cp -r resources/* dist
	mkdir dist/data

build-windows:
	rm -rf dist || true
	mkdir dist
	cd whitetail; env GOOS=windows GOARCH=amd64 go build -v -o whitetail
	mv whitetail/whitetail dist/whitetail
	cp -r resources/* dist
	mkdir dist/data

build-docker:
	make build-linux
	docker build -t johncarterodg/whitetail:$(TAG) .
	docker push johncarterodg/whitetail:$(TAG)

run: 
	cd dist; ./whitetail