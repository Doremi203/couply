import PhotoCameraIcon from '@mui/icons-material/PhotoCamera';
import React from 'react';

import styles from '../EditProfile/editProfile.module.css';

interface ProfilePhotoEditProps {
  profilePhoto: string;
  onCameraClick: (isAvatar: boolean) => void;
}

export const ProfilePhotoEdit: React.FC<ProfilePhotoEditProps> = ({
  profilePhoto,
  onCameraClick,
}) => {
  return (
    <div className={styles.photoEditSection}>
      <h3>Profile Photo</h3>
      <div className={styles.profileImageEdit}>
        <img src={profilePhoto || '/photo1.png'} alt="Profile" className={styles.profilePic} />
        <div className={styles.photoEditIcon} onClick={() => onCameraClick(true)}>
          <PhotoCameraIcon />
        </div>
      </div>
    </div>
  );
};
