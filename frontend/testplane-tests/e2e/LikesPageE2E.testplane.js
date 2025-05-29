describe('Likes Page Functionality', () => {
  it('should display users who liked the current user', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(390, 840);

    // Navigate to the likes page (assuming user is already logged in)
    await browser.url('https://testing.couply.ru/profile');

    // Wait for the page to load completely
    await browser.pause(3000);

    // Log current URL for debugging
    const currentUrl = await browser.getUrl();
    console.log('Current URL:', currentUrl);

    // Check if matches section is displayed
    const matchesSection = await browser.$('.matchesSection');
    const sectionExists = await matchesSection.isExisting();
    console.log('Matches section exists:', sectionExists);

    // Check if match cards are displayed
    const matchCards = await browser.$$('.matchCard');
    console.log('Number of match cards found:', matchCards.length);

    if (matchCards.length > 0) {
      // Check details of the first match card
      const firstCard = matchCards[0];

      // Check if profile image is displayed
      const profileImage = await firstCard.$('img');
      const imageExists = await profileImage.isExisting();
      console.log('Profile image exists in first card:', imageExists);

      // Check if user name is displayed
      const userName = await firstCard.$('h1, h2, h3, .userName');
      const nameExists = await userName.isExisting();
      console.log('User name exists in first card:', nameExists);

      if (nameExists) {
        const nameText = await userName.getText();
        console.log('User name in first card:', nameText);
      }
    } else {
      console.log('No match cards found, checking for empty state');

      // Check if empty state is displayed
      const emptyState = await browser.$('.emptyState, .noLikes');
      const emptyStateExists = await emptyState.isExisting();
      console.log('Empty state exists:', emptyStateExists);
    }
  });

  it('should allow matching with users who liked the current user', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // Navigate to the likes page
    await browser.url('https://moonlit-valkyrie-fdbfc4.netlify.app/likes');

    // Wait for the page to load completely
    await browser.pause(3000);

    // Check if match cards are displayed
    const matchCards = await browser.$$('.matchCard');
    console.log('Number of match cards found:', matchCards.length);

    if (matchCards.length > 0) {
      // Click on the first match card to view details
      await matchCards[0].click();
      console.log('Clicked on first match card');

      // Wait for details to load
      await browser.pause(2000);

      // Find and click the like/match button
      const likeButton = await browser.$('.likeButton, .matchButton');
      const likeButtonExists = await likeButton.isExisting();
      console.log('Like button exists:', likeButtonExists);

      if (likeButtonExists) {
        await likeButton.click();
        console.log('Clicked like button');

        // Wait for match modal to appear
        await browser.pause(2000);

        // Check if match modal appeared
        const matchModal = await browser.$('.matchModal, .modal');
        const matchModalExists = await matchModal.isExisting();
        console.log('Match modal exists:', matchModalExists);

        if (matchModalExists) {
          console.log('Match modal detected');

          // Check for message button in the match modal
          const messageButton = await matchModal.$('.messageButton, .sendButton');
          const messageButtonExists = await messageButton.isExisting();
          console.log('Message button exists:', messageButtonExists);

          // If there's a close button, close the match modal
          const closeButton = await matchModal.$('.closeButton, .continueButton');
          if (await closeButton.isExisting()) {
            await closeButton.click();
            console.log('Closed match modal');
            await browser.pause(1000);
          }
        }
      }
    } else {
      console.log('No match cards found to test matching functionality');
    }
  });

  it('should navigate between likes and matches tabs', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // Navigate to the likes page
    await browser.url('https://moonlit-valkyrie-fdbfc4.netlify.app/likes');

    // Wait for the page to load completely
    await browser.pause(3000);

    // Check if tabs section exists
    const tabsSection = await browser.$('.tabsSection');
    const tabsSectionExists = await tabsSection.isExisting();
    console.log('Tabs section exists:', tabsSectionExists);

    if (tabsSectionExists) {
      // Find all tabs
      const tabs = await tabsSection.$$('button, .tab');
      console.log('Number of tabs found:', tabs.length);

      if (tabs.length >= 2) {
        // Click on the second tab (usually "Matches")
        await tabs[1].click();
        console.log('Clicked on second tab');

        // Wait for content to update
        await browser.pause(2000);

        // Check if matches are displayed
        const matchesContent = await browser.$('.matchesSection, .matchesTab');
        const matchesContentExists = await matchesContent.isExisting();
        console.log('Matches content exists:', matchesContentExists);

        // Click back on the first tab (usually "Likes")
        await tabs[0].click();
        console.log('Clicked back on first tab');

        // Wait for content to update
        await browser.pause(2000);

        // Check if likes are displayed
        const likesContent = await browser.$('.likesSection, .likesTab');
        const likesContentExists = await likesContent.isExisting();
        console.log('Likes content exists:', likesContentExists);
      }
    } else {
      console.log('No tabs section found');
    }
  });

  it('should open match modal when clicking on a match', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // Navigate to the likes page
    await browser.url('https://moonlit-valkyrie-fdbfc4.netlify.app/likes');

    // Wait for the page to load completely
    await browser.pause(3000);

    // Check if tabs section exists and click on matches tab
    const tabsSection = await browser.$('.tabsSection');
    if (await tabsSection.isExisting()) {
      const tabs = await tabsSection.$$('button, .tab');
      if (tabs.length >= 2) {
        // Click on the second tab (usually "Matches")
        await tabs[1].click();
        console.log('Clicked on matches tab');
        await browser.pause(2000);
      }
    }

    // Check if match cards are displayed
    const matchCards = await browser.$$('.matchCard');
    console.log('Number of match cards found:', matchCards.length);

    if (matchCards.length > 0) {
      // Click on the first match card
      await matchCards[0].click();
      console.log('Clicked on first match card');

      // Wait for match modal to appear
      await browser.pause(2000);

      // Check if match modal appeared
      const matchModal = await browser.$('.matchModal, .modal');
      const matchModalExists = await matchModal.isExisting();
      console.log('Match modal exists:', matchModalExists);

      if (matchModalExists) {
        console.log('Match modal detected');

        // Check for user name in the modal
        const userName = await matchModal.$('h1, h2, h3, .userName');
        if (await userName.isExisting()) {
          const nameText = await userName.getText();
          console.log('User name in match modal:', nameText);
        }

        // Check for message button
        const messageButton = await matchModal.$('.messageButton, .sendButton');
        const messageButtonExists = await messageButton.isExisting();
        console.log('Message button exists:', messageButtonExists);

        // If there's a close button, close the match modal
        const closeButton = await matchModal.$('.closeButton, .continueButton');
        if (await closeButton.isExisting()) {
          await closeButton.click();
          console.log('Closed match modal');
          await browser.pause(1000);
        }
      }
    } else {
      console.log('No match cards found to test match modal');
    }
  });
});
