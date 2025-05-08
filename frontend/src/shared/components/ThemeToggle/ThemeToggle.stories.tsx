import type { Meta, StoryObj } from '@storybook/react';

import { ThemeToggle } from './ThemeToggle';

const meta = {
  title: 'Shared/ThemeToggle',
  component: ThemeToggle,
  parameters: {
    layout: 'centered',
    docs: {
      description: {
        component: 'A toggle button that switches between light and dark themes.',
      },
    },
  },
  tags: ['autodocs'],
  decorators: [
    Story => (
      <div style={{ padding: '20px' }}>
        <Story />
      </div>
    ),
  ],
} satisfies Meta<typeof ThemeToggle>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    className: '',
  },
};
