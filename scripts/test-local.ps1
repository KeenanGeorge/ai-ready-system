# Local Testmo Integration Test Script (PowerShell)
# This script helps test the CI workflow locally before pushing

param(
    [switch]$SkipTestmo
)

Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"

Write-Host "🧪 Testing local CI workflow..." -ForegroundColor Green

# Check if required tools are installed
function Test-Requirements {
    Write-Host "Checking requirements..." -ForegroundColor Yellow
    
    if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
        Write-Host "❌ Go is not installed" -ForegroundColor Red
        exit 1
    }
    
    if (-not (Get-Command gotestsum -ErrorAction SilentlyContinue)) {
        Write-Host "📦 Installing gotestsum..." -ForegroundColor Yellow
        go install gotest.tools/gotestsum@latest
    }
    
    if (-not (Get-Command testmo -ErrorAction SilentlyContinue)) {
        Write-Host "📦 Installing Testmo CLI..." -ForegroundColor Yellow
        npm install -g @testmo/testmo-cli
    }
    
    Write-Host "✅ Requirements met" -ForegroundColor Green
}

# Run tests and generate reports
function Invoke-Tests {
    Write-Host "Running tests..." -ForegroundColor Yellow
    
    # Create reports directory
    New-Item -ItemType Directory -Force -Path "reports" | Out-Null
    
    # Run tests with coverage and JUnit output
    gotestsum --format=standard-verbose `
        --junitfile=reports/unit-tests.xml `
        --jsonfile=reports/test-results.json `
        --coverprofile=coverage.out `
        --covermode=atomic `
        ./...
    
    Write-Host "✅ Tests completed" -ForegroundColor Green
}

# Generate coverage reports
function New-CoverageReports {
    Write-Host "Generating coverage reports..." -ForegroundColor Yellow
    
    go tool cover -html=coverage.out -o reports/coverage.html
    go tool cover -func=coverage.out > reports/coverage.txt
    
    # Display coverage summary
    Write-Host "📊 Coverage Summary:" -ForegroundColor Cyan
    go tool cover -func=coverage.out | Select-String "total"
    
    Write-Host "✅ Coverage reports generated" -ForegroundColor Green
}

# Check coverage threshold
function Test-CoverageThreshold {
    Write-Host "Checking coverage threshold..." -ForegroundColor Yellow
    
    $coverageOutput = go tool cover -func=coverage.out | Select-String "total"
    $coverage = [regex]::Match($coverageOutput, '(\d+\.?\d*)%').Groups[1].Value
    
    Write-Host "Total coverage: ${coverage}%" -ForegroundColor Cyan
    
    if ([double]$coverage -lt 80) {
        Write-Host "❌ Coverage below 80% threshold: ${coverage}%" -ForegroundColor Red
        exit 1
    }
    
    Write-Host "✅ Coverage threshold met: ${coverage}%" -ForegroundColor Green
}

# Simulate Testmo integration (if credentials are available)
function Test-TestmoIntegration {
    if ($SkipTestmo) {
        Write-Host "⚠️  Skipping Testmo integration test" -ForegroundColor Yellow
        return
    }
    
    if (-not $env:TESTMO_TOKEN -or -not $env:TESTMO_INSTANCE -or -not $env:TESTMO_PROJECT_ID) {
        Write-Host "⚠️  Testmo credentials not set, skipping integration test" -ForegroundColor Yellow
        Write-Host "   Set TESTMO_TOKEN, TESTMO_INSTANCE, and TESTMO_PROJECT_ID environment variables" -ForegroundColor Yellow
        return
    }
    
    Write-Host "Testing Testmo integration..." -ForegroundColor Yellow
    
    # Get git info
    $branch = git rev-parse --abbrev-ref HEAD
    $commit = git rev-parse --short HEAD
    $goVersion = (go version) -split ' ' | Select-Object -Last 1
    
    # Create a test run
    $runId = testmo run create `
        --name "Local Test: $branch - $commit" `
        --source "go-ci" `
        --milestone "Local Testing" `
        --config "Go $goVersion"
    
    Write-Host "Created Testmo run: $runId" -ForegroundColor Cyan
    
    # Submit results
    testmo run submit `
        --run-id "$runId" `
        --results reports/unit-tests.xml `
        --coverage reports/coverage.txt
    
    Write-Host "Submitted results to Testmo" -ForegroundColor Cyan
    
    # Complete the run
    testmo run complete --run-id "$runId"
    
    Write-Host "✅ Testmo integration test completed" -ForegroundColor Green
}

# Main execution
function Main {
    Test-Requirements
    Invoke-Tests
    New-CoverageReports
    Test-CoverageThreshold
    Test-TestmoIntegration
    
    Write-Host ""
    Write-Host "🎉 Local CI workflow test completed successfully!" -ForegroundColor Green
    Write-Host "📁 Reports generated in: reports/" -ForegroundColor Cyan
    Write-Host "📊 Coverage report: reports/coverage.html" -ForegroundColor Cyan
    Write-Host "📋 JUnit results: reports/unit-tests.xml" -ForegroundColor Cyan
}

# Run main function
Main
