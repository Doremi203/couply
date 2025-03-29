import type { Meta, StoryObj } from '@storybook/react';
import { ButtonWithIcon } from './ButtonWithIcon';

const meta = {
  title: 'shared/components/ButtonWithIcon',
  component: ButtonWithIcon,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
} satisfies Meta<typeof ButtonWithIcon>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {},
};