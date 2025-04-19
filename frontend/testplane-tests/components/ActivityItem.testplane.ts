describe('ActivityItem Component', () => {
  it('should match the reference screenshot for view activity', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the ActivityItem component with view activity
    await browser.url(
      'http://localhost:6006/?path=/story/widgets-activityhistory-activityitem--view-activity',
    );

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('view-activity', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot for like activity', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the ActivityItem component with like activity
    await browser.url(
      'http://localhost:6006/?path=/story/widgets-activityhistory-activityitem--like-activity',
    );

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('like-activity', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot for message activity', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the ActivityItem component with message activity
    await browser.url(
      'http://localhost:6006/?path=/story/widgets-activityhistory-activityitem--message-activity',
    );

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('message-activity', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });
});
