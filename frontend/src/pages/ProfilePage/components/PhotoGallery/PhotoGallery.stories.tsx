import type { Meta, StoryObj } from '@storybook/react';
import { PhotoGallery } from './PhotoGallery';

const meta = {
  title: 'Components/PhotoGallery',
  component: PhotoGallery,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
} satisfies Meta<typeof PhotoGallery>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    photos: [
      '/photo1.png',
      '/woman1.jpg',
      '/man1.jpg',
      '/photo1.png',
      '/woman1.jpg',
      '/man1.jpg',
    ],
  },
};

export const SinglePhoto: Story = {
  args: {
    photos: ['/photo1.png'],
  },
};

export const TwoPhotos: Story = {
  args: {
    photos: ['/photo1.png', '/woman1.jpg'],
  },
};