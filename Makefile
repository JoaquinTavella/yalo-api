# Variables
BINARY_NAME=app
GO_FILES=$(shell find . -type f -name '*.go')
DOCKER_IMAGE=yalo-app
DOCKER_TAG=latest

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

# Install all dependencies, build the binary, and run tests
.PHONY: install-all
install-all: deps build test
	@echo "All dependencies installed, binary built, and tests run successfully"

# Build the Docker image
.PHONY: docker-build
docker-build:
	@docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .
	@echo "Docker image $(DOCKER_IMAGE):$(DOCKER_TAG) built successfully"

# Run the Docker container
.PHONY: docker-run
docker-run:
	@docker run -p 8080:8080 $(DOCKER_IMAGE):$(DOCKER_TAG)
	@echo "Docker container for $(DOCKER_IMAGE):$(DOCKER_TAG) is running"

# Clean Docker artifacts (optional)
.PHONY: docker-clean
docker-clean:
	@docker rmi $(DOCKER_IMAGE):$(DOCKER_TAG)
	@echo "Docker image $(DOCKER_IMAGE):$(DOCKER_TAG) removed"


# Run Docker Compose
.PHONY: docker-compose-up
docker-compose-up:
	@docker-compose up -d
	@echo "Docker Compose services are up and running"

# Stop Docker Compose
.PHONY: docker-compose-down
docker-compose-down:
	@docker-compose down
	@echo "Docker Compose services are stopped"

# Run Python script
.PHONY: populate-db
run-python-script:
	@python3 script.py
	@echo "Python script executed"

# Help message
.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  run          - Run the application"
	@echo "  build        - Build the binary"
	@echo "  test         - Run tests"
	@echo "  test-race    - Run tests with race detection"
	@echo "  clean        - Clean build artifacts"
	@echo "  deps         - Install dependencies"
	@echo "  lint         - Run golangci-lint"
	@echo "  fmt          - Format the code"
	@echo "  check        - Run all checks (fmt, lint, test)"
	@echo "  docker-build - Build the Docker image"
	@echo "  docker-run   - Run the Docker container"
	@echo "  docker-clean - Remove the Docker image"
	@echo "  docker-compose-up   - Run Docker Compose"
	@echo "  docker-compose-down - Stop Docker Compose"
	@echo "  populate-db  - Run Python script"
	@echo "  help         - Show this help message"
