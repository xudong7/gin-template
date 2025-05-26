.PHONY: clean build run dev test

clean:
	@rm -rf bin
	@go mod tidy

build:
	@go build -o bin/gin-template ./cmd/server/main.go

run: clean build
	@./bin/gin-template

dev:
	@go run ./cmd/server/main.go

test:
	@go test ./...