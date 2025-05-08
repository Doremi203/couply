import { configureStore } from '@reduxjs/toolkit';
import type { Meta, StoryObj } from '@storybook/react';
import { Provider } from 'react-redux';

import userReducer from '../../../../entities/user/model/userSlice';
import { baseApi } from '../../../../shared/api/baseApi';

import { ProfileSlider } from './ProfileSlider';

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

const meta = {
  title: 'Features/ProfileSlider',
  component: ProfileSlider,
  parameters: {
    layout: 'centered',
    docs: {
      description: {
        component:
          'A slider component for browsing through user profiles with swipe functionality.',
      },
    },
  },
  tags: ['autodocs'],
  decorators: [
    Story => (
      <Provider store={mockStore}>
        <div style={{ width: '100%', maxWidth: '350px', height: '600px' }}>
          <Story />
        </div>
      </Provider>
    ),
  ],
} satisfies Meta<typeof ProfileSlider>;

export default meta;
type Story = StoryObj<typeof ProfileSlider>;

// Since ProfileSlider has its own internal state and data,
// we don't need to pass any props to it
export const Default: Story = {};

// Note: The ProfileSlider component uses internal state and hardcoded profiles,
// so we can't easily customize it with different stories without modifying the component.
// In a real-world scenario, you might want to refactor the component to accept profiles as props
// to make it more flexible for testing and reuse.
