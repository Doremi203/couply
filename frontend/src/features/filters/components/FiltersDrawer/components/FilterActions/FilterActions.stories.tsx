import type { Meta, StoryObj } from '@storybook/react';

import FilterActions from './FilterActions';

const meta = {
  title: 'Features/Filters/FilterActions',
  component: FilterActions,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
} satisfies Meta<typeof FilterActions>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    onContinue: () => console.log('Continue clicked'),
    buttonText: 'Continue',
  },
};

export const CustomText: Story = {
  args: {
    onContinue: () => console.log('Apply clicked'),
    buttonText: 'Apply Filters',
  },
};
