describe('ProfilePhotoEdit Component', () => {
  it('should match the reference screenshot with photo', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the ProfilePhotoEdit component with photo
    await browser.url(
      'http://localhost:6006/?path=/story/features-profileedit-profilephotoedit--with-photo',
    );

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('with-photo', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot without photo', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the ProfilePhotoEdit component without photo
    await browser.url(
      'http://localhost:6006/?path=/story/features-profileedit-profilephotoedit--without-photo',
    );

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('without-photo', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot in loading state', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the ProfilePhotoEdit component in loading state
    await browser.url(
      'http://localhost:6006/?path=/story/features-profileedit-profilephotoedit--loading',
    );

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('loading', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });
});
