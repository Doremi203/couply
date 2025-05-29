import type { Meta, StoryObj } from '@storybook/react';

import PWAGeolocationHelper from './PWAGeolocationHelper';

const meta = {
  title: 'Pages/EnterInfoPage/PWAGeolocationHelper',
  component: PWAGeolocationHelper,
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
} satisfies Meta<typeof PWAGeolocationHelper>;

export default meta;
type Story = StoryObj<typeof meta>;

export const IOSInstructions: Story = {
  args: {
    open: true,
    onClose: () => console.log('Dialog closed'),
    isIOS: true,
  },
  parameters: {
    docs: {
      description: {
        story: 'Инструкции для включения геолокации на iOS устройствах',
      },
    },
  },
};

export const AndroidInstructions: Story = {
  args: {
    open: true,
    onClose: () => console.log('Dialog closed'),
    isIOS: false,
  },
  parameters: {
    docs: {
      description: {
        story: 'Инструкции для включения геолокации на Android устройствах',
      },
    },
  },
};

export const Closed: Story = {
  args: {
    open: false,
    onClose: () => console.log('Dialog closed'),
    isIOS: true,
  },
  parameters: {
    docs: {
      description: {
        story: 'Закрытый диалог (не отображается)',
      },
    },
  },
};
