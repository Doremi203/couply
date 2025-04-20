describe('RegistrationPage', () => {
  it('should match the reference screenshot for phone registration', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the RegistrationPage component with phone method
    await browser.url(
      'http://localhost:6006/?path=/story/pages-registrationpage--phone-registration',
    );

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('phone-registration', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot for email registration', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the RegistrationPage component with email method
    await browser.url(
      'http://localhost:6006/?path=/story/pages-registrationpage--email-registration',
    );

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('email-registration', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should validate form fields correctly', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the RegistrationPage component
    await browser.url(
      'http://localhost:6006/?path=/story/pages-registrationpage--phone-registration',
    );

    // Wait for the component to render
    await browser.pause(1000);

    // Get the iframe element
    const iframe = await browser.$('#storybook-preview-iframe');

    // Switch to the iframe context
    await browser.switchToFrame(iframe);

    // Click the submit button without filling in any fields
    const submitButton = await browser.$('button');
    await submitButton.click();

    // Wait for validation errors to appear
    await browser.pause(500);

    // Take a screenshot of the validation errors
    await browser.assertView('validation-errors', 'body', {
      allowViewportOverflow: true,
    });

    // Fill in the fields with invalid data
    const phoneInput = await browser.$('input[type="tel"]');
    await phoneInput.setValue('123'); // Invalid phone number

    const passwordInput = await browser.$('input[type="password"]');
    await passwordInput.setValue('123'); // Password too short

    const confirmPasswordInput = await browser.$('input[type="password"]:nth-of-type(2)');
    await confirmPasswordInput.setValue('456'); // Passwords don't match

    // Click the submit button again
    await submitButton.click();

    // Wait for validation errors to update
    await browser.pause(500);

    // Take a screenshot of the updated validation errors
    await browser.assertView('invalid-data-errors', 'body', {
      allowViewportOverflow: true,
    });

    // Switch back to the main context
    await browser.switchToFrame(null);
  });
});
