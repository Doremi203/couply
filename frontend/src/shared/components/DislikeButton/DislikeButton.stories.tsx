import { Meta, StoryObj } from '@storybook/react';

import { DislikeButton } from './DislikeButton';

const meta: Meta<typeof DislikeButton> = {
  title: 'Shared/DislikeButton',
  component: DislikeButton,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof DislikeButton>;

export const Default: Story = {
  args: {
    onClick: () => console.log('DislikeButton clicked'),
  },
};

export const WithCustomClass: Story = {
  args: {
    onClick: () => console.log('DislikeButton clicked'),
    className: 'customDislikeButton',
  },
  decorators: [
    Story => (
      <div style={{ padding: '1rem' }}>
        <style>
          {`
            .customDislikeButton {
              width: 70px;
              height: 70px;
              background-color: #f5f5f5;
            }
          `}
        </style>
        <Story />
      </div>
    ),
  ],
};
