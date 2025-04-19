import type { Meta, StoryObj } from '@storybook/react';
import { useRef } from 'react';

import { ProfileData } from '../../../../shared/components/ProfileCard';

import { ProfileInfo } from './ProfileInfo';

// Define the ProfileDetailsType interface
interface ProfileDetailsType {
  bio: string;
  location: string;
  lifestyle: { [key: string]: string };
  passion: string[];
  photos: string[];
}

// Define the props type for the wrapper component
interface WrapperProps {
  profile: ProfileData;
  profileDetails: ProfileDetailsType;
  menuPosition: 'collapsed' | 'expanded';
  handleToggleClick: () => void;
  handleTouchStart: (e: React.TouchEvent) => void;
  handleTouchMove: (e: React.TouchEvent) => void;
  handleTouchEnd: () => void;
  isCommonInterest: (interest: string) => boolean;
}

// Create a wrapper component to provide the necessary props and refs
const ProfileInfoWrapper = (props: WrapperProps) => {
  const profileInfoRef = useRef<HTMLDivElement>(null);

  return <ProfileInfo {...props} profileInfoRef={profileInfoRef} />;
};

const meta = {
  title: 'Widgets/ProfileView/ProfileInfo',
  component: ProfileInfoWrapper,
  parameters: {
    layout: 'centered',
    docs: {
      description: {
        component:
          'A component that displays detailed profile information including bio, interests, and photos.',
      },
    },
  },
  tags: ['autodocs'],
  decorators: [
    Story => (
      <div style={{ width: '100%', maxWidth: '500px' }}>
        <Story />
      </div>
    ),
  ],
} satisfies Meta<typeof ProfileInfoWrapper>;

export default meta;
type Story = StoryObj<typeof ProfileInfoWrapper>;

// Sample profile data
const sampleProfile = {
  id: 1,
  name: 'Anna Smith',
  age: 28,
  imageUrl: 'https://randomuser.me/api/portraits/women/68.jpg',
  hasLikedYou: false,
};

// Sample profile details
const sampleProfileDetails = {
  bio: 'I love hiking, photography, and exploring new places. I work as a software engineer and enjoy solving complex problems.',
  location: 'Moscow, Russia',
  lifestyle: {
    exercise: 'Regular',
    drinking: 'Social',
    smoking: 'Never',
  },
  passion: ['Music', 'Travel', 'Photography', 'Art', 'Fashion'],
  photos: [
    'https://randomuser.me/api/portraits/women/68.jpg',
    'https://randomuser.me/api/portraits/women/69.jpg',
    'https://randomuser.me/api/portraits/women/70.jpg',
  ],
};

// Common interests checker function
const isCommonInterest = (interest: string) => {
  const commonInterests = ['Music', 'Travel', 'Photography'];
  return commonInterests.includes(interest);
};

// Mock touch event handlers
const handleTouchStart = (_e: React.TouchEvent) => {};
const handleTouchMove = (_e: React.TouchEvent) => {};
const handleTouchEnd = () => {};

export const Collapsed: Story = {
  args: {
    profile: sampleProfile,
    profileDetails: sampleProfileDetails,
    menuPosition: 'collapsed',
    handleToggleClick: () => console.log('Toggle clicked'),
    handleTouchStart,
    handleTouchMove,
    handleTouchEnd,
    isCommonInterest,
  },
};

export const Expanded: Story = {
  args: {
    profile: sampleProfile,
    profileDetails: sampleProfileDetails,
    menuPosition: 'expanded',
    handleToggleClick: () => console.log('Toggle clicked'),
    handleTouchStart,
    handleTouchMove,
    handleTouchEnd,
    isCommonInterest,
  },
};

export const MinimalProfile: Story = {
  args: {
    profile: {
      id: 2,
      name: 'John Doe',
      age: 32,
      imageUrl: 'https://randomuser.me/api/portraits/men/44.jpg',
      hasLikedYou: false,
    },
    profileDetails: {
      bio: 'Just a simple bio.',
      location: 'New York, USA',
      lifestyle: {
        status: 'Single',
      },
      passion: ['Sports', 'Movies'],
      photos: ['https://randomuser.me/api/portraits/men/44.jpg'],
    },
    menuPosition: 'collapsed',
    handleToggleClick: () => console.log('Toggle clicked'),
    handleTouchStart,
    handleTouchMove,
    handleTouchEnd,
    isCommonInterest,
  },
};
