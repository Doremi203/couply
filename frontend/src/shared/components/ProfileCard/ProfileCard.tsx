import SmsOutlinedIcon from '@mui/icons-material/SmsOutlined';
import React, { useEffect, useState } from 'react';
import { Like } from '../../../entities/matches/types';
import { useGetUserMutation } from '../../../entities/user';
import { UserData } from '../../../entities/user/types';
import { MatchModal } from '../../../pages/LikesPage/components/MatchModal';
import { LikeButton } from '../LikeButton';
import { MessageModal } from '../MessageModal/MessageModal';

import styles from './profileCard.module.css';

export interface ProfileCardProps {
  profile: UserData;
  onClick?: () => void;
  onLike?: (id: string) => void;
  className?: string;
  like: Like;
}

export const ProfileCard: React.FC<ProfileCardProps> = ({
  profile,
  like,
  onClick,
  onLike,
  className,
}) => {
  const [messageModalOpen, setMessageModalOpen] = useState(false);
  const [showMatchModal, setShowMatchModal] = useState(false);
  const [isLiked, setIsLiked] = useState(false);

  const [getUser] = useGetUserMutation();

  const [myData, setMyData] = useState<UserData>();

  useEffect(() => {
    const fetchData = async () => {
      try {
        const data = await getUser({}).unwrap();
        setMyData(data.user);
      } catch (error) {
        console.error('Failed to fetch user:', error);
      }
    };

    fetchData();
  }, [getUser]);

  // Don't return early if like is undefined, just set message to undefined
  const message = like?.message;

  const handleLikeClick = async (e: React.MouseEvent) => {
    console.log('like');
    e.stopPropagation();
    e.preventDefault();

    const userId = profile.id;
    console.log('userId', userId);

    if (userId) {
      try {
        // Set local state for immediate feedback
        setIsLiked(true);
        setShowMatchModal(true);
        document.body.style.overflow = 'hidden';

        // Call the onLike function from useLikesAndMatches
        if (onLike) {
          console.log('onLike', userId);
          onLike(userId);
        }
      } catch (error) {
        console.error('Error liking user:', error);
        setIsLiked(false);
        setShowMatchModal(false);
        document.body.style.overflow = '';
      }
    }
  };

  const handleCardClick = (_e: React.MouseEvent) => {
    if (!messageModalOpen && !showMatchModal) {
      onClick?.();
    }
  };

  const handleMessageClick = (e: React.MouseEvent) => {
    e.stopPropagation();
    e.preventDefault();
    setMessageModalOpen(true);
    return false;
  };

  const handleMessageModalClose = () => {
    setMessageModalOpen(false);
  };

  const handleKeepSwiping = () => {
    setIsLiked(false);
    document.body.style.overflow = '';

    // Call onLike's handleKeepSwiping to ensure lists are reloaded
    if (onLike) {
      // This is a bit of a hack, but it ensures the lists are reloaded
      setTimeout(() => {
        onLike(profile.id);
      }, 100);
    }
  };

  console.log('State:', { isLiked, showMatchModal, myData });

  // If the match modal should be shown
  if (showMatchModal && myData) {
    return (
      <MatchModal
        //@ts-ignore
        userImage={myData.photos?.[0].url}
        matchImage={profile.photos?.[0]?.url}
        matchName={profile.name}
        onKeepSwiping={handleKeepSwiping}
      />
    );
  }

  return (
    <>
      <div className={`${styles.profileCard} ${className || ''}`} onClick={handleCardClick}>
        <div className={styles.imageCont}>
          <img src={profile.photos?.[0]?.url} alt={profile.name} className={styles.profileImage} />
          {message !== undefined && message !== '' && (
            <div
              onClick={e => e.stopPropagation()}
              style={{
                position: 'absolute',
                zIndex: 10,
                top: '10px',
                right: '10px',
              }}
            >
              <SmsOutlinedIcon
                className={styles.iconOverlay}
                sx={{
                  fontSize: 30,
                  animation: `${styles.messageNotification} 1s infinite`,
                }}
                onClick={handleMessageClick}
              />
            </div>
          )}
        </div>
        <div className={styles.profileInfo}>
          <div className={styles.profileName}>
            {profile.name}, {profile.age}
          </div>
        </div>

        <div
          onClick={e => {
            e.stopPropagation();
            e.preventDefault();
          }}
          onMouseDown={e => {
            e.stopPropagation();
            e.preventDefault();
          }}
        >
          <LikeButton
            onClick={handleLikeClick}
            className={styles.likeButton}
            likeClassName={styles.like}
          />
        </div>

        <MessageModal
          isOpen={messageModalOpen}
          onClose={handleMessageModalClose}
          message={message}
          senderName={profile.name}
        />
      </div>

      {/* Match modal is now only rendered in the conditional above */}
    </>
  );
};

export default ProfileCard;
