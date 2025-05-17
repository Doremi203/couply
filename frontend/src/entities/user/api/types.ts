import { Gender, Goal, Zodiac, Education, Children, Alcohol, Smoking, FoodDrink, Hobby, MoviesTV, Music, PersonalityTraits, Pets, Selfdevelopment, Sport } from './constants';

//TODO
export interface UserRequest {
  id?: string;
  name: string;
  age: number;
  gender: Gender;
  location: string;
  bio: string;
  goal: Goal;
  interest: {
    sport: [Sport],
    selfDevelopment: [Selfdevelopment],
    hobby: [Hobby],
    music: [Music],
    moviesTv: [MoviesTV],
    foodDrink: [FoodDrink],
    personalityTraits: [PersonalityTraits],
    pets: [Pets],
  };
  zodiac: Zodiac;
  height: number;
  education: Education;
  children: Children;
  alcohol: Alcohol;
  smoking: Smoking;
  hidden: boolean;
  verified: boolean;
  photos: null;
}

export interface UserResponse {
  user: {
    id: string;
    name: string;
    age: number;
    gender: Gender;
    location: string;
    bio: string;
    goal: Goal;
    interest: {
      sport: [Sport],
      selfDevelopment: [Selfdevelopment],
      hobby: [Hobby],
      music: [Music],
      moviesTv: [MoviesTV],
      foodDrink: [FoodDrink],
      personalityTraits: [PersonalityTraits],
      pets: [Pets],
    };
    zodiac: Zodiac;
    height: number;
    education: Education;
    children: Children;
    alcohol: Alcohol;
    smoking: Smoking;
    hidden: boolean;
    verified: boolean;
    photos: null;
  };
}

export interface UpdateUserParams {
  id: string;
  data: Partial<UserRequest>;
}

export interface GetUserRequest {
  id: string;
}
export { Smoking };

export interface PhotoParams {
  orderNumbers: Array<number>;
}

export interface UsersResponse {
  users: Array<UserResponse>;
}

export interface GetUsersRequest {
  userIds: Array<string>;
}
