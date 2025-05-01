# Makefile for cross-compiling Go program to multiple platforms

APP_NAME := pddns
SRC := main.go
BUILD_DIR := build

# List of target OS/Architecture combinations
TARGETS := \
	freebsd/amd64 \
	darwin/amd64 \
	darwin/arm64 \
	linux/amd64 \
	linux/arm64 \
	windows/amd64

# Default target
all: $(TARGETS)

# Pattern rule to build for each target
$(TARGETS):
	@echo "Building for $@"
	@mkdir -p $(BUILD_DIR)
	@OS=$(word 1, $(subst /, ,$@)); \
	ARCH=$(word 2, $(subst /, ,$@)); \
	EXT=$$( [ "$$OS" = "windows" ] && echo ".exe" || echo "" ); \
	GOOS=$$OS GOARCH=$$ARCH go build -o $(BUILD_DIR)/$(APP_NAME)-$$OS-$$ARCH$$EXT $(SRC)

# Clean up build artifacts
clean:
	@rm -rf $(BUILD_DIR)
