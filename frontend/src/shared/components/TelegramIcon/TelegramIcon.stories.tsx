import { Meta, StoryObj } from '@storybook/react';

import { TelegramIcon } from './TelegramIcon';

const meta: Meta<typeof TelegramIcon> = {
  title: 'Shared/TelegramIcon',
  component: TelegramIcon,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof TelegramIcon>;

export const Default: Story = {
  args: {},
};
