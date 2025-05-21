import { configureStore } from '@reduxjs/toolkit';
import { matcherApi } from '../api/baseApi';
import matchesReducer from '../../entities/matches/model/matchesSlice';

export const store = configureStore({
  reducer: {
    [matcherApi.reducerPath]: matcherApi.reducer,
    matches: matchesReducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(matcherApi.middleware),
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch; 