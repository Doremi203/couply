import type { Meta, StoryObj } from '@storybook/react';
import { MemoryRouter, Routes, Route } from 'react-router-dom';

import { RegistrationPage } from './Page';

const meta = {
  title: 'Pages/RegistrationPage',
  component: RegistrationPage,
  parameters: {
    layout: 'fullscreen',
    docs: {
      description: {
        component: 'Registration page for entering phone/email and password.',
      },
    },
  },
  tags: ['autodocs'],
} satisfies Meta<typeof RegistrationPage>;

export default meta;
type Story = StoryObj<typeof RegistrationPage>;

export const Registration: Story = {
  decorators: [
    Story => (
      <MemoryRouter initialEntries={[{ pathname: '/registration' }]}>
        <Routes>
          <Route path="/registration" element={<Story />} />
        </Routes>
      </MemoryRouter>
    ),
  ],
};
