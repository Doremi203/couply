import './App.css';
import { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';

import { getUserId, setUserId } from './entities/user/model/userSlice';
import { AuthPage } from './pages/AuthPage';
import { EnterInfoPage } from './pages/EnterInfoPage';
import { HomePage } from './pages/HomePage';
import { LikesPage } from './pages/LikesPage';
import LoginPage from './pages/LoginPage';
import { ProfilePage } from './pages/ProfilePage';
import RegistrationPage from './pages/RegistrationPage';
import { SplashPage } from './pages/SplashPage';
import { ThemeProvider } from './shared/lib/context/ThemeContext';
import { usePushNotifications } from './shared/lib/hooks/usePushNotifications';

const router = createBrowserRouter([
  {
    path: '/auth',
    element: <AuthPage />,
  },
  {
    path: 'login',
    element: <LoginPage />,
  },
  {
    path: 'registration',
    element: <RegistrationPage />,
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
    element: <HomePage key="home-page" />,
  },
  {
    path: 'profile',
    element: <ProfilePage key="profile-page" />,
  },
  {
    path: 'likes',
    element: <LikesPage key="likes-page" />,
  },
]);

function App() {
  const dispatch = useDispatch();

  // Get userId directly from Redux store
  const userId = useSelector(getUserId);
  const { isSupported, permission, initialize, isInitializing } = usePushNotifications();

  // Load userId from localStorage on initial load
  useEffect(() => {
    if (!userId) {
      const storedUserId = localStorage.getItem('userId');
      if (storedUserId) {
        dispatch(setUserId(storedUserId));
      }
    }
  }, [userId, dispatch]);

  // Инициализация push-уведомлений только если разрешение уже получено
  useEffect(() => {
    if (isSupported && permission === 'granted' && !isInitializing) {
      initialize(userId || 'user123').then(success => {
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
