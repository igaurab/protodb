# Makefile

# Variables
GO_CMD := go
MAIN_FILE := server.go
BINARY_NAME := server
BUILD_DIR := bin

# Default goal
.PHONY: default
default: run

# Run the server
.PHONY: run
run: $(MAIN_FILE)
	@echo "Running the server..."
	$(GO_CMD) run *.go

# Build the application
.PHONY: build
build: $(MAIN_FILE)
	@echo "Building the application..."
	mkdir -p $(BUILD_DIR)
	$(GO_CMD) build -o $(BUILD_DIR)/$(BINARY_NAME) .

# Clean the build artifacts
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)

# List available commands
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make run    - Run the server"
	@echo "  make build  - Build the application"
	@echo "  make clean  - Clean build artifacts"
	@echo "  make help   - Show this help message"