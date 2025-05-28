import type { Meta, StoryObj } from '@storybook/react';

import {
  Gender,
  Goal,
  Zodiac,
  Education,
  Children,
  Alcohol,
  Smoking,
  Sport,
  Selfdevelopment,
  Hobby,
  Music,
  MoviesTV,
  FoodDrink,
  PersonalityTraits,
  Pets,
} from '../../../../entities/user/api/constants';

import { LikesSection } from './LikesSection';

const meta = {
  title: 'Features/Matches/LikesSection',
  component: LikesSection,
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
      <div style={{ width: '375px', marginTop: '40px' }}>
        <Story />
      </div>
    ),
  ],
} satisfies Meta<typeof LikesSection>;

export default meta;
type Story = StoryObj<typeof meta>;

// Create mock users with proper enum types
const mockUsers = [
  {
    id: '123',
    name: 'Анна',
    age: 28,
    gender: Gender.female,
    location: 'Москва',
    bio: 'Люблю путешествовать и пробовать новые блюда. Ищу человека для совместных приключений!',
    goal: Goal.relationship,
    interest: {
      sport: [Sport.running] as [Sport],
      selfDevelopment: [Selfdevelopment.languages] as [Selfdevelopment],
      hobby: [Hobby.travel] as [Hobby],
      music: [Music.pop] as [Music],
      moviesTv: [MoviesTV.comedy] as [MoviesTV],
      foodDrink: [FoodDrink.coffee] as [FoodDrink],
      personalityTraits: [PersonalityTraits.adventurous] as [PersonalityTraits],
      pets: [Pets.cats] as [Pets],
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
  },
  {
    id: '456',
    name: 'Мария',
    age: 25,
    gender: Gender.female,
    location: 'Санкт-Петербург',
    bio: 'Обожаю музыку и искусство. Ищу интересного собеседника.',
    goal: Goal.friendship,
    interest: {
      sport: [Sport.dancing] as [Sport],
      selfDevelopment: [Selfdevelopment.reading] as [Selfdevelopment],
      hobby: [Hobby.painting] as [Hobby],
      music: [Music.classical] as [Music],
      moviesTv: [MoviesTV.drama] as [MoviesTV],
      foodDrink: [FoodDrink.wine] as [FoodDrink],
      personalityTraits: [PersonalityTraits.creative] as [PersonalityTraits],
      pets: [Pets.dogs] as [Pets],
    },
    zodiac: Zodiac.gemini,
    height: 165,
    education: Education.higher,
    children: Children.no,
    alcohol: Alcohol.positively,
    smoking: Smoking.negatively,
    isPremium: true,
    isBlocked: false,
    isVerified: false,
    isHidden: false,
    photos: [
      {
        orderNumber: 1,
        url: '/cactus.jpg',
      },
    ],
  },
];

const mockLikes = [
  {
    senderId: '789',
    receiverId: '123',
    message: 'Привет! Как дела?',
  },
  {
    senderId: '101',
    receiverId: '456',
    message: '',
  },
];

export const WithLikes: Story = {
  args: {
    likes: [{ likes: mockLikes }],
    likesUsers: [{ users: mockUsers }],
    onProfileClick: profile => console.log('Profile clicked:', profile),
    onLike: id => console.log('Liked user with id:', id),
  },
};

export const Empty: Story = {
  args: {
    likes: [{ likes: [] }],
    likesUsers: [{ users: [] }],
    onProfileClick: profile => console.log('Profile clicked:', profile),
    onLike: id => console.log('Liked user with id:', id),
  },
};

export const WithMessage: Story = {
  args: {
    likes: [
      {
        likes: [
          {
            senderId: '789',
            receiverId: '123',
            message: 'Привет! Хотел бы познакомиться поближе. Как твои дела?',
          },
        ],
      },
    ],
    likesUsers: [{ users: [mockUsers[0]] }],
    onProfileClick: profile => console.log('Profile clicked:', profile),
    onLike: id => console.log('Liked user with id:', id),
  },
};

export const WithoutMessage: Story = {
  args: {
    likes: [
      {
        likes: [
          {
            senderId: '101',
            receiverId: '456',
            message: '',
          },
        ],
      },
    ],
    likesUsers: [{ users: [mockUsers[1]] }],
    onProfileClick: profile => console.log('Profile clicked:', profile),
    onLike: id => console.log('Liked user with id:', id),
  },
};
