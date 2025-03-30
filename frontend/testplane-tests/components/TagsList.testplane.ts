describe("TagsList Component", () => {
  it("should match the reference screenshot for default state", async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(375, 667); // Mobile size since it's a mobile component
    
    // Navigate to the Storybook page for the TagsList component
    await browser.url("http://localhost:6006/?path=/story/components-tagslist--default");
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView("default", "#storybook-preview-iframe", {
      allowViewportOverflow: true
    });
  });

  it("should match the reference screenshot with common interests", async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(375, 667); // Mobile size since it's a mobile component
    
    // Navigate to the Storybook page for the TagsList component with common interests
    await browser.url("http://localhost:6006/?path=/story/components-tagslist--with-common-interests");
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView("with-common-interests", "#storybook-preview-iframe", {
      allowViewportOverflow: true
    });
  });

  it("should match the reference screenshot with single item", async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(375, 667); // Mobile size since it's a mobile component
    
    // Navigate to the Storybook page for the TagsList component with single item
    await browser.url("http://localhost:6006/?path=/story/components-tagslist--single-item");
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView("single-item", "#storybook-preview-iframe", {
      allowViewportOverflow: true
    });
  });

  it("should match the reference screenshot with many items", async ({ browser }) => {
    // Set window size to ensure consistent screenshots
    await browser.setWindowSize(375, 667); // Mobile size since it's a mobile component
    
    // Navigate to the Storybook page for the TagsList component with many items
    await browser.url("http://localhost:6006/?path=/story/components-tagslist--many-items");
    
    // Wait for the component to render
    await browser.pause(1000);
    
    // Take a screenshot and compare it with the reference
    await browser.assertView("many-items", "#storybook-preview-iframe", {
      allowViewportOverflow: true
    });
  });
});