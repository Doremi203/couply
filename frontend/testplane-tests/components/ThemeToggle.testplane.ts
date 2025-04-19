describe('ThemeToggle Component', () => {
  it('should match the reference screenshot for light theme', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the ThemeToggle component in light theme
    await browser.url('http://localhost:6006/?path=/story/components-themetoggle--light-theme');

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('light-theme', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot for dark theme', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the ThemeToggle component in dark theme
    await browser.url('http://localhost:6006/?path=/story/components-themetoggle--dark-theme');

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('dark-theme', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });
});
