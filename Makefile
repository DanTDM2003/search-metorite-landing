include .env
export
build:
	@go build -o bin/main

run-api:
	@go run cmd/app/main.go

test:
	@go test -v ./...