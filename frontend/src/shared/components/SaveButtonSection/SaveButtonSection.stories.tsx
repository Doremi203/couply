import type { Meta, StoryObj } from '@storybook/react';

import { SaveButtonSection } from './SaveButtonSection';

const meta = {
  title: 'Shared/SaveButtonSection',
  component: SaveButtonSection,
  parameters: {
    layout: 'centered',
    docs: {
      description: {
        component: 'A section with a save button, typically used at the bottom of forms.',
      },
    },
  },
  tags: ['autodocs'],
  decorators: [
    Story => (
      <div style={{ width: '100%', maxWidth: '400px' }}>
        <Story />
      </div>
    ),
  ],
} satisfies Meta<typeof SaveButtonSection>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    onSave: () => console.log('Save button clicked'),
  },
};
