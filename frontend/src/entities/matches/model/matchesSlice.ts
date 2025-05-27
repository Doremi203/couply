import { createSelector, createSlice, PayloadAction } from '@reduxjs/toolkit';

import { RootState } from '../../../app/store';
import { UserData } from '../../user/types';
import { Like } from '../types';

interface MatchesState {
  likes: Like[];
  likesUsers: { users: UserData[] };
  matches: { users: UserData[] };
  matchesOffset: number;
  likesOffset: number;
  hasMoreMatches: boolean;
  hasMoreLikes: boolean;
  isLoading: boolean;
  showMatchModal: boolean;
  matchedProfile: any | null;
  showChatMessage: number | null;
}

const initialState: MatchesState = {
  likes: [],
  likesUsers: { users: [] },
  matches: { users: [] },
  matchesOffset: 0,
  likesOffset: 0,
  hasMoreMatches: true,
  hasMoreLikes: true,
  isLoading: false,
  showMatchModal: false,
  matchedProfile: null,
  showChatMessage: null,
};

const matchesSlice = createSlice({
  name: 'matches',
  initialState,
  reducers: {
    setLikes: (state, action: PayloadAction<Like[]>) => {
      state.likes = action.payload;
    },
    setLikesUsers: (state, action: PayloadAction<UserData[]>) => {
      state.likesUsers.users = action.payload;
    },
    setMatches: (state, action: PayloadAction<UserData[]>) => {
      state.matches.users = action.payload;
    },
    appendMatches: (state, action: PayloadAction<UserData[]>) => {
      state.matches.users = [...state.matches.users, ...action.payload];
    },
    appendLikes: (state, action: PayloadAction<{ likes: Like[]; users: UserData[] }>) => {
      state.likes = [...state.likes, ...action.payload.likes];
      state.likesUsers.users = [...state.likesUsers.users, ...action.payload.users];
    },
    setMatchesOffset: (state, action: PayloadAction<number>) => {
      state.matchesOffset = action.payload;
    },
    setLikesOffset: (state, action: PayloadAction<number>) => {
      state.likesOffset = action.payload;
    },
    setHasMoreMatches: (state, action: PayloadAction<boolean>) => {
      state.hasMoreMatches = action.payload;
    },
    setHasMoreLikes: (state, action: PayloadAction<boolean>) => {
      state.hasMoreLikes = action.payload;
    },
    setIsLoading: (state, action: PayloadAction<boolean>) => {
      state.isLoading = action.payload;
    },
    setShowMatchModal: (state, action: PayloadAction<boolean>) => {
      state.showMatchModal = action.payload;
    },
    setMatchedProfile: (state, action: PayloadAction<any | null>) => {
      state.matchedProfile = action.payload;
    },
    setShowChatMessage: (state, action: PayloadAction<number | null>) => {
      state.showChatMessage = action.payload;
    },
    removeLike: (state, action: PayloadAction<string>) => {
      state.likes = state.likes.filter(like => like.senderId !== action.payload);
      //@ts-ignore
      state.likesUsers.users.users = state.likesUsers.users.users.filter(
        //@ts-ignore
        user => user.id !== action.payload,
      );
    },
    addMatch: (state, action: PayloadAction<UserData>) => {
      //@ts-ignore
      state.matches.users.users.push(action.payload);
    },
    removeMatch: (state, action: PayloadAction<string>) => {
      //@ts-ignore
      state.matches.users.users = state.matches.users.users.filter(
        //@ts-ignore
        match => match.id !== action.payload,
      );
    },
  },
});

// Selectors
export const selectMatchesState = (state: RootState) => state.matches;

export const selectLikes = createSelector(selectMatchesState, matchesState => matchesState.likes);

export const selectLikesUsers = createSelector(
  selectMatchesState,
  matchesState => matchesState.likesUsers.users,
);

export const selectMatches = createSelector(
  selectMatchesState,
  matchesState => matchesState.matches.users,
);

export const selectMatchesOffset = createSelector(
  selectMatchesState,
  matchesState => matchesState.matchesOffset,
);

export const selectLikesOffset = createSelector(
  selectMatchesState,
  matchesState => matchesState.likesOffset,
);

export const selectHasMoreMatches = createSelector(
  selectMatchesState,
  matchesState => matchesState.hasMoreMatches,
);

export const selectHasMoreLikes = createSelector(
  selectMatchesState,
  matchesState => matchesState.hasMoreLikes,
);

export const selectIsLoading = createSelector(
  selectMatchesState,
  matchesState => matchesState.isLoading,
);

export const selectShowMatchModal = createSelector(
  selectMatchesState,
  matchesState => matchesState.showMatchModal,
);

export const selectMatchedProfile = createSelector(
  selectMatchesState,
  matchesState => matchesState.matchedProfile,
);

export const selectShowChatMessage = createSelector(
  selectMatchesState,
  matchesState => matchesState.showChatMessage,
);

// Combined selectors
export const selectLikesWithUsers = createSelector(
  [selectLikes, selectLikesUsers],
  (likes, users) => ({
    likes,
    users,
  }),
);

export const selectMatchesWithLikes = createSelector(
  [selectMatches, selectLikes],
  (matches, likes) => ({
    matches,
    likes,
  }),
);

export const {
  setLikes,
  setLikesUsers,
  setMatches,
  appendMatches,
  appendLikes,
  setMatchesOffset,
  setLikesOffset,
  setHasMoreMatches,
  setHasMoreLikes,
  setIsLoading,
  setShowMatchModal,
  setMatchedProfile,
  setShowChatMessage,
  removeLike,
  addMatch,
  removeMatch,
} = matchesSlice.actions;

export default matchesSlice.reducer;
