import type { Meta, StoryObj } from '@storybook/react';

import { ProfileView } from './ProfileView';

const meta = {
  title: 'Widgets/ProfileView',
  component: ProfileView,
  parameters: {
    layout: 'centered',
    docs: {
      description: {
        component:
          'A widget that displays a detailed view of a user profile with photo and information.',
      },
    },
  },
  tags: ['autodocs'],
  decorators: [
    Story => (
      <div style={{ width: '100%', maxWidth: '500px', height: '800px' }}>
        <Story />
      </div>
    ),
  ],
} satisfies Meta<typeof ProfileView>;

export default meta;
type Story = StoryObj<typeof ProfileView>;

// Sample profile data
const sampleProfile = {
  user: {
    id: 1,
    name: 'Anna Smith',
    age: 28,
    imageUrl: 'man1.jpg',
    location: 'Moscow, Russia',
    bio: 'I love hiking, photography, and exploring new places. I work as a software engineer and enjoy solving complex problems.',
    interests: ['Hiking', 'Photography', 'Travel', 'Music', 'Art'],
    hasLikedYou: false,
    lifestyle: {
      exercise: 'Regular',
      drinking: 'Social',
      smoking: 'Never',
    },
    passion: ['Music', 'Travel', 'Photography', 'Art', 'Fashion'],
    photos: ['man1.jpg'],
  },
};

export const Default: Story = {
  args: {
    profile: sampleProfile,
    onClose: () => console.log('Profile view closed'),
    onLike: id => console.log(`Liked profile with id: ${id}`),
  },
};

export const HasLikedYou: Story = {
  args: {
    profile: {
      ...sampleProfile,
      // user: {hasLikedYou: true},
    },
    onClose: () => console.log('Profile view closed'),
    onLike: id => console.log(`Liked profile with id: ${id}`),
  },
};

export const MinimalProfile: Story = {
  args: {
    profile: {
      user: {
        id: 2,
        name: 'John Doe',
        age: 32,
        imageUrl: 'https://randomuser.me/api/portraits/men/44.jpg',
        hasLikedYou: false,
      },
    },
    onClose: () => console.log('Profile view closed'),
    onLike: id => console.log(`Liked profile with id: ${id}`),
  },
};

//TODO
