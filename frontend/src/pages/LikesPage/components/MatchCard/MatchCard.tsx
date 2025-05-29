import React, { useEffect, useState } from 'react';
import { useDispatch } from 'react-redux';

import { useDeleteMatchMutation } from '../../../../entities/matches';
import { removeMatch } from '../../../../entities/matches/model/matchesSlice';
import { useGetTelegramMutation } from '../../../../entities/telegram/api/telegramApi';
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
}

export const MatchCard: React.FC<MatchCardProps> = ({
  match,
  onClick,
  // onSocialClick,
  showChatMessage,
}) => {
  const dispatch = useDispatch();
  const [getTelegram] = useGetTelegramMutation();
  const [telegram, setTelegram] = useState<string>('');

  const [deleteMatch] = useDeleteMatchMutation();
  const handleDeleteMatch = async () => {
    try {
      await deleteMatch({ targetUserId: String(match.id) });
      //@ts-ignore
      dispatch(removeMatch(match.id));
    } catch (error) {
      console.error('Error deleting match:', error);
    }
  };

  useEffect(() => {
    const fetchTg = async () => {
      try {
        //@ts-ignore
        console.log('match.id', match.id);

        //@ts-ignore
        const tg = await getTelegram(match.id).unwrap();

        //@ts-ignore
        setTelegram(tg.telegramUrl);
      } catch (err) {
        console.error('Error fetching users:', err);
      }
    };

    fetchTg();
  }, [getTelegram, match.id]);

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
        <div className={styles.telegram} onClick={() => (window.location.href = telegram)}>
          <TelegramIcon />
        </div>
        <DislikeButton onClick={handleDeleteMatch} className={styles.dislikeButton} />
      </div>
    </div>
  );
};

export default MatchCard;
