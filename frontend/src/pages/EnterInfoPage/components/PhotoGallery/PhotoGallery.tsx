import AddIcon from '@mui/icons-material/Add';
import CloseIcon from '@mui/icons-material/Close';
import React from 'react';

import styles from './photoGallery.module.css';

interface FixedPhotoGalleryProps {
  photos: string[];
  onPhotoRemove: (index: number) => void;
  onAddPhotoClick: () => void;
  title?: string;
}

export const FixedPhotoGallery: React.FC<FixedPhotoGalleryProps> = ({
  photos,
  onPhotoRemove,
  onAddPhotoClick,
  title = 'Photos',
}) => {
  // Создаем массив из 6 элементов
  const items = Array.from({ length: 6 }).map((_, index) => ({
    photo: photos[index] || null,
    index,
  }));

  return (
    <div className={styles.gallery}>
      <h3>{title}</h3>
      <div className={styles.grid}>
        {items.map(({ photo, index }) => (
          <div key={index} className={styles.photoItem}>
            {photo ? (
              <>
                <img src={photo} alt={`Photo ${index + 1}`} />
                <div className={styles.removeIcon} onClick={() => onPhotoRemove(index)}>
                  <CloseIcon />
                </div>
              </>
            ) : (
              photos.length < 6 && (
                <div
                  className={styles.addItem}
                  onClick={onAddPhotoClick}
                  data-disabled={photos.length >= 6}
                >
                  <AddIcon />
                </div>
              )
            )}
          </div>
        ))}
      </div>
    </div>
  );
};
