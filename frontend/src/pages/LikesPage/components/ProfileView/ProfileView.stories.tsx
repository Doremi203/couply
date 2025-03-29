import { Meta, StoryObj } from '@storybook/react';
import { ProfileView } from './ProfileView';

const meta: Meta<typeof ProfileView> = {
  title: 'Pages/LikesPage/components/ProfileView',
  component: ProfileView,
  parameters: {
    layout: 'fullscreen',
    backgrounds: {
      default: 'light',
    },
  },
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof ProfileView>;

const sampleProfile = {
  id: 1,
  name: 'Emma',
  age: 26,
  imageUrl: '/woman1.jpg',
  hasLikedYou: true,
  bio: 'Photographer and coffee enthusiast based in New York. Love to travel and capture moments.',
  location: 'New York',
  interests: ['Photography', 'Travel', 'Coffee', 'Art', 'Music'],
  lifestyle: {
    kids: "I don't have kids",
    pets: "Dog lover",
    drinking: "Social drinker",
    smoking: "Non-smoker",
  },
  passion: ['Photography', 'Travel', 'Coffee', 'Art', 'Music', 'Hiking', 'Yoga'],
  photos: [
    '/woman1.jpg',
    '/photo1.png',
    '/woman1.jpg',
    '/photo1.png',
    '/woman1.jpg',
    '/photo1.png',
  ],
};

export const Default: Story = {
  args: {
    profile: sampleProfile,
    onClose: () => console.log('Profile view closed'),
    onLike: (id) => console.log(`Profile ${id} liked`),
  },
  parameters: {
    viewport: {
      defaultViewport: 'mobile1',
    },
  },
};