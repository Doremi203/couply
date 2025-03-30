import type { Meta, StoryObj } from '@storybook/react';
import { InfoGrid } from './InfoGrid';

const meta = {
  title: 'Components/InfoGrid',
  component: InfoGrid,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
} satisfies Meta<typeof InfoGrid>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    infoItems: [
      { label: 'Age', value: '28' },
      { label: 'Height', value: '175 cm' },
      { label: 'Location', value: 'Moscow' },
      { label: 'Education', value: 'University' },
      { label: 'Occupation', value: 'Software Engineer' },
    ],
  },
};

export const ShortList: Story = {
  args: {
    infoItems: [
      { label: 'Age', value: '28' },
      { label: 'Location', value: 'Moscow' },
    ],
  },
};