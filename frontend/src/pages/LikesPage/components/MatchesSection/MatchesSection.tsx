import React from 'react';

import EmptyState from '../../../../shared/components/EmptyState';
import { MatchProfile } from '../../types';
import { MatchCard } from '../MatchCard';

import styles from './matchesSection.module.css';

interface MatchesSectionProps {
  matches: MatchProfile[];
  onMatchClick: (match: MatchProfile) => void;
  onSocialClick: (matchId: number, type: 'telegram' | 'instagram') => void;
  showChatMessage: number | null;
}

export const MatchesSection: React.FC<MatchesSectionProps> = ({
  matches,
  onMatchClick,
  onSocialClick,
  showChatMessage,
}) => {
  //@ts-ignore
  const users = matches.users;
  if (users === undefined || users.length === 0) {
    return (
      <EmptyState title="У вас пока нет мэтчей" subtitle="Лайкайте профили, чтобы найти мэтчи" />
    );
  }

  return (
    <div className={styles.section}>
      <div className={styles.matchesContainer}>
        {/**@ts-ignore */}
        {users.map(match => (
          <MatchCard
            key={match.id}
            // @ts-ignore
            match={match}
            // @ts-ignore
            onClick={onMatchClick}
            onSocialClick={onSocialClick}
            showChatMessage={showChatMessage}
          />
        ))}
      </div>
    </div>
  );
};

export default MatchesSection;
