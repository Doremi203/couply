import React, { useEffect, useState } from 'react';

import {
  goalFromApi,
  childrenFromApi,
  educationFromApi,
  alcoholFromApi,
  smokingFromApi,
  zodiacFromApi,
} from '../../../../features/filters/components/constants';
import { mapInterestsFromApiFormat } from '../../../../features/filters/helpers/mapInterestsFromApiFormat';
import { CommonInterest } from '../../../../shared/components/CommonInterest';
import { getLocationFromCoordinates } from '../../../../shared/lib/location';
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
    latitude?: number;
    longitude?: number;
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
  isCommonInterest: (interest: string) => boolean;
}

export const ProfileInfo: React.FC<ProfileInfoProps> = ({
  profile,
  profileDetails: _profileDetails,
  isCommonInterest: _isCommonInterest,
}) => {
  const [location, setLocation] = useState<string>('Загрузка...');

  useEffect(() => {
    if (profile?.latitude && profile?.longitude) {
      getLocationFromCoordinates(profile.latitude, profile.longitude)
        .then(setLocation)
        .catch(() => setLocation('Ошибка определения локации'));
    } else if (profile?.location) {
      setLocation(profile.location);
    } else {
      setLocation('Локация не указана');
    }
  }, [profile?.latitude, profile?.longitude, profile?.location]);

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
          key: 'Дети',
          value: childrenFromApi[profile.children as keyof typeof childrenFromApi],
        }
      : null,
    profile.education && educationFromApi[profile.education as keyof typeof educationFromApi]
      ? {
          key: 'Образование',
          value: educationFromApi[profile.education as keyof typeof educationFromApi],
        }
      : null,
    profile.alcohol && alcoholFromApi[profile.alcohol as keyof typeof alcoholFromApi]
      ? { key: 'Алкоголь', value: alcoholFromApi[profile.alcohol as keyof typeof alcoholFromApi] }
      : null,
    profile.smoking && smokingFromApi[profile.smoking as keyof typeof smokingFromApi]
      ? { key: 'Курение', value: smokingFromApi[profile.smoking as keyof typeof smokingFromApi] }
      : null,
    profile.zodiac && zodiacFromApi[profile.zodiac as keyof typeof zodiacFromApi]
      ? { key: 'Знак зодиака', value: zodiacFromApi[profile.zodiac as keyof typeof zodiacFromApi] }
      : null,
  ].filter(Boolean) as { key: string; value: string }[];

  const hasBio = !!profile.bio;
  const hasBasicInfo = basicInfoFields.length > 0;

  const hasPhotos = profile.photos && profile.photos.length > 0;
  const hasGoal = !!profile.goal && !!goalFromApi[profile.goal as keyof typeof goalFromApi];

  const interest = mapInterestsFromApiFormat(
    (profile.interest as unknown as Record<string, string[]>) || {},
  );

  const hasInterest = interest.length > 0;

  return (
    <div className={styles.profileInfo}>
      <div className={styles.profileInfoHandle}>
        <div className={styles.profileInfoHandleBar} />
      </div>

      <div className={styles.profileNameContainer}>
        <div>
          <h2 className={styles.profileName}>{profile.name}</h2>
          <p className={styles.profileAge}>
            {profile.age} | {location}
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
              <React.Fragment key={index}>
                <h4>{field.key}</h4>
                <p>{field.value}</p>
              </React.Fragment>
            ))}
          </section>
        )}

        {hasInterest && (
          <section className={styles.infoSection}>
            <h3>Интересы</h3>
            <div className={styles.interestTags}>
              {interest.map((interest: string, index: number) => (
                <CommonInterest key={index} text={interest} className={styles.interestTag} />
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
