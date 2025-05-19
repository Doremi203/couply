import { Meta, StoryObj } from '@storybook/react';

import { MatchCard, MatchProfile } from './MatchCard';

const meta: Meta<typeof MatchCard> = {
  title: 'Pages/LikesPage/components/MatchCard',
  component: MatchCard,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
  decorators: [
    Story => (
      <div style={{ width: '350px', marginTop: '20px' }}>
        <Story />
      </div>
    ),
  ],
};

export default meta;
type Story = StoryObj<typeof MatchCard>;

//@ts-ignore
const sampleMatch: MatchProfile = {
  // @ts-ignore
  id: 1,
  name: 'Anna',
  age: 28,
  imageUrl: '/woman1.jpg',
  telegram: '@anna_28',
};

export const Default: Story = {
  args: {
    match: sampleMatch,
    onClick: match => console.log('Match clicked:', match),
    onSocialClick: (matchId, type) => console.log(`Social clicked: ${matchId}, ${type}`),
    showChatMessage: null,
  },
};

//TODO - delete later
export const WithChatMessage: Story = {
  args: {
    match: sampleMatch,
    onClick: match => console.log('Match clicked:', match),
    onSocialClick: (matchId, type) => console.log(`Social clicked: ${matchId}, ${type}`),
    showChatMessage: 1,
  },
};
