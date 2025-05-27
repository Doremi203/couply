import VerifiedIcon from '@mui/icons-material/Verified';
import React from 'react';
import { useNavigate } from 'react-router-dom';

import { ProfileData } from '../../types';

import styles from './profileInfo.module.css';

interface ProfileInfoProps {
  profileData: ProfileData;
  isVerified: boolean;
  _onVerificationRequest: () => void;
  onPreviewClick: () => void;
  isPremium: boolean;
}

export const ProfileInfo: React.FC<ProfileInfoProps> = ({
  profileData,
  isVerified,
  // _onVerificationRequest,
  onPreviewClick,
  isPremium,
}) => {
  const navigate = useNavigate();

  const getProfileImageUrl = () => {
    if (!profileData.photos || profileData.photos.length === 0) {
      return '/photo1.png';
    }

    const firstPhoto = profileData.photos[0];
    if (typeof firstPhoto === 'string') {
      return firstPhoto;
    } else if (typeof firstPhoto === 'object' && firstPhoto !== null) {
      return (firstPhoto as any).url || '/photo1.png';
    }

    return '/photo1.png';
  };

  const handleVerifyClick = () => {
    navigate('/verification');
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
        {isPremium && (
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
        <>
          <button className={styles.verifyButton} onClick={handleVerifyClick}>
            Верифицировать профиль
          </button>
        </>
      )}
    </div>
  );
};

export default ProfileInfo;
