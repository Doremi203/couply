import { ProfileSlider } from './ProfileSlider';
import { StoryObj, Meta } from '@storybook/react';

const meta: Meta = {
  title: 'Features/ProfileSlider',
  component: ProfileSlider,
  parameters: {
    layout: 'centered',
  },
};

export default meta;
type Story = StoryObj;

export const Default: Story = {
  args: {},
};