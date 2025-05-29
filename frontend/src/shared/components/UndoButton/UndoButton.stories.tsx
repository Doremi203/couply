import type { Meta, StoryObj } from '@storybook/react';

import UndoButton from './UndoButton';

const meta = {
  title: 'Shared/UndoButton',
  component: UndoButton,
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
} satisfies Meta<typeof UndoButton>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    onClick: () => console.log('Undo button clicked'),
  },
};

export const WithCustomClass: Story = {
  args: {
    onClick: () => console.log('Undo button clicked'),
    className: 'custom-undo-button',
  },
  decorators: [
    Story => (
      <div style={{ padding: '20px', background: '#e0e0e0' }}>
        <Story />
      </div>
    ),
  ],
};
