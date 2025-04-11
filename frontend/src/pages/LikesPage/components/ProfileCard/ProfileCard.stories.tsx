import { Meta, StoryObj } from '@storybook/react';

import { ProfileCard } from './ProfileCard';

const meta: Meta<typeof ProfileCard> = {
  title: 'Pages/LikesPage/components/ProfileCard',
  component: ProfileCard,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof ProfileCard>;

const sampleProfile = {
  id: 1,
  name: 'Emma',
  age: 26,
  imageUrl: '/woman1.jpg',
  hasLikedYou: true,
  bio: 'Photographer and coffee enthusiast',
  location: 'New York',
  interests: ['Photography', 'Travel', 'Coffee'],
};

export const Default: Story = {
  args: {
    profile: sampleProfile,
    onClick: () => console.log('Profile card clicked'),
    onLike: (id) => console.log(`Profile ${id} liked`),
  },
};

export const WithoutLikeButton: Story = {
  args: {
    profile: sampleProfile,
    onClick: () => console.log('Profile card clicked'),
  },
};

export const WithCustomClass: Story = {
  args: {
    profile: sampleProfile,
    onClick: () => console.log('Profile card clicked'),
    onLike: (id) => console.log(`Profile ${id} liked`),
    className: 'customProfileCard',
  },
  decorators: [
    (Story) => (
      <div style={{ padding: '1rem' }}>
        <style>
          {`
            .customProfileCard {
              border: 2px solid #FF4B91;
              border-radius: 16px;
            }
          `}
        </style>
        <Story />
      </div>
    ),
  ],
};