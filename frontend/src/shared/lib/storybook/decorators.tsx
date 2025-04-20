import { configureStore } from '@reduxjs/toolkit';
import React, { ComponentType, ReactElement } from 'react';
import { Provider } from 'react-redux';

import { baseApi } from '../../api/baseApi';

// Create a mock store for Storybook
export const mockStore = configureStore({
  reducer: {
    [baseApi.reducerPath]: baseApi.reducer,
  },
  middleware: getDefaultMiddleware => getDefaultMiddleware().concat(baseApi.middleware),
});

// Redux Provider decorator for Storybook
export const withReduxProvider = (Story: ComponentType): ReactElement => (
  <Provider store={mockStore}>
    <Story />
  </Provider>
);
