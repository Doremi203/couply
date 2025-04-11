import { StoryObj, Meta } from '@storybook/react';

import { ProfileSlider } from './ProfileSlider';

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