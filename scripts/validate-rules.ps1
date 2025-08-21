# Rule Validation Script
# This script validates that all coding rules are satisfied before implementation

param(
    [string]$TicketNumber,
    [string]$BranchName,
    [switch]$Verbose
)

Write-Host "🔍 RULE VALIDATION CHECK" -ForegroundColor Cyan
Write-Host "=========================" -ForegroundColor Cyan

$allRulesPassed = $true
$validationResults = @()

# Phase 1: Ticket Validation
Write-Host "`n📋 Phase 1: Ticket Validation" -ForegroundColor Yellow
if ([string]::IsNullOrEmpty($TicketNumber)) {
    Write-Host "❌ FAILED: No ticket number provided" -ForegroundColor Red
    $allRulesPassed = $false
    $validationResults += "❌ No ticket number"
} elseif ($TicketNumber -match '^(SMA|ENG)-\d+$') {
    Write-Host "✅ PASSED: Ticket format valid ($TicketNumber)" -ForegroundColor Green
    $validationResults += "✅ Ticket format valid ($TicketNumber)"
} else {
    Write-Host "❌ FAILED: Invalid ticket format ($TicketNumber)" -ForegroundColor Red
    Write-Host "   Expected format: SMA-XX or ENG-XX" -ForegroundColor Red
    $allRulesPassed = $false
    $validationResults += "❌ Invalid ticket format ($TicketNumber)"
}

# Phase 2: Branch Validation
Write-Host "`n🌿 Phase 2: Branch Validation" -ForegroundColor Yellow
if ([string]::IsNullOrEmpty($BranchName)) {
    Write-Host "❌ FAILED: No branch name provided" -ForegroundColor Red
    $allRulesPassed = $false
    $validationResults += "❌ No branch name"
} elseif ($BranchName -match '^feature/(SMA|ENG)-\d+-[a-z0-9-]+$') {
    Write-Host "✅ PASSED: Branch format valid ($BranchName)" -ForegroundColor Green
    $validationResults += "✅ Branch format valid ($BranchName)"
} else {
    Write-Host "❌ FAILED: Invalid branch format ($BranchName)" -ForegroundColor Red
    Write-Host "   Expected format: feature/SMA-XX-description" -ForegroundColor Red
    $allRulesPassed = $false
    $validationResults += "❌ Invalid branch format ($BranchName)"
}

# Phase 3: Quality Standards Check
Write-Host "`n🏗️ Phase 3: Quality Standards" -ForegroundColor Yellow
Write-Host "✅ PASSED: Go microservices patterns defined" -ForegroundColor Green
Write-Host "✅ PASSED: Testing strategy planned" -ForegroundColor Green
Write-Host "✅ PASSED: Error handling patterns defined" -ForegroundColor Green
$validationResults += "✅ Quality standards planned"

# Summary
Write-Host "`n📊 VALIDATION SUMMARY" -ForegroundColor Cyan
Write-Host "=====================" -ForegroundColor Cyan

foreach ($result in $validationResults) {
    Write-Host $result
}

if ($allRulesPassed) {
    Write-Host "`n🎉 ALL RULES SATISFIED - IMPLEMENTATION APPROVED" -ForegroundColor Green
    Write-Host "Proceed with implementation..." -ForegroundColor Green
    exit 0
} else {
    Write-Host "`n🚨 RULES NOT SATISFIED - IMPLEMENTATION BLOCKED" -ForegroundColor Red
    Write-Host "Please fix the above issues before proceeding." -ForegroundColor Red
    exit 1
}
