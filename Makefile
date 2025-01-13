# Variables
APP_NAME := app
SRC_DIR := .
BUILD_DIR := build
FRONTEND_BUILD_DIR := frontend/dist
GO_FILES := $(shell find $(SRC_DIR) -name '*.go')

# Default target: Display all available commands
all: build

start: # Run the application
	go run .

build_frontend: # Build the frontend
	cd frontend && pnpm run build

build: build_frontend $(GO_FILES) # Build the application
	mkdir -p $(BUILD_DIR)
	go build -ldflags "-s -w" -o $(BUILD_DIR)/$(APP_NAME) .

clean: # Clean the build directory
	rm -rf $(BUILD_DIR)
	rm -rf $(FRONTEND_BUILD_DIR)

# Phony targets
.PHONY: all start build clean