import React from 'react';

import { ProfileData } from '../../types';
import { InfoGrid } from '../InfoGrid';
import { PhotoGallery } from '../PhotoGallery';
import { ProfileHeader } from '../ProfileHeader';
import { ProfileInfo } from '../ProfileInfo';
import { ProfileSection } from '../ProfileSection';
import { TagsList } from '../TagsList';

import styles from './profileView.module.css';

interface ProfileViewProps {
  profileData: ProfileData;
  isVerified: boolean;
  isProfileHidden: boolean;
  onEditToggle: () => void;
  onVisibilityToggle: () => void;
  onActivityClick: () => void;
  onPreviewClick: () => void;
  onVerificationRequest: () => void;
}

export const ProfileView: React.FC<ProfileViewProps> = ({
  profileData,
  isVerified,
  isProfileHidden,
  onEditToggle,
  onVisibilityToggle,
  onActivityClick,
  onPreviewClick,
  onVerificationRequest,
}) => {
  // Create info items for the basic information section
  const basicInfoItems = [
    { label: 'Name', value: profileData.name },
    { label: 'Age', value: profileData.age },
    { label: 'Gender', value: profileData.gender === 'female' ? 'Female' : 'Male' },
    { label: 'Email', value: profileData.email },
    { label: 'Phone', value: profileData.phone },
  ];

  return (
    <div className={styles.profileContent}>
      <ProfileHeader
        isProfileHidden={isProfileHidden}
        onEditToggle={onEditToggle}
        onVisibilityToggle={onVisibilityToggle}
        onActivityClick={onActivityClick}
        onPreviewClick={onPreviewClick}
      />

      <ProfileInfo
        profileData={profileData}
        isVerified={isVerified}
        onVerificationRequest={onVerificationRequest}
      />

      <ProfileSection title="Photos">
        <PhotoGallery photos={profileData.photos} />
      </ProfileSection>

      <ProfileSection title="About Me">
        <p>{profileData.about}</p>
      </ProfileSection>

      <ProfileSection title="Basic Information">
        <InfoGrid infoItems={basicInfoItems} />
      </ProfileSection>

      <ProfileSection title="Interests">
        <TagsList items={profileData.interests} />
      </ProfileSection>

      <ProfileSection title="Music">
        <TagsList items={profileData.music} />
      </ProfileSection>

      <ProfileSection title="Movies">
        <TagsList items={profileData.movies} />
      </ProfileSection>

      <ProfileSection title="Books">
        <TagsList items={profileData.books} />
      </ProfileSection>

      <ProfileSection title="Hobbies">
        <TagsList items={profileData.hobbies} />
      </ProfileSection>
    </div>
  );
};

export default ProfileView;