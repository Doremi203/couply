import React from 'react';
import { useNavigate } from 'react-router-dom';

import { ProfileData } from '../../types';
import { ProfileHeader } from '../ProfileHeader';
import { ProfileInfo } from '../ProfileInfo';
import { ProfileMenu } from '../ProfileMenu';

import styles from './profile.module.css';

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

export const Profile: React.FC<ProfileViewProps> = ({
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
  //@ts-ignore
  // const basicInfoItems = [
  //   { label: 'Name', value: profileData.name },
  //   { label: 'Age', value: profileData.age },
  //   { label: 'Gender', value: profileData.gender === 'female' ? 'Female' : 'Male' },
  //   { label: 'Email', value: profileData.email },
  //   { label: 'Phone', value: profileData.phone },
  // ];

  const navigate = useNavigate();

  const handleMyStatsClick = () => {
    // Handle My Stats click
    console.log('My Stats clicked');
  };

  const handleSettingsClick = () => {
    navigate('/settings');
  };

  const handleInviteFriendClick = () => {
    // Handle Invite a friend click
    console.log('Invite a friend clicked');
  };

  const handleHelpClick = () => {
    // Handle Help click
    console.log('Help clicked');
  };

  return (
    <div>
      <ProfileHeader
        isProfileHidden={isProfileHidden}
        onEditToggle={onEditToggle}
        onVisibilityToggle={onVisibilityToggle}
        onActivityClick={onActivityClick}
        onPreviewClick={onPreviewClick}
      />

      <div className={styles.profileContent}>
        <ProfileInfo
          profileData={profileData}
          isVerified={isVerified}
          onVerificationRequest={onVerificationRequest}
          onPreviewClick={onPreviewClick}
        />

        <div className={styles.premium}>
          <div>Оформите премиум подписку</div>
          <div className={styles.text}>чтобы повысить шансы найти свою каплю</div>
          <button className={styles.premiumButton} onClick={() => navigate('/premium')}>
            Подробнее
          </button>
        </div>

        <ProfileMenu
          onEditProfileClick={onEditToggle}
          onMyStatsClick={handleMyStatsClick}
          onSettingsClick={handleSettingsClick}
          onInviteFriendClick={handleInviteFriendClick}
          onHelpClick={handleHelpClick}
        />
      </div>
    </div>
  );
};

export default Profile;
