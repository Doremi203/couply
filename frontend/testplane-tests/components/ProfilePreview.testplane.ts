describe('ProfilePreview Component', () => {
  it('should match the reference screenshot for profile preview', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(375, 667); // Mobile size since it's a mobile component
    
    // Navigate to the Storybook page for the ProfilePreview component
    await browser.url('http://localhost:6006/?path=/story/components-profilepreview--default');
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView('default', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
    
    // Try to click on the scroll indicator to see the collapsed state
    await browser.execute(() => {
      const iframe = document.querySelector('#storybook-preview-iframe') as HTMLIFrameElement;
      if (iframe && iframe.contentDocument) {
        const scrollIndicator = iframe.contentDocument.querySelector('.photoScrollIndicator') as HTMLElement;
        if (scrollIndicator) {
          scrollIndicator.click();
        }
      }
    });
    
    // Wait for the animation to complete
    await browser.pause(1000);
    
    // Take a screenshot of the collapsed state
    await browser.assertView('collapsed', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });
});