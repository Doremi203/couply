import { Meta, StoryObj } from '@storybook/react';
import { InstagramIcon } from './InstagramIcon';

const meta: Meta<typeof InstagramIcon> = {
  title: 'Shared/components/InstagramIcon',
  component: InstagramIcon,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof InstagramIcon>;

export const Default: Story = {
  args: {},
};