describe("CustomInput Component", () => {
  it("should match the reference screenshot for text input", async ({ browser }) => {
    // Navigate to the Storybook page for the text CustomInput
    await browser.url("http://localhost:6006/?path=/story/components-custominput--text");
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView("text", ".sb-show-main");
  });

  it("should match the reference screenshot for password input", async ({ browser }) => {
    // Navigate to the Storybook page for the password CustomInput
    await browser.url("http://localhost:6006/?path=/story/components-custominput--password");
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView("password", ".sb-show-main");
  });

  it("should match the reference screenshot for email input", async ({ browser }) => {
    // Navigate to the Storybook page for the email CustomInput
    await browser.url("http://localhost:6006/?path=/story/components-custominput--email");
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView("email", ".sb-show-main");
  });
});