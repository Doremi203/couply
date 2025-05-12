import { StoryObj, Meta } from '@storybook/react';
import { BrowserRouter } from 'react-router-dom';

import Page from './Page';

const meta: Meta = {
  title: 'Pages/ProfilePage',
  component: Page,
  decorators: [
    Story => (
      <BrowserRouter>
        <div style={{ width: '350px', maxWidth: '400px' }}>
          <Story />
        </div>
      </BrowserRouter>
    ),
  ],
};

export default meta;
type Story = StoryObj;

export const Default: Story = {
  args: {
    initialTab: 'profile',
    initialEditMode: false,
    initialVerified: false,
  },
  name: 'Profile View',
  parameters: {
    docs: {
      description: {
        story: 'Default profile view showing user information, photos, and interests.',
      },
    },
  },
};

export const EditMode: Story = {
  args: {
    initialTab: 'edit',
    initialEditMode: true,
    initialVerified: false,
  },
  name: 'Edit Mode View',
  parameters: {
    docs: {
      description: {
        story: 'Shows the profile in edit mode where users can modify their information.',
      },
    },
  },
};

export const ProfilePreview: Story = {
  args: {
    initialTab: 'preview',
    initialEditMode: false,
    initialVerified: false,
  },
  name: 'Profile Preview Mode',
  parameters: {
    docs: {
      description: {
        story: 'Shows how the profile appears to other users, with common interests highlighted.',
      },
    },
  },
};

export const VerifiedProfile: Story = {
  args: {
    initialTab: 'profile',
    initialEditMode: false,
    initialVerified: true,
  },
  name: 'Verified Profile View',
  parameters: {
    docs: {
      description: {
        story: 'Shows a profile with verification badge.',
      },
    },
  },
};
