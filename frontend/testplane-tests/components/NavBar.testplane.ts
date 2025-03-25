describe("NavBar Component", () => {
  it("should match the reference screenshot", async ({ browser }) => {
    // Navigate to the Storybook page for the NavBar component
    await browser.url("http://localhost:6006/?path=/story/components-navbar--default");
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView("default", ".sb-show-main");
  });
});