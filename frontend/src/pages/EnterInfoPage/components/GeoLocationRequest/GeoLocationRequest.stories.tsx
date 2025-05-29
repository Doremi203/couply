import type { Meta, StoryObj } from '@storybook/react';

import GeoLocationRequest from './GeoLocationRequest';

const meta = {
  title: 'Pages/EnterInfoPage/GeoLocationRequest',
  component: GeoLocationRequest,
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
} satisfies Meta<typeof GeoLocationRequest>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    onLocationReceived: coords => console.log('Location received:', coords),
  },
};

// Note: This component has complex interactions with browser geolocation APIs
// and device settings that are difficult to fully simulate in Storybook.
// The stories provide a visual representation, but full functionality
// requires testing in a real browser environment.
