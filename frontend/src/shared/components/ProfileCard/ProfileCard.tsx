import SmsOutlinedIcon from '@mui/icons-material/SmsOutlined';
import React, { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import {
  selectShowMatchModal,
  setShowMatchModal,
} from '../../../entities/matches/model/matchesSlice';
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
  const dispatch = useDispatch();
  const [messageModalOpen, setMessageModalOpen] = useState(false);

  const [getUser] = useGetUserMutation();
  const [myData, setMyData] = useState<UserData>();

  // Get the match modal state from Redux
  const show = useSelector(selectShowMatchModal);

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
    e.stopPropagation();
    e.preventDefault();

    const userId = profile.id;

    if (userId) {
      try {
        dispatch(setShowMatchModal(true));

        document.body.style.overflow = 'hidden';

        if (onLike) {
          onLike(userId);
        }
      } catch (error) {
        console.error('ProfileCard: Error liking user:', error);
        document.body.style.overflow = '';
      }
    }
  };

  const handleCardClick = (_e: React.MouseEvent) => {
    if (!messageModalOpen && !show) {
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
    document.body.style.overflow = '';
    dispatch(setShowMatchModal(false));
  };

  // Log state changes for debugging
  useEffect(() => {
    // console.log('ProfileCard state changed:', {
    //   showModal: show,
    //   profileId: profile.id,
    //   profileName: profile.name,
    //   myData: myData?.name,
    // });

    if (show) {
      document.body.style.overflow = 'hidden';
    }
  }, [show, profile.id, profile.name, myData]);

  if (show) {
    return (
      <MatchModal
        userImage={myData?.photos?.[0]?.url || ''}
        matchImage={profile.photos?.[0]?.url || ''}
        matchName={profile.name}
        onKeepSwiping={handleKeepSwiping}
        isOpen={show}
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
    </>
  );
};

export default ProfileCard;
