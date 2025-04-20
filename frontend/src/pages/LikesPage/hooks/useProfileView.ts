import { useState } from 'react';

import { LikeProfile, MatchProfile } from '../types';

export const useProfileView = () => {
  const [selectedProfile, setSelectedProfile] = useState<LikeProfile | null>(null);

  const handleProfileClick = (profile: LikeProfile) => {
    // When clicking on a profile, open the ProfileView
    setSelectedProfile(profile);
  };

  const handleMatchClick = (match: MatchProfile) => {
    // Convert match to a profile format that ProfileView can use
    const matchAsProfile = {
      id: match.id,
      name: match.name,
      age: match.age,
      imageUrl: match.imageUrl,
      hasLikedYou: true,
      bio: 'This is a match! You can contact them via social media.',
      location: 'Matched User',
      interests: ['Match'],
      lifestyle: {
        contact: `Telegram: ${match.telegram}`,
        social: `Instagram: ${match.instagram}`,
      },
      // Add photos property to ensure the ProfileView component can display photos
      photos: [match.imageUrl, match.imageUrl],
      // Add passion property to ensure the ProfileView component can display interests
      passion: ['Match', 'Connection'],
    };
    setSelectedProfile(matchAsProfile);
  };

  const handleCloseProfile = () => {
    setSelectedProfile(null);
  };

  return {
    selectedProfile,
    handleProfileClick,
    handleMatchClick,
    handleCloseProfile,
  };
};
