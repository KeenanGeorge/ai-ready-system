# SMA-10 Implementation Setup Guide

This guide walks you through setting up the complete CI → Testmo automation + Linear PR linking system as specified in SMA-10.

## 🎯 Overview

SMA-10 implements a comprehensive CI/CD pipeline that:
- Runs Go tests with coverage
- Generates JUnit XML reports
- Integrates with Testmo for test result management
- Automatically links Linear issues to GitHub PRs
- Enforces coverage thresholds (80% minimum)
- Provides automated PR comments with Testmo links

## 📋 Prerequisites

### Required Accounts & Access
- [ ] GitHub repository with admin access
- [ ] Linear workspace with GitHub integration enabled
- [ ] Testmo workspace with API access
- [ ] Go 1.22+ development environment

### Required Tools
- [ ] Go 1.22+
- [ ] Node.js (for Testmo CLI)
- [ ] Git
- [ ] Make (optional, for build commands)

## 🚀 Step-by-Step Setup

### 1. Repository Configuration

#### 1.1 Enable GitHub Actions
1. Go to your repository on GitHub
2. Navigate to **Actions** tab
3. Click **Enable Actions** if not already enabled

#### 1.2 Configure Branch Protection
1. Go to **Settings > Branches**
2. Click **Add rule** for `main` branch
3. Enable these settings:
   - ✅ **Require a pull request before merging**
   - ✅ **Require status checks to pass before merging**
   - ✅ **Require branches to be up to date before merging**
   - ✅ **Require linear history**
4. In **Status checks that are required**:
   - Add `ci` (our CI workflow)
   - Add `coverage` (coverage threshold check)
5. Click **Create**

### 2. Testmo Configuration

#### 2.1 Get Testmo Credentials
1. Log into your Testmo workspace
2. Go to **Settings > API**
3. Generate a new API token
4. Note your:
   - **Instance URL** (e.g., `https://yourcompany.testmo.net`)
   - **API Token**
   - **Project ID** (numeric ID of your project)

#### 2.2 Configure GitHub Secrets
1. Go to **Settings > Secrets and variables > Actions**
2. Click **New repository secret**
3. Add these secrets:
   ```
   TESTMO_TOKEN=your_api_token_here
   TESTMO_INSTANCE=https://yourcompany.testmo.net
   TESTMO_PROJECT_ID=12345
   ```

### 3. Linear Integration

#### 3.1 Enable GitHub Integration
1. In Linear, go to **Settings > Integrations**
2. Find **GitHub** and click **Configure**
3. Select your repository
4. Enable **Pull Request** integration
5. Configure team automations if desired

#### 3.2 Verify Issue Linking
1. Create a test issue in Linear (e.g., `SMA-10`)
2. Create a feature branch: `feature/sma-10`
3. Make changes and create a PR with title: `Fixes SMA-10`
4. Verify Linear automatically links the PR to the issue

### 4. Local Development Setup

#### 4.1 Install Dependencies
```bash
# Install Go tools
go mod tidy
go install gotest.tools/gotestsum@latest

# Install Testmo CLI
npm install -g @testmo/testmo-cli
```

#### 4.2 Test Local Workflow
```bash
# Run the local test script
# On Linux/Mac:
./scripts/test-local.sh

# On Windows:
.\scripts\test-local.ps1

# Or use Make commands
make test-coverage
make test-verbose
```

### 5. CI Pipeline Verification

#### 5.1 Test the Pipeline
1. Create a feature branch: `feature/test-sma-10`
2. Make a small change (e.g., add a comment)
3. Push and create a PR
4. Watch the GitHub Actions tab
5. Verify the CI workflow runs successfully

#### 5.2 Check Testmo Integration
1. After CI completes, check Testmo
2. Look for automation runs with source `go-ci`
3. Verify test results and coverage are uploaded
4. Check PR comments for Testmo links

## 🔧 Configuration Files

### GitHub Actions Workflow
- **File**: `.github/workflows/ci.yml`
- **Purpose**: Defines the CI pipeline
- **Key Features**:
  - Test execution with coverage
  - JUnit XML generation
  - Testmo integration
  - PR commenting
  - Coverage threshold enforcement

### Testmo Configuration
- **File**: `testmo.config.yml`
- **Purpose**: Testmo CLI configuration
- **Key Features**:
  - Default run settings
  - Result patterns
  - Metadata configuration

