import type { Meta, StoryObj } from '@storybook/react';

import ConfirmModal from './ConfirmModal';

const meta = {
  title: 'Shared/ConfirmModal',
  component: ConfirmModal,
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
} satisfies Meta<typeof ConfirmModal>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    isOpen: true,
    onClose: () => console.log('Modal closed'),
    onConfirm: () => console.log('Action confirmed'),
    title: 'Подтверждение',
    message: 'Вы уверены, что хотите выполнить это действие?',
    confirmText: 'Подтвердить',
    cancelText: 'Отмена',
  },
};

export const DeleteConfirmation: Story = {
  args: {
    isOpen: true,
    onClose: () => console.log('Modal closed'),
    onConfirm: () => console.log('Delete confirmed'),
    title: 'Удаление профиля',
    message: 'Вы уверены, что хотите удалить свой профиль? Это действие нельзя отменить.',
    confirmText: 'Удалить',
    cancelText: 'Отмена',
  },
};

export const CustomButtons: Story = {
  args: {
    isOpen: true,
    onClose: () => console.log('Modal closed'),
    onConfirm: () => console.log('Action confirmed'),
    title: 'Выход из аккаунта',
    message: 'Вы действительно хотите выйти из своего аккаунта?',
    confirmText: 'Выйти',
    cancelText: 'Остаться',
  },
};
