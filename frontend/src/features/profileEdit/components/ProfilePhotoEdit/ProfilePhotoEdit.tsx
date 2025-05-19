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
}) => {
  // Default profile photo if none is provided
  const defaultPhoto = '/photo1.png';

  return (
    <div className={styles.section}>
      <div className={styles.imageEdit}>
        <img
          src={profilePhoto || defaultPhoto}
          alt="Profile"
          className={styles.profilePic}
          onError={e => {
            // Fallback to default if image fails to load
            (e.target as HTMLImageElement).src = defaultPhoto;
          }}
        />
        <div className={styles.editIcon} onClick={() => onCameraClick(true)}>
          <PhotoCameraIcon />
        </div>
      </div>
    </div>
  );
};
