import type { Meta, StoryObj } from '@storybook/react';

import { PageHeader } from './PageHeader';

const meta = {
  title: 'Shared/PageHeader',
  component: PageHeader,
  parameters: {
    layout: 'centered',
    docs: {
      description: {
        component: 'A header component for pages with a back button and title.',
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
} satisfies Meta<typeof PageHeader>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    onBack: () => console.log('Back button clicked'),
  },
};
