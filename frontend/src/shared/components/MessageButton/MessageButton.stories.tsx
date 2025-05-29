import type { Meta, StoryObj } from '@storybook/react';

import MessageButton from './MessageButton';

const meta = {
  title: 'Shared/MessageButton',
  component: MessageButton,
  parameters: {
    layout: 'centered',
    screenshot: {
      viewport: {
        width: 375,
        height: 812,
      },
    },
  },
  tags: ['autodocs'],
  decorators: [
    Story => (
      <div style={{ padding: '20px', background: '#f5f5f5' }}>
        <Story />
      </div>
    ),
  ],
} satisfies Meta<typeof MessageButton>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    onClick: () => console.log('Message button clicked'),
  },
};

export const WithCustomClass: Story = {
  args: {
    onClick: () => console.log('Message button clicked'),
    className: 'custom-message-button',
  },
  decorators: [
    Story => (
      <div style={{ padding: '20px', background: '#e0e0e0' }}>
        <Story />
      </div>
    ),
  ],
};
