import type { Meta, StoryObj } from '@storybook/react';

import {
  sportOptions,
  selfdevelopmentOptions,
  hobbyOptions,
  musicOptions,
  moviesTVOptions,
  foodDrinkOptions,
  personalityTraitsOptions,
  petsOptions,
} from '../../../constants';
import InterestFilter from './InterestFilter';

const meta = {
  title: 'Features/Filters/InterestFilter',
  component: InterestFilter,
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
} satisfies Meta<typeof InterestFilter>;

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
    title: 'Интересы',
    selectedOptions: [],
    onSelect: selected => console.log('Selected interests:', selected),
  },
};

export const WithSelectedInterests: Story = {
  args: {
    title: 'Интересы',
    selectedOptions: sampleSelectedInterests,
    onSelect: selected => console.log('Selected interests:', selected),
  },
};

export const ManyInterests: Story = {
  args: {
    title: 'Интересы',
    selectedOptions: [
      sportOptions.running,
      sportOptions.swimming,
      sportOptions.cycling,
      selfdevelopmentOptions.languages,
      hobbyOptions.travel,
      hobbyOptions.photography,
      musicOptions.rock,
      musicOptions.jazz,
      moviesTVOptions.comedy,
      foodDrinkOptions.coffee,
      personalityTraitsOptions.creative,
      petsOptions.dogs,
    ],
    onSelect: selected => console.log('Selected interests:', selected),
  },
};
