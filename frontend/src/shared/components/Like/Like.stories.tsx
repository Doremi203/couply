import type { Meta, StoryObj } from '@storybook/react';

import { Like } from './Like';

const meta = {
  title: 'Shared/Like',
  component: Like,
  parameters: {
    // layout: 'centered',
    backgrounds: {
      default: 'dark',
      values: [{ name: 'dark', value: '#202C83' }],
    },
  },
  tags: ['autodocs'],
} satisfies Meta<typeof Like>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {},
};
