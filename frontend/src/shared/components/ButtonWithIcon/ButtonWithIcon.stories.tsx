import EmailIcon from '@mui/icons-material/Email';
import type { Meta, StoryObj } from '@storybook/react';

import { ButtonWithIcon } from './ButtonWithIcon';

interface ButtonWithIconStoryProps {
  icon: React.ReactNode;
  text: string;
  onClick: () => void;
  className?: string;
  disabled?: boolean;
}

const meta = {
  title: 'Shared/ButtonWithIcon',
  component: ButtonWithIcon,
  parameters: {
    layout: 'centered',
  },
  argTypes: {
    onClick: { action: 'clicked' },
  },
  tags: ['autodocs'],
  decorators: [
    Story => (
      <div style={{ width: '350px', maxWidth: '500px', marginTop: '20px' }}>
        <Story />
      </div>
    ),
  ],
} satisfies Meta<ButtonWithIconStoryProps>;

export default meta;
type Story = StoryObj<typeof meta>;

export const VkLogin: Story = {
  args: {
    text: 'Continue with vk',
    onClick: () => console.log('Facebook button clicked'),
    icon: <img src="/vk.png" alt="vk Icon" style={{ width: '20px', height: '20px' }} />,
  },
};

export const EmailLogin: Story = {
  args: {
    text: 'LOGIN WITH EMAIL',
    onClick: () => console.log('Phone button clicked'),
    icon: <EmailIcon style={{ width: '20px', height: '20px' }} />,
  },
};
