import { StoryObj, Meta } from '@storybook/react';
import { BrowserRouter } from 'react-router-dom';

import Page from './Page';

const meta: Meta = {
  title: 'Pages/SplashPage',
  component: Page,
  decorators: [
    Story => (
      <BrowserRouter>
        <Story />
      </BrowserRouter>
    ),
  ],
};

export default meta;
type Story = StoryObj;

export const Default: Story = {};
