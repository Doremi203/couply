import { StoryObj, Meta } from '@storybook/react';

import CustomInput from './CustomInput';

interface CustomInputStoryProps {
  placeholder: string;
  type: string;
  className?: string;
  value?: string;
  onChange?: (e: React.ChangeEvent<HTMLInputElement>) => void;
}

const meta: Meta<CustomInputStoryProps> = {
  title: 'Shared/CustomInput',
  component: CustomInput,
  argTypes: {
    onChange: { action: 'changed' },
  },
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<CustomInputStoryProps>;

export const Text: Story = {
  args: {
    placeholder: 'Enter text',
    type: 'text',
    value: '',
  },
};

export const Password: Story = {
  args: {
    placeholder: 'Enter password',
    type: 'password',
    value: '',
  },
};

export const Email: Story = {
  args: {
    placeholder: 'Enter email',
    type: 'email',
    value: '',
  },
};
