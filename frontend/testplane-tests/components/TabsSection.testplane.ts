describe('TabsSection Component', () => {
  it('should match the reference screenshot for default state', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the TabsSection component
    await browser.url('http://localhost:6006/?path=/story/components-tabssection--default');

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('default', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot with multiple tabs', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the TabsSection component with multiple tabs
    await browser.url('http://localhost:6006/?path=/story/components-tabssection--multiple-tabs');

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('multiple-tabs', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot with custom tab labels', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the TabsSection component with custom tab labels
    await browser.url('http://localhost:6006/?path=/story/components-tabssection--custom-labels');

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('custom-labels', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });
});
