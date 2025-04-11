import type { Meta, StoryObj } from '@storybook/react';

import { ProfileData } from '../../types';

import { ProfileInfo } from './ProfileInfo';

const meta = {
  title: 'Components/ProfileInfo',
  component: ProfileInfo,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
} satisfies Meta<typeof ProfileInfo>;

export default meta;
type Story = StoryObj<typeof meta>;

const mockProfileData: ProfileData = {
  name: 'Anna',
  age: 28,
  phone: '+7 999 123 45 67',
  dateOfBirth: '1997-05-15',
  email: 'anna@example.com',
  gender: 'female',
  interests: ['Travel', 'Photography', 'Cooking'],
  about: 'I love exploring new places and trying new cuisines.',
  music: ['Pop', 'Rock', 'Jazz'],
  movies: ['Drama', 'Comedy'],
  books: ['Fiction', 'Biography'],
  hobbies: ['Hiking', 'Cooking', 'Photography'],
  isHidden: false,
  photos: ['/woman1.jpg', '/photo1.png'],
};

export const Verified: Story = {
  args: {
    profileData: mockProfileData,
    isVerified: true,
    onVerificationRequest: () => console.log('Verification requested'),
  },
};

export const Unverified: Story = {
  args: {
    profileData: mockProfileData,
    isVerified: false,
    onVerificationRequest: () => console.log('Verification requested'),
  },
};