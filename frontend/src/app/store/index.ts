import { configureStore } from '@reduxjs/toolkit';
import { setupListeners } from '@reduxjs/toolkit/query';
import { baseApi } from '../../shared/api/baseApi';
import filtersReducer from '../../features/filters/model/filtersSlice';

export const store = configureStore({
  reducer: {
    // API reducers
    [baseApi.reducerPath]: baseApi.reducer,
    // Feature reducers
    filters: filtersReducer,
    // Feature reducers будут добавляться здесь
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(baseApi.middleware),
});

// Необходимо для refetchOnFocus/refetchOnReconnect
setupListeners(store.dispatch);

// Типы для использования в приложении
export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;