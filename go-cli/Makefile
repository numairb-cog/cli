# Makefile for npm CLI (Go implementation)

.PHONY: all build install clean test help

# Variables
BINARY_NAME=npm
NPX_BINARY_NAME=npx
VERSION=11.6.2
BUILD_DIR=build
INSTALL_DIR=/usr/local/bin

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# Build flags
LDFLAGS=-ldflags "-X main.version=$(VERSION)"

all: build

## build: Build the npm and npx binaries
build:
	@echo "Building npm CLI..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) main.go
	@echo "Building npx wrapper..."
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(NPX_BINARY_NAME) npx.go
	@echo "Build complete! Binaries are in $(BUILD_DIR)/"

## install: Install the binaries to system
install: build
	@echo "Installing npm and npx to $(INSTALL_DIR)..."
	@sudo cp $(BUILD_DIR)/$(BINARY_NAME) $(INSTALL_DIR)/
	@sudo cp $(BUILD_DIR)/$(NPX_BINARY_NAME) $(INSTALL_DIR)/
	@echo "Installation complete!"

## clean: Clean build artifacts
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete!"

## test: Run tests
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

## deps: Download dependencies
deps:
	@echo "Downloading dependencies..."
	$(GOGET) -v ./...
	$(GOMOD) tidy

## run: Run npm CLI directly
run:
	$(GOCMD) run main.go

## help: Show this help message
help:
	@echo "npm CLI (Go implementation) - Makefile commands:"
	@echo ""
	@sed -n 's/^##//p' Makefile | column -t -s ':' | sed -e 's/^/ /'
