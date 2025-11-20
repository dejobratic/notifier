.PHONY: build docs run dev clean test

# Build the application
build:
	@echo "Building notifier..."
	@go build -o notifier ./main.go

# Generate Swagger documentation
docs:
	@echo "Generating Swagger docs..."
	@go run github.com/swaggo/swag/cmd/swag@latest init

# Run the application
run: build
	@echo "Starting notifier..."
	@./notifier

# Development: generate docs and run
dev: docs
	@echo "Starting notifier in development mode..."
	@go run ./main.go

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -f notifier
	@rm -rf docs/*.go docs/*.json docs/*.yaml

# Run tests
test:
	@echo "Running tests..."
	@go test ./...

