import FavoriteIcon from '@mui/icons-material/Favorite';
import HomeIcon from '@mui/icons-material/Home';
import MessageIcon from '@mui/icons-material/Message';
import SettingsIcon from '@mui/icons-material/Settings';
import type { Meta, StoryObj } from '@storybook/react';

import { IconButton } from './IconButton';

const meta = {
  title: 'Shared/IconButton',
  component: IconButton,
  parameters: {
    layout: 'centered',
    docs: {
      description: {
        component: 'A circular button that contains an icon.',
      },
    },
  },
  tags: ['autodocs'],
} satisfies Meta<typeof IconButton>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    onClick: () => console.log('Icon button clicked'),
    children: <HomeIcon />,
  },
};

export const WithCustomClass: Story = {
  args: {
    onClick: () => console.log('Icon button clicked'),
    className: 'custom-button-class',
    children: <FavoriteIcon />,
  },
};

export const WithIconClass: Story = {
  args: {
    onClick: () => console.log('Icon button clicked'),
    iconClassName: 'custom-icon-class',
    children: <MessageIcon />,
  },
};

export const TouchFriendly: Story = {
  args: {
    onClick: () => console.log('Icon button clicked'),
    touchFriendly: true,
    children: <SettingsIcon />,
  },
};

export const CustomContent: Story = {
  args: {
    onClick: () => console.log('Icon button clicked'),
    children: <span style={{ fontSize: '24px' }}>🚀</span>,
  },
};
