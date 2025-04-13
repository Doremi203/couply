import PhotoCameraIcon from '@mui/icons-material/PhotoCamera';
import React from 'react';

import styles from './profilePhotoEdit.module.css';

interface ProfilePhotoEditProps {
  profilePhoto: string;
  onCameraClick: (isAvatar: boolean) => void;
  title?: string;
}

export const ProfilePhotoEdit: React.FC<ProfilePhotoEditProps> = ({
  profilePhoto,
  onCameraClick,
  title = 'Profile Photo',
}) => {
  return (
    <div className={styles.section}>
      <h3>{title}</h3>
      <div className={styles.imageEdit}>
        <img src={profilePhoto || '/photo1.png'} alt="Profile" className={styles.profilePic} />
        <div className={styles.editIcon} onClick={() => onCameraClick(true)}>
          <PhotoCameraIcon />
        </div>
      </div>
    </div>
  );
};
