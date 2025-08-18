#!/bin/bash

# Local Testmo Integration Test Script
# This script helps test the CI workflow locally before pushing

set -e

echo "🧪 Testing local CI workflow..."

# Check if required tools are installed
check_requirements() {
    echo "Checking requirements..."
    
    if ! command -v go &> /dev/null; then
        echo "❌ Go is not installed"
        exit 1
    fi
    
    if ! command -v gotestsum &> /dev/null; then
        echo "📦 Installing gotestsum..."
        go install gotest.tools/gotestsum@latest
    fi
    
    if ! command -v testmo &> /dev/null; then
        echo "📦 Installing Testmo CLI..."
        npm install -g @testmo/testmo-cli
    fi
    
    echo "✅ Requirements met"
}

# Run tests and generate reports
run_tests() {
    echo "Running tests..."
    
    # Create reports directory
    mkdir -p reports
    
    # Run tests with coverage and JUnit output
    gotestsum --format=standard-verbose \
        --junitfile=reports/unit-tests.xml \
        --jsonfile=reports/test-results.json \
        --coverprofile=coverage.out \
        --covermode=atomic \
        ./...
    
    echo "✅ Tests completed"
}

# Generate coverage reports
generate_coverage() {
    echo "Generating coverage reports..."
    
    go tool cover -html=coverage.out -o reports/coverage.html
    go tool cover -func=coverage.out > reports/coverage.txt
    
    # Display coverage summary
    echo "📊 Coverage Summary:"
    go tool cover -func=coverage.out | grep total
    
    echo "✅ Coverage reports generated"
}

# Check coverage threshold
check_coverage() {
    echo "Checking coverage threshold..."
    
    COVERAGE=$(go tool cover -func=coverage.out | grep total | awk '{print $3}' | sed 's/%//')
    echo "Total coverage: ${COVERAGE}%"
    
    if (( $(echo "$COVERAGE < 80" | bc -l) )); then
        echo "❌ Coverage below 80% threshold: ${COVERAGE}%"
        exit 1
    fi
    
    echo "✅ Coverage threshold met: ${COVERAGE}%"
}

# Simulate Testmo integration (if credentials are available)
test_testmo_integration() {
    if [ -z "$TESTMO_TOKEN" ] || [ -z "$TESTMO_INSTANCE" ] || [ -z "$TESTMO_PROJECT_ID" ]; then
        echo "⚠️  Testmo credentials not set, skipping integration test"
        echo "   Set TESTMO_TOKEN, TESTMO_INSTANCE, and TESTMO_PROJECT_ID environment variables"
        return 0
    fi
    
    echo "Testing Testmo integration..."
    
    # Create a test run
    RUN_ID=$(testmo run create \
        --name "Local Test: $(git rev-parse --abbrev-ref HEAD) - $(git rev-parse --short HEAD)" \
        --source "go-ci" \
        --milestone "Local Testing" \
        --config "Go $(go version | awk '{print $3}')")
    
    echo "Created Testmo run: $RUN_ID"
    
    # Submit results
    testmo run submit \
        --run-id "$RUN_ID" \
        --results reports/unit-tests.xml \
        --coverage reports/coverage.txt
    
    echo "Submitted results to Testmo"
    
    # Complete the run
    testmo run complete --run-id "$RUN_ID"
    
    echo "✅ Testmo integration test completed"
}

# Main execution
main() {
    check_requirements
    run_tests
    generate_coverage
    check_coverage
    test_testmo_integration
    
    echo ""
    echo "🎉 Local CI workflow test completed successfully!"
    echo "📁 Reports generated in: reports/"
    echo "📊 Coverage report: reports/coverage.html"
    echo "📋 JUnit results: reports/unit-tests.xml"
}

# Run main function
main "$@"
