// @ts-nocheck
import SmsOutlinedIcon from '@mui/icons-material/SmsOutlined';
import React, { useState } from 'react';

import { useLikeUserMutation } from '../../../entities/matches';
import { MatchModal } from '../../../pages/LikesPage/components/MatchModal';
import { PremiumModal } from '../../../widgets/PremiumModal';
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

// export const ProfileCard: React.FC<ProfileCardProps> = ({
//   profile,
//   onClick,
//   onLike,
//   className,
// }) => {
//   const [premiumOpen, setPremiumOpen] = useState(false);

//   const handleLikeClick = () => {
//     if (onLike) {
//       onLike(profile.id);
//     }
//   };

//   const handleCardClick = () => {
//     if (onClick) {
//       onClick();
//     }
//   };

//   const handleMessageClick = () => {
//     setPremiumOpen(true);
//   };
//   //TODO вернуть profile.user
//   //TODO photo
//   return (
//     <div className={`${styles.profileCard} ${className || ''}`} onClick={handleCardClick}>
//       <div className={styles.imageCont}>
//         <img src={profile.imageUrl} alt={profile.name} className={styles.profileImage} />
//         <SmsOutlinedIcon
//           className={styles.iconOverlay}
//           sx={{
//             fontSize: 30,
//             animation: `${styles.messageNotification} 1s infinite`,
//           }}
//           onClick={handleMessageClick}
//         />
//       </div>
//       <div className={styles.profileInfo}>
//         <div className={styles.profileName}>
//           {profile.name}, {profile.age}
//         </div>
//       </div>
//       {onLike && (
//         <div onClick={e => e.stopPropagation()}>
//           <LikeButton
//             onClick={handleLikeClick}
//             className={styles.likeButton}
//             likeClassName={styles.like}
//           />
//         </div>
//       )}
//       <PremiumModal isOpen={premiumOpen} onClose={() => setPremiumOpen(false)} />
//     </div>
//   );
// };

// export default ProfileCard;

export const ProfileCard: React.FC<ProfileCardProps> = ({
  profile,
  onClick,
  onLike,
  className,
}) => {
  const [premiumOpen, setPremiumOpen] = useState(false);
  const [showMatchModal, setShowMatchModal] = useState(false);
  const [isLiked, setIsLiked] = useState(false);
  const [likeUser] = useLikeUserMutation();

  console.log('Profile:', profile);

  const handleLikeClick = async (e: React.MouseEvent) => {
    // Important: Stop propagation and prevent default behavior
    e.stopPropagation();
    e.preventDefault();

    console.log('Like button clicked!');

    // Сразу отмечаем карточку как лайкнутую для мгновенной визуальной обратной связи
    setIsLiked(true);

    let userId = null;
    if (profile.id) {
      userId = profile.id;
    } else if (profile.user && profile.user.id) {
      userId = profile.user.id;
    }

    if (userId) {
      try {
        // Вызываем API likeUser непосредственно здесь
        const response = await likeUser({
          targetUserId: userId,
          message: '', // Добавляем пустое сообщение для соответствия интерфейсу LikeRequest
        }).unwrap();

        console.log('API response:', response);

        // Для тестирования - всегда показываем модальное окно по клику на лайк
        // В реальной работе раскомментировать проверку ниже:

        // if (response && response.isMatch === true) {
        console.log("It's a match! Showing modal...");
        setShowMatchModal(true);
        // Добавляем предотвращение скролла страницы
        document.body.style.overflow = 'hidden';
        // } else {
        //   console.log("No match, continuing without showing modal");
        // }

        // Вызываем onLike, если он предоставлен
        if (onLike) {
          onLike(userId);
        }
      } catch (error) {
        console.error('Error liking user:', error);
        // В случае ошибки можно вернуть состояние в исходное
        setIsLiked(false);
      }
    }
  };

  const handleCardClick = (_e: React.MouseEvent) => {
    // Only trigger the onClick if premiumOpen is false
    // This prevents the card click from being triggered when closing the modal
    if (!premiumOpen && !showMatchModal) {
      onClick?.();
    }
  };

  const handleMessageClick = (e: React.MouseEvent) => {
    // Stop propagation to prevent the card click handler from being triggered
    e.stopPropagation();
    // Prevent the default behavior
    e.preventDefault();
    // Open the premium modal
    setPremiumOpen(true);
    // Return false to ensure the event doesn't bubble up
    return false;
  };

  const handleModalClose = () => {
    setPremiumOpen(false);
  };

  const handleKeepSwiping = () => {
    setShowMatchModal(false);
    // Восстанавливаем скролл страницы
    document.body.style.overflow = '';
  };

  // Если карточка получила лайк, скрываем её
  if (isLiked) {
    return (
      <>
        {showMatchModal && (
          <MatchModal
            userImage="/photo1.png"
            matchImage={
              profile.photos?.[0]?.url ||
              profile.imageUrl ||
              profile.user?.imageUrl ||
              '/photo1.png'
            }
            matchName={profile.name || profile.user?.name || ''}
            onKeepSwiping={handleKeepSwiping}
          />
        )}
      </>
    );
  }

  return (
    <>
      <div className={`${styles.profileCard} ${className || ''}`} onClick={handleCardClick}>
        <div className={styles.imageCont}>
          <img
            src={profile.photos?.[0]?.url || profile.imageUrl || profile.user?.imageUrl}
            alt={profile.name}
            className={styles.profileImage}
          />
          {/* Wrap the icon in a div with stopPropagation for better isolation */}
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
        </div>
        <div className={styles.profileInfo}>
          <div className={styles.profileName}>
            {profile.name}, {profile.age}
          </div>
        </div>
        {onLike && (
          <div
            onClick={e => {
              e.stopPropagation();
              e.preventDefault();
            }}
            onMouseDown={e => {
              e.stopPropagation();
              e.preventDefault();
            }}
            // style={{
            //   position: 'absolute',
            //   bottom: '10px',
            //   right: '10px',
            //   zIndex: 20,
            //   cursor: 'pointer',
            // }}
          >
            <LikeButton
              onClick={handleLikeClick}
              className={styles.likeButton}
              likeClassName={styles.like}
            />
          </div>
        )}
        <PremiumModal isOpen={premiumOpen} onClose={handleModalClose} />
      </div>

      {/* Выносим модальное окно на уровень выше для корректного отображения */}
      {showMatchModal && (
        <MatchModal
          userImage="/photo1.png"
          matchImage={
            profile.photos?.[0]?.url || profile.imageUrl || profile.user?.imageUrl || '/photo1.png'
          }
          matchName={profile.name || profile.user?.name || ''}
          onKeepSwiping={handleKeepSwiping}
        />
      )}
    </>
  );
};

export default ProfileCard;
