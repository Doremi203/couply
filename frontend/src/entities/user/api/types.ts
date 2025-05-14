import { Gender, Goal, Zodiac, Education, Children, Alcohol, Smoking } from './constants';

//TODO
export interface UserRequest {
  id?: string;
  name: string;
  age: number;
  gender: Gender;
  location: string;
  bio: string;
  goal: Goal;
  interest: null;
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
    interest: null;
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
