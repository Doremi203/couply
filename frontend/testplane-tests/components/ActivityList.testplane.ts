describe('ActivityList Component', () => {
  it('should match the reference screenshot for default state', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the ActivityList component with default items
    await browser.url(
      'http://localhost:6006/?path=/story/widgets-activityhistory-activitylist--default',
    );

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('default', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot for empty list', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the ActivityList component with empty list
    await browser.url(
      'http://localhost:6006/?path=/story/widgets-activityhistory-activitylist--empty',
    );

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('empty', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot with many items', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the ActivityList component with many items
    await browser.url(
      'http://localhost:6006/?path=/story/widgets-activityhistory-activitylist--many-items',
    );

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('many-items', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });
});
