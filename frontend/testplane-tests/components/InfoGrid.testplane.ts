describe('InfoGrid Component', () => {
  it('should match the reference screenshot for default state', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(375, 667); // Mobile size since it's a mobile component
    
    // Navigate to the Storybook page for the InfoGrid component
    await browser.url('http://localhost:6006/?path=/story/components-infogrid--default');
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView('default', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot for short list', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(375, 667); // Mobile size since it's a mobile component
    
    // Navigate to the Storybook page for the InfoGrid component with short list
    await browser.url('http://localhost:6006/?path=/story/components-infogrid--short-list');
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView('short-list', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });
});