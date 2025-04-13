import type { Meta, StoryObj } from '@storybook/react';

import { EditProfile } from './EditProfile';

const meta = {
  title: 'Components/EditProfile',
  component: EditProfile,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
} satisfies Meta<typeof EditProfile>;

export default meta;
type Story = StoryObj<typeof meta>;

const mockProfileData = {
  name: 'Jane Doe',
  age: 28,
  phone: '+1 (555) 123-4567',
  dateOfBirth: '1997-05-15',
  email: 'jane.doe@example.com',
  gender: 'female',
  interests: ['Travel', 'Photography', 'Cooking'],
  about:
    'I love exploring new places and trying different cuisines. Photography is my passion, and I enjoy capturing moments during my travels.',
  music: ['Pop', 'Rock', 'Jazz'],
  movies: ['Inception', 'The Godfather', 'La La Land'],
  books: ['1984', 'To Kill a Mockingbird', 'The Great Gatsby'],
  hobbies: ['Hiking', 'Cooking', 'Reading'],
  isHidden: false,
  photos: ['/woman1.jpg', '/photo1.png'],
};

export const Default: Story = {
  args: {
    profileData: mockProfileData,
    onBack: () => console.log('Back button clicked'),
    onSave: () => console.log('Save button clicked'),
    onInputChange: (field, value) => console.log(`Field ${field} changed to ${value}`),
    onArrayInputChange: (field, value) => console.log(`Array field ${field} changed to ${value}`),
    onPhotoAdd: (_file, isAvatar) => console.log(`Photo added, isAvatar: ${isAvatar}`),
    onPhotoRemove: index => console.log(`Photo at index ${index} removed`),
  },
};
