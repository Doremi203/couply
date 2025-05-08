import type { Meta, StoryObj } from '@storybook/react';

import { ProfileData } from '../../features/profileEdit';

import { EditProfile } from './EditProfile';

const meta = {
  title: 'Widgets/EditProfile',
  component: EditProfile,
  parameters: {
    layout: 'centered',
    docs: {
      description: {
        component: 'A widget for editing user profile information.',
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
} satisfies Meta<typeof EditProfile>;

export default meta;
type Story = StoryObj<typeof EditProfile>;

// Sample profile data
const sampleProfileData: ProfileData = {
  name: 'Anna Smith',
  age: 28,
  phone: '+7 (999) 123-4567',
  dateOfBirth: '1997-05-15',
  email: 'anna.smith@example.com',
  gender: 'female',
  about:
    'I love hiking, photography, and exploring new places. I work as a software engineer and enjoy solving complex problems.',
  interests: ['hiking', 'photography', 'travel'],
  music: ['rock', 'indie', 'classical'],
  movies: ['sci-fi', 'drama', 'documentaries'],
  books: ['fiction', 'science', 'biographies'],
  hobbies: ['cooking', 'yoga', 'painting'],
  photos: [
    'https://randomuser.me/api/portraits/women/68.jpg',
    'https://randomuser.me/api/portraits/women/69.jpg',
    'https://randomuser.me/api/portraits/women/70.jpg',
  ],
  isHidden: false,
};

export const Default: Story = {
  args: {
    profileData: sampleProfileData,
    onBack: () => console.log('Back button clicked'),
    onSave: () => console.log('Save button clicked'),
    onInputChange: (field, value) => console.log(`Field ${field} changed to: ${value}`),
    onArrayInputChange: (field, value) => console.log(`Array field ${field} changed to: ${value}`),
    onPhotoAdd: (_file, isAvatar) => console.log(`Photo added, isAvatar: ${isAvatar}`),
    onPhotoRemove: index => console.log(`Photo removed at index: ${index}`),
  },
};

export const EmptyProfile: Story = {
  args: {
    profileData: {
      name: '',
      age: 0,
      phone: '',
      dateOfBirth: '',
      email: '',
      gender: '',
      about: '',
      interests: [],
      music: [],
      movies: [],
      books: [],
      hobbies: [],
      photos: [],
      isHidden: false,
    },
    onBack: () => console.log('Back button clicked'),
    onSave: () => console.log('Save button clicked'),
    onInputChange: (field, value) => console.log(`Field ${field} changed to: ${value}`),
    onArrayInputChange: (field, value) => console.log(`Array field ${field} changed to: ${value}`),
    onPhotoAdd: (_file, isAvatar) => console.log(`Photo added, isAvatar: ${isAvatar}`),
    onPhotoRemove: index => console.log(`Photo removed at index: ${index}`),
  },
};

export const HiddenProfile: Story = {
  args: {
    profileData: {
      ...sampleProfileData,
      isHidden: true,
    },
    onBack: () => console.log('Back button clicked'),
    onSave: () => console.log('Save button clicked'),
    onInputChange: (field, value) => console.log(`Field ${field} changed to: ${value}`),
    onArrayInputChange: (field, value) => console.log(`Array field ${field} changed to: ${value}`),
    onPhotoAdd: (_file, isAvatar) => console.log(`Photo added, isAvatar: ${isAvatar}`),
    onPhotoRemove: index => console.log(`Photo removed at index: ${index}`),
  },
};
