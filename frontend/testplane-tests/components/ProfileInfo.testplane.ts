describe('ProfileInfo Component', () => {
  it('should match the reference screenshot for collapsed state', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the ProfileInfo component in collapsed state
    await browser.url(
      'http://localhost:6006/?path=/story/widgets-profileview-profileinfo--collapsed',
    );

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('collapsed', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot for expanded state', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the ProfileInfo component in expanded state
    await browser.url(
      'http://localhost:6006/?path=/story/widgets-profileview-profileinfo--expanded',
    );

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('expanded', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot for minimal profile', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the ProfileInfo component with minimal profile
    await browser.url(
      'http://localhost:6006/?path=/story/widgets-profileview-profileinfo--minimal-profile',
    );

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('minimal-profile', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });
});
