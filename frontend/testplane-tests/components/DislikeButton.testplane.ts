describe('DislikeButton Component', () => {
  it('should match the reference screenshot for default variant', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);
    
    // Navigate to the Storybook page for the DislikeButton component
    await browser.url('http://localhost:6006/?path=/story/shared-dislikebutton--default');
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView('default', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot for custom class variant', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);
    
    // Navigate to the Storybook page for the DislikeButton component with custom class
    await browser.url('http://localhost:6006/?path=/story/shared-dislikebutton--with-custom-class');
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView('with-custom-class', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });
});