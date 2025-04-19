describe('EditProfile Component', () => {
  it('should match the reference screenshot for default state', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the EditProfile component with default state
    await browser.url('http://localhost:6006/?path=/story/widgets-editprofile--default');

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('default', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot for empty profile', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the EditProfile component with empty profile
    await browser.url('http://localhost:6006/?path=/story/widgets-editprofile--empty-profile');

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('empty-profile', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot for hidden profile', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the EditProfile component with hidden profile
    await browser.url('http://localhost:6006/?path=/story/widgets-editprofile--hidden-profile');

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('hidden-profile', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });
});
