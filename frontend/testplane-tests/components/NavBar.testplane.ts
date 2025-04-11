describe('NavBar Component', () => {
  it('should match the reference screenshot for Home active state', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);
    
    // Navigate to the Storybook page for the NavBar component with Home active
    await browser.url('http://localhost:6006/?path=/story/components-navbar--home-active');
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView('home-active', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot for Likes active state', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);
    
    // Navigate to the Storybook page for the NavBar component with Likes active
    await browser.url('http://localhost:6006/?path=/story/components-navbar--likes-active');
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView('likes-active', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot for Profile active state', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);
    
    // Navigate to the Storybook page for the NavBar component with Profile active
    await browser.url('http://localhost:6006/?path=/story/components-navbar--profile-active');
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView('profile-active', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });
});