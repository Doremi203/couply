import type { Meta, StoryObj } from '@storybook/react';

import { LikesSection } from './LikesSection';

const meta = {
  title: 'Features/Likes/LikesSection',
  component: LikesSection,
  parameters: {
    layout: 'centered',
    docs: {
      description: {
        component: 'A section that displays profiles that have liked the user.',
      },
    },
  },
  tags: ['autodocs'],
  decorators: [
    Story => (
      <div style={{ width: '100%', maxWidth: '800px' }}>
        <Story />
      </div>
    ),
  ],
} satisfies Meta<typeof LikesSection>;

export default meta;
type Story = StoryObj<typeof LikesSection>;

// Sample likes data
const sampleLikes = [
  {
    id: 1,
    name: 'Anna Smith',
    age: 28,
    imageUrl: 'https://randomuser.me/api/portraits/women/68.jpg',
    hasLikedYou: true,
  },
  {
    id: 2,
    name: 'John Doe',
    age: 32,
    imageUrl: 'https://randomuser.me/api/portraits/men/44.jpg',
    hasLikedYou: true,
  },
  {
    id: 3,
    name: 'Maria Garcia',
    age: 26,
    imageUrl: 'https://randomuser.me/api/portraits/women/45.jpg',
    hasLikedYou: true,
  },
  {
    id: 4,
    name: 'Alex Johnson',
    age: 30,
    imageUrl: 'https://randomuser.me/api/portraits/men/22.jpg',
    hasLikedYou: true,
  },
];

export const WithLikes: Story = {
  args: {
    likes: sampleLikes,
    onProfileClick: profile => console.log('Profile clicked:', profile),
    onLike: id => console.log(`Liked profile with id: ${id}`),
  },
};

export const Empty: Story = {
  args: {
    likes: [],
    onProfileClick: profile => console.log('Profile clicked:', profile),
    onLike: id => console.log(`Liked profile with id: ${id}`),
  },
};

export const SingleLike: Story = {
  args: {
    likes: [sampleLikes[0]],
    onProfileClick: profile => console.log('Profile clicked:', profile),
    onLike: id => console.log(`Liked profile with id: ${id}`),
  },
};

export const ManyLikes: Story = {
  args: {
    likes: [
      ...sampleLikes,
      {
        id: 5,
        name: 'Sarah Williams',
        age: 27,
        imageUrl: 'https://randomuser.me/api/portraits/women/33.jpg',
        hasLikedYou: true,
      },
      {
        id: 6,
        name: 'Michael Brown',
        age: 31,
        imageUrl: 'https://randomuser.me/api/portraits/men/55.jpg',
        hasLikedYou: true,
      },
      {
        id: 7,
        name: 'Emily Davis',
        age: 25,
        imageUrl: 'https://randomuser.me/api/portraits/women/22.jpg',
        hasLikedYou: true,
      },
      {
        id: 8,
        name: 'David Wilson',
        age: 33,
        imageUrl: 'https://randomuser.me/api/portraits/men/33.jpg',
        hasLikedYou: true,
      },
    ],
    onProfileClick: profile => console.log('Profile clicked:', profile),
    onLike: id => console.log(`Liked profile with id: ${id}`),
  },
};
