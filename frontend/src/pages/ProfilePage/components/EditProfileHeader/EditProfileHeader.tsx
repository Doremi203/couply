import KeyboardBackspaceIcon from '@mui/icons-material/KeyboardBackspace';
import React from 'react';

import styles from '../EditProfile/editProfile.module.css';

interface ProfileHeaderProps {
  onBack: () => void;
}

export const EditProfileHeader: React.FC<ProfileHeaderProps> = ({ onBack }) => {
  return (
    <div className={styles.profileHeader}>
      <div className={styles.backButton} onClick={onBack}>
        <KeyboardBackspaceIcon />
      </div>
      <div className={styles.header}>edit profile</div>
    </div>
  );
};
