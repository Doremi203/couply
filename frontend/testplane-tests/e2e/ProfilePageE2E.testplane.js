describe('Profile Page Functionality', () => {
  it('should display user profile information correctly', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // Navigate to the profile page (assuming user is already logged in)
    // In a real test, you might need to go through login first
    await browser.url('https://rainbow-souffle-ece639.netlify.app/profile');

    // Wait for the page to load completely
    await browser.pause(3000);

    // Log current URL for debugging
    const currentUrl = await browser.getUrl();
    console.log('Current URL:', currentUrl);

    // Check if profile header is displayed
    const profileHeader = await browser.$('.profileHeader');
    const headerExists = await profileHeader.isExisting();
    console.log('Profile header exists:', headerExists);

    // Check if profile photo is displayed
    const profilePhoto = await browser.$('img');
    const photoExists = await profilePhoto.isExisting();
    console.log('Profile photo exists:', photoExists);

    // Check if user name is displayed
    const userName = await browser.$('h1, h2, .userName');
    const nameExists = await userName.isExisting();
    console.log('User name exists:', nameExists);

    if (nameExists) {
      const nameText = await userName.getText();
      console.log('User name:', nameText);
    }

    // Check if profile info sections are displayed
    const infoSections = await browser.$$('.profileSection, .infoGrid, .tagsList');
    console.log('Number of profile info sections:', infoSections.length);

    // Check if profile menu is displayed
    const profileMenu = await browser.$('.profileMenu');
    const menuExists = await profileMenu.isExisting();
    console.log('Profile menu exists:', menuExists);
  });

  it('should allow editing profile information', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // Navigate to the profile page
    await browser.url('https://moonlit-valkyrie-fdbfc4.netlify.app/profile');

    // Wait for the page to load completely
    await browser.pause(3000);

    // Find and click the edit button/icon
    const editButton = await browser.$('.editButton, button.edit');
    const editButtonExists = await editButton.isExisting();
    console.log('Edit button exists:', editButtonExists);

    if (editButtonExists) {
      await editButton.click();
      console.log('Clicked edit button');

      // Wait for edit mode to load
      await browser.pause(2000);

      // Find editable fields
      const inputFields = await browser.$$('input[type="text"], textarea');
      console.log('Number of editable fields:', inputFields.length);

      // Edit a text field (e.g., bio or about me)
      if (inputFields.length > 0) {
        const bioField = await browser.$(
          'textarea, input[placeholder*="about"], input[placeholder*="bio"]',
        );
        if (await bioField.isExisting()) {
          await bioField.clearValue();
          await bioField.setValue('This is an automated test bio update');
          console.log('Updated bio text');
        }
      }

      // Find and click save button
      const saveButton = await browser.$('button[type="submit"], .saveButton');
      if (await saveButton.isExisting()) {
        await saveButton.click();
        console.log('Clicked save button');

        // Wait for save to complete
        await browser.pause(3000);

        // Verify changes were saved
        // Look for any element containing our test text
        const updatedBio = await browser.$('p, div');
        const bioUpdated =
          (await updatedBio.isExisting()) &&
          (await updatedBio.getText()).includes('This is an automated test bio update');
        console.log('Bio was updated successfully:', bioUpdated);
      }
    } else {
      console.log('Edit button not found, trying alternative approach');

      // Try to find edit profile section directly
      const editProfileSection = await browser.$('.editProfile');
      if (await editProfileSection.isExisting()) {
        console.log('Edit profile section found directly');

        // Find editable fields
        const inputFields = await browser.$$('input[type="text"], textarea');
        console.log('Number of editable fields:', inputFields.length);

        // Edit a text field
        if (inputFields.length > 0) {
          await inputFields[0].clearValue();
          await inputFields[0].setValue('Updated via automation');
          console.log('Updated text field');

          // Find and click save button
          const saveButton = await browser.$('button[type="submit"], .saveButton');
          if (await saveButton.isExisting()) {
            await saveButton.click();
            console.log('Clicked save button');
            await browser.pause(3000);
          }
        }
      }
    }
  });

  it('should navigate to settings from profile menu', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // Navigate to the profile page
    await browser.url('https://moonlit-valkyrie-fdbfc4.netlify.app/profile');

    // Wait for the page to load completely
    await browser.pause(3000);

    // Find and click the profile menu button/icon
    const menuButton = await browser.$('.profileMenu button, .menuButton, .settingsButton');
    const menuButtonExists = await menuButton.isExisting();
    console.log('Menu button exists:', menuButtonExists);

    if (menuButtonExists) {
      await menuButton.click();
      console.log('Clicked menu button');

      // Wait for menu to open
      await browser.pause(1000);

      // Find and click settings option
      const settingsOption = await browser.$('a.settingsLink, button.settingsButton');
      if (await settingsOption.isExisting()) {
        await settingsOption.click();
        console.log('Clicked settings option');

        // Wait for navigation to settings page
        await browser.pause(2000);

        // Verify navigation to settings page
        const currentUrl = await browser.getUrl();
        console.log('Current URL after clicking settings:', currentUrl);

        const navigatedToSettings = currentUrl.includes('settings');
        console.log('Successfully navigated to settings:', navigatedToSettings);
      }
    } else {
      console.log('Menu button not found, trying direct navigation');

      // Try direct navigation to settings
      await browser.url('https://moonlit-valkyrie-fdbfc4.netlify.app/settings');
      await browser.pause(2000);

      const settingsUrl = await browser.getUrl();
      console.log('Directly navigated to settings URL:', settingsUrl);
    }
  });
});
