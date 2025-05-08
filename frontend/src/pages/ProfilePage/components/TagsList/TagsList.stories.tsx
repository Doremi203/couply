import type { Meta, StoryObj } from '@storybook/react';

import { TagsList } from './TagsList';

const meta = {
  title: 'Components/TagsList',
  component: TagsList,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
  decorators: [
    Story => (
      <div style={{ width: '350px', marginTop: '20px' }}>
        <Story />
      </div>
    ),
  ],
} satisfies Meta<typeof TagsList>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    items: ['Travel', 'Photography', 'Cooking', 'Reading', 'Hiking', 'Music'],
  },
};

export const WithCommonInterests: Story = {
  args: {
    items: ['Travel', 'Photography', 'Cooking', 'Reading', 'Hiking', 'Music'],
    commonItems: ['Photography', 'Music'],
  },
};
