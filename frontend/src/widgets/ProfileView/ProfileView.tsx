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
    id: number;
    name: string;
    age: number;
    imageUrl: string;
    hasLikedYou?: boolean;
    verified?: boolean;
    bio?: string;
    location?: string;
    interests?: string[];
    lifestyle?: { [key: string]: string };
    passion?: string[];
    photos?: string[];
  };
  onClose: () => void;
  onLike: (id: number) => void;
}

//@ts-ignore
export const ProfileView: React.FC<ProfileViewProps> = ({ profile, onClose, onLike }) => {
  const [likeUser] = useLikeUserMutation();

  const [menuPosition, setMenuPosition] = useState<'collapsed' | 'expanded'>('expanded');

  const profileInfoRef = useRef<HTMLDivElement>(null);
  const containerRef = useRef<HTMLDivElement>(null);

  //TODO delete
  const samplePhotos = [
    profile.imageUrl,
    'woman1.jpg',
    'man1.jpg',
    'photo1.png',
    'woman1.jpg',
    'man1.jpg',
  ];

  //TODO delete
  const profileDetails = {
    bio: profile.user.bio || 'Hello, I am a fashion designer based in Florida.',
    location: profile.user.location || 'Miami Beach, Florida',
    lifestyle: profile.user.lifestyle || {
      kids: "I don't have kids",
    },
    passion: profile.user.passion ||
      profile.user.interests || [
        'Music',
        'Travel',
        'Tea',
        'Photography',
        'Fashion',
        'House Parties',
      ],
    photos: profile.photos || samplePhotos,
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
    likeUser({ targetUserId: profile.user.id });
    onLike(profile.id);
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

        <img src={profile.imageUrl} alt={profile.name} className={styles.profileImage} />
        <div className={styles.profileGradient} />

        <div className={styles.photoContent}>
          <div className={styles.nameAndButtons}>
            <h2 className={styles.photoName}>
              {profile.user.name}
              {profile.user.verified && (
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
            {profile.user.age} | {profileDetails.location}
          </p>

          <div className={styles.photoTags}>
            {profileDetails.passion.slice(0, 5).map((interest, index) => (
              <span key={index} className={styles.photoTag}>
                {interest}
              </span>
            ))}
          </div>

          <div
            className={`${styles.photoScrollIndicator} ${
              menuPosition === 'collapsed' ? styles.showScrollBack : ''
            }`}
            onClick={handleToggleClick}
          >
            {menuPosition === 'collapsed' ? 'Scroll back to photo' : 'Scroll up to view details'}
            <svg
              width="16"
              height="16"
              viewBox="0 0 24 24"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
              style={{
                transform: menuPosition === 'collapsed' ? 'rotate(180deg)' : 'none',
              }}
            >
              <path
                d="M7 10L12 15L17 10"
                stroke="white"
                strokeWidth="2"
                strokeLinecap="round"
                strokeLinejoin="round"
              />
            </svg>
          </div>
        </div>
      </div>

      <ProfileInfo
        profile={profile}
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
