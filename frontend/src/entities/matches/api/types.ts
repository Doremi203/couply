export interface Match {
  match: {
    mainUserId: string;
    chosenUserId: string;
    approved: boolean;
  };
}

export interface MatchRequest {
  mainUserId: string;
  chosenUserId: string;
  approved: boolean;
}

export interface CreateMatchRequest {
  mainUserId: string;
  chosenUserId: string;
}

export interface FetchMatchesRequest {
  mainUserId: string;
  limit: number;
  offset: number;
}

export interface FetchInMatchesRequest {
  chosenUserId: string;
  limit: number;
  offset: number;
}

export interface MatchesResponse {
  match: Array<{
    mainUserId: string;
    chosenUserId: string;
    approved: boolean;
  }>;
}
