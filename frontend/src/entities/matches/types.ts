import { UserData } from '../user/types';

export interface Like {
  senderId: string;
  receiverId: string;
  message: string;
}

export interface LikesData {
  likes: Array<Like>;
}

export interface LikesUsersData {
  users: Array<UserData>;
}
