import React from 'react';

import { LikesData, LikesUsersData } from '../../../../entities/matches/types';
import EmptyState from '../../../../shared/components/EmptyState';
import { ProfileCard, ProfileData } from '../../../../shared/components/ProfileCard';

import styles from './likesSection.module.css';

interface LikesSectionProps {
  likes: LikesData[];
  likesUsers: LikesUsersData[];
  onProfileClick: (profile: ProfileData) => void;
  onLike: (id: number) => void;
}

export const LikesSection: React.FC<LikesSectionProps> = ({
  likes,
  likesUsers,
  onProfileClick,
  onLike,
}) => {
  const users = likesUsers.users;

  if (users === undefined || users.length === 0) {
    return (
      <EmptyState title="У вас пока нет лайков" subtitle="Продолжайте искать новые знакомства" />
    );
  }

  return (
    <div className={styles.section}>
      <div className={styles.profilesGrid}>
        {users.map((profile, index) => (
          <ProfileCard
            key={profile.id}
            profile={profile}
            onClick={() => onProfileClick(profile)}
            onLike={onLike}
            className={styles.profileCard}
            like={likes[index]}
          />
        ))}
      </div>
    </div>
  );
};

export default LikesSection;
