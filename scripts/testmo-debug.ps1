# Testmo CLI Debug Script (PowerShell)
# This script helps troubleshoot Testmo CLI integration issues

Write-Host "🔍 Testmo CLI Debug Script" -ForegroundColor Cyan
Write-Host "==========================" -ForegroundColor Cyan

# Check if Testmo CLI is installed
Write-Host "`n1. Checking Testmo CLI installation..." -ForegroundColor Yellow
try {
    $testmoVersion = testmo --version 2>$null
    if ($testmoVersion) {
        Write-Host "✅ Testmo CLI is installed" -ForegroundColor Green
        Write-Host $testmoVersion
    } else {
        throw "Testmo CLI not found"
    }
} catch {
    Write-Host "❌ Testmo CLI is not installed" -ForegroundColor Red
    Write-Host "Installing Testmo CLI..." -ForegroundColor Yellow
    npm install -g @testmo/testmo-cli
}

Write-Host ""

# Check Testmo CLI help
Write-Host "2. Checking Testmo CLI help..." -ForegroundColor Yellow
try {
    testmo --help
} catch {
    Write-Host "❌ Failed to get Testmo CLI help" -ForegroundColor Red
}

Write-Host ""

# Check available commands
Write-Host "3. Checking available commands..." -ForegroundColor Yellow
try {
    testmo automation --help
} catch {
    Write-Host "❌ Failed to get automation help" -ForegroundColor Red
}

Write-Host ""

# Test Testmo CLI commands
Write-Host "4. Testing Testmo CLI commands..." -ForegroundColor Yellow

# Test run create command
Write-Host "Testing 'testmo automation run create'..." -ForegroundColor Yellow
try {
    testmo automation run create --help | Out-Null
    Write-Host "✅ 'testmo automation run create' command works" -ForegroundColor Green
} catch {
    Write-Host "❌ 'testmo automation run create' command failed" -ForegroundColor Red
}

Write-Host ""

# Test run submit command
Write-Host "Testing 'testmo automation run submit'..." -ForegroundColor Yellow
try {
    testmo automation run submit --help | Out-Null
    Write-Host "✅ 'testmo automation run submit' command works" -ForegroundColor Green
} catch {
    Write-Host "❌ 'testmo automation run submit' command failed" -ForegroundColor Red
}

Write-Host ""

# Test run complete command
Write-Host "Testing 'testmo automation run complete'..." -ForegroundColor Yellow
try {
    testmo automation run complete --help | Out-Null
    Write-Host "✅ 'testmo automation run complete' command works" -ForegroundColor Green
} catch {
    Write-Host "❌ 'testmo automation run complete' command failed" -ForegroundColor Red
}

Write-Host ""

# Check environment variables
Write-Host "5. Checking environment variables..." -ForegroundColor Yellow
Write-Host "TESTMO_INSTANCE: $(if ($env:TESTMO_INSTANCE) { $env:TESTMO_INSTANCE } else { 'Not set' })"
Write-Host "TESTMO_TOKEN: $(if ($env:TESTMO_TOKEN) { 'Set' } else { 'Not set' })"
Write-Host "TESTMO_PROJECT_ID: $(if ($env:TESTMO_PROJECT_ID) { $env:TESTMO_PROJECT_ID } else { 'Not set' })"

Write-Host ""

# Test configuration file
Write-Host "6. Checking Testmo configuration..." -ForegroundColor Yellow
if (Test-Path "testmo.config.yml") {
    Write-Host "✅ testmo.config.yml exists" -ForegroundColor Green
    Write-Host "Configuration contents:" -ForegroundColor Yellow
    Get-Content "testmo.config.yml"
} else {
    Write-Host "❌ testmo.config.yml not found" -ForegroundColor Red
}

Write-Host ""
Write-Host "🔍 Debug script completed!" -ForegroundColor Cyan
Write-Host "If you see any ❌ errors above, those need to be fixed before the CI pipeline will work." -ForegroundColor Yellow
