import { configureStore } from '@reduxjs/toolkit';
import type { Meta, StoryObj } from '@storybook/react';
import { Provider } from 'react-redux';
import { MemoryRouter } from 'react-router-dom';

import userReducer from '../../../../entities/user/model/userSlice';
import { baseApi } from '../../../../shared/api/baseApi';

import { ProfileSlider } from './ProfileSlider';

// Фикс для модальных окон
const ModalRoot = () => <div id="modal-root" style={{ position: 'relative', zIndex: 999 }} />;

const mockStore = configureStore({
  reducer: {
    [baseApi.reducerPath]: baseApi.reducer,
    user: userReducer,
  },
  middleware: getDefaultMiddleware =>
    getDefaultMiddleware({
      serializableCheck: false, // Отключаем проверку для Storybook
    }).concat(baseApi.middleware),
  preloadedState: {
    user: {
      id: 'mock-user-id',
      isAuthenticated: true,
      profile: null,
      status: 'idle',
      error: null,
      // Добавляем все обязательные поля из userSlice
    },
    [baseApi.reducerPath]: {
      queries: {},
      mutations: {},
      provided: {},
      subscriptions: {},
      config: baseApi.reducerPath,
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
        <MemoryRouter>
          <ModalRoot /> {/* Добавляем корень для модалок */}
          <div
            style={{
              maxWidth: '350px',
              marginTop: '20px',
              height: '600px',
              position: 'relative',
            }}
          >
            <Story />
          </div>
        </MemoryRouter>
      </Provider>
    ),
  ],
} satisfies Meta<typeof ProfileSlider>;

export default meta;
type Story = StoryObj<typeof ProfileSlider>;

export const Default: Story = {};
