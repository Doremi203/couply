import { configureStore } from '@reduxjs/toolkit';
import { StoryObj, Meta } from '@storybook/react';
import { Provider } from 'react-redux';
import { BrowserRouter } from 'react-router-dom';

import userReducer from '../../../../entities/user/model/userSlice';
import { baseApi } from '../../../../shared/api/baseApi';

import HomePage from './Page';

// Create a mock store with the user reducer
const mockStore = configureStore({
  reducer: {
    [baseApi.reducerPath]: baseApi.reducer,
    user: userReducer,
  },
  middleware: getDefaultMiddleware => getDefaultMiddleware().concat(baseApi.middleware),
  preloadedState: {
    user: {
      id: 'mock-user-id',
      isAuthenticated: true,
    },
  },
});

const meta: Meta = {
  title: 'Pages/HomePage',
  component: HomePage,
  decorators: [
    Story => (
      <Provider store={mockStore}>
        <BrowserRouter>
          <Story />
        </BrowserRouter>
      </Provider>
    ),
  ],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj;

export const Default: Story = {
  args: {},
};
