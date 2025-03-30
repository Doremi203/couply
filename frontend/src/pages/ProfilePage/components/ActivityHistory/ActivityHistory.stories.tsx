import type { Meta, StoryObj } from '@storybook/react';
import { ActivityHistory } from './ActivityHistory';

const meta = {
  title: 'Components/ActivityHistory',
  component: ActivityHistory,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
} satisfies Meta<typeof ActivityHistory>;

export default meta;
type Story = StoryObj<typeof meta>;

const mockActivityHistory = [
  {
    type: 'view',
    user: 'Alex Smith',
    date: '2025-03-30T12:00:00Z',
  },
  {
    type: 'like',
    user: 'Emma Johnson',
    date: '2025-03-29T15:30:00Z',
  },
  {
    type: 'message',
    user: 'Michael Brown',
    date: '2025-03-28T09:45:00Z',
  },
];

const formatDate = (dateString: string) => {
  const date = new Date(dateString);
  return date.toLocaleDateString('en-US', {
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  });
};

export const Default: Story = {
  args: {
    activityHistory: mockActivityHistory,
    onBack: () => console.log('Back button clicked'),
    formatDate: formatDate,
  },
};