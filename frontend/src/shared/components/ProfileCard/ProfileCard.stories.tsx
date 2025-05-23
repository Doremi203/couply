import type { Meta, StoryObj } from '@storybook/react';
import { MemoryRouter } from 'react-router-dom'; // Добавляем импорт

import { ProfileCard } from './ProfileCard';

const meta = {
  title: 'Shared/ProfileCard',
  component: ProfileCard,
  parameters: {
    layout: 'centered',
    docs: {
      description: {
        component:
          'A card component that displays a user profile with image, name, age, and optional like button.',
      },
    },
  },
  tags: ['autodocs'],
  decorators: [
    Story => (
      <MemoryRouter>
        {' '}
        {/* Добавляем обертку роутера */}
        <div style={{ width: '100%', maxWidth: '350px' }}>
          <Story />
        </div>
      </MemoryRouter>
    ),
  ],
} satisfies Meta<typeof ProfileCard>;

export default meta;
type Story = StoryObj<typeof meta>;

const sampleProfile = {
  user: {
    id: 1,
    name: 'Anna',
    age: 28,
    imageUrl: 'https://randomuser.me/api/portraits/women/68.jpg',
    location: 'Moscow, Russia',
    bio: 'Love hiking and photography',
    interests: ['hiking', 'photography', 'travel'],
    lifestyle: {
      exercise: 'Regular',
      drinking: 'Social',
      smoking: 'Never',
    },
    passion: ['Nature', 'Art'],
    photos: [
      'https://randomuser.me/api/portraits/women/68.jpg',
      'https://randomuser.me/api/portraits/women/69.jpg',
    ],
  },
};

export const Default: Story = {
  args: {
    //@ts-ignore
    profile: sampleProfile,
    onClick: () => console.log('Profile card clicked'),
  },
};

export const WithLikeButton: Story = {
  args: {
    //@ts-ignore
    profile: sampleProfile,
    onClick: () => console.log('Profile card clicked'),
    onLike: id => console.log(`Liked profile with id: ${id}`),
  },
};

export const HasLikedYou: Story = {
  args: {
    //@ts-ignore
    profile: {
      ...sampleProfile,
      //hasLikedYou: true,
    },
    onClick: () => console.log('Profile card clicked'),
    onLike: id => console.log(`Liked profile with id: ${id}`),
  },
};

export const WithCustomClass: Story = {
  args: {
    //@ts-ignore
    profile: sampleProfile,
    onClick: () => console.log('Profile card clicked'),
    className: 'custom-profile-card',
  },
};

export const MinimalProfile: Story = {
  args: {
    profile: {
      //@ts-ignore
      user: {
        id: 2,
        name: 'John',
        age: 32,
        imageUrl: 'https://randomuser.me/api/portraits/men/44.jpg',
      },
    },
    onClick: () => console.log('Profile card clicked'),
  },
};

// TODO
