import React from 'react';

import EmptyState from '../../../../shared/components/EmptyState';
import { ProfileCard, ProfileData } from '../../../../shared/components/ProfileCard';

import styles from './likesSection.module.css';

interface LikesSectionProps {
  likes: ProfileData[];
  onProfileClick: (profile: ProfileData) => void;
  onLike: (id: number) => void;
}

export const LikesSection: React.FC<LikesSectionProps> = ({ likes, onProfileClick, onLike }) => {
  console.log('3', likes);
  if (likes.length === 0) {
    return (
      <EmptyState title="У вас пока нет лайков" subtitle="Продолжайте искать новые знакомства" />
    );
  }

  //TODO вернуть profile.user

  return (
    <div className={styles.section}>
      <div className={styles.profilesGrid}>
        {likes.map(profile => (
          <ProfileCard
            //@ts-ignore
            key={profile.id}
            profile={profile}
            onClick={() => onProfileClick(profile)}
            onLike={onLike}
            className={styles.profileCard}
          />
        ))}
      </div>
    </div>
  );
};

export default LikesSection;
