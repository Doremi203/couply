import { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';

import { getUserId, setUserId } from './entities/user/model/userSlice';
import { AuthPage } from './pages/AuthPage';
import { EnterInfoPage } from './pages/EnterInfoPage';
import { EnterPhonePage } from './pages/EnterPhonePage';
import { HomePage } from './pages/HomePage';
import { LikesPage } from './pages/LikesPage';
import LoginPage from './pages/LoginPage';
import { PremiumPage } from './pages/PremiumPage';
import { ProfilePage } from './pages/ProfilePage';
import RegistrationPage from './pages/RegistrationPage';
import { SettingsPage } from './pages/SettingsPage';
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
    path: 'enterPhone',
    element: <EnterPhonePage />,
  },
  {
    path: 'enterInfo',
    element: <EnterInfoPage />,
  },
  {
    path: 'settings',
    element: <SettingsPage />,
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
  {
    path: 'premium',
    element: <PremiumPage key="premium-page" />,
  },
]);

// Компонент для отображения сообщения о повороте устройства
const OrientationMessage: React.FC = () => {
  return (
    <div className="orientation-message">
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="var(--primary-color)">
        <path d="M0 0h24v24H0z" fill="none" />
        <path d="M17 1.01L7 1C5.9 1 5 1.9 5 3v18c0 1.1.9 2 2 2h10c1.1 0 2-.9 2-2V3c0-1.1-.9-1.99-2-1.99zM17 19H7V5h10v14z" />
      </svg>
      <p>Пожалуйста, поверните устройство в портретную ориентацию для лучшего взаимодействия</p>
    </div>
  );
};

function App() {
  const dispatch = useDispatch();
  const [isLandscape, setIsLandscape] = useState(false);

  // Get userId directly from Redux store
  const userId = useSelector(getUserId);
  const { isSupported, permission, subscribe, isInitializing } = usePushNotifications();

  // Load userId from localStorage on initial load
  useEffect(() => {
    if (!userId) {
      const storedUserId = localStorage.getItem('userId');
      if (storedUserId) {
        dispatch(setUserId(storedUserId));
      }
    }
  }, [userId, dispatch]);

  // Функция проверки ориентации экрана
  const checkOrientation = () => {
    // Проверяем, находимся ли мы на мобильном устройстве
    const isMobile = /Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(
      navigator.userAgent,
    );

    // Если это мобильное устройство и ширина экрана больше высоты, значит альбомная ориентация
    if (isMobile && window.innerWidth > window.innerHeight) {
      setIsLandscape(true);
    } else {
      setIsLandscape(false);
    }
  };

  // Проверка ориентации при загрузке и изменении размеров окна
  useEffect(() => {
    checkOrientation();
    window.addEventListener('resize', checkOrientation);
    window.addEventListener('orientationchange', checkOrientation);

    return () => {
      window.removeEventListener('resize', checkOrientation);
      window.removeEventListener('orientationchange', checkOrientation);
    };
  }, []);

  // Check for existing geolocation permission
  useEffect(() => {
    // Only check if geolocation is available
    if (navigator.geolocation) {
      // Use a one-time permission check to set the initial state
      navigator.permissions
        .query({ name: 'geolocation' })
        .then(permissionStatus => {
          if (permissionStatus.state === 'granted') {
            localStorage.setItem('userLocationAllowed', 'true');
          } else if (permissionStatus.state === 'denied') {
            localStorage.setItem('userLocationAllowed', 'false');
          }

          // Set up a listener for future changes
          permissionStatus.onchange = () => {
            localStorage.setItem(
              'userLocationAllowed',
              permissionStatus.state === 'granted' ? 'true' : 'false',
            );
          };
        })
        .catch(() => {
          // If permissions API fails, we'll rely on the manual permission setting
          // in GeoLocationRequest component
        });
    }
  }, []);

  // Подписка на push-уведомления только если разрешение уже получено
  useEffect(() => {
    if (isSupported && permission === 'granted' && !isInitializing) {
      subscribe().then(success => {
        console.log('Push notifications subscription status:', success);
      });
    }
  }, [isSupported, permission, isInitializing, subscribe]);

  return (
    <ThemeProvider>
      {isLandscape && <OrientationMessage />}
      <RouterProvider router={router} />
    </ThemeProvider>
  );
}

export default App;
