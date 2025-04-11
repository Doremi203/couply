import type { Meta, StoryObj } from '@storybook/react';

import { Dislike } from './Dislike';

const meta = {
  title: 'shared/components/Dislike',
  component: Dislike,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
} satisfies Meta<typeof Dislike>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {},
};