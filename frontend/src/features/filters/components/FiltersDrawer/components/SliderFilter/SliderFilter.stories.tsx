import type { Meta, StoryObj } from '@storybook/react';

import SliderFilter from './SliderFilter';

const meta = {
  title: 'Features/Filters/SliderFilter',
  component: SliderFilter,
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
} satisfies Meta<typeof SliderFilter>;

export default meta;
type Story = StoryObj<typeof meta>;

export const SingleValue: Story = {
  args: {
    title: 'Distance',
    value: 40,
    min: 1,
    max: 100,
    onChange: () => console.log('Distance changed'),
    unit: 'km',
  },
};

export const RangeValue: Story = {
  args: {
    title: 'Age',
    value: [20, 28],
    min: 18,
    max: 65,
    onChange: () => console.log('Age range changed'),
    valueLabelDisplay: 'auto',
  },
};

export const WithValueLabels: Story = {
  args: {
    title: 'Price',
    value: [500, 2000],
    min: 0,
    max: 5000,
    onChange: () => console.log('Price range changed'),
    unit: '₽',
    valueLabelDisplay: 'on',
  },
};
