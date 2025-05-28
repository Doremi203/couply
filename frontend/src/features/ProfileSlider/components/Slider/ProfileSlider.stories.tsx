import { configureStore } from '@reduxjs/toolkit';
import type { Meta, StoryObj } from '@storybook/react';
import { Provider } from 'react-redux';

// Mock the required hooks and APIs
import * as matchesApi from '../../../../entities/matches';
import * as searchApi from '../../../../entities/search/api/searchApi';
import * as subscriptionApi from '../../../../entities/subscription/api/subscriptionApi';
import * as userActions from '../../../../entities/user';

import { ProfileSlider } from './ProfileSlider';

// Mock Redux store
const mockStore = configureStore({
  reducer: {
    user: (state = {}) => state,
    matches: (state = {}) => state,
  },
});

// Mock API hooks
(matchesApi as any).useLikeUserMutation = () => [
  () => Promise.resolve({ unwrap: () => ({}) }),
  { isLoading: false },
];

(searchApi as any).useCreateFilterMutation = () => [
  () => Promise.resolve({ unwrap: () => ({}) }),
  { isLoading: false },
];

(searchApi as any).useSearchUsersMutation = () => [
  () =>
    Promise.resolve({
      unwrap: () => ({
        usersSearchInfo: [
          {
            user: {
              id: '123',
              name: 'Анна',
              age: 28,
              bio: 'Люблю путешествовать и пробовать новые блюда. Ищу человека для совместных приключений!',
              verified: true,
              interests: ['Путешествия', 'Фотография', 'Кулинария'],
              photos: [
                { url: '/photo1.png', orderNumber: 1 },
                { url: '/cactus.jpg', orderNumber: 2 },
              ],
              distanceToUser: 5,
            },
            distanceToUser: 5,
          },
          {
            user: {
              id: '456',
              name: 'Мария',
              age: 25,
              bio: 'Обожаю музыку и искусство. Ищу интересного собеседника.',
              verified: false,
              interests: ['Музыка', 'Искусство', 'Театр'],
              photos: [{ url: '/cactus2.jpg', orderNumber: 1 }],
              distanceToUser: 10,
            },
            distanceToUser: 10,
          },
        ],
      }),
    }),
  { isLoading: false },
];

(subscriptionApi as any).useGetSubscriptionMutation = () => [
  () =>
    Promise.resolve({
      unwrap: () => ({ status: 'SUBSCRIPTION_STATUS_ACTIVE', subscriptionId: '123' }),
    }),
  { isLoading: false },
];

// Mock user actions
(userActions as any).setUserVerified = () => ({ type: 'user/setUserVerified' });

const meta = {
  title: 'Features/ProfileSlider',
  component: ProfileSlider,
  parameters: {
    layout: 'centered',
    screenshot: {
      viewport: {
        width: 375,
        height: 812,
      },
    },
  },
  tags: ['autodocs'],
  decorators: [
    Story => (
      <Provider store={mockStore}>
        <div style={{ width: '375px', height: '700px', marginTop: '40px' }}>
          <Story />
        </div>
      </Provider>
    ),
  ],
} satisfies Meta<typeof ProfileSlider>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {};

// Mock for no users left scenario
export const NoUsersLeft: Story = {
  decorators: [
    Story => {
      // Override the search users mock for this story
      (searchApi as any).useSearchUsersMutation = () => [
        () =>
          Promise.resolve({
            unwrap: () => ({
              usersSearchInfo: [],
            }),
          }),
        { isLoading: false },
      ];
      return (
        <Provider store={mockStore}>
          <div style={{ width: '375px', height: '700px', marginTop: '40px' }}>
            <Story />
          </div>
        </Provider>
      );
    },
  ],
};

// Mock for loading state
export const Loading: Story = {
  decorators: [
    Story => {
      // Override the search users mock for this story to simulate loading
      (searchApi as any).useSearchUsersMutation = () => [
        () => new Promise(resolve => setTimeout(resolve, 10000)),
        { isLoading: true },
      ];
      return (
        <Provider store={mockStore}>
          <div style={{ width: '375px', height: '700px', marginTop: '40px' }}>
            <Story />
          </div>
        </Provider>
      );
    },
  ],
};

// Mock for non-premium user
export const NonPremiumUser: Story = {
  decorators: [
    Story => {
      // Override the subscription API mock for this story
      (subscriptionApi as any).useGetSubscriptionMutation = () => [
        () =>
          Promise.resolve({
            unwrap: () => ({ status: 'SUBSCRIPTION_STATUS_INACTIVE', subscriptionId: '123' }),
          }),
        { isLoading: false },
      ];
      return (
        <Provider store={mockStore}>
          <div style={{ width: '375px', height: '700px', marginTop: '40px' }}>
            <Story />
          </div>
        </Provider>
      );
    },
  ],
};
