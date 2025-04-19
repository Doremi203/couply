import type { Meta, StoryObj } from '@storybook/react';

import { ProfilePhotoEdit } from './ProfilePhotoEdit';

const meta = {
  title: 'Features/ProfileEdit/ProfilePhotoEdit',
  component: ProfilePhotoEdit,
  parameters: {
    layout: 'centered',
    docs: {
      description: {
        component: 'A component for editing the profile photo with camera icon for upload.',
      },
    },
  },
  tags: ['autodocs'],
  decorators: [
    Story => (
      <div style={{ width: '100%', maxWidth: '400px' }}>
        <Story />
      </div>
    ),
  ],
} satisfies Meta<typeof ProfilePhotoEdit>;

export default meta;
type Story = StoryObj<typeof ProfilePhotoEdit>;

export const WithPhoto: Story = {
  args: {
    profilePhoto: 'https://randomuser.me/api/portraits/women/68.jpg',
    onCameraClick: isAvatar => console.log(`Camera clicked, isAvatar: ${isAvatar}`),
  },
};

export const NoPhoto: Story = {
  args: {
    profilePhoto: '',
    onCameraClick: isAvatar => console.log(`Camera clicked, isAvatar: ${isAvatar}`),
  },
};

export const MalePhoto: Story = {
  args: {
    profilePhoto: 'https://randomuser.me/api/portraits/men/44.jpg',
    onCameraClick: isAvatar => console.log(`Camera clicked, isAvatar: ${isAvatar}`),
  },
};

export const CustomTitle: Story = {
  args: {
    profilePhoto: 'https://randomuser.me/api/portraits/women/68.jpg',
    onCameraClick: isAvatar => console.log(`Camera clicked, isAvatar: ${isAvatar}`),
    title: 'Your Avatar',
  },
};
