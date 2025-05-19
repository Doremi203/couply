import VerifiedIcon from '@mui/icons-material/Verified';
import React from 'react';

import { ProfileData } from '../../types';

import styles from './profileInfo.module.css';

interface ProfileInfoProps {
  profileData: ProfileData;
  isVerified: boolean;
  onVerificationRequest: () => void;
  onPreviewClick: () => void;
}

export const ProfileInfo: React.FC<ProfileInfoProps> = ({
  profileData,
  isVerified,
  onVerificationRequest,
  onPreviewClick,
}) => {
  const isPro = false;

  // Get profile image URL safely
  const getProfileImageUrl = () => {
    if (!profileData.photos || profileData.photos.length === 0) {
      return '/photo1.png'; // Default image
    }

    const firstPhoto = profileData.photos[0];
    if (typeof firstPhoto === 'string') {
      return firstPhoto;
    } else if (typeof firstPhoto === 'object' && firstPhoto !== null) {
      return (firstPhoto as any).url || '/photo1.png';
    }

    return '/photo1.png'; // Fallback
  };

  return (
    <div className={styles.profileInfo}>
      <div className={styles.profileImageContainer} onClick={onPreviewClick}>
        <img
          src={getProfileImageUrl()}
          alt="Profile"
          className={styles.profilePic}
          onError={e => {
            (e.target as HTMLImageElement).src = '/photo1.png';
          }}
        />
        {isPro && (
          <div className={styles.proBadge}>
            <span>PRO</span>
          </div>
        )}
        {isVerified && (
          <div className={styles.verificationBadge}>
            <VerifiedIcon />
          </div>
        )}
      </div>
      <h2 className={styles.name}>
        {profileData.name}, {profileData.age}
      </h2>
      {!isVerified && (
        <button className={styles.verifyButton} onClick={onVerificationRequest}>
          Верифицировать профиль
        </button>
      )}
    </div>
  );
};

export default ProfileInfo;
