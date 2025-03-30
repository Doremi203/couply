import type { Meta, StoryObj } from '@storybook/react';
import { TagsList } from './TagsList';

const meta = {
  title: 'Components/TagsList',
  component: TagsList,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
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

export const SingleItem: Story = {
  args: {
    items: ['Photography'],
  },
};

export const ManyItems: Story = {
  args: {
    items: [
      'Travel', 'Photography', 'Cooking', 'Reading', 'Hiking', 'Music',
      'Movies', 'Art', 'Dancing', 'Swimming', 'Yoga', 'Running',
      'Gaming', 'Painting', 'Writing', 'Gardening'
    ],
    commonItems: ['Photography', 'Music', 'Art', 'Gaming'],
  },
};