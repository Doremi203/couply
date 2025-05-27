import React from 'react';

import { LikesData, LikesUsersData } from '../../../../entities/matches/types';
import EmptyState from '../../../../shared/components/EmptyState';
import { ProfileCard } from '../../../../shared/components/ProfileCard';

import styles from './likesSection.module.css';

interface LikesSectionProps {
  likes: LikesData[];
  likesUsers: LikesUsersData[];
  //@ts-ignore
  onProfileClick: (profile) => void;
  onLike: (id: string) => void;
}

export const LikesSection: React.FC<LikesSectionProps> = ({
  likes,
  likesUsers,
  onProfileClick,
  onLike,
}) => {
  //@ts-ignore
  const users = likesUsers.users;

  if (users === undefined || users.length === 0) {
    return (
      <EmptyState title="У вас пока нет лайков" subtitle="Продолжайте искать новые знакомства" />
    );
  }

  return (
    <div className={styles.section}>
      <div className={styles.profilesGrid}>
        {/*@ts-ignore */}
        {users.map((profile, index) => (
          <ProfileCard
            key={profile.id}
            profile={profile}
            onClick={() => onProfileClick(profile)}
            //@ts-ignore
            onLike={onLike}
            className={styles.profileCard}
            //@ts-ignore
            like={likes[index]}
          />
        ))}
      </div>
    </div>
  );
};

export default LikesSection;
