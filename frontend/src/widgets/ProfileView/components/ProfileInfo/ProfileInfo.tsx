import React from 'react';

import {
  goalFromApi,
  childrenFromApi,
  educationFromApi,
  alcoholFromApi,
  smokingFromApi,
  zodiacFromApi,
} from '../../../../features/filters/components/constants';
import { CommonInterest } from '../../../../shared/components/CommonInterest';
import styles from '../../profileView.module.css';

interface ProfileInfoProps {
  profile: {
    id?: number;
    name: string;
    age: number;
    imageUrl?: string;
    hasLikedYou?: boolean;
    bio?: string;
    location?: string;
    interests?: string[];
    interest?: string[];
    lifestyle?: { [key: string]: string };
    passion?: string[];
    photos?: (string | { url: string })[];
    children?: string;
    education?: string;
    alcohol?: string;
    smoking?: string;
    zodiac?: string;
    goal?: string;
  };
  profileDetails: {
    bio: string;
    location: string;
    lifestyle: { [key: string]: string };
    passion: string[];
    photos: (string | { url: string })[];
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
  profileDetails: _profileDetails,
  menuPosition,
  handleToggleClick,
  handleTouchStart,
  handleTouchMove,
  handleTouchEnd,
  isCommonInterest,
  profileInfoRef,
}) => {
  // Helper function to get photo URL whether it's a string or object
  const getPhotoUrl = (photo: string | { url: string }): string => {
    if (typeof photo === 'string') {
      return photo;
    } else if (typeof photo === 'object' && photo !== null) {
      return photo.url || '';
    }
    return '';
  };

  const basicInfoFields = [
    profile.children && childrenFromApi[profile.children as keyof typeof childrenFromApi]
      ? {
          key: 'children',
          value: childrenFromApi[profile.children as keyof typeof childrenFromApi],
        }
      : null,
    profile.education && educationFromApi[profile.education as keyof typeof educationFromApi]
      ? {
          key: 'education',
          value: educationFromApi[profile.education as keyof typeof educationFromApi],
        }
      : null,
    profile.alcohol && alcoholFromApi[profile.alcohol as keyof typeof alcoholFromApi]
      ? { key: 'alcohol', value: alcoholFromApi[profile.alcohol as keyof typeof alcoholFromApi] }
      : null,
    profile.smoking && smokingFromApi[profile.smoking as keyof typeof smokingFromApi]
      ? { key: 'smoking', value: smokingFromApi[profile.smoking as keyof typeof smokingFromApi] }
      : null,
    profile.zodiac && zodiacFromApi[profile.zodiac as keyof typeof zodiacFromApi]
      ? { key: 'zodiac', value: zodiacFromApi[profile.zodiac as keyof typeof zodiacFromApi] }
      : null,
  ].filter(Boolean) as { key: string; value: string }[];

  console.log(basicInfoFields);

  // Check if sections have content
  const hasBio = !!profile.bio;
  const hasBasicInfo = basicInfoFields.length > 0;
  const hasInterests = profile.interest && profile.interest.length > 0;
  const hasPhotos = profile.photos && profile.photos.length > 0;
  const hasGoal = !!profile.goal && !!goalFromApi[profile.goal as keyof typeof goalFromApi];

  console.log('INTERST', profile.interest);

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
            {profile.age} | {profile.location}
          </p>
        </div>
      </div>

      <div className={styles.detailedInfo}>
        {hasBio && (
          <section className={styles.infoSection}>
            <h3>Био</h3>
            <p>{profile.bio}</p>
          </section>
        )}

        {hasBasicInfo && (
          <section className={styles.infoSection}>
            <h3>Основное</h3>
            {basicInfoFields.map((field, index) => (
              <p key={index}>{field.value}</p>
            ))}
          </section>
        )}

        {hasInterests && (
          <section className={styles.infoSection}>
            <h3>Интересы</h3>
            <div className={styles.interestTags}>
              {profile.interest?.map((interest: string, index: number) => (
                <CommonInterest
                  key={index}
                  text={interest}
                  isCommon={isCommonInterest(interest)}
                  className={styles.interestTag}
                />
              ))}
            </div>
          </section>
        )}

        {hasGoal && profile.goal && (
          <section className={styles.infoSection}>
            <h3>Цель</h3>
            <p>{goalFromApi[profile.goal as keyof typeof goalFromApi]}</p>
          </section>
        )}

        {hasPhotos && profile.photos && (
          <section className={styles.infoSection}>
            <h3>Фото</h3>
            <div className={styles.photosGrid}>
              {profile.photos.map((photo, index) => (
                <div key={index} className={styles.photoItem}>
                  <img
                    src={getPhotoUrl(photo)}
                    alt={`Photo ${index + 1}`}
                    onError={e => {
                      (e.target as HTMLImageElement).src = '/photo1.png';
                    }}
                  />
                </div>
              ))}
            </div>
          </section>
        )}

        <section className={styles.infoSection}>
          <div className={styles.empty} />
        </section>
      </div>
    </div>
  );
};

export default ProfileInfo;
