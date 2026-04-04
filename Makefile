# Variables
BINARY_NAME=myapp
GO_FILES=$(shell find . -type f -name '*.go')

.PHONY: all dev generate build run schema test lint clean help docs

# Default target
all: help

dev: ## Run the project with hot reload
	go tool air & (cd ui && bun dev)

generate: ## Generate ent code
	go tool ent generate ./internal/ent/schema

build: ## Compile the binary
	go build -o bin/$(BINARY_NAME) ./cmd/api

run: build ## Build and run the application
	./bin/$(BINARY_NAME) & (cd ui && bun run dev)

schema: ## Generate new ent schema (usage: make schema schema_name=User)
	go tool ent new $(schema_name) --target ./internal/ent/schema

test: ## Run unit tests with race detection
	go test -v -race ./...

docs: ## Generate OpenAPI spec from annotations
	go tool swag init -g cmd/api/main.go -o internal/docs/

clean: ## Remove build artifacts
	go clean
	rm -rf bin/

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
