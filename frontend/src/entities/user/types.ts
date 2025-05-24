import {
  Gender,
  Goal,
  Sport,
  Selfdevelopment,
  Hobby,
  Music,
  MoviesTV,
  FoodDrink,
  PersonalityTraits,
  Pets,
  Zodiac,
  Education,
  Children,
  Alcohol,
  Smoking,
} from './api/constants';

export interface User {
  user: {
    id: string;
    name: string;
    age: number;
    gender: Gender;
    location: string;
    bio: string;
    goal: Goal;
    interest: {
      sport: [Sport];
      selfDevelopment: [Selfdevelopment];
      hobby: [Hobby];
      music: [Music];
      moviesTv: [MoviesTV];
      foodDrink: [FoodDrink];
      personalityTraits: [PersonalityTraits];
      pets: [Pets];
    };
    zodiac: Zodiac;
    height: number;
    education: Education;
    children: Children;
    alcohol: Alcohol;
    smoking: Smoking;
    isPremium: boolean;
    isBlocked: boolean;
    isVerified: boolean;
    isHidden: boolean;
    photos: Array<{
      orderNumber: number;
      url: string;
    }>;
  };
}

export interface UserData {
  id: string;
  name: string;
  age: number;
  gender: Gender;
  location: string;
  bio: string;
  goal: Goal;
  interest: {
    sport: [Sport];
    selfDevelopment: [Selfdevelopment];
    hobby: [Hobby];
    music: [Music];
    moviesTv: [MoviesTV];
    foodDrink: [FoodDrink];
    personalityTraits: [PersonalityTraits];
    pets: [Pets];
  };
  zodiac: Zodiac;
  height: number;
  education: Education;
  children: Children;
  alcohol: Alcohol;
  smoking: Smoking;
  isPremium: boolean;
  isBlocked: boolean;
  isVerified: boolean;
  isHidden: boolean;
  photos: Array<{
    orderNumber: number;
    url: string;
  }>;
}
