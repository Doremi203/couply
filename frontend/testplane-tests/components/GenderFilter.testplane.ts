describe('GenderFilter Component', () => {
  it('should match the reference screenshot with Girls selected', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the GenderFilter component with Girls selected
    await browser.url('http://localhost:6006/?path=/story/features-filters-genderfilter--girls');

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('girls', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot with Boys selected', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the GenderFilter component with Boys selected
    await browser.url('http://localhost:6006/?path=/story/features-filters-genderfilter--boys');

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('boys', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot with Both selected', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the GenderFilter component with Both selected
    await browser.url('http://localhost:6006/?path=/story/features-filters-genderfilter--both');

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('both', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });
});
