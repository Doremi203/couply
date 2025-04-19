describe('ActivityIcon Component', () => {
  it('should match the reference screenshot for online state', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the ActivityIcon component in online state
    await browser.url('http://localhost:6006/?path=/story/components-activityicon--online');

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('online', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot for offline state', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the ActivityIcon component in offline state
    await browser.url('http://localhost:6006/?path=/story/components-activityicon--offline');

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('offline', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });
});
