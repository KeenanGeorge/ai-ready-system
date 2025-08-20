#!/bin/bash

# Testmo CLI Debug Script
# This script helps troubleshoot Testmo CLI integration issues

set -e

echo "🔍 Testmo CLI Debug Script"
echo "=========================="

# Check if Testmo CLI is installed
echo "1. Checking Testmo CLI installation..."
if command -v testmo &> /dev/null; then
    echo "✅ Testmo CLI is installed"
    testmo --version
else
    echo "❌ Testmo CLI is not installed"
    echo "Installing Testmo CLI..."
    npm install -g @testmo/testmo-cli
fi

echo ""

# Check Testmo CLI help
echo "2. Checking Testmo CLI help..."
testmo --help

echo ""

# Check available commands
echo "3. Checking available commands..."
testmo automation --help

echo ""

# Test Testmo CLI commands
echo "4. Testing Testmo CLI commands..."

# Test run create command
echo "Testing 'testmo automation run create'..."
if testmo automation run create --help &> /dev/null; then
    echo "✅ 'testmo automation run create' command works"
else
    echo "❌ 'testmo automation run create' command failed"
fi

echo ""

# Test run submit command
echo "Testing 'testmo automation run submit'..."
if testmo automation run submit --help &> /dev/null; then
    echo "✅ 'testmo automation run submit' command works"
else
    echo "❌ 'testmo automation run submit' command failed"
fi

echo ""

# Test run complete command
echo "Testing 'testmo automation run complete'..."
if testmo automation run complete --help &> /dev/null; then
    echo "✅ 'testmo automation run complete' command works"
else
    echo "❌ 'testmo automation run complete' command failed"
fi

echo ""

# Check environment variables
echo "5. Checking environment variables..."
echo "TESTMO_INSTANCE: ${TESTMO_INSTANCE:-'Not set'}"
echo "TESTMO_TOKEN: ${TESTMO_TOKEN:+'Set'}"
echo "TESTMO_PROJECT_ID: ${TESTMO_PROJECT_ID:-'Not set'}"

echo ""

# Test configuration file
echo "6. Checking Testmo configuration..."
if [ -f "testmo.config.yml" ]; then
    echo "✅ testmo.config.yml exists"
    echo "Configuration contents:"
    cat testmo.config.yml
else
    echo "❌ testmo.config.yml not found"
fi

echo ""

echo "🔍 Debug script completed!"
echo "If you see any ❌ errors above, those need to be fixed before the CI pipeline will work."
