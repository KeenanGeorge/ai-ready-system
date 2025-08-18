# Smart-View AI Ready System

A Go service with integrated CI/CD automation, Testmo test result reporting, and Linear PR linking.

## 🚀 Features

- **Health endpoint** (`/health`) for service monitoring
- **Automated testing** with coverage reporting
- **CI/CD pipeline** via GitHub Actions
- **Testmo integration** for test result management
- **Linear PR linking** with automated issue tracking
- **Coverage thresholds** (80% minimum required)

## 🏗️ Architecture

```
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│   GitHub PR     │───▶│  GitHub Actions  │───▶│     Testmo      │
│                 │    │       CI         │    │   Test Results  │
└─────────────────┘    └──────────────────┘    └─────────────────┘
         │                       │                       │
         │                       ▼                       │
         │              ┌──────────────────┐            │
         │              │   JUnit XML      │            │
         │              │   Coverage       │            │
         │              │   Reports        │            │
         │              └──────────────────┘            │
         │                       │                       │
         ▼                       ▼                       ▼
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│  Linear Issue   │    │  PR Comments     │    │  Branch         │
│   Auto-linking  │    │  with Testmo     │    │  Protection     │
│                 │    │  Results         │    │                 │
└─────────────────┘    └──────────────────┘    └─────────────────┘
```

## 🧪 Testing

### Test Naming Convention

Following the SMA-10 specification, tests use descriptive subtest names that appear in Testmo:

```go
func TestHealthEndpoint(t *testing.T) {
    t.Run("SMA-7: Health endpoint returns ok status", func(t *testing.T) {
        // Test implementation
    })
    
    t.Run("SMA-7: Health endpoint handles GET method correctly", func(t *testing.T) {
        // Test implementation
    })
}
```

### Running Tests

```bash
# Basic test run
make test

# Tests with coverage
make test-coverage

# Tests with verbose output and JUnit XML generation
make test-verbose

# Generate coverage reports
make coverage
```

## 🔄 CI/CD Pipeline

### GitHub Actions Workflow

The CI pipeline automatically:

1. **Runs tests** with coverage
2. **Generates JUnit XML** reports
3. **Creates Testmo runs** with metadata
4. **Submits results** to Testmo
5. **Comments on PRs** with Testmo links
6. **Enforces coverage** thresholds (80% minimum)
7. **Uploads artifacts** for review

### Workflow Triggers

- **Push to main**: Full CI run
- **Pull Request**: CI run + PR comment + Testmo integration

### Required Secrets

Configure these in your GitHub repository settings:

```bash
TESTMO_TOKEN=your_testmo_api_token
TESTMO_INSTANCE=your_testmo_instance_url
TESTMO_PROJECT_ID=your_testmo_project_id
```

## 📊 Testmo Integration

### Automation Run Details

Each CI execution creates a Testmo run with:

- **Name**: `CI: {branch} - {commit_sha}`
- **Source**: `go-ci`
- **Milestone**: `CI Automation`
- **Config**: `Go 1.22`

### Test Results

- **JUnit XML**: Test execution results
- **Coverage**: Code coverage metrics
- **Timing**: Test execution times
- **Logs**: Console output and errors

## 🔗 Linear PR Linking

### Automatic Issue Linking

Linear automatically links PRs to issues when:

- **PR title** contains issue key (e.g., `Fixes SMA-10`)
- **PR description** references issue keys
- **Branch names** follow convention (e.g., `feature/sma-10`)

### PR Comments

Each PR receives an automated comment with:

- ✅ CI status
- 📊 Testmo run link
- 📈 Coverage information
- 🔗 Run ID for reference

## 🛡️ Branch Protection

### Required Checks

- **CI workflow** must pass
- **Coverage threshold** (80%) must be met
- **Branches** must be up to date
- **Linear history** required

### Configuration

Enable branch protection in GitHub:

1. Go to **Settings > Branches**
2. Add rule for `main` branch
3. Enable required status checks
4. Configure protection rules

## 🚀 Development

### Prerequisites

- Go 1.22+
- Node.js (for Testmo CLI)
- Make (for build commands)

### Setup

```bash
# Clone repository
git clone <repository-url>
cd ai-ready-system

# Install dependencies
make deps

# Run tests
make test

# Start service
make run
```

### Available Commands

```bash
make help          # Show available commands
make test          # Run tests
make coverage      # Generate coverage report
make build         # Build binary
make run           # Run service
make clean         # Clean generated files
make fmt           # Format code
make lint          # Run linter
```

## 📁 Project Structure

```
ai-ready-system/
├── .github/
│   ├── workflows/
│   │   └── ci.yml              # CI/CD pipeline
│   └── branch-protection.yml   # Branch protection config
├── main.go                     # Main application
├── main_test.go               # Test suite
├── Makefile                   # Build commands
├── go.mod                     # Go module file
└── README.md                  # This file
```

## 🔧 Configuration

### Testmo CLI

The CI pipeline uses Testmo CLI for:

- Creating automation runs
- Submitting test results
- Managing run metadata
- Completing test executions

### Coverage Thresholds

- **Minimum**: 80%
- **Enforcement**: Blocks merges if below threshold
- **Reporting**: HTML and text formats generated

## 📈 Monitoring

### Health Endpoint

```
GET /health
Response: "ok"
```

### CI Status

Monitor CI status via:

- GitHub Actions tab
- PR status checks
- Testmo automation runs
- Linear issue updates

## 🚨 Troubleshooting

### Common Issues

1. **CI fails on coverage**: Ensure tests cover 80% of code
2. **Testmo integration fails**: Check API tokens and configuration
3. **Linear linking issues**: Verify issue keys in PR titles/descriptions
4. **Branch protection**: Ensure required checks are configured

### Debug Commands

```bash
# Check test coverage
make test-coverage

# Run verbose tests
make test-verbose

# Clean and rebuild
make clean && make deps
```

## 📚 References

- [SMA-10 Issue](https://linear.app/smart-view-technology/issue/SMA-10)
- [Testmo Documentation](https://docs.testmo.com/)
- [Linear Integration Guide](https://linear.app/docs)
- [GitHub Actions](https://docs.github.com/en/actions)
- [Go Testing](https://golang.org/pkg/testing/)

## 🤝 Contributing

1. Create feature branch: `feature/sma-{issue-number}`
2. Include issue key in PR title/description
3. Ensure tests pass and coverage meets threshold
4. Wait for CI approval before merging

---

**Status**: ✅ SMA-10 Implementation Complete

*This implementation provides full CI/CD automation with Testmo integration and Linear PR linking as specified in the requirements.*
