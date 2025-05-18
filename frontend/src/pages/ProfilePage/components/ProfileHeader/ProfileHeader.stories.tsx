import type { Meta, StoryObj } from '@storybook/react';

import { ProfileHeader } from './ProfileHeader';

const meta = {
  title: 'Components/ProfileHeader',
  component: ProfileHeader,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
} satisfies Meta<typeof ProfileHeader>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    isProfileHidden: false,
    onEditToggle: () => console.log('Edit toggle clicked'),
    onVisibilityToggle: () => console.log('Visibility toggle clicked'),
    onActivityClick: () => console.log('Activity clicked'),
    onPreviewClick: () => console.log('Preview clicked'),
  },
};

export const Hidden: Story = {
  args: {
    isProfileHidden: true,
    onEditToggle: () => console.log('Edit toggle clicked'),
    onVisibilityToggle: () => console.log('Visibility toggle clicked'),
    onActivityClick: () => console.log('Activity clicked'),
    onPreviewClick: () => console.log('Preview clicked'),
  },
};
