export interface User {
  id: string;
  email?: string;
  phone?: string;
  name?: string;
  createdAt: string;
}

export interface RegisterParams {
  email?: string;
  phone?: string;
  password: string;
}

export interface LoginParams {
  email?: string;
  phone?: string;
  password: string;
}

export interface AuthResponse {
  user: User;
  token: string;
}

export interface LoginResponse {
  token: string;
  expiresIn: number;
  refreshToken: {
    token: string;
    expiresIn: number;
  };
}

export interface RefreshResponse {
  accessToken: {
    token: string;
    expiresIn: number;
  };
  refreshToken: {
    token: string;
    expiresIn: number;
  };
}

export interface RefreshRequest {
  refreshToken: string;
}
