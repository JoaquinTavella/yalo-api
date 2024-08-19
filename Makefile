# Variables
BINARY_NAME=app
GO_FILES=$(shell find . -type f -name '*.go')

# Default target: run the application
.PHONY: run
run:
	@go run cmd/main.go

# Build the binary
.PHONY: build
build:
	@go build -o $(BINARY_NAME) cmd/main.go
	@echo "Built $(BINARY_NAME)"

# Run tests
.PHONY: test
test:
	@go test ./... -v
	@echo "Tests completed"

# Run tests with race detection
.PHONY: test-race
test-race:
	@go test -race ./... -v
	@echo "Race tests completed"

# Clean build artifacts
.PHONY: clean
clean:
	@rm -f $(BINARY_NAME)
	@echo "Cleaned up build artifacts"

# Install dependencies (Go modules)
.PHONY: deps
deps:
	@go mod tidy
	@go mod download
	@echo "Dependencies installed"

# Run golangci-lint
.PHONY: lint
lint:
	@golangci-lint run ./...
	@echo "Linting completed"

# Format the code
.PHONY: fmt
fmt:
	@go fmt ./...
	@echo "Code formatted"

# Run all checks (lint, format, and test)
.PHONY: check
check: fmt lint test

# Help message
.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  run        - Run the application"
	@echo "  build      - Build the binary"
	@echo "  test       - Run tests"
	@echo "  test-race  - Run tests with race detection"
	@echo "  clean      - Clean build artifacts"
	@echo "  deps       - Install dependencies"
	@echo "  lint       - Run golangci-lint"
	@echo "  fmt        - Format the code"
	@echo "  check      - Run all checks (fmt, lint, test)"
	@echo "  help       - Show this help message"
