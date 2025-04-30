# Makefile for cross-compiling Go program to multiple platforms

APP_NAME := pddns
SRC := main.go
BUILD_DIR := build

# List of target OS/Architecture combinations
TARGETS := \
	freebsd/amd64 \
	darwin/amd64 \
	linux/amd64 \
	linux/arm64

# Default target
all: $(TARGETS)

# Pattern rule to build for each target
$(TARGETS):
	@echo "Building for $@"
	@mkdir -p $(BUILD_DIR)
	@GOOS=$(word 1, $(subst /, ,$@)) GOARCH=$(word 2, $(subst /, ,$@)) \
	go build -o $(BUILD_DIR)/$(APP_NAME)-$(word 1, $(subst /, ,$@))-$(word 2, $(subst /, ,$@)) $(SRC)

# Clean up build artifacts
clean:
	@rm -rf $(BUILD_DIR)
