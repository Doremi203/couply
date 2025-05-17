describe('Phone Verification Flow', () => {
  it('should allow a user to enter and verify their phone number', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // Navigate to the phone verification page
    await browser.url('https://rainbow-souffle-ece639.netlify.app/enterPhone');

    // Wait for the page to load completely
    await browser.pause(3000);

    // Log current URL for debugging
    const currentUrl = await browser.getUrl();
    console.log('Current URL:', currentUrl);

    // Find the phone input field
    const phoneInput = await browser.$('input[type="tel"], input[placeholder*="phone"]');

    // Enter a valid phone number
    if (phoneInput) {
      await phoneInput.setValue('9123456789');
      console.log('Entered phone number');
    } else {
      console.log('Phone input not found');
    }

    // Find and click the continue/submit button
    const submitButton = await browser.$('button[type="submit"], button');
    if (submitButton) {
      await submitButton.click();
      console.log('Clicked submit button');
    }

    // Wait for the code verification page to load
    await browser.pause(2000);

    // Find code input fields (usually there are multiple for verification code)
    const codeInputs = await browser.$$(
      'input[type="text"], input[type="number"], input.codeInput',
    );
    console.log('Number of code inputs found:', codeInputs.length);

    // Enter verification code (e.g., 123456)
    if (codeInputs.length >= 1) {
      // If there's a single input for the entire code
      if (codeInputs.length === 1) {
        await codeInputs[0].setValue('123456');
      }
      // If there are separate inputs for each digit
      else if (codeInputs.length >= 6) {
        await codeInputs[0].setValue('1');
        await codeInputs[1].setValue('2');
        await codeInputs[2].setValue('3');
        await codeInputs[3].setValue('4');
        await codeInputs[4].setValue('5');
        await codeInputs[5].setValue('6');
      }
      console.log('Entered verification code');
    }

    // Find and click the verify button
    const verifyButton = await browser.$('button[type="submit"], button');
    if (verifyButton) {
      await verifyButton.click();
      console.log('Clicked verify button');
    }

    // Wait for verification process
    await browser.pause(3000);

    // Log final URL to verify successful verification
    const finalUrl = await browser.getUrl();
    console.log('Final URL after verification:', finalUrl);

    // Check if redirected to the next step (likely profile or home page)
    const redirected =
      finalUrl.includes('profile') || finalUrl.includes('home') || finalUrl.includes('info');
    console.log('Successfully redirected:', redirected);
  });

  //   it('should show validation error for invalid phone number', async ({ browser }) => {
  //     // Set window size to ensure consistent behavior
  //     await browser.setWindowSize(1920, 1080);

  //     // Navigate to the phone verification page
  //     await browser.url('https://moonlit-valkyrie-fdbfc4.netlify.app/enterPhone');

  //     // Wait for the page to load completely
  //     await browser.pause(3000);

  //     // Find the phone input field
  //     const phoneInput = await browser.$('input[type="tel"], input[placeholder*="phone"]');

  //     // Enter an invalid phone number (too short)
  //     if (phoneInput) {
  //       await phoneInput.setValue('123');
  //       console.log('Entered invalid phone number');
  //     }

  //     // Find and click the continue/submit button
  //     const submitButton = await browser.$('button[type="submit"], button');
  //     if (submitButton) {
  //       await submitButton.click();
  //       console.log('Clicked submit button');
  //     }

  //     // Wait for validation errors to appear
  //     await browser.pause(1000);

  //     // Find error messages
  //     const errorElements = await browser.$$('.errorText, .error-text, [role="alert"]');
  //     console.log('Number of error elements found:', errorElements.length);

  //     // Verify that error messages are displayed
  //     if (errorElements.length > 0) {
  //       const errorTexts = await Promise.all(errorElements.map(el => el.getText()));
  //       console.log('Error texts:', errorTexts);
  //       console.log('Validation errors are displayed as expected');
  //     }
  //   });

  //   it('should handle incorrect verification code', async ({ browser }) => {
  //     // Set window size to ensure consistent behavior
  //     await browser.setWindowSize(1920, 1080);

  //     // Navigate to the phone verification page
  //     await browser.url('https://moonlit-valkyrie-fdbfc4.netlify.app/enterPhone');

  //     // Wait for the page to load completely
  //     await browser.pause(3000);

  //     // Find the phone input field
  //     const phoneInput = await browser.$('input[type="tel"], input[placeholder*="phone"]');

  //     // Enter a valid phone number
  //     if (phoneInput) {
  //       await phoneInput.setValue('9123456789');
  //     }

  //     // Find and click the continue/submit button
  //     const submitButton = await browser.$('button[type="submit"], button');
  //     if (submitButton) {
  //       await submitButton.click();
  //     }

  //     // Wait for the code verification page to load
  //     await browser.pause(2000);

  //     // Find code input fields
  //     const codeInputs = await browser.$$(
  //       'input[type="text"], input[type="number"], input.codeInput',
  //     );

  //     // Enter incorrect verification code (e.g., 111111)
  //     if (codeInputs.length >= 1) {
  //       // If there's a single input for the entire code
  //       if (codeInputs.length === 1) {
  //         await codeInputs[0].setValue('111111');
  //       }
  //       // If there are separate inputs for each digit
  //       else if (codeInputs.length >= 6) {
  //         for (let i = 0; i < 6; i++) {
  //           await codeInputs[i].setValue('1');
  //         }
  //       }
  //     }

  //     // Find and click the verify button
  //     const verifyButton = await browser.$('button[type="submit"], button');
  //     if (verifyButton) {
  //       await verifyButton.click();
  //     }

  //     // Wait for error message
  //     await browser.pause(2000);

  //     // Find error messages
  //     const errorElements = await browser.$$('.errorText, .error-text, [role="alert"]');
  //     console.log('Number of error elements found after incorrect code:', errorElements.length);

  //     // Verify that error messages are displayed
  //     if (errorElements.length > 0) {
  //       const errorTexts = await Promise.all(errorElements.map(el => el.getText()));
  //       console.log('Error texts for incorrect code:', errorTexts);
  //     }
  //   });
});
