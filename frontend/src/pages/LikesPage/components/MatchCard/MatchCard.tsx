import React from 'react';

import { IconButton } from '../../../../shared/components/IconButton';
import { TelegramIcon } from '../../../../shared/components/TelegramIcon';

import styles from './matchCard.module.css';
// import { IconButton } from "@mui/material";

// Define the match profile interface
export interface MatchProfile {
  match: {
    id: number;
    name: string;
    age: number;
    imageUrl: string;
    telegram: string;
    instagram: string;
  };
}

interface MatchCardProps {
  match: MatchProfile;
  onClick: (match: MatchProfile) => void;
  onSocialClick: (matchId: number, type: 'telegram' | 'instagram') => void;
  showChatMessage: number | null;
}

//TODO вернуть profile.user

export const MatchCard: React.FC<MatchCardProps> = ({
  match,
  onClick,
  onSocialClick,
  showChatMessage,
}) => {
  return (
    <div className={styles.matchCard} onClick={() => onClick(match)}>
      {/* @ts-ignore */}
      <img src={match.imageUrl} alt={match.name} className={styles.matchImage} />
      <div className={styles.matchInfo}>
        <div className={styles.matchName}>
          {/* @ts-ignore */}
          {match.name}, {match.age}
        </div>
        {/* @ts-ignore */}
        {showChatMessage === match.id && (
          <div className={styles.chatMessage}>Открыто в новой вкладке</div>
        )}
      </div>
      <div onClick={e => e.stopPropagation()}>
        <IconButton
          // size="small"
          className={styles.socialButton}
          // @ts-ignore
          onClick={() => onSocialClick(match.id, 'telegram')}
        >
          <TelegramIcon />
        </IconButton>
      </div>
    </div>
  );
};

export default MatchCard;
