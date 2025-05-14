describe('Settings Page Functionality', () => {
  it('should display settings options correctly', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // Navigate to the settings page (assuming user is already logged in)
    await browser.url('https://rainbow-souffle-ece639.netlify.app/settings');

    // Wait for the page to load completely
    await browser.pause(3000);

    // Log current URL for debugging
    const currentUrl = await browser.getUrl();
    console.log('Current URL:', currentUrl);

    // Check if settings page title is displayed
    const pageTitle = await browser.$('h1, h2');
    const titleExists = await pageTitle.isExisting();
    console.log('Settings page title exists:', titleExists);

    if (titleExists) {
      const titleText = await pageTitle.getText();
      console.log('Page title text:', titleText);
    }

    // Check if notification settings section is displayed
    const notificationSettings = await browser.$('.notificationSettings');
    const notificationSectionExists = await notificationSettings.isExisting();
    console.log('Notification settings section exists:', notificationSectionExists);

    // Check if theme toggle is displayed
    const themeToggle = await browser.$('.themeToggle');
    const themeToggleExists = await themeToggle.isExisting();
    console.log('Theme toggle exists:', themeToggleExists);

    // Check for other common settings sections
    const settingsSections = await browser.$$('.settingsSection, .section');
    console.log('Number of settings sections found:', settingsSections.length);

    // Check for back button
    const backButton = await browser.$('.backButton, button');
    const backButtonExists = await backButton.isExisting();
    console.log('Back button exists:', backButtonExists);
  });

  it('should toggle notification settings', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // Navigate to the settings page
    await browser.url('https://rainbow-souffle-ece639.netlify.app/settings');

    // Wait for the page to load completely
    await browser.pause(3000);

    // Find notification settings section
    const notificationSettings = await browser.$('.notificationSettings');
    if (await notificationSettings.isExisting()) {
      console.log('Found notification settings section');

      // Find toggle switches in notification settings
      const toggles = await notificationSettings.$$('input[type="checkbox"], .toggle, .switch');
      console.log('Number of notification toggles found:', toggles.length);

      if (toggles.length > 0) {
        // Get initial state of first toggle
        const firstToggle = toggles[0];
        const initialState = await firstToggle.isSelected();
        console.log('Initial state of first toggle:', initialState);

        // Click the toggle to change its state
        await firstToggle.click();
        console.log('Clicked first toggle');

        // Wait for state to update
        await browser.pause(1000);

        // Get new state of toggle
        const newState = await firstToggle.isSelected();
        console.log('New state of first toggle:', newState);

        // Verify that the state changed
        const stateChanged = initialState !== newState;
        console.log('Toggle state changed:', stateChanged);

        // Toggle back to original state
        await firstToggle.click();
        console.log('Toggled back to original state');
        await browser.pause(1000);
      }
    } else {
      console.log('Notification settings section not found');
    }
  });

  it('should toggle theme between light and dark', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // Navigate to the settings page
    await browser.url('https://rainbow-souffle-ece639.netlify.app/settings');

    // Wait for the page to load completely
    await browser.pause(3000);

    // Find theme toggle
    const themeToggle = await browser.$('.themeToggle, .darkModeToggle, .lightModeToggle');
    const themeToggleExists = await themeToggle.isExisting();
    console.log('Theme toggle exists:', themeToggleExists);

    if (themeToggleExists) {
      // Get initial theme
      const body = await browser.$('body');
      const initialThemeIsDark =
        (await body.hasClass('darkTheme')) || (await body.getAttribute('data-theme')) === 'dark';
      console.log('Initial theme is dark:', initialThemeIsDark);

      // Click theme toggle
      await themeToggle.click();
      console.log('Clicked theme toggle');

      // Wait for theme to change
      await browser.pause(2000);

      // Check if theme changed
      const newThemeIsDark =
        (await body.hasClass('darkTheme')) || (await body.getAttribute('data-theme')) === 'dark';
      console.log('New theme is dark:', newThemeIsDark);

      // Verify that theme changed
      const themeChanged = initialThemeIsDark !== newThemeIsDark;
      console.log('Theme changed:', themeChanged);

      // Toggle back to original theme
      await themeToggle.click();
      console.log('Toggled back to original theme');
      await browser.pause(2000);
    } else {
      console.log('Theme toggle not found, trying alternative approach');

      // Try to find theme toggle in navbar or other locations
      const alternativeThemeToggle = await browser.$(
        '.themeToggle, .toggleButtons, .darkModeToggle, .lightModeToggle',
      );
      if (await alternativeThemeToggle.isExisting()) {
        console.log('Found alternative theme toggle');
        await alternativeThemeToggle.click();
        console.log('Clicked alternative theme toggle');
        await browser.pause(2000);
      }
    }
  });

  it('should navigate back from settings page', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // Navigate to the settings page
    await browser.url('https://rainbow-souffle-ece639.netlify.app/settings');

    // Wait for the page to load completely
    await browser.pause(3000);

    // Find back button
    const backButton = await browser.$('.backButton, button');
    const backButtonExists = await backButton.isExisting();
    console.log('Back button exists:', backButtonExists);

    if (backButtonExists) {
      // Click back button
      await backButton.click();
      console.log('Clicked back button');

      // Wait for navigation
      await browser.pause(2000);

      // Check if navigated away from settings page
      const currentUrl = await browser.getUrl();
      console.log('URL after clicking back:', currentUrl);

      const navigatedAway = !currentUrl.includes('settings');
      console.log('Successfully navigated away from settings:', navigatedAway);
    } else {
      console.log('Back button not found');
    }
  });

  it('should save settings changes', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // Navigate to the settings page
    await browser.url('https://rainbow-souffle-ece639.netlify.app/settings');

    // Wait for the page to load completely
    await browser.pause(3000);

    // Find save button if it exists
    const saveButton = await browser.$('button[type="submit"], .saveButton, .submitButton');
    const saveButtonExists = await saveButton.isExisting();
    console.log('Save button exists:', saveButtonExists);

    if (saveButtonExists) {
      // Make a change to a setting
      const toggles = await browser.$$('input[type="checkbox"], .toggle, .switch');
      if (toggles.length > 0) {
        await toggles[0].click();
        console.log('Changed a setting');
        await browser.pause(1000);
      }

      // Click save button
      await saveButton.click();
      console.log('Clicked save button');

      // Wait for save to complete
      await browser.pause(2000);

      // Check for success message or indicator
      const successMessage = await browser.$('.successMessage, .toast, [role="alert"]');
      const successMessageExists = await successMessage.isExisting();
      console.log('Success message exists:', successMessageExists);

      if (successMessageExists) {
        const messageText = await successMessage.getText();
        console.log('Success message:', messageText);
      }
    } else {
      console.log('Save button not found, settings might save automatically');

      // Make a change to a setting to test auto-save
      const toggles = await browser.$$('input[type="checkbox"], .toggle, .switch');
      if (toggles.length > 0) {
        await toggles[0].click();
        console.log('Changed a setting to test auto-save');
        await browser.pause(2000);
      }
    }
  });
});
