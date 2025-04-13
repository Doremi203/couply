import AddIcon from '@mui/icons-material/Add';
import CloseIcon from '@mui/icons-material/Close';
import React from 'react';

import styles from '../EditProfile/editProfile.module.css';

interface PhotoGalleryEditProps {
  photos: string[];
  onPhotoRemove: (index: number) => void;
  onAddPhotoClick: () => void;
}

export const PhotoGalleryEdit: React.FC<PhotoGalleryEditProps> = ({
  photos,
  onPhotoRemove,
  onAddPhotoClick,
}) => {
  return (
    <div className={styles.photoGalleryEdit}>
      <h3>Photos</h3>
      <div className={styles.photosGrid}>
        {photos.map((photo, index) => (
          <div key={index} className={styles.photoItemEdit}>
            <img src={photo} alt={`Photo ${index + 1}`} />
            <div className={styles.removePhotoIcon} onClick={() => onPhotoRemove(index)}>
              <CloseIcon />
            </div>
          </div>
        ))}
        <div className={styles.addPhotoItem} onClick={onAddPhotoClick}>
          <AddIcon />
        </div>
      </div>
    </div>
  );
};
