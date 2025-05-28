import type { Meta, StoryObj } from '@storybook/react';

import { ProfileInfo } from './ProfileInfo';

const meta = {
  title: 'Widgets/ProfileView/ProfileInfo',
  component: ProfileInfo,
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
      <div style={{ width: '350px', marginTop: '40px' }}>
        <Story />
      </div>
    ),
  ],
} satisfies Meta<typeof ProfileInfo>;

export default meta;
type Story = StoryObj<typeof meta>;

const mockProfile = {
  id: 1,
  name: 'Anna',
  age: 28,
  bio: 'I love exploring new places and trying new cuisines. Looking for someone to share adventures with!',
  location: 'Moscow',
  latitude: 55.7558,
  longitude: 37.6173,
  interests: ['Travel', 'Photography', 'Cooking'],
  children: 'dont_want',
  education: 'higher',
  alcohol: 'socially',
  smoking: 'never',
  zodiac: 'leo',
  goal: 'relationship',
  photos: ['/woman1.jpg', '/photo1.png'],
};

const mockProfileDetails = {
  bio: mockProfile.bio,
  location: mockProfile.location,
  lifestyle: {},
  passion: [],
  photos: mockProfile.photos,
};

const isCommonInterest = (interest: string) => ['Travel', 'Photography'].includes(interest);

export const Complete: Story = {
  args: {
    profile: mockProfile,
    profileDetails: mockProfileDetails,
    isCommonInterest,
  },
};

export const Minimal: Story = {
  args: {
    profile: {
      name: 'Anna',
      age: 28,
    },
    profileDetails: {
      bio: '',
      location: '',
      lifestyle: {},
      passion: [],
      photos: [],
    },
    isCommonInterest: () => false,
  },
};

export const WithLocationError: Story = {
  args: {
    profile: {
      ...mockProfile,
      latitude: undefined,
      longitude: undefined,
      location: undefined,
    },
    profileDetails: mockProfileDetails,
    isCommonInterest,
  },
};
