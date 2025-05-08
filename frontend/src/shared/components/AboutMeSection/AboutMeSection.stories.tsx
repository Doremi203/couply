import type { Meta, StoryObj } from '@storybook/react';

import { AboutMeSection } from './AboutMeSection';

const meta = {
  title: 'Shared/AboutMeSection',
  component: AboutMeSection,
  parameters: {
    layout: 'centered',
    docs: {
      description: {
        component: 'A section for users to input and display their bio or about me information.',
      },
    },
  },
  tags: ['autodocs'],
  decorators: [
    Story => (
      <div style={{ width: '100%', maxWidth: '500px' }}>
        <Story />
      </div>
    ),
  ],
} satisfies Meta<typeof AboutMeSection>;

export default meta;
type Story = StoryObj<typeof AboutMeSection>;

export const Empty: Story = {
  args: {
    about: '',
    onInputChange: (field, value) => console.log(`${field} changed to: ${value}`),
  },
};

export const WithContent: Story = {
  args: {
    about:
      'I love hiking, photography, and exploring new places. I work as a software engineer and enjoy solving complex problems.',
    onInputChange: (field, value) => console.log(`${field} changed to: ${value}`),
  },
};

export const CustomTitle: Story = {
  args: {
    about: 'I am a passionate chef who loves to experiment with new recipes and flavors.',
    onInputChange: (field, value) => console.log(`${field} changed to: ${value}`),
    title: 'Bio',
  },
};

export const CustomPlaceholder: Story = {
  args: {
    about: '',
    onInputChange: (field, value) => console.log(`${field} changed to: ${value}`),
    placeholder: 'Share your interests, hobbies, and what makes you unique...',
  },
};

export const SmallerMaxLength: Story = {
  args: {
    about: 'This is a short bio with a smaller character limit.',
    onInputChange: (field, value) => console.log(`${field} changed to: ${value}`),
    maxLength: 100,
  },
};

export const LongText: Story = {
  args: {
    about:
      'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam auctor, nisl eget ultricies tincidunt, nisl nisl aliquam nisl, eget ultricies nisl nisl eget nisl. Nullam auctor, nisl eget ultricies tincidunt, nisl nisl aliquam nisl, eget ultricies nisl nisl eget nisl. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam auctor, nisl eget ultricies tincidunt, nisl nisl aliquam nisl, eget ultricies nisl nisl eget nisl. Nullam auctor, nisl eget ultricies tincidunt, nisl nisl aliquam nisl, eget ultricies nisl nisl eget nisl.',
    onInputChange: (field, value) => console.log(`${field} changed to: ${value}`),
  },
};
