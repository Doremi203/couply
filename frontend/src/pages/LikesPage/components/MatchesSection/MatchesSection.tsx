import React, { useState } from 'react';

import EmptyState from '../../../../shared/components/EmptyState';
import { MatchProfile } from '../../types';
import { MatchCard } from '../MatchCard';

import styles from './matchesSection.module.css';
import { selectMatches } from '../../../../entities/matches/model/matchesSlice';
import { useSelector } from 'react-redux';

interface MatchesSectionProps {
  matches: MatchProfile[];
  onMatchClick: (match: MatchProfile) => void;
  onSocialClick: (matchId: number, type: 'telegram' | 'instagram') => void;
  showChatMessage: number | null;
}

export const MatchesSection: React.FC<MatchesSectionProps> = ({
  // matches,
  onMatchClick,
  onSocialClick,
  showChatMessage,
}) => {
  const matches = useSelector(selectMatches);

  if (matches.users.length === 0) {
    return (
      <EmptyState title="У вас пока нет мэтчей" subtitle="Лайкайте профили, чтобы найти мэтчи" />
    );
  }

  return (
    <div className={styles.section}>
      <div className={styles.matchesContainer}>
        {/**@ts-ignore */}
        {matches.users.map(match => (
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
