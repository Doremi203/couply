describe("ToggleButtons Component", () => {
  it("should match the reference screenshot for default state", async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);
    
    // Navigate to the Storybook page for the default ToggleButtons
    await browser.url("http://localhost:6006/?path=/story/components-togglebuttons--default");
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView("default", "#storybook-preview-iframe", {
      allowViewportOverflow: true
    });
  });

  it("should match the reference screenshot for three options", async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);
    
    // Navigate to the Storybook page for the ToggleButtons with three options
    await browser.url("http://localhost:6006/?path=/story/components-togglebuttons--three-options");
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView("three-options", "#storybook-preview-iframe", {
      allowViewportOverflow: true
    });
  });

  it("should match the reference screenshot for no selection", async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);
    
    // Navigate to the Storybook page for the ToggleButtons with no selection
    await browser.url("http://localhost:6006/?path=/story/components-togglebuttons--no-selection");
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView("no-selection", "#storybook-preview-iframe", {
      allowViewportOverflow: true
    });
  });
});