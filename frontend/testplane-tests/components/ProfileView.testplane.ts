describe('ProfileView Component', () => {
  it('should match the reference screenshot for default state', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the ProfileView component with default state
    await browser.url('http://localhost:6006/?path=/story/widgets-profileview--default');

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('default', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot for profile that has liked you', async ({
    browser,
  }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the ProfileView component with a profile that has liked you
    await browser.url('http://localhost:6006/?path=/story/widgets-profileview--has-liked-you');

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('has-liked-you', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot for minimal profile', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the ProfileView component with minimal profile data
    await browser.url('http://localhost:6006/?path=/story/widgets-profileview--minimal-profile');

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('minimal-profile', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });
});
