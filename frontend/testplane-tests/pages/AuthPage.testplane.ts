describe("AuthPage Component", () => {
  it("should match the reference screenshot", async ({ browser }) => {
    // Navigate to the Storybook page for the AuthPage component
    await browser.url("http://localhost:6006/?path=/story/pages-authpage--default");
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView("default", ".sb-show-main");
    
    // Test button hover state
    const loginButton = await browser.$("button:nth-child(1)");
    await loginButton.moveTo();
    
    // Wait for hover effect
    await browser.pause(500);
    
    // Take a screenshot of the hover state
    await browser.assertView("button-hover", ".sb-show-main");
  });
});