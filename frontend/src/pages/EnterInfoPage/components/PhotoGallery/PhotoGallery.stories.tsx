import type { Meta, StoryObj } from '@storybook/react';

import { FixedPhotoGallery } from './PhotoGallery';

const meta = {
  title: 'Pages/EnterInfoPage/FixedPhotoGallery',
  component: FixedPhotoGallery,
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
} satisfies Meta<typeof FixedPhotoGallery>;

export default meta;
type Story = StoryObj<typeof meta>;

// Sample photo URLs
const samplePhotos = [
  'https://via.placeholder.com/150',
  'https://via.placeholder.com/150',
  'https://via.placeholder.com/150',
];

export const Default: Story = {
  args: {
    photos: [],
    onPhotoRemove: index => console.log('Photo removed at index:', index),
    onAddPhotoClick: () => console.log('Add photo clicked'),
    title: 'Фотографии',
  },
};

export const WithPhotos: Story = {
  args: {
    photos: samplePhotos,
    onPhotoRemove: index => console.log('Photo removed at index:', index),
    onAddPhotoClick: () => console.log('Add photo clicked'),
    title: 'Фотографии',
  },
};

export const FullGallery: Story = {
  args: {
    photos: [...samplePhotos, ...samplePhotos],
    onPhotoRemove: index => console.log('Photo removed at index:', index),
    onAddPhotoClick: () => console.log('Add photo clicked'),
    title: 'Фотографии',
  },
};
