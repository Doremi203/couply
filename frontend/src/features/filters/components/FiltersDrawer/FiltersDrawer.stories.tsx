import type { Meta, StoryObj } from '@storybook/react';

import FiltersDrawer from './FiltersDrawer';

const meta = {
  title: 'Components/FiltersDrawer',
  component: FiltersDrawer,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
} satisfies Meta<typeof FiltersDrawer>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    open: true,
    onClose: () => console.log('Drawer closed'),
  },
};