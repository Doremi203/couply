import { Status } from '../constants';

export interface FetchMatchesRequest {
  limit: number;
  offset: number;
}

export interface MatchRequest {
  targetUserId: string;
}

export interface FetchMatchesResponse {
  likes: Array<{
    senderId: string;
    receiverId: string;
    message: string;
    status: Status;
  }>;
}

export interface FetchMatchesUserIdsResponse {
  userIds: Array<string>;
}

export interface LikeRequest {
  targetUserId: string;
  message: string;
}

export interface LikeResponse {
  isMatch: boolean;
  match: {
    firstUserId: string;
    secondUserId: string;
    createdAt: string;
  };
}
