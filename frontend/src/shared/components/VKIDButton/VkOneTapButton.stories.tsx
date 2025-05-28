import type { Meta, StoryObj } from '@storybook/react';
import React from 'react';

// Create a mock component for Storybook that visually represents the VK button
// This avoids issues with the actual SDK in the Storybook environment
const MockVkButton = ({ scheme = 'light', lang = 'ru' }) => {
  const buttonStyle = {
    display: 'inline-flex',
    alignItems: 'center',
    justifyContent: 'center',
    width: '280px',
    height: '44px',
    borderRadius: '22px',
    backgroundColor: scheme === 'light' ? '#fff' : '#222',
    color: scheme === 'light' ? '#000' : '#fff',
    border: scheme === 'light' ? '1px solid #ddd' : '1px solid #444',
    fontFamily: 'Arial, sans-serif',
    fontSize: '14px',
    cursor: 'pointer',
    padding: '0 16px',
    boxShadow: '0 2px 4px rgba(0, 0, 0, 0.1)',
  };

  const logoStyle = {
    width: '24px',
    height: '24px',
    marginRight: '8px',
    borderRadius: '50%',
    backgroundColor: '#4680C2',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    color: '#fff',
    fontWeight: 'bold',
    fontSize: '16px',
  };

  return (
    <div style={buttonStyle}>
      <div style={logoStyle}>VK</div>
      {lang === 'ru' ? 'Войти через VK ID' : 'Log in with VK ID'}
    </div>
  );
};

// Use our mock component for Storybook
const VkOneTapButton = MockVkButton;

const meta = {
  title: 'Shared/VkOneTapButton',
  component: VkOneTapButton,
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
      <div style={{ padding: '20px' }}>
        <Story />
      </div>
    ),
  ],
} satisfies Meta<typeof VkOneTapButton>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    scheme: 'light',
    lang: 'ru',
  },
  parameters: {
    docs: {
      description: {
        story: 'Кнопка авторизации через VK ID в светлой теме',
      },
    },
  },
};

export const DarkTheme: Story = {
  args: {
    scheme: 'dark',
    lang: 'ru',
  },
  parameters: {
    docs: {
      description: {
        story: 'Кнопка авторизации через VK ID в темной теме',
      },
    },
    backgrounds: {
      default: 'dark',
    },
  },
  decorators: [
    Story => (
      <div style={{ padding: '20px', background: '#333' }}>
        <Story />
      </div>
    ),
  ],
};

export const EnglishLanguage: Story = {
  args: {
    scheme: 'light',
    lang: 'en',
  },
  parameters: {
    docs: {
      description: {
        story: 'Кнопка авторизации через VK ID на английском языке',
      },
    },
  },
};
