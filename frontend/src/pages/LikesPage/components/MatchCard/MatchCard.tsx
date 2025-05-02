import React from 'react';

import { IconButton } from '../../../../shared/components/IconButton';
import { InstagramIcon } from '../../../../shared/components/InstagramIcon';
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

export const MatchCard: React.FC<MatchCardProps> = ({
  match,
  onClick,
  onSocialClick,
  showChatMessage,
}) => {
  return (
    <div className={styles.matchCard} onClick={() => onClick(match)}>
      {/* @ts-ignore */}
      <img src={match.user.imageUrl} alt={match.user.name} className={styles.matchImage} />
      <div className={styles.matchInfo}>
        <div className={styles.matchName}>
          {/* @ts-ignore */}
          {match.user.name}, {match.user.age}
        </div>
        {/* @ts-ignore */}
        {showChatMessage === match.id && (
          <div className={styles.chatMessage}>Открыто в новой вкладке</div>
        )}
      </div>
      <div className={styles.socialButtons} onClick={e => e.stopPropagation()}>
        <IconButton
          // size="small"
          className={styles.socialButton}
          // @ts-ignore
          onClick={() => onSocialClick(match.user.id, 'telegram')}
        >
          <TelegramIcon />
        </IconButton>
        <IconButton
          // size="small"
          className={styles.socialButton}
          // @ts-ignore
          onClick={() => onSocialClick(match.user.id, 'instagram')}
        >
          <InstagramIcon />
        </IconButton>
      </div>
    </div>
  );
};

export default MatchCard;
