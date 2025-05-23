import { configureStore } from '@reduxjs/toolkit';
import { ComponentType, ReactElement } from 'react';
import { Provider } from 'react-redux';

import { baseApi } from '../../api/baseApi';

export const mockStore = configureStore({
  reducer: {
    [baseApi.reducerPath]: baseApi.reducer,
  },
  middleware: getDefaultMiddleware => getDefaultMiddleware().concat(baseApi.middleware),
});

export const withReduxProvider = (Story: ComponentType): ReactElement => (
  <Provider store={mockStore}>
    <Story />
  </Provider>
);
