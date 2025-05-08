import type { Meta, StoryObj } from '@storybook/react';

import { ProfileSlider } from './ProfileSlider';

const meta = {
  title: 'Features/ProfileSlider',
  component: ProfileSlider,
  parameters: {
    layout: 'centered',
    docs: {
      description: {
        component:
          'A slider component for browsing through user profiles with swipe functionality.',
      },
    },
  },
  tags: ['autodocs'],
  decorators: [
    Story => (
      <div style={{ width: '100%', maxWidth: '400px', height: '600px' }}>
        <Story />
      </div>
    ),
  ],
} satisfies Meta<typeof ProfileSlider>;

export default meta;
type Story = StoryObj<typeof ProfileSlider>;

// Since ProfileSlider has its own internal state and data,
// we don't need to pass any props to it
export const Default: Story = {};

// Note: The ProfileSlider component uses internal state and hardcoded profiles,
// so we can't easily customize it with different stories without modifying the component.
// In a real-world scenario, you might want to refactor the component to accept profiles as props
// to make it more flexible for testing and reuse.
