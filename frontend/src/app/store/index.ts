import { configureStore } from '@reduxjs/toolkit';
import { setupListeners } from '@reduxjs/toolkit/query';

import matchesReducer from '../../entities/matches/model/matchesSlice';
import userReducer from '../../entities/user/model/userSlice';
import filtersReducer from '../../features/filters/model/filtersSlice';
import { baseApi, matcherApi } from '../../shared/api/baseApi';

export const store = configureStore({
  reducer: {
    [baseApi.reducerPath]: baseApi.reducer,
    [matcherApi.reducerPath]: matcherApi.reducer,
    matches: matchesReducer,
    filters: filtersReducer,
    user: userReducer,
  },
  middleware: getDefaultMiddleware =>
    getDefaultMiddleware().concat(baseApi.middleware, matcherApi.middleware),
});

setupListeners(store.dispatch);

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
