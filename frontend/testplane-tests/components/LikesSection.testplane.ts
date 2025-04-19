describe('LikesSection Component', () => {
  it('should match the reference screenshot with likes', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the LikesSection component with likes
    await browser.url('http://localhost:6006/?path=/story/features-likes-likessection--with-likes');

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('with-likes', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot with empty state', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the LikesSection component with empty state
    await browser.url(
      'http://localhost:6006/?path=/story/features-likes-likessection--empty-state',
    );

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('empty-state', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });
});
