import { createSlice, PayloadAction } from '@reduxjs/toolkit';

import { RootState } from '../../../app/store';

interface UserState {
  id: string | null;
  isAuthenticated: boolean;
}

const initialState: UserState = {
  id: null,
  isAuthenticated: false,
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
  },
});

export const { setUserId, clearUserId } = userSlice.actions;

export const getUserId = (state: RootState) => state.user.id;
export const getIsAuthenticated = (state: RootState) => state.user.isAuthenticated;

export default userSlice.reducer;
