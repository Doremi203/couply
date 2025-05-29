import type { Meta, StoryObj } from '@storybook/react';

import LocationSelector from './LocationSelector';

const meta = {
  title: 'Pages/EnterInfoPage/LocationSelector',
  component: LocationSelector,
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
  decorators: [
    Story => (
      <div style={{ width: '375px', padding: '20px' }}>
        <Story />
      </div>
    ),
  ],
} satisfies Meta<typeof LocationSelector>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    onLocationSelected: location => console.log('Location selected:', location),
  },
};

// Mock implementation for Storybook
// In a real environment, this component makes API calls to Nominatim
// which cannot be fully simulated in Storybook

// To test the component with mock data, we could use MSW (Mock Service Worker)
// or other mocking libraries, but that's beyond the scope of this story file
