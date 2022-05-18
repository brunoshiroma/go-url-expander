clean:
	go clean -cache -testcache -modcache -fuzzcache -i
	rm go-url-expander
build:
	go build -o go-url-expander cmd/main.go