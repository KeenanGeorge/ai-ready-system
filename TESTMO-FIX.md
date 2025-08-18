# Testmo CLI Integration Fix

## Issue Description

The CI pipeline was failing with the error:
```
error: unknown command 'run'
Error: Process completed with exit code 1
```

This occurred because the Testmo CLI commands were using incorrect syntax.

## Root Cause

The original CI workflow used:
```bash
testmo run create
testmo run submit
testmo run complete
```

However, the correct Testmo CLI syntax is:
```bash
testmo automation run create
testmo automation run submit
testmo automation run complete
```

## Solution Implemented

### 1. Fixed CI Workflow (`.github/workflows/ci.yml`)

- **Updated Testmo commands** to use correct `automation run` syntax
- **Added error handling** for Testmo CLI failures
- **Enhanced workflow triggers** to include `features` branch
- **Added conditional PR comments** for both success and failure scenarios
- **Improved error reporting** with detailed failure information

### 2. Added Debug Scripts

#### Linux/Mac (`scripts/testmo-debug.sh`)
```bash
make testmo-debug
```

#### Windows (`scripts/testmo-debug.ps1`)
```powershell
make testmo-debug-win
```

These scripts help troubleshoot Testmo CLI integration by:
- Checking Testmo CLI installation
- Verifying command syntax
- Testing all required commands
- Checking environment variables
- Validating configuration files

### 3. Enhanced Makefile

Added new targets:
- `testmo-debug` - Debug Testmo CLI on Linux/Mac
- `testmo-debug-win` - Debug Testmo CLI on Windows
- `ci-test` - Test CI workflow locally

## Usage

### Debug Testmo Integration

**Linux/Mac:**
```bash
make testmo-debug
```

**Windows:**
```bash
make testmo-debug-win
```

### Test CI Workflow Locally

```bash
make ci-test
```

This will:
1. Run tests with coverage
2. Generate JUnit XML reports
3. Create coverage reports
4. Verify all CI outputs

### Manual Testmo CLI Testing

```bash
# Install Testmo CLI
npm install -g @testmo/testmo-cli

# Test commands
testmo automation run create --help
testmo automation run submit --help
testmo automation run complete --help
```

## Environment Variables Required

Ensure these are set in your GitHub repository secrets:

```bash
TESTMO_TOKEN=your_testmo_api_token
TESTMO_INSTANCE=your_testmo_instance_url
TESTMO_PROJECT_ID=your_testmo_project_id
```

## Configuration

The Testmo configuration is in `testmo.config.yml`:

```yaml
instance: ${TESTMO_INSTANCE}
token: ${TESTMO_TOKEN}
project_id: ${TESTMO_PROJECT_ID}

run:
  source: go-ci
  milestone: CI Automation
  config: Go 1.22
```

## Troubleshooting

### Common Issues

1. **"unknown command 'run'"**
   - **Solution**: Use `testmo automation run` instead of `testmo run`

2. **Testmo CLI not found**
   - **Solution**: Run `npm install -g @testmo/testmo-cli`

3. **Authentication failed**
   - **Solution**: Verify `TESTMO_TOKEN` is set correctly

4. **Invalid project ID**
   - **Solution**: Check `TESTMO_PROJECT_ID` value

### Debug Steps

1. **Run debug script**:
   ```bash
   make testmo-debug  # Linux/Mac
   make testmo-debug-win  # Windows
   ```

2. **Check environment variables**:
   ```bash
   echo $TESTMO_INSTANCE
   echo $TESTMO_TOKEN
   echo $TESTMO_PROJECT_ID
   ```

3. **Test Testmo CLI manually**:
   ```bash
   testmo --version
   testmo automation --help
   ```

4. **Verify configuration file**:
   ```bash
   cat testmo.config.yml
   ```

## CI Workflow Changes

### Before (Broken)
```yaml
- name: Create Testmo run
  run: |
    RUN_ID=$(testmo run create \
      --name "CI: ${{ steps.commit.outputs.branch }} - ${{ steps.commit.outputs.sha }}" \
      --source "${{ env.TESTMO_SOURCE }}" \
      --milestone "CI Automation" \
      --config "Go ${{ env.GO_VERSION }}")
```

### After (Fixed)
```yaml
- name: Create Testmo run
  run: |
    RUN_ID=$(testmo automation run create \
      --name "CI: ${{ steps.commit.outputs.branch }} - ${{ steps.commit.outputs.sha }}" \
      --source "${{ env.TESTMO_SOURCE }}" \
      --milestone "CI Automation" \
      --config "Go ${{ env.GO_VERSION }}" 2>&1) || {
      echo "Failed to create Testmo run: $RUN_ID"
      echo "run_id=failed" >> $GITHUB_OUTPUT
      exit 1
    }
```

## Benefits of the Fix

1. **Resolves CI failures** - Fixes the "unknown command 'run'" error
2. **Better error handling** - Provides clear feedback when Testmo integration fails
3. **Enhanced debugging** - Debug scripts help troubleshoot issues locally
4. **Improved reliability** - Conditional execution prevents cascading failures
5. **Better user experience** - Clear PR comments for both success and failure cases

## Testing the Fix

1. **Push changes** to trigger CI workflow
2. **Monitor CI execution** for Testmo integration steps
3. **Check PR comments** for Testmo run links or failure details
4. **Verify artifacts** are uploaded correctly
5. **Confirm coverage threshold** enforcement works

## Future Improvements

- Add Testmo CLI version pinning for consistency
- Implement retry logic for transient failures
- Add Testmo integration health checks
- Create Testmo run templates for different test types
- Add Testmo metrics and reporting

---

**Status**: ✅ Implemented and Ready for Testing

This fix addresses the core Testmo CLI integration issue and provides comprehensive debugging tools for future troubleshooting.
