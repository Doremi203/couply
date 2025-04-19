import type { Meta, StoryObj } from '@storybook/react';

import { ActivityHeader } from './ActivityHeader';

const meta = {
  title: 'Widgets/ActivityHistory/ActivityHeader',
  component: ActivityHeader,
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
} satisfies Meta<typeof ActivityHeader>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    onBack: () => console.log('Back clicked'),
    title: 'Activity History',
  },
};

export const CustomTitle: Story = {
  args: {
    onBack: () => console.log('Back clicked'),
    title: 'Recent Activity',
  },
};
