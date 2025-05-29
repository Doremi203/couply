import type { Meta, StoryObj } from '@storybook/react';

import { CustomSlider } from './CustomSlider';

const meta = {
  title: 'Features/Filters/CustomSlider',
  component: CustomSlider,
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
      <div style={{ width: '300px', padding: '20px' }}>
        <Story />
      </div>
    ),
  ],
} satisfies Meta<typeof CustomSlider>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    defaultValue: 50,
    min: 0,
    max: 100,
    valueLabelDisplay: 'auto',
  },
};

export const Range: Story = {
  args: {
    defaultValue: [20, 80],
    min: 0,
    max: 100,
    valueLabelDisplay: 'auto',
  },
};

export const Disabled: Story = {
  args: {
    defaultValue: 50,
    min: 0,
    max: 100,
    disabled: true,
  },
};

export const WithMarks: Story = {
  args: {
    defaultValue: 50,
    min: 0,
    max: 100,
    step: 20,
    marks: [
      { value: 0, label: '0' },
      { value: 20, label: '20' },
      { value: 40, label: '40' },
      { value: 60, label: '60' },
      { value: 80, label: '80' },
      { value: 100, label: '100' },
    ],
    valueLabelDisplay: 'auto',
  },
};
