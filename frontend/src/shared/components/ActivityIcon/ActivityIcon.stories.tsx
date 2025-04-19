import type { Meta, StoryObj } from '@storybook/react';

import { ActivityIcon } from './ActivityIcon';

const meta = {
  title: 'Shared/ActivityIcon',
  component: ActivityIcon,
  parameters: {
    layout: 'centered',
    docs: {
      description: {
        component: 'An icon that represents different types of activity (view, like, message).',
      },
    },
  },
  tags: ['autodocs'],
} satisfies Meta<typeof ActivityIcon>;

export default meta;
type Story = StoryObj<typeof meta>;

export const View: Story = {
  args: {
    type: 'view',
  },
};

export const Like: Story = {
  args: {
    type: 'like',
  },
};

export const Message: Story = {
  args: {
    type: 'message',
  },
};

export const Unknown: Story = {
  args: {
    type: 'unknown',
  },
};
