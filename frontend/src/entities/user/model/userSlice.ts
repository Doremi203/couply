import { createSlice, PayloadAction } from '@reduxjs/toolkit';

import { RootState } from '../../../app/store';

interface UserState {
  id: string | null;
  isAuthenticated: boolean;
  isPremium: boolean;
  isVerified: boolean;
}

const initialState: UserState = {
  id: null,
  isAuthenticated: false,
  isPremium: false,
  isVerified: false,
};

export const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {
    setUserId: (state, action: PayloadAction<string>) => {
      state.id = action.payload;
      state.isAuthenticated = true;
    },
    clearUserId: state => {
      state.id = null;
      state.isAuthenticated = false;
    },
    setUserPremium: state => {
      state.isPremium = true;
    },
    setUserVerified: state => {
      state.isVerified = true;
    },
  },
});

export const { setUserId, clearUserId, setUserPremium, setUserVerified } = userSlice.actions;

export const getUserId = (state: RootState) => state.user.id;
export const getIsAuthenticated = (state: RootState) => state.user.isAuthenticated;

export const getIsPremium = (state: RootState) => state.user.isPremium;
export const getIsVerified = (state: RootState) => state.user.isVerified;

export default userSlice.reducer;
