import type { Meta, StoryObj } from '@storybook/react';

import { EditSection } from '../../../widgets/EditProfile/components/EditSection/EditSection';

const meta = {
  title: 'Shared/InterestsSection',
  component: EditSection,
  parameters: {
    layout: 'centered',
    docs: {
      description: {
        component: 'A section for inputting and displaying user interests.',
      },
    },
  },
  tags: ['autodocs'],
  decorators: [
    Story => (
      <div style={{ width: '350px', maxWidth: '400px', marginTop: '20px' }}>
        <Story />
      </div>
    ),
  ],
} satisfies Meta<typeof EditSection>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Empty: Story = {
  args: {
    title: 'Interests',
    placeholder: 'Add your interests (e.g. hiking, reading)',
    values: [],
    fieldName: 'interests',
    // onArrayInputChange: (field, value) => console.log(`${field} changed to ${value}`),
  },
};

export const WithValues: Story = {
  args: {
    title: 'Interests',
    placeholder: 'Add your interests (e.g. hiking, reading)',
    values: ['hiking', 'reading', 'cooking', 'photography'],
    fieldName: 'interests',
    // onArrayInputChange: (field, value) => console.log(`${field} changed to ${value}`),
  },
};

export const Hobbies: Story = {
  args: {
    title: 'Hobbies',
    placeholder: 'What do you enjoy doing?',
    values: ['painting', 'cycling', 'gaming'],
    fieldName: 'hobbies',
    // onArrayInputChange: (field, value) => console.log(`${field} changed to ${value}`),
  },
};

export const MusicTastes: Story = {
  args: {
    title: 'Music Tastes',
    placeholder: 'What music do you like?',
    values: ['rock', 'jazz', 'classical'],
    fieldName: 'musicTastes',
    // onArrayInputChange: (field, value) => console.log(`${field} changed to ${value}`),
  },
};
