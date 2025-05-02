# Name of the binary
BINARY_NAME = pddns

# Build command
build:
	go build -o $(BINARY_NAME) .

# Clean up binary
clean:
	rm -f $(BINARY_NAME)

# Install to $GOBIN or $GOPATH/bin
install:
	go install .

# Run the program
run:
	go run .

# Rebuild from scratch
rebuild: clean build
