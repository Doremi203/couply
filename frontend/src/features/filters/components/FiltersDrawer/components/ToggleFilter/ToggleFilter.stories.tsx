import type { Meta, StoryObj } from '@storybook/react';

import ToggleFilter from './ToggleFilter';

const meta = {
  title: 'Features/Filters/ToggleFilter',
  component: ToggleFilter,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
  decorators: [
    Story => (
      <div style={{ width: '100%', maxWidth: '400px' }}>
        <Story />
      </div>
    ),
  ],
} satisfies Meta<typeof ToggleFilter>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Enabled: Story = {
  args: {
    title: 'Verification Status',
    description: 'Only show verified profiles',
    value: true,
    onChange: () => console.log('Toggle changed'),
  },
};

export const Disabled: Story = {
  args: {
    title: 'Verification Status',
    description: 'Only show verified profiles',
    value: false,
    onChange: () => console.log('Toggle changed'),
  },
};

export const CustomText: Story = {
  args: {
    title: 'Online Only',
    description: 'Show only users who are currently online',
    value: true,
    onChange: () => console.log('Toggle changed'),
  },
};
