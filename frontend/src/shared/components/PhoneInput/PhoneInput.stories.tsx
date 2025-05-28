import { useState } from 'react';
import type { Meta, StoryObj } from '@storybook/react';

import { PhoneInput } from './PhoneInput';

const meta = {
  title: 'Shared/PhoneInput',
  component: PhoneInput,
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
} satisfies Meta<typeof PhoneInput>;

export default meta;
type Story = StoryObj<typeof meta>;

// Controlled component wrapper for interactive stories
const PhoneInputWrapper = () => {
  const [phone, setPhone] = useState('+7');
  return <PhoneInput value={phone} onChange={setPhone} />;
};

export const Default: Story = {
  args: {
    value: '+7',
    onChange: value => console.log('Phone changed:', value),
  },
};

export const WithPartialNumber: Story = {
  args: {
    value: '+7 (123)',
    onChange: value => console.log('Phone changed:', value),
  },
};

export const WithFullNumber: Story = {
  args: {
    value: '+7 (123) 456 78 90',
    onChange: value => console.log('Phone changed:', value),
  },
};

// Interactive story that allows typing in Storybook
export const Interactive: Story = {
  render: () => <PhoneInputWrapper />,
  args: {
    value: '+7',
    onChange: () => {},
  },
};
