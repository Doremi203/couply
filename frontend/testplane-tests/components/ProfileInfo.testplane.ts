describe('ProfileInfo Component', () => {
  it('should match the reference screenshot for verified profile', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(375, 667); // Mobile size since it's a mobile component
    
    // Navigate to the Storybook page for the ProfileInfo component
    await browser.url('http://localhost:6006/?path=/story/components-profileinfo--verified');
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView('verified', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot for unverified profile', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(375, 667); // Mobile size since it's a mobile component
    
    // Navigate to the Storybook page for the ProfileInfo component with unverified profile
    await browser.url('http://localhost:6006/?path=/story/components-profileinfo--unverified');
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView('unverified', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });
});