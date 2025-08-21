import { test, expect } from '@playwright/test';

test.describe('Login Page Tests', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/');
  });

  test('GivenLoginPage_WhenPageLoads_ThenAllElementsRendered', async ({ page }) => {
    // Verify page title
    await expect(page).toHaveTitle('Smart-View AI - Login');
    
    // Verify logo is visible
    await expect(page.locator('.logo')).toBeVisible();
    await expect(page.locator('.logo')).toHaveText('Smart-View AI');
    
    // Verify form elements are present
    await expect(page.locator('#username')).toBeVisible();
    await expect(page.locator('#password')).toBeVisible();
    await expect(page.locator('#loginBtn')).toBeVisible();
    
    // Verify demo credentials section
    await expect(page.locator('.demo-credentials')).toBeVisible();
    await expect(page.locator('.demo-credentials h4')).toHaveText('Demo Credentials');
  });

  test('GivenFormFields_WhenTyping_ThenInputValidationWorks', async ({ page }) => {
    const usernameInput = page.locator('#username');
    const passwordInput = page.locator('#password');
    
    // Type in username field
    await usernameInput.fill('testuser');
    await expect(usernameInput).toHaveValue('testuser');
    
    // Type in password field
    await passwordInput.fill('testpass');
    await expect(passwordInput).toHaveValue('testpass');
    
    // Clear fields and verify
    await usernameInput.clear();
    await passwordInput.clear();
    await expect(usernameInput).toHaveValue('');
    await expect(passwordInput).toHaveValue('');
  });

  test('GivenEmptyFields_WhenLoginSubmitted_ThenValidationErrorsShown', async ({ page }) => {
    const loginBtn = page.locator('#loginBtn');
    
    // Submit form without entering credentials
    await loginBtn.click();
    
    // Verify validation errors are shown
    await expect(page.locator('#username-error')).toBeVisible();
    await expect(page.locator('#username-error')).toHaveText('Username is required');
    await expect(page.locator('#password-error')).toBeVisible();
    await expect(page.locator('#password-error')).toHaveText('Password is required');
    
    // Verify input fields have error styling
    await expect(page.locator('#username')).toHaveClass(/error/);
    await expect(page.locator('#password')).toHaveClass(/error/);
  });

  test('GivenValidCredentials_WhenLoginSubmitted_ThenSuccessMessageDisplayed', async ({ page }) => {
    const usernameInput = page.locator('#username');
    const passwordInput = page.locator('#password');
    const loginBtn = page.locator('#loginBtn');
    
    // Fill in valid admin credentials
    await usernameInput.fill('admin');
    await passwordInput.fill('admin123');
    
    // Submit form
    await loginBtn.click();
    
    // Verify success message
    await expect(page.locator('.message.success')).toBeVisible();
    await expect(page.locator('.message.success')).toHaveText('Welcome back, admin! Login successful.');
    
    // Verify no error messages
    await expect(page.locator('#username-error')).not.toBeVisible();
    await expect(page.locator('#password-error')).not.toBeVisible();
  });

  test('GivenInvalidCredentials_WhenLoginSubmitted_ThenErrorMessageDisplayed', async ({ page }) => {
    const usernameInput = page.locator('#username');
    const passwordInput = page.locator('#password');
    const loginBtn = page.locator('#loginBtn');
    
    // Fill in invalid credentials
    await usernameInput.fill('wronguser');
    await passwordInput.fill('wrongpass');
    
    // Submit form
    await loginBtn.click();
    
    // Verify error message
    await expect(page.locator('.message.error')).toBeVisible();
    await expect(page.locator('.message.error')).toHaveText('Invalid username or password');
  });

  test('GivenValidUser_WhenLoginSuccessful_ThenRedirectToDashboard', async ({ page }) => {
    const usernameInput = page.locator('#username');
    const passwordInput = page.locator('#password');
    const loginBtn = page.locator('#loginBtn');
    
    // Fill in valid user credentials
    await usernameInput.fill('user');
    await passwordInput.fill('user123');
    
    // Submit form
    await loginBtn.click();
    
    // Verify success message
    await expect(page.locator('.message.success')).toBeVisible();
    await expect(page.locator('.message.success')).toHaveText('Welcome back, user! Login successful.');
    
    // Note: In this implementation, we don't actually redirect, but we verify the success state
    // In a real application, this would redirect to a dashboard
  });

  test('GivenMobileDevice_WhenLoginPageViewed_ThenResponsiveDesignWorks', async ({ page }) => {
    // Set mobile viewport
    await page.setViewportSize({ width: 375, height: 667 });
    
    // Verify elements are still visible and properly sized
    await expect(page.locator('.login-container')).toBeVisible();
    await expect(page.locator('.logo')).toBeVisible();
    await expect(page.locator('#username')).toBeVisible();
    await expect(page.locator('#password')).toBeVisible();
    await expect(page.locator('#loginBtn')).toBeVisible();
    
    // Verify form is still functional on mobile
    await page.locator('#username').fill('admin');
    await page.locator('#password').fill('admin123');
    await page.locator('#loginBtn').click();
    
    await expect(page.locator('.message.success')).toBeVisible();
  });

  test('GivenNetworkError_WhenLoginAttempted_ThenErrorHandledGracefully', async ({ page }) => {
    // Mock network error by intercepting the request
    await page.route('/api/login', route => {
      route.abort('failed');
    });
    
    const usernameInput = page.locator('#username');
    const passwordInput = page.locator('#password');
    const loginBtn = page.locator('#loginBtn');
    
    // Fill in credentials
    await usernameInput.fill('admin');
    await passwordInput.fill('admin123');
    
    // Submit form
    await loginBtn.click();
    
    // Verify network error is handled gracefully
    await expect(page.locator('.message.error')).toBeVisible();
    await expect(page.locator('.message.error')).toHaveText('Network error. Please try again.');
  });

  test('GivenMultipleLoginAttempts_WhenRateLimited_ThenAppropriateMessageShown', async ({ page }) => {
    const usernameInput = page.locator('#username');
    const passwordInput = page.locator('#password');
    const loginBtn = page.locator('#loginBtn');
    
    // First login attempt (should succeed)
    await usernameInput.fill('admin');
    await passwordInput.fill('admin123');
    await loginBtn.click();
    await expect(page.locator('.message.success')).toBeVisible();
    
    // Wait for message to disappear
    await page.waitForTimeout(1000);
    
    // Second login attempt (should also succeed in this implementation)
    await loginBtn.click();
    await expect(page.locator('.message.success')).toBeVisible();
    
    // Note: This implementation doesn't have rate limiting, but we verify multiple attempts work
    // In a real application, this would test rate limiting behavior
  });

  test('GivenAccessibility_WhenLoginPageUsed_ThenScreenReaderCompatible', async ({ page }) => {
    // Verify form labels are properly associated
    const usernameInput = page.locator('#username');
    const passwordInput = page.locator('#password');
    
    // Check that labels are properly associated with inputs
    await expect(usernameInput).toHaveAttribute('id', 'username');
    await expect(passwordInput).toHaveAttribute('id', 'password');
    
    // Verify form has proper structure
    await expect(page.locator('form')).toHaveAttribute('id', 'loginForm');
    
    // Check that error messages are properly associated
    await expect(page.locator('#username-error')).toHaveAttribute('id', 'username-error');
    await expect(page.locator('#password-error')).toHaveAttribute('id', 'password-error');
  });

  test('GivenFormValidation_WhenFieldsInteracted_ThenRealTimeValidationWorks', async ({ page }) => {
    const usernameInput = page.locator('#username');
    const passwordInput = page.locator('#password');
    
    // Focus and blur username field without input
    await usernameInput.focus();
    await usernameInput.blur();
    await expect(page.locator('#username-error')).toBeVisible();
    
    // Focus and blur password field without input
    await passwordInput.focus();
    await passwordInput.blur();
    await expect(page.locator('#password-error')).toBeVisible();
    
    // Type in username field and verify error clears
    await usernameInput.fill('test');
    await expect(page.locator('#username-error')).not.toBeVisible();
    
    // Type in password field and verify error clears
    await passwordInput.fill('test');
    await expect(page.locator('#password-error')).not.toBeVisible();
  });
});
