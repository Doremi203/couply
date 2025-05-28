import type { Meta, StoryObj } from '@storybook/react';

import CodeInput from './CodeInput';

const meta = {
  title: 'Shared/CodeInput',
  component: CodeInput,
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
      <div style={{ width: '375px', padding: '20px' }}>
        <Story />
      </div>
    ),
  ],
} satisfies Meta<typeof CodeInput>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    length: 6,
    onCodeChange: code => console.log('Code changed:', code),
  },
};

export const FourDigits: Story = {
  args: {
    length: 4,
    onCodeChange: code => console.log('Code changed:', code),
  },
};
