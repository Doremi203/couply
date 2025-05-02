import { MatchProfile } from '../types';

export const generateMatchId = (existingMatches: MatchProfile[]): number => {
  const maxId = existingMatches.reduce(
    (max: number, match: MatchProfile) => Math.max(max, match.id),
    100,
  );
  return maxId + 1;
};

//TODO
