export interface GetBlockResponse {
  blockId: string;
  blockedUserId: string;
  message: string;
  reasons: string[];
  createdAt: string;
}

export interface CreateBlockRequest {
  targetUserId: string;
  reasons: string[];
  message: string;
}
