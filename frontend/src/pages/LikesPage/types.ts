export interface LikeProfile {
  id: number;
  name: string;
  age: number;
  imageUrl: string;
  liked?: boolean;
  hasLikedYou?: boolean;
  bio?: string;
  location?: string;
  interests?: string[];
  lifestyle?: { [key: string]: string };
  passion?: string[];
}

export interface MatchProfile {
  id: number;
  name: string;
  age: number;
  imageUrl: string;
  telegram: string;
}
