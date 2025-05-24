export interface LikeProfile {
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
