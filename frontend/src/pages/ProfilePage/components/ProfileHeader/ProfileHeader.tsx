import React from 'react';

import styles from './profileHeader.module.css';

interface ProfileHeaderProps {
  isProfileHidden: boolean;
  onEditToggle: () => void;
  onVisibilityToggle: () => void;
  onActivityClick: () => void;
  onPreviewClick: () => void;
}

export const ProfileHeader: React.FC<ProfileHeaderProps> = () => {
  return (
    <div className={styles.profileHeader}>
      <div className={styles.header}>Профиль</div>
      <div className={styles.profileActions} />
    </div>
  );
};

export default ProfileHeader;
