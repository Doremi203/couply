describe('EmptyState Component', () => {
  it('should match the reference screenshot for default state', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the EmptyState component
    await browser.url('http://localhost:6006/?path=/story/components-emptystate--default');

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('default', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot with custom content', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the EmptyState component with custom content
    await browser.url('http://localhost:6006/?path=/story/components-emptystate--custom-content');

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('custom-content', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });
});
