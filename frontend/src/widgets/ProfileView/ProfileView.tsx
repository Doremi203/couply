import VerifiedIcon from '@mui/icons-material/Verified';
import React, { useRef, useEffect } from 'react';

import { useLikeUserMutation } from '../../entities/matches';
import { BackButton } from '../../shared/components/BackButton';
import { DislikeButton } from '../../shared/components/DislikeButton';
import { LikeButton } from '../../shared/components/LikeButton';

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
}

export const ProfileView: React.FC<ProfileViewProps> = ({
  profile,
  onClose,
  onLike,
  onDislike,
  isMatchView = false,
}) => {
  const [likeUser] = useLikeUserMutation();
  const containerRef = useRef<HTMLDivElement>(null);

  // Handle the profile data - it might come directly or nested in a user property
  const profileData = profile.user || profile;

  // Get the profile photo safely
  const getProfilePhoto = () => {
    if (!profileData.photos || !profileData.photos.length) {
      return '/photo1.png';
    }

    const firstPhoto = profileData.photos[0];
    if (typeof firstPhoto === 'string') {
      return firstPhoto;
    } else if (firstPhoto && typeof firstPhoto === 'object') {
      return firstPhoto.url || '/photo1.png';
    }

    return '/photo1.png';
  };

  const commonInterests = ['Music', 'Travel', 'Photography'];

  const isCommonInterest = (interest: string) => {
    return commonInterests.includes(interest);
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

        <img
          src={getProfilePhoto()}
          alt={profileData.name}
          className={styles.profileImage}
          onError={e => {
            (e.target as HTMLImageElement).src = '/photo1.png';
          }}
        />
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

            {!isMatchView && (
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

      <ProfileInfo
        profile={profileData}
        profileDetails={profileDetails}
        isCommonInterest={isCommonInterest}
      />
    </div>
  );
};

export default ProfileView;
