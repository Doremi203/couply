describe('User Registration Scenario', () => {
  it('should allow a user to register with email', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // 1. Navigate to the Auth page
    await browser.url('https://testing.couply.ru/auth');

    // Wait for the page to load
    await browser.pause(1000);

    // Take a screenshot of the auth page
    await browser.assertView('auth-page', 'body', {
      allowViewportOverflow: true,
    });

    // 2. Click on the "login with email" button to navigate to registration
    const emailLoginButton = await browser.$('button:nth-of-type(4)');
    await emailLoginButton.click();

    // Wait for navigation to registration page
    await browser.pause(1000);

    // Take a screenshot of the registration page
    await browser.assertView('registration-page', 'body', {
      allowViewportOverflow: true,
    });

    // 3. Fill in the registration form with email
    // Find the email input field
    const emailInput = await browser.$('input[type="email"]');
    await emailInput.setValue('test@example.com');

    // Find the password input fields
    const passwordInputs = await browser.$$('input[type="password"]');

    // Fill in the password
    await passwordInputs[0].setValue('password123');

    // Fill in the confirm password
    await passwordInputs[1].setValue('password123');

    // Take a screenshot of the filled form
    await browser.assertView('registration-form-filled', 'body', {
      allowViewportOverflow: true,
    });

    // 4. Submit the registration form
    const registerButton = await browser.$('button');
    await registerButton.click();

    // Wait for the registration process and navigation
    await browser.pause(2000);

    // 5. Verify redirection to the enter info page
    const currentUrl = await browser.getUrl();
    await expect(currentUrl).toContain('/enterInfo');

    // Take a screenshot of the enter info page
    await browser.assertView('enter-info-page', 'body', {
      allowViewportOverflow: true,
    });
  });

  it('should allow a user to register with phone', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // 1. Navigate to the Auth page
    await browser.url('https://testing.couply.ru/auth');

    // Wait for the page to load
    await browser.pause(1000);

    // 2. Click on the "login with phone" button to navigate to registration
    const phoneLoginButton = await browser.$('button:nth-of-type(3)');
    await phoneLoginButton.click();

    // Wait for navigation to registration page
    await browser.pause(1000);

    // Take a screenshot of the registration page
    await browser.assertView('registration-page-phone', 'body', {
      allowViewportOverflow: true,
    });

    // 3. Fill in the registration form with phone
    // Find the phone input field
    const phoneInput = await browser.$('input[type="tel"]');
    await phoneInput.setValue('+79991234567');

    // Find the password input fields
    const passwordInputs = await browser.$$('input[type="password"]');

    // Fill in the password
    await passwordInputs[0].setValue('password123');

    // Fill in the confirm password
    await passwordInputs[1].setValue('password123');

    // Take a screenshot of the filled form
    await browser.assertView('registration-form-phone-filled', 'body', {
      allowViewportOverflow: true,
    });

    // 4. Submit the registration form
    const registerButton = await browser.$('button');
    await registerButton.click();

    // Wait for the registration process and navigation
    await browser.pause(2000);

    // 5. Verify redirection to the enter info page
    const currentUrl = await browser.getUrl();
    await expect(currentUrl).toContain('/enterInfo');

    // Take a screenshot of the enter info page
    await browser.assertView('enter-info-page-from-phone', 'body', {
      allowViewportOverflow: true,
    });
  });

  it('should show validation errors for invalid registration data', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // 1. Navigate to the Auth page
    await browser.url('https://testing.couply.ru/auth');

    // Wait for the page to load
    await browser.pause(1000);

    // 2. Click on the "login with email" button to navigate to registration
    const emailLoginButton = await browser.$('button:nth-of-type(4)');
    await emailLoginButton.click();

    // Wait for navigation to registration page
    await browser.pause(1000);

    // 3. Submit the form without filling in any fields
    const registerButton = await browser.$('button');
    await registerButton.click();

    // Wait for validation errors to appear
    await browser.pause(500);

    // Take a screenshot of the validation errors
    await browser.assertView('registration-validation-errors', 'body', {
      allowViewportOverflow: true,
    });

    // 4. Fill in invalid email
    const emailInput = await browser.$('input[type="email"]');
    await emailInput.setValue('invalid-email');

    // Find the password input fields
    const passwordInputs = await browser.$$('input[type="password"]');

    // Fill in the password (too short)
    await passwordInputs[0].setValue('123');

    // Fill in the confirm password (doesn't match)
    await passwordInputs[1].setValue('456');

    // Submit the form with invalid data
    await registerButton.click();

    // Wait for validation errors to appear
    await browser.pause(500);

    // Take a screenshot of the validation errors for invalid data
    await browser.assertView('registration-invalid-data-errors', 'body', {
      allowViewportOverflow: true,
    });

    // Verify error messages are displayed
    const errorMessages = await browser.$$('.errorText');
    expect(await errorMessages.length).toBeGreaterThan(0);
  });
});
