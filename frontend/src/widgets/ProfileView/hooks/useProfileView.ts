import { useState } from 'react';

import { MatchProfile } from '../../../features/matches/types';

// TODO
export const useProfileView = () => {
  const [selectedProfile, setSelectedProfile] = useState(null);

  //@ts-ignore
  const handleProfileClick = (profile) => {
    // When clicking on a profile, open the ProfileView
    setSelectedProfile(profile);
  };

  const handleMatchClick = (match: MatchProfile) => {
    // Convert match to a profile format that ProfileView can use
    const matchAsProfile = {
      user: {
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
        },
      },
    };
    //@ts-ignore
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
