import VerifiedIcon from '@mui/icons-material/Verified';
import React from 'react';

import { CustomButton } from '../../../../shared/components/CustomButton';
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
  const isPro = true;
  return (
    <div className={styles.profileInfo}>
      <div className={styles.profileImageContainer} onClick={onPreviewClick}>
        <img
          src={profileData.photos[0] || '/photo1.png'}
          alt="Profile"
          className={styles.profilePic}
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
        // <CustomButton
        //   text="Верифицировать профиль"
        //   onClick={onVerificationRequest}
        //   className={styles.verifyButton}
        // />
        <button className={styles.verifyButton} onClick={onVerificationRequest}>
          Верифицировать профиль
        </button>
      )}
    </div>
  );
};

export default ProfileInfo;
