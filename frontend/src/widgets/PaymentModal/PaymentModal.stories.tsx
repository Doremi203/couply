import type { Meta, StoryObj } from '@storybook/react';

import { Plan } from '../../entities/subscription/types';

import { PaymentModal } from './PaymentModal';

const meta = {
  title: 'Widgets/PaymentModal',
  component: PaymentModal,
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
      <div style={{ width: '100%', height: '100vh', position: 'relative' }}>
        <Story />
      </div>
    ),
  ],
} satisfies Meta<typeof PaymentModal>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    isOpen: true,
    onClose: () => console.log('Modal closed'),
    selectedPlan: Plan.monthly,
    price: '299₽',
  },
};

export const Processing: Story = {
  args: {
    isOpen: true,
    onClose: () => console.log('Modal closed'),
    selectedPlan: Plan.monthly,
    price: '299₽',
  },
  play: async ({ canvasElement }) => {
    const submitButton = canvasElement.querySelector('button[type="submit"]') as HTMLButtonElement;
    if (submitButton) {
      submitButton.click();
    }
  },
};

export const Closed: Story = {
  args: {
    isOpen: false,
    onClose: () => console.log('Modal closed'),
    selectedPlan: Plan.monthly,
    price: '299₽',
  },
};
