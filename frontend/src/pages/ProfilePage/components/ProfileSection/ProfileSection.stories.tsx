import type { Meta, StoryObj } from '@storybook/react';

import { ProfileSection } from './ProfileSection';

const meta = {
  title: 'Components/ProfileSection',
  component: ProfileSection,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
  decorators: [
    Story => (
      <div style={{ width: '350px', marginTop: '20px' }}>
        <Story />
      </div>
    ),
  ],
} satisfies Meta<typeof ProfileSection>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    title: 'About Me',
    children: (
      <p>
        I love exploring new places and trying new cuisines. Photography is my passion, and I enjoy
        capturing moments during my travels.
      </p>
    ),
    showEditLink: false,
  },
};
