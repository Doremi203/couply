import type { Meta, StoryObj } from '@storybook/react';

import MessageModal from './MessageModal';

const meta = {
  title: 'Shared/MessageModal',
  component: MessageModal,
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
} satisfies Meta<typeof MessageModal>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    isOpen: true,
    onClose: () => console.log('Modal closed'),
    message: 'Привет! Хотел бы познакомиться. Как твои дела?',
    senderName: 'Александр',
  },
};

export const LongMessage: Story = {
  args: {
    isOpen: true,
    onClose: () => console.log('Modal closed'),
    message:
      'Привет! Я заметил, что у нас много общих интересов. Я тоже люблю путешествовать и читать книги. Может быть, мы могли бы обсудить наши любимые места и авторов? Буду рад пообщаться!',
    senderName: 'Мария',
  },
};

export const NoMessage: Story = {
  args: {
    isOpen: true,
    onClose: () => console.log('Modal closed'),
    message: '',
    senderName: 'Иван',
  },
};

export const DefaultValues: Story = {
  args: {
    isOpen: true,
    onClose: () => console.log('Modal closed'),
  },
};
