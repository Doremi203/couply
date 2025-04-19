describe('BasicInfoForm Component', () => {
  it('should match the reference screenshot for empty form', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the BasicInfoForm component with empty form
    await browser.url(
      'http://localhost:6006/?path=/story/features-profileedit-basicinfoform--empty',
    );

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('empty', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot with filled data', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the BasicInfoForm component with filled data
    await browser.url(
      'http://localhost:6006/?path=/story/features-profileedit-basicinfoform--filled',
    );

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('filled', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot with validation errors', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the BasicInfoForm component with validation errors
    await browser.url(
      'http://localhost:6006/?path=/story/features-profileedit-basicinfoform--with-errors',
    );

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('with-errors', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });
});
