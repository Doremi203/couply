describe('User Registration Scenario', () => {
  it('should allow a user to register with email', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // 1. Navigate to the AuthPage story in Storybook
    await browser.url('http://localhost:6006/?path=/story/pages-authpage--default');

    // Wait for the page to load completely
    await browser.pause(3000);

    // Switch to the iframe where the story is rendered
    const iframe = await browser.$('#storybook-preview-iframe');
    await browser.switchToFrame(iframe);

    // Find and click the email login button (4th button)
    const buttons = await browser.$$('button');
    console.log('Number of buttons found:', buttons.length);

    // Click the email login button (usually the 4th button)
    if (buttons.length >= 4) {
      await buttons[3].click();
    } else {
      console.log('Not enough buttons found, using direct navigation');
      await browser.url('http://localhost:6006/?path=/story/pages-registrationpage--default');
      await browser.pause(2000);
      await browser.switchToFrame(await browser.$('#storybook-preview-iframe'));
    }

    // Wait for navigation to registration page
    await browser.pause(2000);

    // 3. Fill in the registration form with email
    // Find the email input field
    const inputs = await browser.$$('input');
    console.log('Number of inputs found:', inputs.length);

    // Assuming first input is for email/phone
    if (inputs.length >= 1) {
      await inputs[0].setValue('test@example.com');
    }

    // Find the password input fields (usually the 2nd and 3rd inputs)
    if (inputs.length >= 3) {
      // Fill in the password
      await inputs[1].setValue('password123');

      // Fill in the confirm password
      await inputs[2].setValue('password123');
    }

    // 4. Submit the registration form
    const submitButton = await browser.$('button');
    if (submitButton) {
      await submitButton.click();
    }

    // Wait for the registration process
    await browser.pause(2000);
  });

  it('should allow a user to register with phone', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // 1. Navigate to the AuthPage story in Storybook
    await browser.url('http://localhost:6006/?path=/story/pages-authpage--default');

    // Wait for the page to load completely
    await browser.pause(3000);

    // Switch to the iframe where the story is rendered
    const iframe = await browser.$('#storybook-preview-iframe');
    await browser.switchToFrame(iframe);

    // Find and click the phone login button (3rd button)
    const buttons = await browser.$$('button');
    console.log('Number of buttons found:', buttons.length);

    // Click the phone login button (usually the 3rd button)
    if (buttons.length >= 3) {
      await buttons[2].click();
    } else {
      console.log('Not enough buttons found, using direct navigation');
      await browser.url('http://localhost:6006/?path=/story/pages-registrationpage--default');
      await browser.pause(2000);
      await browser.switchToFrame(await browser.$('#storybook-preview-iframe'));
    }

    // Wait for navigation to registration page
    await browser.pause(2000);

    // 3. Fill in the registration form with phone
    // Find the phone input field
    const inputs = await browser.$$('input');
    console.log('Number of inputs found:', inputs.length);

    // Assuming first input is for email/phone
    if (inputs.length >= 1) {
      await inputs[0].setValue('+79991234567');
    }

    // Find the password input fields (usually the 2nd and 3rd inputs)
    if (inputs.length >= 3) {
      // Fill in the password
      await inputs[1].setValue('password123');

      // Fill in the confirm password
      await inputs[2].setValue('password123');
    }

    // 4. Submit the registration form
    const submitButton = await browser.$('button');
    if (submitButton) {
      await submitButton.click();
    }

    // Wait for the registration process
    await browser.pause(2000);
  });

  it('should show validation errors for invalid registration data', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // 1. Navigate directly to the RegistrationPage story in Storybook
    await browser.url('http://localhost:6006/?path=/story/pages-registrationpage--default');

    // Wait for the page to load completely
    await browser.pause(3000);

    // Switch to the iframe where the story is rendered
    const iframe = await browser.$('#storybook-preview-iframe');
    await browser.switchToFrame(iframe);

    // 2. Submit the form without filling in any fields
    const submitButton = await browser.$('button');
    if (submitButton) {
      await submitButton.click();
    }

    // Wait for validation errors to appear
    await browser.pause(1000);

    // Verify validation errors are displayed
    const errorElements = await browser.$$('.errorText');
    console.log('Number of error elements found:', errorElements.length);
    expect(await errorElements.length).toBeGreaterThan(0);

    // 3. Fill in invalid data
    const inputs = await browser.$$('input');

    // Fill in invalid email
    if (inputs.length >= 1) {
      await inputs[0].setValue('invalid-email');
    }

    // Fill in short password
    if (inputs.length >= 2) {
      await inputs[1].setValue('123');
    }

    // Fill in non-matching confirm password
    if (inputs.length >= 3) {
      await inputs[2].setValue('456');
    }

    // Submit the form with invalid data
    if (submitButton) {
      await submitButton.click();
    }

    // Wait for validation errors to appear
    await browser.pause(1000);

    // Verify validation errors are displayed
    const invalidDataErrors = await browser.$$('.errorText');
    expect(await invalidDataErrors.length).toBeGreaterThan(0);

    // Log error messages for debugging
    const errorTexts = await Promise.all(invalidDataErrors.map(el => el.getText()));
    console.log('Error texts:', errorTexts);
  });
});
