describe("ButtonWithIcon Component", () => {
  it("should match the reference screenshot", async ({ browser }) => {
    // Navigate to the Storybook page for the ButtonWithIcon component
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);
    
    await browser.url("http://localhost:6006/?path=/story/components-buttonwithicon--default");
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    // Use the iframe selector for Storybook v8 with allowViewportOverflow option
    await browser.assertView("default", "#storybook-preview-iframe", {
      allowViewportOverflow: true
    });
  });
});