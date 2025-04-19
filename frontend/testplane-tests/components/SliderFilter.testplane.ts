describe('SliderFilter Component', () => {
  it('should match the reference screenshot for single value slider', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the SliderFilter component with single value
    await browser.url(
      'http://localhost:6006/?path=/story/features-filters-sliderfilter--single-value',
    );

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('single-value', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot for range value slider', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the SliderFilter component with range value
    await browser.url(
      'http://localhost:6006/?path=/story/features-filters-sliderfilter--range-value',
    );

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('range-value', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot for slider with value labels', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);

    // Navigate to the Storybook page for the SliderFilter component with value labels
    await browser.url(
      'http://localhost:6006/?path=/story/features-filters-sliderfilter--with-value-labels',
    );

    // Wait for the component to render
    await browser.pause(1000);

    // Take a screenshot and compare it with the reference
    await browser.assertView('with-value-labels', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });
});
