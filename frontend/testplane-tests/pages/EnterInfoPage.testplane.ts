describe("EnterInfoPage Component", () => {
  it("should match the reference screenshot for the first step", async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(1920, 1080);
    
    // Navigate to the Storybook page for the EnterInfoPage component
    await browser.url("http://localhost:6006/?path=/story/pages-enterinfopage--default");
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView("step1", "#storybook-preview-iframe", {
      allowViewportOverflow: true
    });
    
    // Fill in the name field
    const nameInput = await browser.$("input[type='text']");
    await nameInput.setValue("Test User");
    
    // Take a screenshot of the filled form
    await browser.assertView("step1-filled", "#storybook-preview-iframe", {
      allowViewportOverflow: true
    });
    
    // Click the next button
    const nextButton = await browser.$("button");
    await nextButton.click();
    
    // Wait for the next step to render
    await browser.pause(500);
    
    // Take a screenshot of the second step
    await browser.assertView("step2", "#storybook-preview-iframe", {
      allowViewportOverflow: true
    });
  });
});