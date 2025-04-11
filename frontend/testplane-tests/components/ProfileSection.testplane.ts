describe('ProfileSection Component', () => {
  it('should match the reference screenshot for default state', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(375, 667); // Mobile size since it's a mobile component
    
    // Navigate to the Storybook page for the ProfileSection component
    await browser.url('http://localhost:6006/?path=/story/components-profilesection--default');
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView('default', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot with edit link', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(375, 667); // Mobile size since it's a mobile component
    
    // Navigate to the Storybook page for the ProfileSection component with edit link
    await browser.url('http://localhost:6006/?path=/story/components-profilesection--with-edit-link');
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView('with-edit-link', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });

  it('should match the reference screenshot with list content', async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(375, 667); // Mobile size since it's a mobile component
    
    // Navigate to the Storybook page for the ProfileSection component with list content
    await browser.url('http://localhost:6006/?path=/story/components-profilesection--with-list-content');
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView('with-list-content', '#storybook-preview-iframe', {
      allowViewportOverflow: true,
    });
  });
});