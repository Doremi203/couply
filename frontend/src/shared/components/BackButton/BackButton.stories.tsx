import type { Meta, StoryObj } from '@storybook/react';

import { BackButton } from './BackButton';

const meta = {
  title: 'Components/BackButton',
  component: BackButton,
  parameters: {
    layout: 'centered',
    backgrounds: {
      default: 'dark',
      values: [
        { name: 'dark', value: '#202C83' },
      ],
    },
  },
  tags: ['autodocs'],
} satisfies Meta<typeof BackButton>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    onClose: () => console.log('Back button clicked'),
  },
};