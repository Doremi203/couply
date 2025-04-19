import type { Meta, StoryObj } from '@storybook/react';

import FilterHeader from './FilterHeader';

const meta = {
  title: 'Features/Filters/FilterHeader',
  component: FilterHeader,
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
} satisfies Meta<typeof FilterHeader>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    onBack: () => console.log('Back clicked'),
    onClear: () => console.log('Clear clicked'),
  },
};
