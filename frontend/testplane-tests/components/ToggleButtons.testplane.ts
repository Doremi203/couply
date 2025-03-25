describe("ToggleButtons Component", () => {
  it("should match the reference screenshot for default state", async ({ browser }) => {
    // Navigate to the Storybook page for the default ToggleButtons
    await browser.url("http://localhost:6006/?path=/story/components-togglebuttons--default");
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView("default", ".sb-show-main");
  });

  it("should match the reference screenshot for three options", async ({ browser }) => {
    // Navigate to the Storybook page for the ToggleButtons with three options
    await browser.url("http://localhost:6006/?path=/story/components-togglebuttons--three-options");
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView("three-options", ".sb-show-main");
  });

  it("should match the reference screenshot for no selection", async ({ browser }) => {
    // Navigate to the Storybook page for the ToggleButtons with no selection
    await browser.url("http://localhost:6006/?path=/story/components-togglebuttons--no-selection");
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView("no-selection", ".sb-show-main");
  });
});