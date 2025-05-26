import { configureStore, combineReducers } from '@reduxjs/toolkit';
import { setupListeners } from '@reduxjs/toolkit/query';

import matchesReducer from '../../entities/matches/model/matchesSlice';
import profileReducer from '../../entities/profile/model/profileSlice';
import userReducer from '../../entities/user/model/userSlice';
import filtersReducer from '../../features/filters/model/filtersSlice';
import {
  baseApi,
  blockerApi,
  matcherApi,
  notificatorApi,
  paymentsApi,
} from '../../shared/api/baseApi';

import blockingReducer from './blockingSlice';

const rootReducer = combineReducers({
  [baseApi.reducerPath]: baseApi.reducer,
  [matcherApi.reducerPath]: matcherApi.reducer,
  [blockerApi.reducerPath]: blockerApi.reducer,
  [paymentsApi.reducerPath]: paymentsApi.reducer,
  [notificatorApi.reducerPath]: notificatorApi.reducer,
  matches: matchesReducer,
  filters: filtersReducer,
  user: userReducer,
  profile: profileReducer,
  blocking: blockingReducer,
});

export const store = configureStore({
  reducer: rootReducer,
  middleware: getDefaultMiddleware => {
    return getDefaultMiddleware()
      .concat(baseApi.middleware)
      .concat(matcherApi.middleware)
      .concat(blockerApi.middleware)
      .concat(paymentsApi.middleware)
      .concat(notificatorApi.middleware);
  },
});

setupListeners(store.dispatch);

export type RootState = ReturnType<typeof rootReducer>;
export type AppDispatch = typeof store.dispatch;
