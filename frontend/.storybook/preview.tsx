import { configureStore } from '@reduxjs/toolkit';
import type { Preview } from '@storybook/react';
import React from 'react';
import { Provider } from 'react-redux';

import { baseApi } from '../src/shared/api/baseApi';
import { ThemeProvider } from '../src/shared/lib/context/ThemeContext';

// Import global styles
import '../src/index.css';

// Create a mock store for Storybook
const mockStore = configureStore({
  reducer: {
    [baseApi.reducerPath]: baseApi.reducer,
  },
  middleware: getDefaultMiddleware => getDefaultMiddleware().concat(baseApi.middleware),
});

// Define mobile phone viewport dimensions
const MOBILE_VIEWPORT = {
  name: 'Mobile',
  styles: {
    width: '375px',
    height: '812px',
  },
  type: 'mobile',
};

const preview: Preview = {
  parameters: {
    controls: {
      matchers: {
        color: /(background|color)$/i,
        date: /Date$/i,
      },
    },
    // Set default viewport to mobile
    viewport: {
      viewports: {
        mobile: MOBILE_VIEWPORT,
      },
      defaultViewport: 'mobile',
    },
    // Make the background match a typical phone screen
    backgrounds: {
      default: 'app',
      values: [
        {
          name: 'app',
          value: '#f8f9fa',
        },
      ],
    },
  },
  decorators: [
    (Story: React.ComponentType) => {
      return (
        <Provider store={mockStore}>
          <ThemeProvider>
            <div
              style={{
                width: '375px',
                height: '812px',
                margin: '0 auto',
                border: '1px solid #ddd',
                borderRadius: '32px',
                overflow: 'hidden',
                position: 'relative',
                boxShadow: '0 0 10px rgba(0,0,0,0.1)',
              }}
            >
              <Story />
            </div>
          </ThemeProvider>
        </Provider>
      );
    },
  ],
};

export default preview;
