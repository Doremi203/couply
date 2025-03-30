describe("ProfileHeader Component", () => {
  it("should match the reference screenshot for default state", async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(375, 667); // Mobile size since it's a mobile component
    
    // Navigate to the Storybook page for the ProfileHeader component
    await browser.url("http://localhost:6006/?path=/story/components-profileheader--default");
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView("default", "#storybook-preview-iframe", {
      allowViewportOverflow: true
    });
  });

  it("should match the reference screenshot for hidden state", async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(375, 667); // Mobile size since it's a mobile component
    
    // Navigate to the Storybook page for the ProfileHeader component in hidden state
    await browser.url("http://localhost:6006/?path=/story/components-profileheader--hidden");
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView("hidden", "#storybook-preview-iframe", {
      allowViewportOverflow: true
    });
  });
});