import type { Meta, StoryObj } from '@storybook/react';

import { Dislike } from './Dislike';

const meta = {
  title: 'Shared/Dislike',
  component: Dislike,
  tags: ['autodocs'],
  parameters: {
    // layout: 'centered',
    backgrounds: {
      default: 'dark',
      values: [{ name: 'dark', value: '#202C83' }],
    },
  },
} satisfies Meta<typeof Dislike>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {},
};
