import React from 'react';

import { LikeButton } from '../LikeButton';

import styles from './profileCard.module.css';

export interface ProfileData {
  user: {
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
}

export interface ProfileCardProps {
  profile: ProfileData;
  onClick?: () => void;
  onLike?: (id: number) => void;
  className?: string;
}

export const ProfileCard: React.FC<ProfileCardProps> = ({
  profile,
  onClick,
  onLike,
  className,
}) => {
  // Create a wrapper function to handle the like button click
  const handleLikeClick = () => {
    if (onLike) {
      onLike(profile.user.id);
    }
  };

  // Handle card click with event propagation control
  const handleCardClick = () => {
    if (onClick) {
      onClick();
    }
  };

  // console.log(profile.user.photos[0].url);

  return (
    <div className={`${styles.profileCard} ${className || ''}`} onClick={handleCardClick}>
      <img
        // src={profile.user.photos[0].url}
        alt={profile.user.name}
        className={styles.profileImage}
      />
      <div className={styles.profileInfo}>
        <div className={styles.profileName}>
          {profile.user.name}, {profile.user.age}
        </div>
      </div>
      {onLike && (
        <div onClick={e => e.stopPropagation()}>
          <LikeButton
            onClick={handleLikeClick}
            className={styles.likeButton}
            likeClassName={styles.like}
          />
        </div>
      )}
    </div>
  );
};

export default ProfileCard;
