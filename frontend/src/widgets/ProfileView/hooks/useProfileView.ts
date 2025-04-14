import { useState } from 'react';

import { MatchProfile } from '../../../features/likes/types';
import { ProfileData } from '../../../shared/components/ProfileCard';

export const useProfileView = () => {
  const [selectedProfile, setSelectedProfile] = useState<ProfileData | null>(null);

  const handleProfileClick = (profile: ProfileData) => {
    // When clicking on a profile, open the ProfileView
    setSelectedProfile(profile);
  };

  const handleMatchClick = (match: MatchProfile) => {
    // Convert match to a profile format that ProfileView can use
    const matchAsProfile: ProfileData = {
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

export default useProfileView;
