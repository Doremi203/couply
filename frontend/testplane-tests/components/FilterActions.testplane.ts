describe('FilterActions Component', () => {
  it('should match the reference screenshot for default state', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the FilterActions component with default state
    await browser.url('http://localhost:6006/?path=/story/features-filters-filteractions--default');

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('default', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot with custom button text', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the FilterActions component with custom text
    await browser.url(
      'http://localhost:6006/?path=/story/features-filters-filteractions--custom-text',
    );

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('custom-text', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });
});
