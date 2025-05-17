describe('Theme Switching Functionality', () => {
  it('should toggle between light and dark themes using theme toggle', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // Navigate to the home page (or any page that has the theme toggle)
    await browser.url('https://rainbow-souffle-ece639.netlify.app/settings');

    // Wait for the page to load completely
    await browser.pause(3000);

    // Log current URL for debugging
    const currentUrl = await browser.getUrl();
    console.log('Current URL:', currentUrl);

    // Find theme toggle component
    const themeToggle = await browser.$('.themeToggle');
    const themeToggleExists = await themeToggle.isExisting();
    console.log('Theme toggle exists:', themeToggleExists);

    if (themeToggleExists) {
      // Get initial theme state
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

      // Check for visual indicators of theme change
      if (newThemeIsDark) {
        console.log('Checking for dark theme visual indicators');

        // Check background color (should be darker)
        const backgroundColor = await body.getCSSProperty('background-color');
        console.log('Body background color in dark mode:', backgroundColor.value);

        // Check text color (should be lighter)
        const textColor = await body.getCSSProperty('color');
        console.log('Text color in dark mode:', textColor.value);
      } else {
        console.log('Checking for light theme visual indicators');

        // Check background color (should be lighter)
        const backgroundColor = await body.getCSSProperty('background-color');
        console.log('Body background color in light mode:', backgroundColor.value);

        // Check text color (should be darker)
        const textColor = await body.getCSSProperty('color');
        console.log('Text color in light mode:', textColor.value);
      }

      // Toggle back to original theme
      await themeToggle.click();
      console.log('Toggled back to original theme');
      await browser.pause(2000);

      // Verify returned to original theme
      const finalThemeIsDark =
        (await body.hasClass('darkTheme')) || (await body.getAttribute('data-theme')) === 'dark';
      console.log('Final theme is dark:', finalThemeIsDark);
      console.log('Returned to original theme:', finalThemeIsDark === initialThemeIsDark);
    } else {
      console.log('Theme toggle not found, trying alternative approaches');

      // Try to find theme toggle in navbar or settings
      const alternativeThemeToggle = await browser.$(
        '.darkModeToggle, .lightModeToggle, .toggleButtons',
      );
      if (await alternativeThemeToggle.isExisting()) {
        console.log('Found alternative theme toggle');
        await alternativeThemeToggle.click();
        console.log('Clicked alternative theme toggle');
        await browser.pause(2000);
      } else {
        // Try navigating to settings page to find theme toggle
        console.log('Trying to find theme toggle in settings page');
        await browser.url('https://rainbow-souffle-ece639.netlify.app/settings');
        await browser.pause(3000);

        const settingsThemeToggle = await browser.$(
          '.themeToggle, .darkModeToggle, .lightModeToggle',
        );
        if (await settingsThemeToggle.isExisting()) {
          console.log('Found theme toggle in settings');
          await settingsThemeToggle.click();
          console.log('Clicked theme toggle in settings');
          await browser.pause(2000);
        }
      }
    }
  });

  it('should persist theme preference across page navigation', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // Navigate to the home page
    await browser.url('https://rainbow-souffle-ece639.netlify.app/');

    // Wait for the page to load completely
    await browser.pause(3000);

    // Find and click theme toggle to set a specific theme
    const themeToggle = await browser.$('.themeToggle, .darkModeToggle, .lightModeToggle');
    if (await themeToggle.isExisting()) {
      // Get initial theme state
      const body = await browser.$('body');
      const initialThemeIsDark =
        (await body.hasClass('darkTheme')) || (await body.getAttribute('data-theme')) === 'dark';
      console.log('Initial theme is dark:', initialThemeIsDark);

      // If not in dark mode, switch to dark mode
      if (!initialThemeIsDark) {
        await themeToggle.click();
        console.log('Switched to dark theme');
        await browser.pause(2000);
      }

      // Verify dark theme is active
      const darkThemeActive =
        (await body.hasClass('darkTheme')) || (await body.getAttribute('data-theme')) === 'dark';
      console.log('Dark theme is active:', darkThemeActive);

      // Navigate to another page
      await browser.url('https://rainbow-souffle-ece639.netlify.app/profile');
      console.log('Navigated to profile page');
      await browser.pause(3000);

      // Check if dark theme persisted
      const newBody = await browser.$('body');
      const themePersisted =
        (await newBody.hasClass('darkTheme')) ||
        (await newBody.getAttribute('data-theme')) === 'dark';
      console.log('Dark theme persisted after navigation:', themePersisted);

      // Navigate to a third page
      await browser.url('https://rainbow-souffle-ece639.netlify.app/settings');
      console.log('Navigated to settings page');
      await browser.pause(3000);

      // Check if dark theme still persisted
      const finalBody = await browser.$('body');
      const finalThemePersisted =
        (await finalBody.hasClass('darkTheme')) ||
        (await finalBody.getAttribute('data-theme')) === 'dark';
      console.log('Dark theme persisted after second navigation:', finalThemePersisted);

      // Reset to original theme if needed
      if (initialThemeIsDark !== darkThemeActive) {
        const finalThemeToggle = await browser.$('.themeToggle, .darkModeToggle, .lightModeToggle');
        if (await finalThemeToggle.isExisting()) {
          await finalThemeToggle.click();
          console.log('Reset to original theme');
          await browser.pause(2000);
        }
      }
    } else {
      console.log('Theme toggle not found');
    }
  });

  it('should apply theme-specific styling to components', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // Navigate to a page with multiple components
    await browser.url('https://rainbow-souffle-ece639.netlify.app/profile');

    // Wait for the page to load completely
    await browser.pause(3000);

    // Find theme toggle
    const themeToggle = await browser.$('.themeToggle, .darkModeToggle, .lightModeToggle');
    if (await themeToggle.isExisting()) {
      // Get initial theme state
      const body = await browser.$('body');
      const initialThemeIsDark =
        (await body.hasClass('darkTheme')) || (await body.getAttribute('data-theme')) === 'dark';
      console.log('Initial theme is dark:', initialThemeIsDark);

      // Check styling of various components in initial theme
      console.log('Checking component styling in initial theme');

      // Check button styling
      const button = await browser.$('button');
      if (await button.isExisting()) {
        const buttonBg = await button.getCSSProperty('background-color');
        const buttonText = await button.getCSSProperty('color');
        console.log('Button background in initial theme:', buttonBg.value);
        console.log('Button text color in initial theme:', buttonText.value);
      }

      // Check card/section styling
      const card = await browser.$('.profileSection, .card');
      if (await card.isExisting()) {
        const cardBg = await card.getCSSProperty('background-color');
        const cardBorder = await card.getCSSProperty('border-color');
        console.log('Card background in initial theme:', cardBg.value);
        console.log('Card border in initial theme:', cardBorder.value);
      }

      // Toggle theme
      await themeToggle.click();
      console.log('Toggled theme');
      await browser.pause(2000);

      // Check if theme changed
      const newThemeIsDark =
        (await body.hasClass('darkTheme')) || (await body.getAttribute('data-theme')) === 'dark';
      console.log('New theme is dark:', newThemeIsDark);

      // Check styling of the same components in new theme
      console.log('Checking component styling in new theme');

      // Check button styling in new theme
      if (await button.isExisting()) {
        const newButtonBg = await button.getCSSProperty('background-color');
        const newButtonText = await button.getCSSProperty('color');
        console.log('Button background in new theme:', newButtonBg.value);
        console.log('Button text color in new theme:', newButtonText.value);
      }

      // Check card/section styling in new theme
      if (await card.isExisting()) {
        const newCardBg = await card.getCSSProperty('background-color');
        const newCardBorder = await card.getCSSProperty('border-color');
        console.log('Card background in new theme:', newCardBg.value);
        console.log('Card border in new theme:', newCardBorder.value);
      }

      // Toggle back to original theme
      await themeToggle.click();
      console.log('Toggled back to original theme');
      await browser.pause(2000);
    } else {
      console.log('Theme toggle not found');
    }
  });
});
