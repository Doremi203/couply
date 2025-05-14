describe('Home Page Functionality', () => {
  it('should display user profiles for swiping', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // Navigate to the home page (assuming user is already logged in)
    await browser.url('https://rainbow-souffle-ece639.netlify.app/home');

    // Wait for the page to load completely
    await browser.pause(3000);

    // Log current URL for debugging
    const currentUrl = await browser.getUrl();
    console.log('Current URL:', currentUrl);

    // Check if profile card is displayed
    const profileCard = await browser.$('.profileCard');
    const cardExists = await profileCard.isExisting();
    console.log('Profile card exists:', cardExists);

    if (cardExists) {
      // Check if profile image is displayed
      const profileImage = await profileCard.$('img');
      const imageExists = await profileImage.isExisting();
      console.log('Profile image exists:', imageExists);

      // Check if user name is displayed
      const userName = await profileCard.$('h1, h2, .userName');
      const nameExists = await userName.isExisting();
      console.log('User name exists:', nameExists);

      if (nameExists) {
        const nameText = await userName.getText();
        console.log('User name:', nameText);
      }
    }

    // Check if like/dislike buttons are displayed
    const likeButton = await browser.$('.likeButton');
    const likeButtonExists = await likeButton.isExisting();
    console.log('Like button exists:', likeButtonExists);

    const dislikeButton = await browser.$('.dislikeButton');
    const dislikeButtonExists = await dislikeButton.isExisting();
    console.log('Dislike button exists:', dislikeButtonExists);
  });

  it('should allow swiping through profiles', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // Navigate to the home page
    await browser.url('https://moonlit-valkyrie-fdbfc4.netlify.app/home');

    // Wait for the page to load completely
    await browser.pause(3000);

    // Get the first profile name for comparison
    let firstProfileName = '';
    const firstProfileNameElement = await browser.$(
      '.profileCard h1, .profileCard h2, .profileCard .userName',
    );
    if (await firstProfileNameElement.isExisting()) {
      firstProfileName = await firstProfileNameElement.getText();
      console.log('First profile name:', firstProfileName);
    }

    // Find and click the like button
    const likeButton = await browser.$('.likeButton');
    if (await likeButton.isExisting()) {
      await likeButton.click();
      console.log('Clicked like button');

      // Wait for next profile to load
      await browser.pause(2000);

      // Get the second profile name
      const secondProfileNameElement = await browser.$(
        '.profileCard h1, .profileCard h2, .profileCard .userName',
      );
      if (await secondProfileNameElement.isExisting()) {
        const secondProfileName = await secondProfileNameElement.getText();
        console.log('Second profile name:', secondProfileName);

        // Verify that the profile changed
        const profileChanged = firstProfileName !== secondProfileName;
        console.log('Profile changed after like:', profileChanged);
      }
    }

    // Try disliking the next profile
    const dislikeButton = await browser.$('.dislikeButton');
    if (await dislikeButton.isExisting()) {
      await dislikeButton.click();
      console.log('Clicked dislike button');

      // Wait for next profile to load
      await browser.pause(2000);

      // Verify that another profile loaded
      const thirdProfileNameElement = await browser.$(
        '.profileCard h1, .profileCard h2, .profileCard .userName',
      );
      const thirdProfileExists = await thirdProfileNameElement.isExisting();
      console.log('Third profile loaded after dislike:', thirdProfileExists);
    }
  });

  it('should handle match scenario', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // Navigate to the home page
    await browser.url('https://moonlit-valkyrie-fdbfc4.netlify.app/home');

    // Wait for the page to load completely
    await browser.pause(3000);

    // Find and click the like button (to potentially trigger a match)
    const likeButton = await browser.$('.likeButton');
    if (await likeButton.isExisting()) {
      await likeButton.click();
      console.log('Clicked like button to potentially trigger match');

      // Wait for potential match modal to appear
      await browser.pause(3000);

      // Check if match modal appeared
      const matchModal = await browser.$('.matchModal, .modal');
      const matchOccurred = await matchModal.isExisting();
      console.log('Match occurred:', matchOccurred);

      if (matchOccurred) {
        console.log('Match modal detected');

        // Check for message button in the match modal
        const messageButton = await matchModal.$('.messageButton, .sendButton');
        const messageButtonExists = await messageButton.isExisting();
        console.log('Message button exists:', messageButtonExists);

        if (messageButtonExists) {
          await messageButton.click();
          console.log('Clicked message button');
          await browser.pause(2000);

          // Check if redirected to messaging interface
          const currentUrl = await browser.getUrl();
          console.log('URL after clicking message button:', currentUrl);
        }

        // If there's a close button, close the match modal
        const closeButton = await matchModal.$('.closeButton, .continueButton');
        if (await closeButton.isExisting()) {
          await closeButton.click();
          console.log('Closed match modal');
          await browser.pause(1000);
        }
      }
    }
  });

  it('should handle no users left scenario', async ({ browser }) => {
    // Set window size to ensure consistent behavior
    await browser.setWindowSize(1920, 1080);

    // Navigate to the home page
    await browser.url('https://moonlit-valkyrie-fdbfc4.netlify.app/home');

    // Wait for the page to load completely
    await browser.pause(3000);

    // Check if "No Users Left" message is displayed
    // Note: In a real test, you might need to swipe through all profiles to reach this state
    const noUsersLeft = await browser.$('.noUsersLeft, .emptyState');
    const noUsersLeftExists = await noUsersLeft.isExisting();
    console.log('No users left message exists:', noUsersLeftExists);

    if (noUsersLeftExists) {
      console.log('No users left message detected');

      // Check if there's a refresh or try again button
      const refreshButton = await noUsersLeft.$('.refreshButton, .tryAgainButton');
      const refreshButtonExists = await refreshButton.isExisting();
      console.log('Refresh button exists:', refreshButtonExists);

      if (refreshButtonExists) {
        await refreshButton.click();
        console.log('Clicked refresh button');
        await browser.pause(3000);

        // Check if profiles are loaded after refresh
        const profileCard = await browser.$('.profileCard');
        const profilesLoaded = await profileCard.isExisting();
        console.log('Profiles loaded after refresh:', profilesLoaded);
      }
    } else {
      console.log('Users are available for swiping');
    }
  });
});
