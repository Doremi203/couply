import type { Meta, StoryObj } from '@storybook/react';

import { ActivityList } from './ActivityList';

const meta = {
  title: 'Widgets/ActivityHistory/ActivityList',
  component: ActivityList,
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
} satisfies Meta<typeof ActivityList>;

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

export const Default: Story = {
  args: {
    activities: [
      {
        type: 'view',
        user: 'Anna Smith',
        date: '2025-04-19T15:30:00',
      },
      {
        type: 'like',
        user: 'John Doe',
        date: '2025-04-19T14:45:00',
      },
      {
        type: 'message',
        user: 'Maria Garcia',
        date: '2025-04-19T12:15:00',
      },
    ],
    formatDate,
  },
};

export const Empty: Story = {
  args: {
    activities: [],
    formatDate,
  },
};

export const ManyItems: Story = {
  args: {
    activities: [
      {
        type: 'view',
        user: 'Anna Smith',
        date: '2025-04-19T15:30:00',
      },
      {
        type: 'like',
        user: 'John Doe',
        date: '2025-04-19T14:45:00',
      },
      {
        type: 'message',
        user: 'Maria Garcia',
        date: '2025-04-19T12:15:00',
      },
      {
        type: 'view',
        user: 'Alex Johnson',
        date: '2025-04-18T18:20:00',
      },
      {
        type: 'like',
        user: 'Sarah Williams',
        date: '2025-04-18T16:10:00',
      },
      {
        type: 'message',
        user: 'Michael Brown',
        date: '2025-04-18T10:05:00',
      },
    ],
    formatDate,
  },
};
