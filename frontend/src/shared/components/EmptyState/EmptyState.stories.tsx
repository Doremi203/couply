import type { Meta, StoryObj } from '@storybook/react';

import { EmptyState } from './EmptyState';

const meta = {
  title: 'Shared/EmptyState',
  component: EmptyState,
  parameters: {
    layout: 'centered',
    docs: {
      description: {
        component: 'A component to display when there is no content to show.',
      },
    },
  },
  tags: ['autodocs'],
} satisfies Meta<typeof EmptyState>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    title: 'No matches yet',
    subtitle: 'Keep swiping to find your match!',
  },
};

export const NoLikes: Story = {
  args: {
    title: 'No likes yet',
    subtitle: 'When someone likes you, they will appear here.',
  },
};

export const NoMessages: Story = {
  args: {
    title: 'No messages',
    subtitle: 'Start a conversation with your matches!',
  },
};

export const ShortText: Story = {
  args: {
    title: 'Empty',
    subtitle: 'Nothing to see here',
  },
};

export const LongText: Story = {
  args: {
    title: 'No activity in your feed',
    subtitle:
      'Your activity feed shows who viewed, liked, or messaged you. Check back later for updates!',
  },
};
