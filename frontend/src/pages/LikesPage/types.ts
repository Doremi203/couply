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
  verified?: boolean;
  user?: {
    id: number;
    name: string;
    age: number;
    photos: { url: string }[];
    bio?: string;
    location?: string;
    interests?: string[];
    lifestyle?: { [key: string]: string };
    passion?: string[];
    verified?: boolean;
    hasLikedYou?: boolean;
  };
}

export interface MatchProfile {
  id: number;
  name: string;
  age: number;
  imageUrl: string;
  telegram: string;
}
