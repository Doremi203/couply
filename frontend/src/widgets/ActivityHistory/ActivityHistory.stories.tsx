import type { Meta, StoryObj } from '@storybook/react';

import { ActivityHistory } from './ActivityHistory';
import { ActivityItem } from './types';

const meta = {
  title: 'Widgets/ActivityHistory',
  component: ActivityHistory,
  parameters: {
    layout: 'centered',
    docs: {
      description: {
        component: "A widget that displays a user's activity history.",
      },
    },
  },
  tags: ['autodocs'],
  decorators: [
    Story => (
      <div style={{ width: '100%', maxWidth: '400px' }}>
        <Story />
      </div>
    ),
  ],
} satisfies Meta<typeof ActivityHistory>;

export default meta;
type Story = StoryObj<typeof ActivityHistory>;

// Sample activity history data
const sampleActivities: ActivityItem[] = [
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
];

// Format date function
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
    activityHistory: sampleActivities,
    onBack: () => console.log('Back button clicked'),
    formatDate,
  },
};

export const Empty: Story = {
  args: {
    activityHistory: [],
    onBack: () => console.log('Back button clicked'),
    formatDate,
  },
};

export const ManyItems: Story = {
  args: {
    activityHistory: [
      ...sampleActivities,
      {
        type: 'message',
        user: 'Michael Brown',
        date: '2025-04-18T10:05:00',
      },
      {
        type: 'view',
        user: 'Emily Davis',
        date: '2025-04-17T20:30:00',
      },
      {
        type: 'like',
        user: 'David Wilson',
        date: '2025-04-17T15:45:00',
      },
      {
        type: 'message',
        user: 'Olivia Martinez',
        date: '2025-04-17T09:20:00',
      },
      {
        type: 'view',
        user: 'James Taylor',
        date: '2025-04-16T22:10:00',
      },
    ],
    onBack: () => console.log('Back button clicked'),
    formatDate,
  },
};
