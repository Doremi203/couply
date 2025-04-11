describe('PhotoGallery Component', () => {
  it('should match the reference screenshot for default state', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(375, 667); // Mobile size since it's a mobile component
    
    // Navigate to the Storybook page for the PhotoGallery component
    await browser.url('http://localhost:6006/?path=/story/components-photogallery--default');
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView('default', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot for single photo', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(375, 667); // Mobile size since it's a mobile component
    
    // Navigate to the Storybook page for the PhotoGallery component with single photo
    await browser.url('http://localhost:6006/?path=/story/components-photogallery--single-photo');
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView('single-photo', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot for two photos', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(375, 667); // Mobile size since it's a mobile component
    
    // Navigate to the Storybook page for the PhotoGallery component with two photos
    await browser.url('http://localhost:6006/?path=/story/components-photogallery--two-photos');
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView('two-photos', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });
});