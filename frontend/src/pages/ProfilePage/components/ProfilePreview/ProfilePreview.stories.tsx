import type { Meta, StoryObj } from '@storybook/react';
import { ProfilePreview } from './ProfilePreview';
import { ProfileData } from '../../types';

const meta = {
  title: 'Components/ProfilePreview',
  component: ProfilePreview,
  parameters: {
    layout: 'fullscreen', // Use fullscreen layout since this is a full-page component
  },
  tags: ['autodocs'],
} satisfies Meta<typeof ProfilePreview>;

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
  photos: ['/woman1.jpg', '/photo1.png', '/man1.jpg'],
};

export const Default: Story = {
  args: {
    profileData: mockProfileData,
    onClose: () => console.log('Preview closed'),
  },
};