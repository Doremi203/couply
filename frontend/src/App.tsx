import './App.css';
import { useEffect, useState } from 'react';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';

import { AuthPage } from './pages/AuthPage';
import { EnterInfoPage } from './pages/EnterInfoPage';
import { HomePage } from './pages/HomePage';
import { LikesPage } from './pages/LikesPage';
import { ProfilePage } from './pages/ProfilePage';
import { SplashPage } from './pages/SplashPage';
import { ThemeProvider } from './shared/lib/context/ThemeContext';
import { usePushNotifications } from './shared/lib/hooks/usePushNotifications';

const router = createBrowserRouter([
  {
    path: 'auth',
    element: <AuthPage />,
  },
  {
    path: 'enterInfo',
    element: <EnterInfoPage />,
  },
  {
    path: '/',
    element: <SplashPage />,
  },
  {
    path: 'home',
    element: <HomePage />,
  },
  {
    path: 'profile',
    element: <ProfilePage />,
  },
  {
    path: 'likes',
    element: <LikesPage />,
  },
]);

function App() {
  const [userId] = useState('user123'); // В реальном приложении это должно быть получено из аутентификации
  const { isSupported, permission, initialize, isInitializing } = usePushNotifications();

  // Инициализация push-уведомлений только если разрешение уже получено
  useEffect(() => {
    if (isSupported && permission === 'granted' && !isInitializing) {
      initialize(userId).then(success => {
        console.log('Push notifications initialized:', success);
      });
    }
  }, [isSupported, permission, isInitializing, initialize, userId]);

  return (
    <ThemeProvider>
      <RouterProvider router={router} />
    </ThemeProvider>
  );
}

export default App;
