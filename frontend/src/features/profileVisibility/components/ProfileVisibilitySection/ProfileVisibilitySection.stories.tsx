import type { Meta, StoryObj } from '@storybook/react';

import { ProfileVisibilitySection } from './ProfileVisibilitySection';

const meta = {
  title: 'Features/ProfileVisibility/ProfileVisibilitySection',
  component: ProfileVisibilitySection,
  parameters: {
    layout: 'centered',
    docs: {
      description: {
        component: 'A section that allows users to control the visibility of their profile.',
      },
    },
  },
  tags: ['autodocs'],
  decorators: [
    Story => (
      <div style={{ width: '350px', maxWidth: '400px' }}>
        <Story />
      </div>
    ),
  ],
} satisfies Meta<typeof ProfileVisibilitySection>;

export default meta;
type Story = StoryObj<typeof ProfileVisibilitySection>;

export const Visible: Story = {
  args: {
    isHidden: false,
    onInputChange: (field, value) => console.log(`Field ${field} changed to: ${value}`),
  },
};

export const Hidden: Story = {
  args: {
    isHidden: true,
    onInputChange: (field, value) => console.log(`Field ${field} changed to: ${value}`),
  },
};

export const CustomTitle: Story = {
  args: {
    isHidden: false,
    onInputChange: (field, value) => console.log(`Field ${field} changed to: ${value}`),
    title: 'Privacy Settings',
  },
};
