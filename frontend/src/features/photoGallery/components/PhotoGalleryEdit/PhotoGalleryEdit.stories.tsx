import type { Meta, StoryObj } from '@storybook/react';

import { PhotoGalleryEdit } from './PhotoGalleryEdit';

const meta = {
  title: 'Features/PhotoGallery/PhotoGalleryEdit',
  component: PhotoGalleryEdit,
  parameters: {
    layout: 'centered',
    docs: {
      description: {
        component: 'A component for editing a photo gallery with add and remove functionality.',
      },
    },
  },
  tags: ['autodocs'],
  decorators: [
    Story => (
      <div style={{ width: '100%', maxWidth: '600px' }}>
        <Story />
      </div>
    ),
  ],
} satisfies Meta<typeof PhotoGalleryEdit>;

export default meta;
type Story = StoryObj<typeof PhotoGalleryEdit>;

// Sample photos
const samplePhotos = [
  'https://randomuser.me/api/portraits/women/68.jpg',
  'https://randomuser.me/api/portraits/women/69.jpg',
  'https://randomuser.me/api/portraits/women/70.jpg',
];

export const WithPhotos: Story = {
  args: {
    photos: samplePhotos,
    onPhotoRemove: index => console.log(`Removed photo at index: ${index}`),
    onAddPhotoClick: () => console.log('Add photo clicked'),
  },
};

export const Empty: Story = {
  args: {
    photos: [],
    onPhotoRemove: index => console.log(`Removed photo at index: ${index}`),
    onAddPhotoClick: () => console.log('Add photo clicked'),
  },
};

export const SinglePhoto: Story = {
  args: {
    photos: [samplePhotos[0]],
    onPhotoRemove: index => console.log(`Removed photo at index: ${index}`),
    onAddPhotoClick: () => console.log('Add photo clicked'),
  },
};

export const ManyPhotos: Story = {
  args: {
    photos: [
      ...samplePhotos,
      'https://randomuser.me/api/portraits/men/44.jpg',
      'https://randomuser.me/api/portraits/men/45.jpg',
      'https://randomuser.me/api/portraits/men/46.jpg',
    ],
    onPhotoRemove: index => console.log(`Removed photo at index: ${index}`),
    onAddPhotoClick: () => console.log('Add photo clicked'),
  },
};

export const CustomTitle: Story = {
  args: {
    photos: samplePhotos,
    onPhotoRemove: index => console.log(`Removed photo at index: ${index}`),
    onAddPhotoClick: () => console.log('Add photo clicked'),
    title: 'My Photo Gallery',
  },
};
