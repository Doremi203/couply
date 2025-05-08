import type { Meta, StoryObj } from '@storybook/react';
import { MemoryRouter } from 'react-router-dom';

import { NavBar } from './NavBar';

const meta = {
  title: 'Shared/NavBar',
  component: NavBar,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
} satisfies Meta<typeof NavBar>;

export default meta;
type Story = StoryObj<typeof meta>;

export const HomeActive: Story = {
  decorators: [
    Story => (
      <MemoryRouter initialEntries={['/home']}>
        <div style={{ width: '375px', padding: '20px', backgroundColor: '#f5f5f5' }}>
          <Story />
        </div>
      </MemoryRouter>
    ),
  ],
};

export const LikesActive: Story = {
  decorators: [
    Story => (
      <MemoryRouter initialEntries={['/likes']}>
        <div style={{ width: '375px', padding: '20px', backgroundColor: '#f5f5f5' }}>
          <Story />
        </div>
      </MemoryRouter>
    ),
  ],
};

export const ProfileActive: Story = {
  decorators: [
    Story => (
      <MemoryRouter initialEntries={['/profile']}>
        <div style={{ width: '375px', padding: '20px', backgroundColor: '#f5f5f5' }}>
          <Story />
        </div>
      </MemoryRouter>
    ),
  ],
};
