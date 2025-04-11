describe('MatchCard Component', () => {
  it('should match the reference screenshot for default variant', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);
    
    // Navigate to the Storybook page for the MatchCard component
    await browser.url('http://localhost:6006/?path=/story/pages-likespage-matchcard--default');
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView('default', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot with chat message', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);
    
    // Navigate to the Storybook page for the MatchCard component with chat message
    await browser.url('http://localhost:6006/?path=/story/pages-likespage-matchcard--with-chat-message');
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView('with-chat-message', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });
});