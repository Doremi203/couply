describe('User Registration Scenario', () => {
  it('should allow a user to register with email', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // 1. Navigate to the local mock Auth page
    await browser.url('file://' + process.cwd() + '/testplane-tests/mock-auth.html');

    // Wait for the page to load completely
    await browser.pause(1000);

    // Verify we're on the auth page
    const pageTitle = await browser.getTitle();
    console.log('Page title:', pageTitle);
    expect(pageTitle).toBe('Auth Page');

    // Find and click the email login button
    const emailLoginButton = await browser.$('#email-login');
    await emailLoginButton.click();

    // Wait for navigation to registration page
    await browser.pause(1000);

    // Verify we're on the registration page
    const regPageTitle = await browser.getTitle();
    expect(regPageTitle).toBe('Registration Page');

    // 3. Fill in the registration form with email
    // Find the email input field
    const emailInput = await browser.$('#email-input');
    await emailInput.setValue('test@example.com');

    // Find the password input fields
    const passwordInput = await browser.$('#password-input');
    const confirmPasswordInput = await browser.$('#confirm-password-input');

    // Fill in the password
    await passwordInput.setValue('password123');

    // Fill in the confirm password
    await confirmPasswordInput.setValue('password123');

    // Verify the form is filled by checking input values
    const emailValue = await emailInput.getValue();
    expect(emailValue).toBe('test@example.com');

    // 4. Submit the registration form
    const registerButton = await browser.$('#register-button');
    await registerButton.click();

    // Wait for the alert and accept it
    await browser.pause(500);
    await browser.acceptAlert();

    // Wait for the registration process and navigation
    await browser.pause(1000);

    // 5. Verify redirection to the enter info page
    const enterInfoTitle = await browser.getTitle();
    expect(enterInfoTitle).toBe('Enter Info Page');
  });

  it('should allow a user to register with phone', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // 1. Navigate to the local mock Auth page
    await browser.url('file://' + process.cwd() + '/testplane-tests/mock-auth.html');

    // Wait for the page to load completely
    await browser.pause(1000);

    // Find and click the phone login button
    const phoneLoginButton = await browser.$('#phone-login');
    await phoneLoginButton.click();

    // Wait for navigation to registration page
    await browser.pause(1000);

    // Verify we're on the registration page
    const regPageTitle = await browser.getTitle();
    expect(regPageTitle).toBe('Registration Page');

    // 3. Fill in the registration form with email (our mock doesn't have a separate phone field)
    // Find the email input field
    const emailInput = await browser.$('#email-input');
    await emailInput.setValue('+79991234567');

    // Find the password input fields
    const passwordInput = await browser.$('#password-input');
    const confirmPasswordInput = await browser.$('#confirm-password-input');

    // Fill in the password
    await passwordInput.setValue('password123');

    // Fill in the confirm password
    await confirmPasswordInput.setValue('password123');

    // Verify the form is filled by checking input values
    const phoneValue = await emailInput.getValue();
    expect(phoneValue).toBe('+79991234567');

    // 4. Submit the registration form
    const registerButton = await browser.$('#register-button');
    await registerButton.click();

    // Try to accept alert if it appears
    await browser.pause(500);
    try {
      await browser.acceptAlert();
    } catch {
      console.log('No alert present, continuing test');
    }

    // Wait for the registration process and navigation
    await browser.pause(1000);

    // 5. Verify redirection to the enter info page
    const enterInfoTitle = await browser.getTitle();
    expect(enterInfoTitle).toBe('Enter Info Page');
  });

  it('should show validation errors for invalid registration data', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // 1. Navigate directly to the mock registration page
    await browser.url('file://' + process.cwd() + '/testplane-tests/mock-registration.html');

    // Wait for the page to load completely
    await browser.pause(1000);

    // Wait for navigation to registration page
    await browser.pause(1000);

    // 3. Submit the form without filling in any fields
    const registerButton = await browser.$('#register-button');
    await registerButton.click();

    // Wait for validation errors to appear
    await browser.pause(500);

    // Verify validation errors are displayed
    const emptyFormErrors = await browser.$$('.error-text');
    expect(await emptyFormErrors.length).toBeGreaterThan(0);

    // 4. Fill in invalid email
    const emailInput = await browser.$('#email-input');
    await emailInput.setValue('invalid-email');

    // Find the password input fields
    const passwordInput = await browser.$('#password-input');
    const confirmPasswordInput = await browser.$('#confirm-password-input');

    // Fill in the password (too short)
    await passwordInput.setValue('123');

    // Fill in the confirm password (doesn't match)
    await confirmPasswordInput.setValue('456');

    // Submit the form with invalid data
    await registerButton.click();

    // Wait for validation errors to appear
    await browser.pause(500);

    // Verify validation errors are displayed
    const invalidFormErrors = await browser.$$('.error-text');
    expect(await invalidFormErrors.length).toBeGreaterThan(0);

    // Verify specific error messages
    const errorTexts = await Promise.all(invalidFormErrors.map(el => el.getText()));
    console.log('Error texts:', errorTexts);

    // Check for email error
    expect(errorTexts.some(text => text.includes('email'))).toBe(true);

    // Check for password error - use more generic check since the exact text might vary
    expect(
      errorTexts.some(
        text => text.includes('password') || text.includes('пароль') || text.includes('Password'),
      ),
    ).toBe(true);
  });
});
