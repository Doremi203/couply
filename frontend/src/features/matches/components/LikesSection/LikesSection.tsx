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
  //@ts-ignore
  const users = likes.users;

  if (users === undefined || users.length === 0) {
    return (
      <EmptyState title="У вас пока нет лайков" subtitle="Продолжайте искать новые знакомства" />
    );
  }

  return (
    <div className={styles.section}>
      <div className={styles.profilesGrid}>
        {/**@ts-ignore */}
        {users.map(profile => (
          <ProfileCard
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
