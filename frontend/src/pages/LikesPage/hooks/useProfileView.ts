import { useState, useCallback, useMemo } from 'react';

import { LikeProfile, MatchProfile } from '../types';

// TODO
export const useProfileView = () => {
  const [selectedProfile, setSelectedProfile] = useState<LikeProfile | null>(null);

  const handleProfileClick = useCallback((profile: LikeProfile) => {
    // Format the profile data to match what ProfileView expects
    //@ts-ignore
    const formattedProfile: LikeProfile = {
      user: {
        id: profile.id,
        name: profile.name,
        age: profile.age,
        photos: [{ url: profile.imageUrl }],
        bio: profile.bio || '',
        location: profile.location || '',
        interests: profile.interests || [],
        lifestyle: profile.lifestyle || {},
        passion: profile.passion || [],
        verified: profile.verified || false,
      },
    };
    setSelectedProfile(formattedProfile);
  }, []);

  const handleMatchClick = useCallback((match: MatchProfile) => {
    //@ts-ignore
    const matchAsProfile: LikeProfile = {
      user: {
        id: match.id,
        name: match.name,
        age: match.age,
        //@ts-ignore
        photos: match.photos,
        hasLikedYou: true,
        bio: 'This is a match! You can contact them via social media.',
        location: 'Matched User',
        interests: [],
        lifestyle: {
          contact: `Telegram: ${match.telegram}`,
        },
        passion: ['Match', 'Connection'],
        verified: false,
      },
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
