import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { Provider } from 'react-redux';

import './index.css';
import {registerSW} from 'virtual:pwa-register';

import { store } from './app/store';
import App from './App.tsx';


registerSW({immediate: true});

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <Provider store={store}>
      <App />
    </Provider>
  </StrictMode>,
);
