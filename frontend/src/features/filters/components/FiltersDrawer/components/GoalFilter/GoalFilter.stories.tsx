import type { Meta, StoryObj } from '@storybook/react';

import GoalFilter from './GoalFilter';

const meta = {
  title: 'Features/Filters/GoalFilter',
  component: GoalFilter,
  parameters: {
    layout: 'centered',
    screenshot: {
      viewport: {
        width: 375,
        height: 812,
      },
    },
  },
  tags: ['autodocs'],
  decorators: [
    Story => (
      <div style={{ width: '375px', padding: '20px' }}>
        <Story />
      </div>
    ),
  ],
} satisfies Meta<typeof GoalFilter>;

export default meta;
type Story = StoryObj<typeof meta>;

const goalOptions = [
  { label: 'Отношения', value: 'relationship' },
  { label: 'Дружба', value: 'friendship' },
  { label: 'Общение', value: 'communication' },
];

export const Default: Story = {
  args: {
    value: 'relationship',
    options: goalOptions,
    onChange: value => console.log('Selected goal:', value),
  },
};

export const FriendshipSelected: Story = {
  args: {
    value: 'friendship',
    options: goalOptions,
    onChange: value => console.log('Selected goal:', value),
  },
};

export const CommunicationSelected: Story = {
  args: {
    value: 'communication',
    options: goalOptions,
    onChange: value => console.log('Selected goal:', value),
  },
};
