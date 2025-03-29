import type { Meta, StoryObj } from '@storybook/react';
import { FiltersIcon } from './FiltersIcon';

const meta = {
  title: 'shared/components/FiltersIcon',
  component: FiltersIcon,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
} satisfies Meta<typeof FiltersIcon>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {},
};