describe('ProfileVisibilitySection Component', () => {
  it('should match the reference screenshot for visible profile', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the ProfileVisibilitySection component with visible profile
    await browser.url(
      'http://localhost:6006/?path=/story/features-profilevisibility-profilevisibilitysection--visible',
    );

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('visible', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot for hidden profile', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the ProfileVisibilitySection component with hidden profile
    await browser.url(
      'http://localhost:6006/?path=/story/features-profilevisibility-profilevisibilitysection--hidden',
    );

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('hidden', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot in loading state', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the ProfileVisibilitySection component in loading state
    await browser.url(
      'http://localhost:6006/?path=/story/features-profilevisibility-profilevisibilitysection--loading',
    );

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('loading', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });
});
