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
  title: 'shared/components/ButtonWithIcon',
  component: ButtonWithIcon,
  parameters: {
    layout: 'centered',
  },
  argTypes: {
    onClick: { action: 'clicked' },
  },
  tags: ['autodocs'],
} satisfies Meta<ButtonWithIconStoryProps>;

export default meta;
type Story = StoryObj<typeof meta>;

export const FacebookLogin: Story = {
  args: {
    text: 'Continue with Facebook',
    onClick: () => console.log('Facebook button clicked'),
    icon: (
      <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
        <path d="M12 0C5.373 0 0 5.373 0 12C0 18.627 5.373 24 12 24C18.627 24 24 18.627 24 12C24 5.373 18.627 0 12 0Z" fill="#1877F2"/>
        <path d="M15.0168 12.4697H12.8755V20.0015H9.63947V12.4697H8.08737V9.71846H9.63947V7.92879C9.63947 6.65291 10.2368 4.65918 12.9071 4.65918L15.3059 4.67137V7.33697H13.5644C13.2887 7.33697 12.8757 7.48791 12.8757 8.09072V9.72115H15.3002L15.0168 12.4697Z" fill="white"/>
      </svg>
    ),
  },
};

export const PhoneLogin: Story = {
  args: {
    text: 'LOGIN WITH PHONE',
    onClick: () => console.log('Phone button clicked'),
    icon: (
      <img src="/phone.png" alt="Phone Icon" style={{ width: '20px', height: '20px' }} />
    ),
  },
};

export const Disabled: Story = {
  args: {
    text: 'Continue with Facebook',
    onClick: () => console.log('Button clicked'),
    disabled: true,
    icon: (
      <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
        <path d="M12 0C5.373 0 0 5.373 0 12C0 18.627 5.373 24 12 24C18.627 24 24 18.627 24 12C24 5.373 18.627 0 12 0Z" fill="#1877F2"/>
        <path d="M15.0168 12.4697H12.8755V20.0015H9.63947V12.4697H8.08737V9.71846H9.63947V7.92879C9.63947 6.65291 10.2368 4.65918 12.9071 4.65918L15.3059 4.67137V7.33697H13.5644C13.2887 7.33697 12.8757 7.48791 12.8757 8.09072V9.72115H15.3002L15.0168 12.4697Z" fill="white"/>
      </svg>
    ),
  },
};