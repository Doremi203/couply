import { StoryObj, Meta } from '@storybook/react';

import CustomButton from './CustomButton';

interface CustomButtonStoryProps {
  text: string;
  onClick: () => void;
  className?: string;
  disabled?: boolean;
}

const meta: Meta<CustomButtonStoryProps> = {
  title: 'Shared/CustomButton',
  component: CustomButton,
  argTypes: {
    onClick: { action: 'clicked' },
  },
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<CustomButtonStoryProps>;

export const Default: Story = {
  args: {
    text: 'Button',
    onClick: () => console.log('Button clicked'),
  },
};

export const Disabled: Story = {
  args: {
    text: 'Disabled Button',
    onClick: () => console.log('Button clicked'),
    disabled: true,
  },
};
