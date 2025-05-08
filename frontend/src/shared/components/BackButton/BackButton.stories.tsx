import type { Meta, StoryObj } from '@storybook/react';

import { BackButton } from './BackButton';

const meta = {
  title: 'Shared/BackButton',
  component: BackButton,
  parameters: {
    // layout: 'centered',
    backgrounds: {
      default: 'dark',
      values: [{ name: 'dark', value: '#202C83' }],
    },
  },
  tags: ['autodocs'],
  decorators: [
    Story => (
      <div style={{ width: '100px', height: '100px', marginTop: '20px' }}>
        <Story />
      </div>
    ),
  ],
} satisfies Meta<typeof BackButton>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    onClose: () => console.log('Back button clicked'),
  },
};
