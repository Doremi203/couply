// @ts-nocheck
import SmsOutlinedIcon from '@mui/icons-material/SmsOutlined';
import React, { useState } from 'react';

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

  const handleLikeClick = (e: React.MouseEvent) => {
    e.stopPropagation();
    onLike?.(profile.user.id);
  };

  const handleCardClick = (_e: React.MouseEvent) => {
    // Only trigger the onClick if premiumOpen is false
    // This prevents the card click from being triggered when closing the modal
    if (!premiumOpen) {
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

  return (
    <div className={`${styles.profileCard} ${className || ''}`} onClick={handleCardClick}>
      <div className={styles.imageCont}>
        <img src={profile.imageUrl} alt={profile.name} className={styles.profileImage} />
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
        <div onClick={e => e.stopPropagation()}>
          <LikeButton
            onClick={handleLikeClick}
            className={styles.likeButton}
            likeClassName={styles.like}
          />
        </div>
      )}
      <PremiumModal isOpen={premiumOpen} onClose={handleModalClose} />
    </div>
  );
};

export default ProfileCard;
