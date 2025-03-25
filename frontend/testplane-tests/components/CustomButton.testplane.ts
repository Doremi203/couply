describe("CustomButton Component", () => {
  it("should match the reference screenshot for default state", async ({ browser }) => {
    // Navigate to the Storybook page for the CustomButton component
    await browser.url("http://localhost:6006/?path=/story/components-custombutton--default");
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView("default", ".sb-show-main");
  });

  it("should match the reference screenshot for disabled state", async ({ browser }) => {
    // Navigate to the Storybook page for the disabled CustomButton
    await browser.url("http://localhost:6006/?path=/story/components-custombutton--disabled");
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView("disabled", ".sb-show-main");
  });
});