import type { Meta, StoryObj } from '@storybook/react';

import { CommonInterest } from './CommonInterest';

const meta = {
  title: 'Shared/CommonInterest',
  component: CommonInterest,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
  decorators: [
    Story => (
      <div style={{ width: '350px', marginTop: '20px' }}>
        <Story />
      </div>
    ),
  ],
} satisfies Meta<typeof CommonInterest>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    text: 'Music',
    isCommon: false,
  },
};

export const Common: Story = {
  args: {
    text: 'Travel',
    isCommon: true,
  },
};
