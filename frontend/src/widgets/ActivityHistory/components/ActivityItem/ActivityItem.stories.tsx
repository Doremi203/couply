import type { Meta, StoryObj } from '@storybook/react';

import { ActivityItem } from './ActivityItem';

const meta = {
  title: 'Widgets/ActivityHistory/ActivityItem',
  component: ActivityItem,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
  decorators: [
    Story => (
      <div style={{ width: '100%', maxWidth: '400px' }}>
        <Story />
      </div>
    ),
  ],
} satisfies Meta<typeof ActivityItem>;

export default meta;
type Story = StoryObj<typeof meta>;

const formatDate = (dateString: string) => {
  const date = new Date(dateString);
  return new Intl.DateTimeFormat('en-US', {
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  }).format(date);
};

export const ViewActivity: Story = {
  args: {
    activity: {
      type: 'view',
      user: 'Anna Smith',
      date: '2025-04-19T15:30:00',
    },
    formatDate,
  },
};

export const LikeActivity: Story = {
  args: {
    activity: {
      type: 'like',
      user: 'John Doe',
      date: '2025-04-19T14:45:00',
    },
    formatDate,
  },
};

export const MessageActivity: Story = {
  args: {
    activity: {
      type: 'message',
      user: 'Maria Garcia',
      date: '2025-04-19T12:15:00',
    },
    formatDate,
  },
};
