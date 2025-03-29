describe("IconButton Component", () => {
  it("should match the reference screenshot for default variant", async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);
    
    // Navigate to the Storybook page for the IconButton component
    await browser.url("http://localhost:6006/?path=/story/shared-circleiconbutton--default");
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView("default", "#storybook-preview-iframe", {
      allowViewportOverflow: true
    });
  });

  it("should match the reference screenshot for custom color variant", async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);
    
    // Navigate to the Storybook page for the IconButton component with custom color
    await browser.url("http://localhost:6006/?path=/story/shared-circleiconbutton--custom-color");
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView("custom-color", "#storybook-preview-iframe", {
      allowViewportOverflow: true
    });
  });
});