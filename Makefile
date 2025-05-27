.PHONY: clean build run dev test

clean:
	@rm -rf bin
	@rm -rf logs
	@go mod tidy

build:
	@rm -rf bin
	@go build -o bin/gin-template ./cmd/server/main.go

run: build
	@./bin/gin-template

dev:
	@go run ./cmd/server/main.go

test:
	@go test ./...