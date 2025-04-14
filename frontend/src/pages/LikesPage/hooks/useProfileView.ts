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
      location: '',
      interests: [],
      lifestyle: {
        contact: `Telegram: ${match.telegram}`,
        social: `Instagram: ${match.instagram}`,
      },
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
