import { Meta, StoryObj } from '@storybook/react';

import { LikeButton } from './LikeButton';

const meta: Meta<typeof LikeButton> = {
  title: 'Shared/components/LikeButton',
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

export const WithCustomClass: Story = {
  args: {
    onClick: () => console.log('LikeButton clicked'),
    className: 'customLikeButton',
    likeClassName: 'customLike',
  },
  decorators: [
    (Story) => (
      <div style={{ padding: '1rem' }}>
        <style>
          {`
            .customLikeButton {
              width: 70px;
              height: 70px;
              background-color: #ffebf2;
            }
            .customLike {
              transform: scale(1.2);
            }
          `}
        </style>
        <Story />
      </div>
    ),
  ],
};