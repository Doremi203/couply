import type { Meta, StoryObj } from '@storybook/react';

import { ProfileData } from '../../types';

import { BasicInfoForm } from './BasicInfoForm';

const meta = {
  title: 'Features/ProfileEdit/BasicInfoForm',
  component: BasicInfoForm,
  parameters: {
    layout: 'centered',
    docs: {
      description: {
        component: 'A form for editing basic user profile information.',
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
} satisfies Meta<typeof BasicInfoForm>;

export default meta;
type Story = StoryObj<typeof BasicInfoForm>;

// Sample profile data
//@ts-ignore
const filledProfileData: ProfileData = {
  name: 'Anna Smith',
  age: 28,
  dateOfBirth: '1997-05-15',
  phone: '+7 (999) 123-4567',
  email: 'anna.smith@example.com',
  gender: 'female',
  about: 'I love hiking, photography, and exploring new places.',
  interests: ['hiking', 'photography', 'travel'],
  music: ['rock', 'indie', 'classical'],
  movies: ['sci-fi', 'drama', 'documentaries'],
  books: ['fiction', 'science', 'biographies'],
  hobbies: ['cooking', 'yoga', 'painting'],
  photos: [
    'https://randomuser.me/api/portraits/women/68.jpg',
    'https://randomuser.me/api/portraits/women/69.jpg',
  ],
  isHidden: false,
};

//@ts-ignore
const emptyProfileData: ProfileData = {
  name: '',
  age: 0,
  dateOfBirth: '',
  phone: '',
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
};

export const FilledForm: Story = {
  args: {
    profileData: filledProfileData,
    onInputChange: (field, value) => console.log(`Field ${field} changed to: ${value}`),
  },
};

export const EmptyForm: Story = {
  args: {
    profileData: emptyProfileData,
    onInputChange: (field, value) => console.log(`Field ${field} changed to: ${value}`),
  },
};

export const MaleProfile: Story = {
  args: {
    profileData: {
      ...filledProfileData,
      name: 'John Doe',
      gender: 'male',
    },
    onInputChange: (field, value) => console.log(`Field ${field} changed to: ${value}`),
  },
};

export const PartiallyFilled: Story = {
  args: {
    profileData: {
      ...emptyProfileData,
      name: 'Maria Garcia',
      age: 26,
      gender: 'female',
    },
    onInputChange: (field, value) => console.log(`Field ${field} changed to: ${value}`),
  },
};
