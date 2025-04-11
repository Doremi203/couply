describe('CommonInterest Component', () => {
  it('should match the reference screenshot for default state', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);
    
    // Navigate to the Storybook page for the CommonInterest component
    await browser.url('http://localhost:6006/?path=/story/components-commoninterest--default');
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView('default', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot for common state', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);
    
    // Navigate to the Storybook page for the CommonInterest component with common state
    await browser.url('http://localhost:6006/?path=/story/components-commoninterest--common');
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView('common', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });
});