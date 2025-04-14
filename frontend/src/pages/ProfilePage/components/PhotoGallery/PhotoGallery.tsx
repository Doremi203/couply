import React from 'react';

import styles from './photoGallery.module.css';

interface PhotoGalleryProps {
  photos: string[];
}

export const PhotoGallery: React.FC<PhotoGalleryProps> = ({ photos }) => {
  return (
    <div className={styles.photoGallery}>
      <div className={styles.photosGrid}>
        {photos.map((photo, index) => (
          <div key={index} className={styles.photoItem}>
            <img src={photo} alt={`Photo ${index + 1}`} />
          </div>
        ))}
      </div>
    </div>
  );
};

export default PhotoGallery;
