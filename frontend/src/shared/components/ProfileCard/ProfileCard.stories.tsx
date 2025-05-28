import type { Meta, StoryObj } from '@storybook/react';

import {
  Gender,
  Goal,
  Zodiac,
  Education,
  Children,
  Alcohol,
  Smoking,
} from '../../../entities/user/api/constants';

import { ProfileCard } from './ProfileCard';

const meta = {
  title: 'Shared/ProfileCard',
  component: ProfileCard,
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
} satisfies Meta<typeof ProfileCard>;

export default meta;
type Story = StoryObj<typeof meta>;

const mockUserData = {
  id: '123',
  name: 'Анна',
  age: 28,
  gender: Gender.female,
  location: 'Москва',
  bio: 'Люблю путешествовать и пробовать новые блюда. Ищу человека для совместных приключений!',
  goal: Goal.relationship,
  interest: {
    sport: ['SPORT_RUNNING'] as [string],
    selfDevelopment: ['SELFDEVELOPMENT_LANGUAGES'] as [string],
    hobby: ['HOBBY_TRAVEL', 'HOBBY_PHOTOGRAPHY'] as [string, string],
    music: ['MUSIC_POP', 'MUSIC_ROCK'] as [string, string],
    moviesTv: ['MOVIESTV_COMEDY'] as [string],
    foodDrink: ['FOODDRINK_COFFEE'] as [string],
    personalityTraits: ['TRAIT_ADVENTUROUS'] as [string],
    pets: ['PETS_CATS'] as [string],
  },
  zodiac: Zodiac.leo,
  height: 168,
  education: Education.higher,
  children: Children.no,
  alcohol: Alcohol.neutrally,
  smoking: Smoking.negatively,
  isPremium: false,
  isBlocked: false,
  isVerified: true,
  isHidden: false,
  photos: [
    {
      orderNumber: 1,
      url: '/photo1.png',
    },
  ],
};

const mockLike = {
  senderId: '456',
  receiverId: '123',
  message: 'Привет! Как дела?',
};

export const Default: Story = {
  args: {
    //@ts-ignore
    profile: mockUserData,
    like: mockLike,
    onClick: () => console.log('Card clicked'),
    onLike: id => console.log('Liked user with id:', id),
  },
};

export const WithoutMessage: Story = {
  args: {
    //@ts-ignore
    profile: mockUserData,
    like: {
      ...mockLike,
      message: '',
    },
    onClick: () => console.log('Card clicked'),
    onLike: id => console.log('Liked user with id:', id),
  },
};

export const WithoutLikeButton: Story = {
  args: {
    //@ts-ignore
    profile: mockUserData,
    like: mockLike,
    onClick: () => console.log('Card clicked'),
  },
};

export const WithCustomClassName: Story = {
  args: {
    //@ts-ignore
    profile: mockUserData,
    like: mockLike,
    onClick: () => console.log('Card clicked'),
    onLike: id => console.log('Liked user with id:', id),
    className: 'custom-profile-card',
  },
};
