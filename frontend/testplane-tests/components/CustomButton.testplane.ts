describe("CustomButton Component", () => {
  it("should match the reference screenshot for default state", async ({ browser }) => {
    // Navigate to the Storybook page for the CustomButton component
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);
    
    await browser.url("http://localhost:6006/?path=/story/components-custombutton--default");
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView("default", "#storybook-preview-iframe", {
      allowViewportOverflow: true
    });
  });

  it("should match the reference screenshot for disabled state", async ({ browser }) => {
    // Navigate to the Storybook page for the disabled CustomButton
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);
    
    await browser.url("http://localhost:6006/?path=/story/components-custombutton--disabled");
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView("disabled", "#storybook-preview-iframe", {
      allowViewportOverflow: true
    });
  });
});