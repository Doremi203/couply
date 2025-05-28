import type { Meta, StoryObj } from '@storybook/react';

import { NotificationSettings } from './NotificationSettings';

// Mock the required hooks and context
import * as SubscriptionApi from '../../../../entities/subscription/api/subscriptionApi';
import * as ThemeContext from '../../../../shared/lib/context/ThemeContext';
import * as PushNotificationHook from '../../../../shared/lib/hooks/usePushNotificationPermission';
import * as PushSubscriptionHook from '../../../../shared/lib/hooks/usePushSubscription';

// Mock ThemeContext
(ThemeContext as any).useTheme = () => ({
  theme: 'light',
  toggleTheme: () => console.log('Theme toggled'),
});

// Mock usePushNotificationPermission
(PushNotificationHook as any).usePushNotificationPermission = () => ({
  permission: 'default',
  requestPermission: () => Promise.resolve('granted'),
});

// Mock usePushSubscription
(PushSubscriptionHook as any).usePushSubscription = () => ({
  subscription: null,
  subscribe: () => Promise.resolve({ endpoint: 'https://example.com' }),
  unsubscribe: () => Promise.resolve(),
});

// Mock subscription API hooks
(SubscriptionApi as any).useGetSubscriptionMutation = () => [
  () =>
    Promise.resolve({
      unwrap: () => ({ status: 'SUBSCRIPTION_STATUS_ACTIVE', subscriptionId: '123' }),
    }),
  { isLoading: false },
];

(SubscriptionApi as any).useCancelSubscriptionMutation = () => [
  () => Promise.resolve({ unwrap: () => ({}) }),
  { isLoading: false },
];

// Mock PushNotificationService
import * as PushNotificationService from '../../../../shared/lib/services/PushNotificationService';
(PushNotificationService as any).sendSubscriptionToServer = () => Promise.resolve();
(PushNotificationService as any).unsubscribeFromPushNotifications = () => Promise.resolve();

const meta = {
  title: 'Pages/SettingsPage/NotificationSettings',
  component: NotificationSettings,
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
      <div style={{ width: '375px', marginTop: '40px' }}>
        <Story />
      </div>
    ),
  ],
} satisfies Meta<typeof NotificationSettings>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    className: '',
  },
};

// Light theme variant
export const LightTheme: Story = {
  args: {
    className: '',
  },
  decorators: [
    Story => {
      // Override the theme mock for this story
      (ThemeContext as any).useTheme = () => ({
        theme: 'light',
        toggleTheme: () => console.log('Theme toggled'),
      });
      return <Story />;
    },
  ],
};

// Dark theme variant
export const DarkTheme: Story = {
  args: {
    className: '',
  },
  decorators: [
    Story => {
      // Override the theme mock for this story
      (ThemeContext as any).useTheme = () => ({
        theme: 'dark',
        toggleTheme: () => console.log('Theme toggled'),
      });
      return <Story />;
    },
  ],
};

// With push notifications enabled
export const WithPushEnabled: Story = {
  args: {
    className: '',
  },
  decorators: [
    Story => {
      // Override the push subscription mock for this story
      (PushSubscriptionHook as any).usePushSubscription = () => ({
        subscription: { endpoint: 'https://example.com' },
        subscribe: () => Promise.resolve({ endpoint: 'https://example.com' }),
        unsubscribe: () => Promise.resolve(),
      });
      return <Story />;
    },
  ],
};

// Without premium subscription
export const WithoutPremium: Story = {
  args: {
    className: '',
  },
  decorators: [
    Story => {
      // Override the subscription API mock for this story
      (SubscriptionApi as any).useGetSubscriptionMutation = () => [
        () =>
          Promise.resolve({
            unwrap: () => ({ status: 'SUBSCRIPTION_STATUS_INACTIVE', subscriptionId: '123' }),
          }),
        { isLoading: false },
      ];
      return <Story />;
    },
  ],
};
