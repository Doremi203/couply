import type { Meta, StoryObj } from '@storybook/react';

import ChipFilter from './ChipFilter';

const meta = {
  title: 'Features/Filters/ChipFilter',
  component: ChipFilter,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
  decorators: [
    Story => (
      <div style={{ marginTop: '20px' }}>
        <Story />
      </div>
    ),
  ],
} satisfies Meta<typeof ChipFilter>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    title: 'Interests',
    options: ['Sports', 'Travel', 'Music', 'Art', 'Food'],
    selectedOptions: ['Sports', 'Music'],
    onToggle: () => console.log('Option toggled'),
  },
};

export const Empty: Story = {
  args: {
    title: 'Interests',
    options: ['Sports', 'Travel', 'Music', 'Art', 'Food'],
    selectedOptions: [],
    onToggle: () => console.log('Option toggled'),
  },
};

export const AllSelected: Story = {
  args: {
    title: 'Music Preferences',
    options: ['Rock', 'Pop', 'Hip Hop', 'Jazz', 'Classical'],
    selectedOptions: ['Rock', 'Pop', 'Hip Hop', 'Jazz', 'Classical'],
    onToggle: () => console.log('Option toggled'),
  },
};
