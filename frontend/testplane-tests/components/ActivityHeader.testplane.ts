describe('ActivityHeader Component', () => {
  it('should match the reference screenshot for default state', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the ActivityHeader component
    await browser.url(
      'http://localhost:6006/?path=/story/widgets-activityhistory-activityheader--default',
    );

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('default', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot with custom title', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the ActivityHeader component with custom title
    await browser.url(
      'http://localhost:6006/?path=/story/widgets-activityhistory-activityheader--custom-title',
    );

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('custom-title', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });
});
