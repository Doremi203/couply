import EnterInfoPage from './Page';
import { StoryObj, Meta } from '@storybook/react';
import { BrowserRouter } from 'react-router-dom';

const meta: Meta = {
  title: 'Pages/EnterInfoPage',
  component: EnterInfoPage,
  decorators: [
    (Story) => (
      <BrowserRouter>
        <Story />
      </BrowserRouter>
    ),
  ],
  parameters: {
    layout: 'fullscreen',
  },
};

export default meta;
type Story = StoryObj;

export const Default: Story = {
  args: {},
};