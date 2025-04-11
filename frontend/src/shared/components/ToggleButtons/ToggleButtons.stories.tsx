import { StoryObj, Meta } from '@storybook/react';

import ToggleButtons from './ToggleButtons';

interface ToggleButtonsStoryProps {
  options: Array<{ label: string; value: string }>;
  onSelect: (value: string) => void;
  value?: string;
}

const meta: Meta<ToggleButtonsStoryProps> = {
  title: 'shared/Components/ToggleButtons',
  component: ToggleButtons,
  argTypes: {
    onSelect: { action: 'selected' },
  },
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<ToggleButtonsStoryProps>;

export const Default: Story = {
  args: {
    options: [
      { label: 'Option 1', value: 'option1' },
      { label: 'Option 2', value: 'option2' },
    ],
    value: 'option1',
  },
};

export const ThreeOptions: Story = {
  args: {
    options: [
      { label: 'Option 1', value: 'option1' },
      { label: 'Option 2', value: 'option2' },
      { label: 'Option 3', value: 'option3' },
    ],
    value: 'option2',
  },
};

export const NoSelection: Story = {
  args: {
    options: [
      { label: 'Option 1', value: 'option1' },
      { label: 'Option 2', value: 'option2' },
    ],
    value: undefined,
  },
};