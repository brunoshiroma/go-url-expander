clean:
	go clean -cache -testcache -modcache -fuzzcache -i
	rm go-url-expander
build:
	go build -o go-url-expander cmd/main.go

build-image: build-alpine-x64
	docker build -t go-url-expander .

build-image-multi: 
	docker build -t go-url-expander -f Dockerfile.multistage .
