import React from 'react';

import { CommonInterest } from '../../../../shared/components/CommonInterest';
import styles from '../../profileView.module.css';

interface ProfileInfoProps {
  profile: {
    id: number;
    name: string;
    age: number;
    imageUrl: string;
    hasLikedYou?: boolean;
    bio?: string;
    location?: string;
    interests?: string[];
    lifestyle?: { [key: string]: string };
    passion?: string[];
    photos?: string[];
  };
  profileDetails: {
    bio: string;
    location: string;
    lifestyle: { [key: string]: string };
    passion: string[];
    photos: string[];
  };
  menuPosition: 'collapsed' | 'expanded';
  handleToggleClick: () => void;
  handleTouchStart: (e: React.TouchEvent) => void;
  handleTouchMove: (e: React.TouchEvent) => void;
  handleTouchEnd: () => void;
  isCommonInterest: (interest: string) => boolean;
  profileInfoRef: React.RefObject<HTMLDivElement>;
}

export const ProfileInfo: React.FC<ProfileInfoProps> = ({
  profile,
  profileDetails,
  menuPosition,
  handleToggleClick,
  handleTouchStart,
  handleTouchMove,
  handleTouchEnd,
  isCommonInterest,
  profileInfoRef,
}) => {
  return (
    <div
      ref={profileInfoRef}
      className={`${styles.profileInfo} ${
        menuPosition === 'expanded' ? styles.expanded : styles.collapsed
      }`}
      onTouchStart={handleTouchStart}
      onTouchMove={handleTouchMove}
      onTouchEnd={handleTouchEnd}
      onClick={handleToggleClick}
    >
      <div className={styles.profileNameContainer}>
        <div>
          <h2 className={styles.profileName}>{profile.name}</h2>
          <p className={styles.profileAge}>
            {profile.age} | {profileDetails.location}
          </p>
        </div>
      </div>

      <div className={styles.interestTags}>
        {profileDetails.passion.slice(0, 5).map((interest, index) => (
          <CommonInterest
            key={index}
            text={interest}
            isCommon={isCommonInterest(interest)}
            className={styles.interestTag}
          />
        ))}
      </div>

      <div className={styles.detailedInfo}>
        <section className={styles.infoSection}>
          <h3>Био</h3>
          <p>{profileDetails.bio}</p>
        </section>

        <section className={styles.infoSection}>
          <h3>Основное</h3>
          {Object.values(profileDetails.lifestyle).map((value, index) => (
            <p key={index}>{value}</p>
          ))}
        </section>

        <section className={styles.infoSection}>
          <h3>Интересы</h3>
          <div className={styles.interestTags}>
            {profileDetails.passion.slice(0, 5).map((interest, index) => (
              <CommonInterest
                key={index}
                text={interest}
                isCommon={isCommonInterest(interest)}
                className={styles.interestTag}
              />
            ))}
          </div>
        </section>

        <section className={styles.infoSection}>
          <h3>Фото</h3>
          <div className={styles.photosGrid}>
            {profileDetails.photos.map((photo, index) => (
              <div key={index} className={styles.photoItem}>
                <img src={photo} alt={`Photo ${index + 1}`} />
              </div>
            ))}
          </div>
        </section>

        <section className={styles.infoSection}>
          <div className={styles.empty} />
        </section>
      </div>
    </div>
  );
};

export default ProfileInfo;
