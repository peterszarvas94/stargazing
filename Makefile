.PHONY: dev build run test clean fmt vet lint

# Development with hot reload
dev:
	air

# Build the binary
build:
	go build -o bin/server ./cmd/server

# Run the binary
run: build
	./bin/server

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -rf bin/ tmp/ logs/

# Format code
fmt:
	go fmt ./...

# Vet code
vet:
	go vet ./...

# Lint (requires golangci-lint)
lint:
	golangci-lint run

# All checks
check: fmt vet test
