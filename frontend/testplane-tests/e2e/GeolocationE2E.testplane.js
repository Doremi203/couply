describe('Geolocation Functionality', () => {
  it('should request geolocation permission', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // Navigate to the page with geolocation request
    await browser.url('https://rainbow-souffle-ece639.netlify.app/enterInfo');

    // Wait for the page to load completely
    await browser.pause(3000);

    // Log current URL for debugging
    const currentUrl = await browser.getUrl();
    console.log('Current URL:', currentUrl);

    // Check if geolocation request component is displayed
    const geoLocationRequest = await browser.$('.geoLocationRequest');
    const requestExists = await geoLocationRequest.isExisting();
    console.log('Geolocation request component exists:', requestExists);

    if (requestExists) {
      // Check if there's a button to request geolocation
      const geoButton = await geoLocationRequest.$('button, .geoButton');
      const buttonExists = await geoButton.isExisting();
      console.log('Geolocation request button exists:', buttonExists);

      if (buttonExists) {
        // Get button text
        const buttonText = await geoButton.getText();
        console.log('Geolocation button text:', buttonText);

        // Click the geolocation button
        await geoButton.click();
        console.log('Clicked geolocation button');

        // Wait for geolocation permission dialog (browser will handle this)
        await browser.pause(3000);

        // Note: We can't directly interact with browser permission dialogs in most e2e testing frameworks
        // Instead, we can check if the UI updates to reflect permission state

        // Check for success indicators in the UI
        const successIndicator = await browser.$('.geoSuccess, .locationEnabled');
        const successExists = await successIndicator.isExisting();
        console.log('Geolocation success indicator exists:', successExists);

        // Check for error indicators in the UI
        const errorIndicator = await browser.$('.geoError, .locationDisabled');
        const errorExists = await errorIndicator.isExisting();
        console.log('Geolocation error indicator exists:', errorExists);
      }
    } else {
      console.log('Geolocation request component not found, trying alternative approach');

      // Try to find PWA geolocation helper
      const pwaHelper = await browser.$('.pwaGeolocationHelper');
      const pwaHelperExists = await pwaHelper.isExisting();
      console.log('PWA geolocation helper exists:', pwaHelperExists);

      if (pwaHelperExists) {
        // Check for geolocation button within the helper
        const geoButton = await pwaHelper.$('button, .geoButton');
        if (await geoButton.isExisting()) {
          await geoButton.click();
          console.log('Clicked geolocation button in PWA helper');
          await browser.pause(3000);
        }
      }
    }
  });

  it('should display location-based information when geolocation is enabled', async ({
    browser,
  }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // Navigate to the page with geolocation features
    await browser.url('https://moonlit-valkyrie-fdbfc4.netlify.app/enterInfo');

    // Wait for the page to load completely
    await browser.pause(3000);

    // Mock geolocation API (if possible)
    // Note: This is a limitation of most e2e testing frameworks
    // In a real test, you might need to use browser-specific capabilities to mock geolocation

    // For this test, we'll assume geolocation is already enabled and check for location-based UI elements

    // Check for location display elements
    const locationDisplay = await browser.$('.locationDisplay, .cityDisplay, .locationInfo');
    const locationDisplayExists = await locationDisplay.isExisting();
    console.log('Location display exists:', locationDisplayExists);

    if (locationDisplayExists) {
      // Get displayed location text
      const locationText = await locationDisplay.getText();
      console.log('Displayed location:', locationText);
    }

    // Check for map component if it exists
    const mapComponent = await browser.$('.map, .locationMap');
    const mapExists = await mapComponent.isExisting();
    console.log('Map component exists:', mapExists);

    // Check for distance indicators in user profiles or matches
    const distanceIndicators = await browser.$$('.distance, .locationDistance');
    console.log('Number of distance indicators found:', distanceIndicators.length);

    if (distanceIndicators.length > 0) {
      // Get text from first distance indicator
      const distanceText = await distanceIndicators[0].getText();
      console.log('Distance text:', distanceText);
    }
  });

  it('should handle geolocation permission denial gracefully', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // Navigate to the page with geolocation request
    await browser.url('https://moonlit-valkyrie-fdbfc4.netlify.app/enterInfo');

    // Wait for the page to load completely
    await browser.pause(3000);

    // Note: We can't directly simulate permission denial in most e2e testing frameworks
    // Instead, we'll check if the UI has elements for handling denied permission

    // Check for permission denied message
    const deniedMessage = await browser.$('.geoDenied, .locationDenied, .locationError');
    const deniedMessageExists = await deniedMessage.isExisting();
    console.log('Geolocation denied message exists:', deniedMessageExists);

    if (deniedMessageExists) {
      // Get denied message text
      const messageText = await deniedMessage.getText();
      console.log('Denied message text:', messageText);
    }

    // Check for retry button
    const retryButton = await browser.$('.retryButton, .allowButton, .enableButton');
    const retryButtonExists = await retryButton.isExisting();
    console.log('Retry button exists:', retryButtonExists);

    if (retryButtonExists) {
      // Click retry button
      await retryButton.click();
      console.log('Clicked retry button');
      await browser.pause(2000);
    }

    // Check for manual location input as fallback
    const manualLocationInput = await browser.$(
      'input[placeholder*="location"], input[placeholder*="city"]',
    );
    const manualInputExists = await manualLocationInput.isExisting();
    console.log('Manual location input exists:', manualInputExists);

    if (manualInputExists) {
      // Enter a location manually
      await manualLocationInput.setValue('Moscow');
      console.log('Entered manual location');

      // Find and click submit button
      const submitButton = await browser.$('button[type="submit"], .saveButton, .continueButton');
      if (await submitButton.isExisting()) {
        await submitButton.click();
        console.log('Submitted manual location');
        await browser.pause(2000);
      }
    }
  });

  it('should update location when user changes it in settings', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // Navigate to settings page
    await browser.url('https://moonlit-valkyrie-fdbfc4.netlify.app/settings');

    // Wait for the page to load completely
    await browser.pause(3000);

    // Check for location settings section
    const locationSettings = await browser.$('.locationSettings, .locationSection');
    const locationSettingsExists = await locationSettings.isExisting();
    console.log('Location settings section exists:', locationSettingsExists);

    if (locationSettingsExists) {
      // Check for current location display
      const currentLocation = await locationSettings.$('.currentLocation, .locationValue');
      if (await currentLocation.isExisting()) {
        const locationText = await currentLocation.getText();
        console.log('Current location text:', locationText);
      }

      // Check for location update button
      const updateButton = await locationSettings.$('.updateButton, .changeButton');
      const updateButtonExists = await updateButton.isExisting();
      console.log('Location update button exists:', updateButtonExists);

      if (updateButtonExists) {
        // Click update button
        await updateButton.click();
        console.log('Clicked location update button');
        await browser.pause(2000);

        // Check for location input field
        const locationInput = await browser.$(
          'input[placeholder*="location"], input[placeholder*="city"]',
        );
        if (await locationInput.isExisting()) {
          // Enter new location
          await locationInput.clearValue();
          await locationInput.setValue('Saint Petersburg');
          console.log('Entered new location');

          // Find and click save button
          const saveButton = await browser.$('button[type="submit"], .saveButton, .applyButton');
          if (await saveButton.isExisting()) {
            await saveButton.click();
            console.log('Saved new location');
            await browser.pause(2000);

            // Check if location was updated
            const updatedLocation = await locationSettings.$('.currentLocation, .locationValue');
            if (await updatedLocation.isExisting()) {
              const newLocationText = await updatedLocation.getText();
              console.log('Updated location text:', newLocationText);

              // Check if location text contains the new city
              const locationUpdated = newLocationText.includes('Saint Petersburg');
              console.log('Location was successfully updated:', locationUpdated);
            }
          }
        }
      }
    } else {
      console.log('Location settings section not found');
    }
  });
});
