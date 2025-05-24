import { createSlice, PayloadAction } from '@reduxjs/toolkit';

import { RootState } from '../../../app/store';

export interface ProfileState {
  data: {
    id?: string;
    name: string;
    age: number;
    gender: string;
    location?: string;
    bio?: string;
    goal?: string;
    interest?: {
      sport?: string[];
      selfDevelopment?: string[];
      art?: string[];
      music?: string[];
      movies?: string[];
      books?: string[];
      pets?: string[];
    };
    education?: string;
    children?: string;
    alcohol?: string;
    smoking?: string;
    zodiac?: string;
    isVerified?: boolean;
    isPremium?: boolean;
    isBlocked?: boolean;
    isHidden?: boolean;
    latitude?: number;
    longitude?: number;
    height?: number;
    photos?: Array<{
      url: string;
      isNew?: boolean;
    }>;
  };
  isLoading: boolean;
  error: string | null;
}

const initialState: ProfileState = {
  data: {
    name: '',
    age: 0,
    gender: '',
    location: '',
    bio: '',
    isHidden: false,
    photos: [],
  },
  isLoading: false,
  error: null,
};

const profileSlice = createSlice({
  name: 'profile',
  initialState,
  reducers: {
    setProfileData: (state, action: PayloadAction<ProfileState['data']>) => {
      state.data = action.payload;
    },
    updateProfileField: (
      state,
      action: PayloadAction<{ field: keyof ProfileState['data']; value: any }>,
    ) => {
      const { field, value } = action.payload;
      //@ts-ignore
      state.data[field] = value;
    },
    setLoading: (state, action: PayloadAction<boolean>) => {
      state.isLoading = action.payload;
    },
    setError: (state, action: PayloadAction<string | null>) => {
      state.error = action.payload;
    },
    resetProfile: state => {
      state.data = initialState.data;
      state.isLoading = false;
      state.error = null;
    },
  },
});

// Selectors
export const selectProfileData = (state: RootState) => state.profile.data;
export const selectProfileLoading = (state: RootState) => state.profile.isLoading;
export const selectProfileError = (state: RootState) => state.profile.error;

// Actions
export const { setProfileData, updateProfileField, setLoading, setError, resetProfile } =
  profileSlice.actions;

export default profileSlice.reducer;
