import { useState, useCallback, useMemo } from 'react';

import { LikeProfile, MatchProfile } from '../types';

// TODO
export const useProfileView = () => {
  const [selectedProfile, setSelectedProfile] = useState<LikeProfile | null>(null);

  const handleProfileClick = useCallback((profile: LikeProfile) => {
    // When clicking on a profile, open the ProfileView
    setSelectedProfile(profile);
  }, []);

  const handleMatchClick = useCallback((match: MatchProfile) => {
    // Convert match to a profile format that ProfileView can use
    const matchAsProfile = {
      id: match.id,
      name: match.name,
      age: match.age,
      imageUrl: match.imageUrl,
      hasLikedYou: true,
      bio: 'This is a match! You can contact them via social media.',
      location: 'Matched User',
      photos: [match.imageUrl, match.imageUrl],
      // Add passion property to ensure the ProfileView component can display interests
      passion: ['Match', 'Connection'],
    };
    setSelectedProfile(matchAsProfile);
  }, []);

  const handleCloseProfile = useCallback(() => {
    // First set the profile to null to remove it from the DOM
    setSelectedProfile(null);

    // Then ensure any body styles or other side effects are cleaned up
    document.body.style.overflow = 'auto'; // Re-enable scrolling

    // Force a re-render of the NavBar by triggering a small timeout
    setTimeout(() => {
      const navBar = document.querySelector('.navBarContainer');
      if (navBar) {
        // Temporarily modify a style to force a repaint
        navBar.classList.add('force-repaint');
        setTimeout(() => {
          navBar.classList.remove('force-repaint');
        }, 10);
      }
    }, 100);
  }, []);

  return useMemo(
    () => ({
      selectedProfile,
      handleProfileClick,
      handleMatchClick,
      handleCloseProfile,
    }),
    [selectedProfile, handleProfileClick, handleMatchClick, handleCloseProfile],
  );
};
