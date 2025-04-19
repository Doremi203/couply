describe('InterestsSection Component', () => {
  it('should match the reference screenshot with interests', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the InterestsSection component with interests
    await browser.url(
      'http://localhost:6006/?path=/story/components-interestssection--with-interests',
    );

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('with-interests', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot with no interests', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the InterestsSection component with no interests
    await browser.url(
      'http://localhost:6006/?path=/story/components-interestssection--no-interests',
    );

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('no-interests', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot with many interests', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the InterestsSection component with many interests
    await browser.url(
      'http://localhost:6006/?path=/story/components-interestssection--many-interests',
    );

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('many-interests', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });
});
