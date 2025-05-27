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
  decorators: [
    Story => (
      <div style={{ width: '350px', marginTop: '40px' }}>
        <Story />
      </div>
    ),
  ],
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
  // @ts-ignore
  args: {
    profileData: mockProfileData,
    isVerified: true,
    //@ts-ignore
    onVerificationRequest: () => console.log('Verification requested'),
  },
};

export const Unverified: Story = {
  // @ts-ignore
  args: {
    profileData: mockProfileData,
    isVerified: false,
    //@ts-ignore
    onVerificationRequest: () => console.log('Verification requested'),
  },
};
