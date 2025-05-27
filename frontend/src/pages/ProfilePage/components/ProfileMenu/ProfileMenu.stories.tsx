import type { Meta, StoryObj } from '@storybook/react';
import { BrowserRouter } from 'react-router-dom';
import { ProfileMenu } from './ProfileMenu';

// Mock the useDeleteUserMutation hook
import * as userHooks from '../../../../entities/user';
// Override the hook implementation for Storybook
(userHooks as any).useDeleteUserMutation = () => [() => Promise.resolve(), { isLoading: false }];

const meta = {
  title: 'Pages/ProfilePage/ProfileMenu',
  component: ProfileMenu,
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
      <BrowserRouter>
        <div style={{ width: '375px', marginTop: '40px' }}>
          <Story />
        </div>
      </BrowserRouter>
    ),
  ],
} satisfies Meta<typeof ProfileMenu>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    onEditProfileClick: () => console.log('Edit profile clicked'),
    onSettingsClick: () => console.log('Settings clicked'),
  },
};