### Branch Protection
- **File**: `.github/branch-protection.yml`
- **Purpose**: Documents required protection rules
- **Key Features**:
  - Required status checks
  - Protection rule documentation

## 🧪 Testing the Implementation

### 1. Unit Tests
```bash
# Basic test run
make test

# With coverage
make test-coverage

# With JUnit output
make test-verbose
```

### 2. Coverage Verification
```bash
# Generate coverage reports
make coverage

# Check coverage threshold
make test-coverage
```

### 3. Local CI Simulation
```bash
# Run full local workflow
./scripts/test-local.sh

# Skip Testmo integration
./scripts/test-local.sh --skip-testmo
```

## 📊 Expected Results

### After Successful Implementation

#### GitHub Actions
- ✅ CI workflow runs on every PR and push to main
- ✅ Tests execute with coverage reporting
- ✅ JUnit XML files generated
- ✅ Coverage threshold enforced (80% minimum)

#### Testmo Integration
- ✅ Automation runs created for each CI execution
- ✅ Test results uploaded with metadata
- ✅ Coverage reports available
- ✅ Run IDs captured and logged

#### Linear Integration
- ✅ PRs automatically linked to issues
- ✅ Issue state transitions on PR events
- ✅ Branch names follow convention

#### PR Experience
- ✅ Automated comments with Testmo links
- ✅ Coverage information displayed
- ✅ CI status clearly visible
- ✅ Merge blocked until all checks pass

## 🚨 Troubleshooting

### Common Issues

#### CI Workflow Fails
1. **Check Go version**: Ensure Go 1.22+ is specified
2. **Verify secrets**: Check TESTMO_* secrets are set
3. **Review logs**: Check GitHub Actions logs for specific errors

#### Testmo Integration Issues
1. **API credentials**: Verify token, instance, and project ID
2. **CLI installation**: Ensure Testmo CLI is accessible
3. **Network access**: Check firewall/proxy settings

#### Coverage Threshold Failures
1. **Add more tests**: Increase test coverage above 80%
2. **Check coverage calculation**: Verify coverage.out generation
3. **Review test structure**: Ensure all code paths are tested

#### Linear Linking Problems
1. **GitHub integration**: Verify Linear GitHub integration is enabled
2. **Issue keys**: Ensure PR titles/descriptions contain issue keys
3. **Branch names**: Follow `feature/sma-{issue-number}` convention

### Debug Commands
```bash
# Check test coverage
make test-coverage

# Run verbose tests
make test-verbose

# Clean and rebuild
make clean && make deps

# Test specific components
go test -v ./...
go tool cover -func=coverage.out
```

## 🔄 Maintenance

### Regular Tasks
1. **Update dependencies**: `go mod tidy`
2. **Review coverage**: Ensure tests maintain 80%+ coverage
3. **Monitor CI performance**: Check workflow execution times
4. **Update Testmo CLI**: `npm update -g @testmo/testmo-cli`

### Monitoring
1. **GitHub Actions**: Monitor workflow success rates
2. **Testmo**: Review automation run quality
3. **Linear**: Check issue-PR linking accuracy
4. **Coverage**: Track coverage trends over time

## 📚 Additional Resources

### Documentation
- [SMA-10 Issue](https://linear.app/smart-view-technology/issue/SMA-10)
- [Testmo Documentation](https://docs.testmo.com/)
- [Linear Integration Guide](https://linear.app/docs)
- [GitHub Actions](https://docs.github.com/en/actions)
- [Go Testing](https://golang.org/pkg/testing/)

### Support
- **Linear**: Use `?` in bottom left for help
- **Testmo**: Check documentation or contact support
- **GitHub**: Review Actions documentation
- **Go**: Consult Go documentation and community

## ✅ Completion Checklist

- [ ] GitHub Actions workflow configured
- [ ] Testmo secrets added to repository
- [ ] Branch protection rules enabled
- [ ] Linear GitHub integration active
- [ ] Local development environment set up
- [ ] Tests pass with 80%+ coverage
- [ ] CI pipeline runs successfully
- [ ] Testmo integration working
- [ ] Linear PR linking functional
- [ ] PR comments generated automatically
- [ ] Coverage threshold enforced
- [ ] Documentation updated

---

**Status**: 🎯 SMA-10 Implementation Complete

*This setup guide provides everything needed to implement the complete CI → Testmo automation + Linear PR linking system as specified in SMA-10.*
