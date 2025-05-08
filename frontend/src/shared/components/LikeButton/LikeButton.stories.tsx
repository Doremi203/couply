import { Meta, StoryObj } from '@storybook/react';

import { LikeButton } from './LikeButton';

const meta: Meta<typeof LikeButton> = {
  title: 'Shared/LikeButton',
  component: LikeButton,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof LikeButton>;

export const Default: Story = {
  args: {
    onClick: () => console.log('LikeButton clicked'),
  },
};
