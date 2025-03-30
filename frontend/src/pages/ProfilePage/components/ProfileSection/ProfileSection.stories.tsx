import type { Meta, StoryObj } from '@storybook/react';
import { ProfileSection } from './ProfileSection';

const meta = {
  title: 'Components/ProfileSection',
  component: ProfileSection,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
} satisfies Meta<typeof ProfileSection>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    title: 'About Me',
    children: (
      <p>I love exploring new places and trying new cuisines. Photography is my passion, and I enjoy capturing moments during my travels.</p>
    ),
    showEditLink: false,
  },
};

export const WithEditLink: Story = {
  args: {
    title: 'About Me',
    children: (
      <p>I love exploring new places and trying new cuisines. Photography is my passion, and I enjoy capturing moments during my travels.</p>
    ),
    showEditLink: true,
    onEdit: () => console.log('Edit clicked'),
  },
};

export const WithListContent: Story = {
  args: {
    title: 'Interests',
    children: (
      <ul style={{ margin: 0, paddingLeft: '20px' }}>
        <li>Photography</li>
        <li>Traveling</li>
        <li>Cooking</li>
        <li>Reading</li>
      </ul>
    ),
    showEditLink: true,
    onEdit: () => console.log('Edit clicked'),
  },
};