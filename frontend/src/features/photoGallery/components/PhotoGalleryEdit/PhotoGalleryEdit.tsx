import AddIcon from '@mui/icons-material/Add';
import CloseIcon from '@mui/icons-material/Close';
import React from 'react';

import styles from './photoGalleryEdit.module.css';

interface PhotoGalleryEditProps {
  photos: string[];
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
  return (
    <div className={styles.gallery}>
      <h3>{title}</h3>
      <div className={styles.grid}>
        {photos.map((photo, index) => (
          <div key={index} className={styles.photoItem}>
            <img src={photo} alt={`Photo ${index + 1}`} />
            <div className={styles.removeIcon} onClick={() => onPhotoRemove(index)}>
              <CloseIcon />
            </div>
          </div>
        ))}
        <div className={styles.addItem} onClick={onAddPhotoClick}>
          <AddIcon />
        </div>
      </div>
    </div>
  );
};
