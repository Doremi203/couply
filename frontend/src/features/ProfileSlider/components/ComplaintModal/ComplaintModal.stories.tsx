import type { Meta, StoryObj } from '@storybook/react';

import { ComplaintModal } from './CompliantModal';

// Note: This component requires the Redux store and API hooks to work properly.
// In Storybook, we're just showing the visual representation without the actual API functionality.
// In a real environment, this component would make API calls to submit complaints.
const meta = {
  title: 'Features/ProfileSlider/ComplaintModal',
  component: ComplaintModal,
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
} satisfies Meta<typeof ComplaintModal>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    isOpen: true,
    onClose: () => console.log('Modal closed'),
    targetUserId: 'user123',
  },
};

export const WithSelectedReason: Story = {
  args: {
    isOpen: true,
    onClose: () => console.log('Modal closed'),
    targetUserId: 'user123',
  },
  parameters: {
    docs: {
      description: {
        story: 'Модальное окно с выбранной причиной жалобы',
      },
    },
  },
};

export const WithCustomText: Story = {
  args: {
    isOpen: true,
    onClose: () => console.log('Modal closed'),
    targetUserId: 'user123',
  },
  parameters: {
    docs: {
      description: {
        story: 'Модальное окно с введенным пользовательским текстом',
      },
    },
  },
};
