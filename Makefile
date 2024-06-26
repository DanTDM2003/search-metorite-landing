build:
	@go build -o bin/main

run-api: build
	@./bin/main

test:
	@go test -v ./...