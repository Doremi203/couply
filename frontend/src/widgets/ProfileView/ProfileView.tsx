import VerifiedIcon from '@mui/icons-material/Verified';
import React, { useRef, useEffect, useState } from 'react';

import { useLikeUserMutation } from '../../entities/matches';
import { BackButton } from '../../shared/components/BackButton';
import { DislikeButton } from '../../shared/components/DislikeButton';
import { LikeButton } from '../../shared/components/LikeButton';
import { PremiumModal } from '../../widgets/PremiumModal';

import ProfileInfo from './components/ProfileInfo';
import styles from './profileView.module.css';

interface ProfileViewProps {
  profile: {
    id?: number;
    name: string;
    age: number;
    imageUrl?: string;
    hasLikedYou?: boolean;
    verified?: boolean;
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
    user?: any;
  };
  onClose: () => void;
  onLike?: (id: number) => void;
  onDislike?: (id: number) => void;
  isMatchView?: boolean;
  isProfile?: boolean;
}

export const ProfileView: React.FC<ProfileViewProps> = ({
  profile,
  onClose,
  onLike,
  onDislike,
  isMatchView = false,
  isProfile = false,
}) => {
  const [likeUser] = useLikeUserMutation();
  const containerRef = useRef<HTMLDivElement>(null);
  const [isPremiumModalOpen, setIsPremiumModalOpen] = useState(false);
  const [currentPhotoIndex, setCurrentPhotoIndex] = useState(0);

  const profileData = profile.user || profile;

  const getProfilePhoto = (index: number) => {
    const photo = profileData.photos[index];
    if (typeof photo === 'string') {
      return photo;
    } else if (photo && typeof photo === 'object') {
      return photo.url;
    }
    return '';
  };

  const handlePhotoClick = (e: React.MouseEvent<HTMLDivElement>) => {
    const rect = e.currentTarget.getBoundingClientRect();
    const clickPosition = e.clientX - rect.left;
    const width = rect.width;

    if (clickPosition > width * 0.75) {
      handleNextPhoto();
    } else if (clickPosition < width * 0.25) {
      handlePrevPhoto();
    }
  };

  const handleNextPhoto = () => {
    if (!profileData.photos || profileData.photos.length <= 1) return;
    setCurrentPhotoIndex(prevIndex => (prevIndex + 1) % profileData.photos.length);
  };

  const handlePrevPhoto = () => {
    if (!profileData.photos || profileData.photos.length <= 1) return;
    setCurrentPhotoIndex(
      prevIndex => (prevIndex - 1 + profileData.photos.length) % profileData.photos.length,
    );
  };

  const handleLike = () => {
    const userId = profileData.id;
    if (userId && onLike) {
      likeUser({
        targetUserId: userId,
        message: '',
      });
      onLike(userId);
      onClose();
    }
  };

  const handleDislike = () => {
    //@ts-ignore
    onDislike(profile.id);
    onClose();
  };

  useEffect(() => {
    document.body.style.overflow = 'hidden';

    return () => {
      document.body.style.overflow = 'auto';
    };
  }, []);

  const profileDetails = {
    bio: profileData.bio || '',
    location: profileData.location || '',
    lifestyle: profileData.lifestyle || {},
    passion: profileData.passion || profileData.interests || [],
    photos: profileData.photos || [],
  };

  return (
    <div
      className={styles.profileViewContainer}
      ref={containerRef}
      onClick={e => e.stopPropagation()}
    >
      <div className={styles.profileImageContainer}>
        <BackButton
          onClose={e => {
            e?.stopPropagation?.();
            onClose();
            setTimeout(() => {
              document.body.style.overflow = 'auto';
            }, 100);
          }}
        />

        <div className={styles.profileImageWrapper} onClick={handlePhotoClick}>
          <img
            src={getProfilePhoto(currentPhotoIndex)}
            alt={profileData.name}
            className={styles.profileImage}
            onError={e => {
              (e.target as HTMLImageElement).src = '/photo1.png';
            }}
          />
          {profileData.photos && profileData.photos.length > 1 && (
            <div className={styles.photoCounter}>
              {currentPhotoIndex + 1}/{profileData.photos.length}
            </div>
          )}
        </div>
        <div className={styles.profileGradient} />

        <div className={styles.photoContent}>
          <div className={styles.nameAndButtons}>
            <h2 className={styles.photoName}>
              {profileData.name}
              {profileData.verified && (
                <div className={styles.verifiedBadge}>
                  <VerifiedIcon />
                </div>
              )}
            </h2>

            {!isMatchView && !isProfile && (
              <div className={styles.buttons}>
                <div onClick={e => e.stopPropagation()}>
                  <DislikeButton onClick={handleDislike} className={styles.dislikeButton} />
                </div>
                <div onClick={e => e.stopPropagation()}>
                  <LikeButton
                    onClick={handleLike}
                    className={styles.likeButton}
                    likeClassName={styles.like}
                  />
                </div>
              </div>
            )}
          </div>

          <p className={styles.photoInfo}>
            {profileData.age}
            {profileData.location ? `, ${profileData.location}` : ''}
          </p>
        </div>
      </div>

      {/**@ts-ignore */}
      <ProfileInfo profile={profileData} profileDetails={profileDetails} isCommonInterest={[]} />

      <PremiumModal isOpen={isPremiumModalOpen} onClose={() => setIsPremiumModalOpen(false)} />
    </div>
  );
};

export default ProfileView;
