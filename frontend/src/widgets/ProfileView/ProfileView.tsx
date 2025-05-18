import VerifiedIcon from '@mui/icons-material/Verified';
import React, { useState, useRef, useEffect } from 'react';

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
}

export const ProfileView: React.FC<ProfileViewProps> = ({ profile, onClose, onLike }) => {
  const [likeUser] = useLikeUserMutation();
  const [menuPosition, setMenuPosition] = useState<'collapsed' | 'expanded'>('expanded');
  const profileInfoRef = useRef<HTMLDivElement>(null);
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

  const [touchStart, setTouchStart] = useState<number | null>(null);
  const [touchEnd, setTouchEnd] = useState<number | null>(null);

  const toggleMenuPosition = () => {
    setMenuPosition(menuPosition === 'collapsed' ? 'expanded' : 'collapsed');
  };

  const handleLike = () => {
    const userId = profileData.id;
    if (userId && onLike) {
      likeUser({
        targetUserId: userId,
        message: '', // Add empty message to satisfy LikeRequest interface
      });
      onLike(userId);
    }
  };

  useEffect(() => {
    const handleScroll = () => {
      if (containerRef.current) {
        const scrollTop = containerRef.current.scrollTop;

        if (menuPosition === 'expanded' && scrollTop > 50) {
          setMenuPosition('collapsed');
        }

        if (menuPosition === 'collapsed' && scrollTop < 20) {
          setMenuPosition('expanded');
        }
      }
    };

    const container = containerRef.current;
    if (container) {
      container.addEventListener('scroll', handleScroll);
    }

    return () => {
      if (container) {
        container.removeEventListener('scroll', handleScroll);
      }
    };
  }, [menuPosition]);

  const handleToggleClick = () => {
    toggleMenuPosition();

    if (menuPosition === 'collapsed') {
      setTimeout(() => {
        if (containerRef.current) {
          containerRef.current.scrollTo({
            top: containerRef.current.scrollHeight,
            behavior: 'smooth',
          });
        }
      }, 100);
    } else {
      setTimeout(() => {
        if (containerRef.current) {
          containerRef.current.scrollTo({
            top: 0,
            behavior: 'smooth',
          });
        }
      }, 100);
    }
  };

  const handleTouchStart = (e: React.TouchEvent) => {
    setTouchStart(e.targetTouches[0].clientY);
    setTouchEnd(null);
  };

  const handleTouchMove = (e: React.TouchEvent) => {
    setTouchEnd(e.targetTouches[0].clientY);
  };

  const handleTouchEnd = () => {
    if (!touchStart || !touchEnd) return;

    const distance = touchStart - touchEnd;
    const isDownSwipe = distance < -30;
    const isUpSwipe = distance > 30;

    if (isDownSwipe && menuPosition === 'collapsed') {
      setMenuPosition('expanded');

      setTimeout(() => {
        if (containerRef.current) {
          containerRef.current.scrollTo({
            top: containerRef.current.scrollHeight,
            behavior: 'smooth',
          });
        }
      }, 50);
    } else if (isUpSwipe && menuPosition === 'expanded') {
      setMenuPosition('collapsed');

      setTimeout(() => {
        if (containerRef.current) {
          containerRef.current.scrollTo({
            top: 0,
            behavior: 'smooth',
          });
        }
      }, 50);
    }

    setTouchStart(null);
    setTouchEnd(null);
  };

  // Prepare profile details for ProfileInfo component
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
      <div
        className={`${styles.profileImageContainer} ${
          menuPosition === 'expanded' ? styles.expanded : ''
        }`}
      >
        <BackButton
          onClose={() => {
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

            <div className={styles.buttons}>
              <div onClick={e => e.stopPropagation()}>
                <DislikeButton
                  onClick={() => {
                    onClose();
                  }}
                  className={styles.dislikeButton}
                />
              </div>
              <div onClick={e => e.stopPropagation()}>
                <LikeButton
                  onClick={handleLike}
                  className={styles.likeButton}
                  likeClassName={styles.like}
                />
              </div>
            </div>
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
        menuPosition={menuPosition}
        handleToggleClick={handleToggleClick}
        handleTouchStart={handleTouchStart}
        handleTouchMove={handleTouchMove}
        handleTouchEnd={handleTouchEnd}
        isCommonInterest={isCommonInterest}
        profileInfoRef={profileInfoRef}
      />
    </div>
  );
};

export default ProfileView;
