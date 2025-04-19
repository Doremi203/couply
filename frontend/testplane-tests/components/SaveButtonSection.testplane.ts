describe('SaveButtonSection Component', () => {
  it('should match the reference screenshot for default state', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the SaveButtonSection component
    await browser.url('http://localhost:6006/?path=/story/components-savebuttonsection--default');

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('default', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot for loading state', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the SaveButtonSection component in loading state
    await browser.url('http://localhost:6006/?path=/story/components-savebuttonsection--loading');

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('loading', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot for disabled state', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the SaveButtonSection component in disabled state
    await browser.url('http://localhost:6006/?path=/story/components-savebuttonsection--disabled');

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('disabled', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });
});
