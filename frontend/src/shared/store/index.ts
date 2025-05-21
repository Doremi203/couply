import { configureStore } from '@reduxjs/toolkit';

import matchesReducer from '../../entities/matches/model/matchesSlice';
import { matcherApi } from '../api/baseApi';

export const store = configureStore({
  reducer: {
    [matcherApi.reducerPath]: matcherApi.reducer,
    matches: matchesReducer,
  },
  middleware: getDefaultMiddleware => getDefaultMiddleware().concat(matcherApi.middleware),
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
