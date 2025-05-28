import type { Meta, StoryObj } from '@storybook/react';

import {
  sportOptions,
  hobbyOptions,
  musicOptions,
  foodDrinkOptions,
} from '../../../../features/filters/components/constants';
import { InterestSection } from './InterestsSection';

const meta = {
  title: 'Widgets/EditProfile/InterestSection',
  component: InterestSection,
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
} satisfies Meta<typeof InterestSection>;

export default meta;
type Story = StoryObj<typeof meta>;

// Sample selected interests from different categories
const sampleSelectedInterests = [
  sportOptions.running,
  sportOptions.swimming,
  hobbyOptions.travel,
  musicOptions.rock,
  foodDrinkOptions.coffee,
];

export const Default: Story = {
  args: {
    selectedOptions: [],
    onSelect: selected => console.log('Selected interests:', selected),
  },
};

export const WithSelectedInterests: Story = {
  args: {
    selectedOptions: sampleSelectedInterests,
    onSelect: selected => console.log('Selected interests:', selected),
  },
};

export const ManyInterests: Story = {
  args: {
    selectedOptions: [
      sportOptions.running,
      sportOptions.swimming,
      sportOptions.cycling,
      hobbyOptions.travel,
      hobbyOptions.photography,
      musicOptions.rock,
      musicOptions.jazz,
      foodDrinkOptions.coffee,
    ],
    onSelect: selected => console.log('Selected interests:', selected),
  },
};
