# Simple Test Coverage Analysis Script
# This script analyzes current test coverage and identifies gaps

Write-Host "🧪 TEST COVERAGE ANALYSIS TOOL" -ForegroundColor Cyan
Write-Host "=================================" -ForegroundColor Cyan

Write-Host "`n📊 ANALYZING CURRENT TEST COVERAGE..." -ForegroundColor Yellow

# Run tests with coverage
Write-Host "Running Go tests with coverage..." -ForegroundColor Green
go test -coverprofile=coverage.out ./...

# Generate coverage report
Write-Host "Generating coverage report..." -ForegroundColor Green
go tool cover -func=coverage.out > coverage-analysis.txt

# Parse coverage data
$coverageData = Get-Content coverage-analysis.txt | Select-String "total:"
if ($coverageData) {
    $coveragePercent = ($coverageData -split '\s+')[2] -replace '%', ''
    Write-Host "`n📈 COVERAGE ANALYSIS RESULTS:" -ForegroundColor Cyan
    Write-Host "Current Coverage: $coveragePercent%" -ForegroundColor Yellow
    
    if ([double]$coveragePercent -lt 80) {
        Write-Host "❌ BELOW TARGET: Need to reach 80% coverage" -ForegroundColor Red
        $gap = 80 - [double]$coveragePercent
        Write-Host "Coverage Gap: $gap%" -ForegroundColor Red
    } else {
        Write-Host "✅ TARGET ACHIEVED: 80%+ coverage reached" -ForegroundColor Green
    }
}

# Show detailed coverage by package
Write-Host "`n📋 PACKAGE COVERAGE BREAKDOWN:" -ForegroundColor Cyan
Get-Content coverage-analysis.txt | Where-Object { $_ -match '^\S+\.go:\d+:\s+\w+\s+\d+\.\d+%' } | ForEach-Object {
    $parts = $_ -split '\s+'
    $package = $parts[0]
    $coverage = $parts[2]
    Write-Host "$package : $coverage" -ForegroundColor White
}

Write-Host "`n✅ Coverage analysis completed!" -ForegroundColor Green
Write-Host "Check coverage-analysis.txt for detailed results" -ForegroundColor Cyan
