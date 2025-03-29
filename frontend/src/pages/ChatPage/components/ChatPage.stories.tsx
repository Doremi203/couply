import ChatPage from './ChatPage';
import { StoryObj, Meta } from '@storybook/react';
import { BrowserRouter } from 'react-router-dom';

const meta: Meta = {
  title: 'Pages/ChatPage',
  component: ChatPage,
  decorators: [
    (Story) => (
      <BrowserRouter>
        <Story />
      </BrowserRouter>
    ),
  ],
};

export default meta;
type Story = StoryObj;

export const Default: Story = {};