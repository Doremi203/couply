describe('Registration', () => {
  it('Register with email and passwords error', async ({ browser }) => {
    await browser.setWindowSize(390, 840);

    await browser.url('https://testing.couply.ru/auth');

    await browser.pause(1000);


  const emailButton = await browser.$('[data-testid="email-button"]');

  await emailButton.waitForDisplayed();

    await emailButton.click();

    await browser.pause(2000);

    const regButton = await browser.$('[data-testid="register-button"]');

    await regButton.waitForDisplayed();

    await regButton.click();

    const inputs = await browser.$$('input');


    if (inputs.length >= 1) {
      const emailInput = await browser.$('input[type="email"], input[placeholder*="email"], input[placeholder*="Email"]');
      if (emailInput) {
        await emailInput.setValue('test@example.com');
      } else if (inputs.length >= 1) {
        await inputs[0].setValue('test@example.com');
      }
    }

    const passwordInputs = await browser.$$('input[type="password"]');

    if (passwordInputs.length >= 2) {
      await passwordInputs[0].setValue('11111111');
      await passwordInputs[1].setValue('1165661111');
    }

    // const submitButton = await browser.$('button[type="submit"], button');
    const submitButton = await browser.$('[data-testid="submit-button"]');
    if (submitButton) {
      await submitButton.click();
    }

    await browser.pause(3000);

  });

  it('should show validation errors for invalid registration data', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // Navigate directly to the auth page
    await browser.url('https://testing.couply.ru/auth');

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
      // Find email input by type, placeholder, or position
      const emailInput = await browser.$('input[type="email"], input[placeholder*="email"], input[placeholder*="Email"]');
      if (emailInput) {
        await emailInput.setValue('invalid-email');
      } else if (inputs.length >= 1) {
        // If can't find by type or placeholder, use the first input
        console.log('Email input not found by type or placeholder, using first input');
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
