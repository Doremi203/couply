import type { Meta, StoryObj } from '@storybook/react';
import { BrowserRouter } from 'react-router-dom';
import { LikesPage } from './Page';

// Mock the required hooks and components
import * as matchesHooks from '../../../../features/matches';
import * as profileViewHooks from '../../hooks/useProfileView';
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

// Mock the useLikesAndMatches hook
(matchesHooks as any).useLikesAndMatches = () => ({
  showChatMessage: () => console.log('Show chat message'),
  handleLike: (id: string) => console.log('Like user with id:', id),
  handleSocialClick: () => console.log('Social click'),
  handleDislike: (id: string) => console.log('Dislike user with id:', id),
  matchesUsers: {
    users: [
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
    ],
  },
  likesUsers: [
    {
      users: [
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
      ],
    },
  ],
  likes: [
    {
      likes: [
        {
          senderId: '789',
          receiverId: '456',
          message: 'Привет! Как дела?',
        },
      ],
    },
  ],
});

// Mock the useProfileView hook
(profileViewHooks as any).useProfileView = () => ({
  selectedProfile: null,
  handleProfileClick: (profile: any) => console.log('Profile clicked:', profile),
  handleMatchClick: (match: any) => console.log('Match clicked:', match),
  handleCloseProfile: () => console.log('Close profile'),
});

// Mock NavBar component
import * as NavBarModule from '../../../../shared/components/NavBar';
(NavBarModule as any).NavBar = () => <div data-testid="navbar">NavBar</div>;

const meta = {
  title: 'Pages/LikesPage/Page',
  component: LikesPage,
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
        <div style={{ width: '375px', height: '700px', marginTop: '40px' }}>
          <Story />
        </div>
      </BrowserRouter>
    ),
  ],
} satisfies Meta<typeof LikesPage>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {};

// With selected profile
export const WithSelectedProfile: Story = {
  decorators: [
    Story => {
      // Override the useProfileView mock for this story
      (profileViewHooks as any).useProfileView = () => ({
        selectedProfile: {
          user: {
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
        },
        handleProfileClick: (profile: any) => console.log('Profile clicked:', profile),
        handleMatchClick: (match: any) => console.log('Match clicked:', match),
        handleCloseProfile: () => console.log('Close profile'),
      });
      return (
        <BrowserRouter>
          <div style={{ width: '375px', height: '700px', marginTop: '40px' }}>
            <Story />
          </div>
        </BrowserRouter>
      );
    },
  ],
};

// We can't easily mock the useState hook for the MatchesTab story in Storybook
// without using jest, so we'll create a simpler version that just shows
// what the matches tab would look like by modifying our hook mock
export const MatchesTab: Story = {
  decorators: [
    Story => {
      // Override the useLikesAndMatches hook to make the component render the matches tab
      const originalHook = matchesHooks.useLikesAndMatches;
      (matchesHooks as any).useLikesAndMatches = () => {
        const result = originalHook();
        // Force the activeTab to be 'мэтчи' by making the component think
        // it's already been set to that value
        (result as any).activeTab = 'мэтчи';
        return result;
      };

      return (
        <BrowserRouter>
          <div style={{ width: '375px', height: '700px', marginTop: '40px' }}>
            <Story />
          </div>
        </BrowserRouter>
      );
    },
  ],
};
