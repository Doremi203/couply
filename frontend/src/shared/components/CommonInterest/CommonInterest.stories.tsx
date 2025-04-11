import type { Meta, StoryObj } from '@storybook/react';

import { CommonInterest } from './CommonInterest';

const meta = {
  title: 'Components/CommonInterest',
  component: CommonInterest,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
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