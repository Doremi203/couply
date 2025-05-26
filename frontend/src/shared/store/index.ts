// import { configureStore } from '@reduxjs/toolkit';

// import matchesReducer from '../../entities/matches/model/matchesSlice';
// import { matcherApi } from '../api/baseApi';

// export const store = configureStore({
//   reducer: {
//     [matcherApi.reducerPath]: matcherApi.reducer,
//     matches: matchesReducer,
//   },
//   middleware: getDefaultMiddleware => getDefaultMiddleware().concat(matcherApi.middleware),
// });

// export type RootState = ReturnType<typeof store.getState>;
// export type AppDispatch = typeof store.dispatch;


import { configureStore } from '@reduxjs/toolkit';

import { baseApi, blockerApi, matcherApi, notificatorApi, paymentsApi } from '../api/baseApi';


export const store = configureStore({
  reducer: {
    [baseApi.reducerPath]: baseApi.reducer,
    [matcherApi.reducerPath]: matcherApi.reducer,
    [blockerApi.reducerPath]: blockerApi.reducer,
    [paymentsApi.reducerPath]: paymentsApi.reducer,
    [notificatorApi.reducerPath]: notificatorApi.reducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware()
      .concat(baseApi.middleware)
      .concat(matcherApi.middleware)
      .concat(blockerApi.middleware)
      .concat(paymentsApi.middleware)
      .concat(notificatorApi.middleware),
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;