.Phony: dependencies run build

dependencies:
	go get -u github.com/gin-gonic/gin
	go get k8s.io/client-go

build-darwin:
	rm -rf dist || true
	mkdir dist
	cd whitetail; env GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -v -o whitetail
	mv whitetail/whitetail dist/whitetail
	cp -r resources/* dist
	mkdir dist/data
	mkdir -p dist/config/custom/logo ||true
	mkdir -p dist/config/custom/icon ||true

build-linux:
	# if building from a Mac you must install this first:
	# brew install FiloSottile/musl-cross/musl-cross
	rm -rf dist || true
	mkdir dist
	cd whitetail; env GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -ldflags "-linkmode external -extldflags -static" -v -o whitetail
	mv whitetail/whitetail dist/whitetail
	cp -r resources/* dist
	mkdir dist/data
	mkdir -p dist/config/custom/logo ||true
	mkdir -p dist/config/custom/icon ||true

build-docker:
	make build-linux
	docker build -t johncarterodg/whitetail:$(TAG) .
	docker push johncarterodg/whitetail:$(TAG)

run: 
	cd dist; ./whitetail