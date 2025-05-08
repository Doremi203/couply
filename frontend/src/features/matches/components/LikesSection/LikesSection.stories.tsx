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
      <div style={{ width: '350px', marginTop: '20px' }}>
        <Story />
      </div>
    ),
  ],
} satisfies Meta<typeof LikesSection>;

export default meta;
type Story = StoryObj<typeof LikesSection>;

const sampleLikes = [
  {
    id: 1,
    name: 'Олег',
    age: 28,
    imageUrl: 'man1.jpg',
    hasLikedYou: true,
  },
  {
    id: 2,
    name: 'Анна',
    age: 32,
    imageUrl: 'woman1.jpg',
    hasLikedYou: true,
  },
  {
    id: 3,
    name: 'Анастасия',
    age: 26,
    imageUrl: 'woman1.jpg',
    hasLikedYou: true,
  },
  {
    id: 4,
    name: 'Павел',
    age: 30,
    imageUrl: 'man1.jpg',
    hasLikedYou: true,
  },
];

export const WithLikes: Story = {
  args: {
    // @ts-ignore
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
    // @ts-ignore
    likes: [sampleLikes[0]],
    onProfileClick: profile => console.log('Profile clicked:', profile),
    onLike: id => console.log(`Liked profile with id: ${id}`),
  },
};
