import type { Meta, StoryObj } from '@storybook/react';

import GenderFilter from './GenderFilter';

const meta = {
  title: 'Features/Filters/GenderFilter',
  component: GenderFilter,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
  decorators: [
    Story => (
      <div style={{ width: '100%', maxWidth: '400px', marginTop: '20px' }}>
        <Story />
      </div>
    ),
  ],
} satisfies Meta<typeof GenderFilter>;

export default meta;
type Story = StoryObj<typeof meta>;

const genderOptions = [
  { label: 'Girls', value: 'Girls' },
  { label: 'Boys', value: 'Boys' },
  { label: 'Both', value: 'Both' },
];

export const Girls: Story = {
  args: {
    value: 'Girls',
    //@ts-ignore
    options: genderOptions,
    onChange: () => console.log('Gender changed'),
  },
};

export const Boys: Story = {
  args: {
    value: 'Boys',
    //@ts-ignore
    options: genderOptions,
    onChange: () => console.log('Gender changed'),
  },
};

export const Both: Story = {
  args: {
    value: 'Both',
    //@ts-ignore
    options: genderOptions,
    onChange: () => console.log('Gender changed'),
  },
};
