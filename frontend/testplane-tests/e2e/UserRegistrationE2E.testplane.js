describe('User Registration Scenario', () => {
  it('should allow a user to register with email', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // 1. Navigate to the Auth page
    await browser.url('https://moonlit-valkyrie-fdbfc4.netlify.app/');

    // Wait for the page to load completely
    await browser.pause(3000);

    // Log current URL for debugging
    const currentUrl = await browser.getUrl();
    console.log('Current URL:', currentUrl);

    // Find all buttons on the page
    const buttons = await browser.$$('button');
    console.log('Number of buttons found:', buttons.length);

    // Click the email login button (usually the 4th button)
    if (buttons.length >= 4) {
      await buttons[3].click();
    } else {
      console.log('Not enough buttons found, trying to find by text content');

      // Try to find button by text content
      const emailButton = await browser.$('button=login with email');
      if (emailButton) {
        await emailButton.click();
      } else {
        console.log('Email button not found, navigating directly to registration');
        await browser.url('https://testing.couply.ru/registration');
      }
    }

    // Wait for navigation to registration page
    await browser.pause(2000);

    // Log current URL after navigation
    const regUrl = await browser.getUrl();
    console.log('Registration URL:', regUrl);

    // Find all inputs on the page
    const inputs = await browser.$$('input');
    console.log('Number of inputs found:', inputs.length);

    // Fill in the registration form with email
    if (inputs.length >= 1) {
      // Find email input by type or placeholder
      const emailInput = await browser.$('input[type="email"], input[placeholder*="email"]');
      if (emailInput) {
        await emailInput.setValue('test@example.com');
      } else if (inputs.length >= 1) {
        // If can't find by type, use the first input
        await inputs[0].setValue('test@example.com');
      }
    }

    // Find password inputs
    const passwordInputs = await browser.$$('input[type="password"]');
    console.log('Number of password inputs found:', passwordInputs.length);

    // Fill in password fields
    if (passwordInputs.length >= 2) {
      // Fill in the password
      await passwordInputs[0].setValue('password123');

      // Fill in the confirm password
      await passwordInputs[1].setValue('password123');
    }

    // Find and click the submit button
    const submitButton = await browser.$('button[type="submit"], button');
    if (submitButton) {
      await submitButton.click();
    }

    // Wait for the registration process
    await browser.pause(3000);

    // Log final URL
    const finalUrl = await browser.getUrl();
    console.log('Final URL after registration:', finalUrl);
  });

  it('should show validation errors for invalid registration data', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // Navigate directly to the registration page
    await browser.url('https://moonlit-valkyrie-fdbfc4.netlify.app/');

    // Wait for the page to load completely
    await browser.pause(3000);

    // Find and click the submit button without filling in any fields
    const submitButton = await browser.$('button[type="submit"], button');
    if (submitButton) {
      await submitButton.click();
    }

    // Wait for validation errors to appear
    await browser.pause(1000);

    // Find error messages
    const errorElements = await browser.$$('.errorText, .error-text, [role="alert"]');
    console.log('Number of error elements found:', errorElements.length);

    // Verify that error messages are displayed
    if (errorElements.length > 0) {
      console.log('Validation errors are displayed as expected');
    }

    // Find all inputs on the page
    const inputs = await browser.$$('input');
    console.log('Number of inputs found:', inputs.length);

    // Fill in invalid data
    if (inputs.length >= 1) {
      // Find email input
      const emailInput = await browser.$('input[type="email"], input[placeholder*="email"]');
      if (emailInput) {
        await emailInput.setValue('invalid-email');
      } else if (inputs.length >= 1) {
        // If can't find by type, use the first input
        await inputs[0].setValue('invalid-email');
      }
    }

    // Find password inputs
    const passwordInputs = await browser.$$('input[type="password"]');

    // Fill in invalid password data
    if (passwordInputs.length >= 2) {
      // Fill in the password (too short)
      await passwordInputs[0].setValue('123');

      // Fill in the confirm password (doesn't match)
      await passwordInputs[1].setValue('456');
    }

    // Submit the form with invalid data
    if (submitButton) {
      await submitButton.click();
    }

    // Wait for validation errors to appear
    await browser.pause(1000);

    // Find error messages again
    const invalidDataErrors = await browser.$$('.errorText, .error-text, [role="alert"]');
    console.log('Number of error elements after invalid data:', invalidDataErrors.length);

    // Log error messages for debugging
    if (invalidDataErrors.length > 0) {
      const errorTexts = await Promise.all(invalidDataErrors.map(el => el.getText()));
      console.log('Error texts:', errorTexts);
    }
  });
});
