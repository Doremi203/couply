import type { Meta, StoryObj } from '@storybook/react';
import { BrowserRouter } from 'react-router-dom';
import { Profile } from './Profile';

const meta = {
  title: 'Pages/ProfilePage/Profile',
  component: Profile,
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
      <BrowserRouter>
        <div style={{ width: '375px', marginTop: '40px' }}>
          <Story />
        </div>
      </BrowserRouter>
    ),
  ],
} satisfies Meta<typeof Profile>;

export default meta;
type Story = StoryObj<typeof meta>;

const mockProfileData = {
  name: 'Анна',
  age: 28,
  phone: '+7 (999) 123-45-67',
  dateOfBirth: '1997-05-15',
  email: 'anna@example.com',
  gender: 'GENDER_FEMALE',
  interests: ['Путешествия', 'Фотография', 'Кулинария'],
  about: 'Люблю путешествовать и пробовать новые блюда. Ищу человека для совместных приключений!',
  music: ['Поп', 'Рок'],
  movies: ['Комедии', 'Драмы'],
  books: ['Фантастика', 'Детективы'],
  hobbies: ['Фотография', 'Готовка'],
  isHidden: false,
  photos: ['/photo1.png', '/cactus.jpg'],
  bio: 'Люблю путешествовать и пробовать новые блюда.',
};

export const Default: Story = {
  args: {
    profileData: mockProfileData,
    isVerified: true,
    isProfileHidden: false,
    onEditToggle: () => console.log('Edit toggled'),
    onVisibilityToggle: () => console.log('Visibility toggled'),
    onActivityClick: () => console.log('Activity clicked'),
    onPreviewClick: () => console.log('Preview clicked'),
    onVerificationRequest: () => console.log('Verification requested'),
    isPremium: false,
  },
};

export const Hidden: Story = {
  args: {
    profileData: mockProfileData,
    isVerified: true,
    isProfileHidden: true,
    onEditToggle: () => console.log('Edit toggled'),
    onVisibilityToggle: () => console.log('Visibility toggled'),
    onActivityClick: () => console.log('Activity clicked'),
    onPreviewClick: () => console.log('Preview clicked'),
    onVerificationRequest: () => console.log('Verification requested'),
    isPremium: false,
  },
};

export const NotVerified: Story = {
  args: {
    profileData: mockProfileData,
    isVerified: false,
    isProfileHidden: false,
    onEditToggle: () => console.log('Edit toggled'),
    onVisibilityToggle: () => console.log('Visibility toggled'),
    onActivityClick: () => console.log('Activity clicked'),
    onPreviewClick: () => console.log('Preview clicked'),
    onVerificationRequest: () => console.log('Verification requested'),
    isPremium: false,
  },
};

export const Premium: Story = {
  args: {
    profileData: mockProfileData,
    isVerified: true,
    isProfileHidden: false,
    onEditToggle: () => console.log('Edit toggled'),
    onVisibilityToggle: () => console.log('Visibility toggled'),
    onActivityClick: () => console.log('Activity clicked'),
    onPreviewClick: () => console.log('Preview clicked'),
    onVerificationRequest: () => console.log('Verification requested'),
    isPremium: true,
  },
};
