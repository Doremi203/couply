describe('PageHeader Component', () => {
  it('should match the reference screenshot with title only', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the PageHeader component with title only
    await browser.url('http://localhost:6006/?path=/story/components-pageheader--title-only');

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('title-only', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot with back button', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the PageHeader component with back button
    await browser.url('http://localhost:6006/?path=/story/components-pageheader--with-back-button');

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('with-back-button', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot with action button', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the PageHeader component with action button
    await browser.url(
      'http://localhost:6006/?path=/story/components-pageheader--with-action-button',
    );

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('with-action-button', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });
});
