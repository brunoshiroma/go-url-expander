clean:
	go clean -cache -testcache -modcache -fuzzcache -i
	rm go-url-expander
	
build:
	go build -o go-url-expander cmd/main.go

build-image: build
	docker build -t go-url-expander .

build-image-multi: 
	docker build -t go-url-expander -f Dockerfile.multistage .

test-cover:
	go test -cover ./...