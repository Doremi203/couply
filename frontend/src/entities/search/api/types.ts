import {
  Alcohol,
  Children,
  Education,
  FoodDrink,
  Goal,
  Hobby,
  MoviesTV,
  Music,
  PersonalityTraits,
  Pets,
  Selfdevelopment,
  Smoking,
  Sport,
  Zodiac,
} from '../../user';

import { GenderPriority } from './constants';

export interface CreateFilterRequest {
  genderPriority: GenderPriority;
  minAge: number;
  maxAge: number;
  minHeight: number;
  maxHeight: number;
  distance: number;
  goal: Goal;
  zodiac: Zodiac;
  education: Education;
  children: Children;
  alcohol: Alcohol;
  smoking: Smoking;
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
  onlyVerified: boolean;
  onlyPremium: boolean;
}

export interface FilterResponse {
  filter: {
    genderPriority: GenderPriority;
    ageRange: {
      min: number;
      max: number;
    };
    heightRange: {
      min: number;
      max: number;
    };
    distance: number;
    goal: Goal;
    zodiac: Zodiac;
    education: Education;
    children: Children;
    alcohol: Alcohol;
    smoking: Smoking;
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
    onlyVerified: boolean;
    onlyPremium: boolean;
    createdAt: string;
    updatedAt: string;
  };
}

export interface SearchRequest {
  offset: number;
  limit: number;
}
export interface SearchResponse {
  users: [
    {
      id: boolean;
      name: string;
      age: number;
      gender: GenderPriority;
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
      hidden: boolean;
      verified: boolean;
      photos: [
        {
          orderNumber: number;
          url: string;
          mimeType: string;
          uploadedAt: string;
          updatedAt: string;
        },
      ];
      createdAt: string;
      updatedAt: string;
    },
  ];
}
