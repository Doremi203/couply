describe('CustomInput Component', () => {
  it('should match the reference screenshot for text input', async ({ browser }) => {
    // Navigate to the Storybook page for the text CustomInput
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);
    
    await browser.url('http://localhost:6006/?path=/story/components-custominput--text');
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView('text', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot for password input', async ({ browser }) => {
    // Navigate to the Storybook page for the password CustomInput
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);
    
    await browser.url('http://localhost:6006/?path=/story/components-custominput--password');
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView('password', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot for email input', async ({ browser }) => {
    // Navigate to the Storybook page for the email CustomInput
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);
    
    await browser.url('http://localhost:6006/?path=/story/components-custominput--email');
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView('email', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });
});