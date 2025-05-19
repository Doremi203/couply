import React from 'react';

import { useDeleteMatchMutation } from '../../../../entities/matches';
import { DislikeButton } from '../../../../shared/components/DislikeButton';
import { TelegramIcon } from '../../../../shared/components/TelegramIcon';

import styles from './matchCard.module.css';

export interface MatchProfile {
  id: number;
  name: string;
  age: number;
  imageUrl: string;
  telegram: string;
  instagram: string;
  photos: { url: string }[];
}

interface MatchCardProps {
  match: MatchProfile;
  onClick: (match: MatchProfile) => void;
  onSocialClick: (matchId: number, type: 'telegram' | 'instagram') => void;
  showChatMessage: number | null;
  onRemove: (id: number) => void;
}

//TODO вернуть profile.user

export const MatchCard: React.FC<MatchCardProps> = ({
  match,
  onClick,
  // onSocialClick,
  showChatMessage,
  onRemove,
}) => {
  const [deleteMatch] = useDeleteMatchMutation();
  const handleDeleteMatch = async () => {
    try {
      await deleteMatch({ targetUserId: String(match.id) });
      onRemove(match.id);
    } catch (error) {
      console.error('Error deleting match:', error);
    }
  };

  return (
    <div className={styles.matchCard} onClick={() => onClick(match)}>
      <img src={match.photos[0].url} alt={match.name} className={styles.matchImage} />
      <div className={styles.matchInfo}>
        <div className={styles.matchName}>
          {match.name}, {match.age}
        </div>
        {showChatMessage === match.id && (
          <div className={styles.chatMessage}>Открыто в новой вкладке</div>
        )}
      </div>
      <div onClick={e => e.stopPropagation()} className={styles.buttons}>
        <div className={styles.telegram}>
          <TelegramIcon />
        </div>
        <DislikeButton onClick={handleDeleteMatch} className={styles.dislikeButton} />
      </div>
    </div>
  );
};

export default MatchCard;
