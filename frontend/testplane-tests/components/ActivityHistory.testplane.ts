describe('ActivityHistory Component', () => {
  it('should match the reference screenshot with activity items', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the ActivityHistory component with activity items
    await browser.url('http://localhost:6006/?path=/story/widgets-activityhistory--default');

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('with-items', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot with empty state', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the ActivityHistory component with empty state
    await browser.url('http://localhost:6006/?path=/story/widgets-activityhistory--empty');

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('empty', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot with many activity items', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the ActivityHistory component with many items
    await browser.url('http://localhost:6006/?path=/story/widgets-activityhistory--many-items');

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('many-items', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });
});
