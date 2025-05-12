import { Meta, StoryObj } from '@storybook/react';

import { MatchModal } from './MatchModal';

const meta: Meta<typeof MatchModal> = {
  title: 'Pages/LikesPage/components/MatchModal',
  component: MatchModal,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
  decorators: [
    Story => (
      <div style={{ width: '350px', maxWidth: '400px' }}>
        <Story />
      </div>
    ),
  ],
};

export default meta;
type Story = StoryObj<typeof MatchModal>;

export const Default: Story = {
  args: {
    userImage: '/woman1.jpg',
    matchImage: '/man1.jpg',
    // @ts-ignore
    onSendMessage: () => console.log('Send message clicked'),
    onKeepSwiping: () => console.log('Keep swiping clicked'),
  },
  parameters: {
    backgrounds: {
      default: 'dark',
    },
  },
};
