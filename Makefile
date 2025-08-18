.PHONY: help test test-coverage test-verbose coverage clean build run testmo-debug testmo-debug-win ci-test

# Default target
help:
	@echo "Available commands:"
	@echo "  test          - Run tests"
	@echo "  test-coverage - Run tests with coverage"
	@echo "  test-verbose  - Run tests with verbose output"
	@echo "  coverage      - Generate coverage report"
	@echo "  clean         - Clean generated files"
	@echo "  build         - Build the application"
	@echo "  run           - Run the application"
	@echo "  testmo-debug  - Debug Testmo CLI integration (Linux/Mac)"
	@echo "  testmo-debug-win - Debug Testmo CLI integration (Windows)"
	@echo "  ci-test       - Test CI workflow locally"

# Run tests
test:
	go test ./...

# Run tests with coverage
test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

# Run tests with verbose output and JUnit XML
test-verbose:
	mkdir -p reports
	go install gotest.tools/gotestsum@latest
	gotestsum --format=standard-verbose \
		--junitfile=reports/unit-tests.xml \
		--jsonfile=reports/test-results.json \
		--coverprofile=coverage.out \
		--covermode=atomic \
		./...

# Generate coverage report
coverage: test-coverage
	go tool cover -html=coverage.out -o reports/coverage.html
	go tool cover -func=coverage.out > reports/coverage.txt
	@echo "Coverage report generated:"
	@echo "  HTML: reports/coverage.html"
	@echo "  Text: reports/coverage.txt"

# Clean generated files
clean:
	rm -rf reports/
	rm -f coverage.out
	go clean -cache

# Build the application
build:
	go build -o bin/server main.go

# Run the application
run:
	go run main.go

# Install dependencies
deps:
	go mod tidy
	go install gotest.tools/gotestsum@latest

# Check code quality
lint:
	golangci-lint run

# Format code
fmt:
	go fmt ./...
	go vet ./...

# Debug Testmo CLI integration (Linux/Mac)
testmo-debug:
	@echo "🔍 Debugging Testmo CLI integration..."
	@chmod +x scripts/testmo-debug.sh
	@./scripts/testmo-debug.sh

# Debug Testmo CLI integration (Windows)
testmo-debug-win:
	@echo "🔍 Debugging Testmo CLI integration (Windows)..."
	@powershell -ExecutionPolicy Bypass -File scripts/testmo-debug.ps1

# Test CI workflow locally
ci-test: test-verbose
	@echo "🧪 Testing CI workflow locally..."
	@echo "✅ Tests completed successfully"
	@echo "📊 Coverage report generated"
	@echo "📁 Check reports/ directory for outputs"
	@echo "🔍 Run 'make testmo-debug' or 'make testmo-debug-win' to debug Testmo integration"
