import { ProfileData } from '../../shared/components/ProfileCard';

// Define types for the likes feature
export interface LikeProfile extends ProfileData {
  liked?: boolean;
  hasLikedYou?: boolean;
}

export interface MatchProfile {
  id: number;
  name: string;
  age: number;
  imageUrl: string;
  telegram: string;
  instagram: string;
}
