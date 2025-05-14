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
  const navigate = useNavigate();

  const handleSettingsClick = () => {
    navigate('/settings');
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

        <ProfileMenu onEditProfileClick={onEditToggle} onSettingsClick={handleSettingsClick} />
      </div>
    </div>
  );
};

export default Profile;
