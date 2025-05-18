import AddIcon from '@mui/icons-material/Add';
import CloseIcon from '@mui/icons-material/Close';
import React from 'react';

import styles from './photoGalleryEdit.module.css';

interface PhotoGalleryEditProps {
  photos: (string | { url: string })[];
  onPhotoRemove: (index: number) => void;
  onAddPhotoClick: () => void;
  title?: string;
}

export const PhotoGalleryEdit: React.FC<PhotoGalleryEditProps> = ({
  photos,
  onPhotoRemove,
  onAddPhotoClick,
  title = 'Photos',
}) => {
  // Handle photo URL whether it's a string or an object with url property
  const getPhotoUrl = (photo: string | { url: string }): string => {
    return typeof photo === 'string' ? photo : photo.url;
  };

  // Maximum of 6 photos allowed
  const MAX_PHOTOS = 6;

  return (
    <div className={styles.gallery}>
      <h3>{title}</h3>
      <div className={styles.grid}>
        {photos.map((photo, index) => (
          <div key={index} className={styles.photoItem}>
            <img src={getPhotoUrl(photo)} alt={`Photo ${index + 1}`} />
            <div className={styles.removeIcon} onClick={() => onPhotoRemove(index)}>
              <CloseIcon />
            </div>
          </div>
        ))}
        {photos.length < MAX_PHOTOS && (
          <div className={styles.addItem} onClick={onAddPhotoClick}>
            <AddIcon />
          </div>
        )}
      </div>
    </div>
  );
};
