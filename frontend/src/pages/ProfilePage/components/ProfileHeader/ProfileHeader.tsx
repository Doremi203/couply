import EditIcon from '@mui/icons-material/Edit';
import VisibilityIcon from '@mui/icons-material/Visibility';
import VisibilityOffIcon from '@mui/icons-material/VisibilityOff';
import React from 'react';

import { CustomButton } from '../../../../shared/components/CustomButton';
import { IconButton } from '../../../../shared/components/IconButton';
import { ThemeToggle } from '../../../../shared/components/ThemeToggle';

import styles from './profileHeader.module.css';

interface ProfileHeaderProps {
  isProfileHidden: boolean;
  onEditToggle: () => void;
  onVisibilityToggle: () => void;
  onActivityClick: () => void;
  onPreviewClick: () => void;
}

// @ts-nocheck

export const ProfileHeader: React.FC<ProfileHeaderProps> = ({
  isProfileHidden,
  onEditToggle,
  onVisibilityToggle,
  // onActivityClick,
  onPreviewClick,
}) => {
  return (
    <div className={styles.profileHeader}>
      <div className={styles.header}>Профиль</div>
      <div className={styles.profileActions}>
        {/* <IconButton onClick={onEditToggle} touchFriendly={true}>
          <EditIcon />
        </IconButton>
        <IconButton onClick={onVisibilityToggle}>
          {isProfileHidden ? <VisibilityOffIcon /> : <VisibilityIcon />}
        </IconButton> */}
        {/* <ThemeToggle /> */}
      </div>
      {/* <div className={styles.previewButtonContainer}>
        <CustomButton text="preview" onClick={onPreviewClick} className={styles.previewButton} />
      </div> */}
    </div>
  );
};

export default ProfileHeader;
