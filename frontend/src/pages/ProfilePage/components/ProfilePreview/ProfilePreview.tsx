import React from 'react';

import { ProfileView } from '../../../../pages/LikesPage/components/ProfileView/ProfileView';
import { ProfileData } from '../../types';

import styles from './profilePreview.module.css';

interface ProfilePreviewProps {
  profileData: ProfileData;
  onClose: () => void;
}

export const ProfilePreview: React.FC<ProfilePreviewProps> = ({
  profileData,
  onClose,
}) => {
  // Create a profile object that matches the ProfileView component's expected props
  const profile = {
    id: 1, // Dummy ID
    name: profileData.name,
    age: profileData.age,
    imageUrl: profileData.photos[0] || '/photo1.png',
    bio: profileData.about,
    location: 'Your Location', // You can add location to profileData if needed
    interests: profileData.interests,
    passion: [...profileData.interests, ...profileData.hobbies],
    photos: profileData.photos,
    lifestyle: {
      kids: "I don't have kids", // Example lifestyle data
    },
  };

  return (
    <div className={styles.previewContainer}>
      {/* Preview badge overlay */}
      <div className={styles.previewBadge}>
        <span>Preview Mode</span>
      </div>
      
      {/* Use the ProfileView component from LikesPage */}
      <ProfileView
        profile={profile}
        onClose={onClose}
        onLike={() => {}} // Empty function since we don't need like functionality in preview
      />
    </div>
  );
};

export default ProfilePreview;