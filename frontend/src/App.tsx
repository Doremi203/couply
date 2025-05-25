import { Config, ConfigResponseMode } from '@vkid/sdk';
import { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { createBrowserRouter, Navigate, RouterProvider } from 'react-router-dom';

import { setBlocking } from './app/store/blockingSlice';
import { useGetBlockInfoMutation } from './entities/blocker/index.ts';
import { getUserId, setUserId } from './entities/user/index.ts';
import { AuthPage } from './pages/AuthPage';
import { BlockerPage } from './pages/BlockerPage/BlockerPage';
import { EnterInfoPage } from './pages/EnterInfoPage';
import { EnterPhonePage } from './pages/EnterPhonePage';
import { HomePage } from './pages/HomePage';
import { LikesPage } from './pages/LikesPage';
import LoginPage from './pages/LoginPage';
import OAuthCallback from './pages/OAuthCallback/OAuthCallback.tsx';
import { PremiumPage } from './pages/PremiumPage';
import { ProfilePage } from './pages/ProfilePage';
import RegistrationPage from './pages/RegistrationPage';
import { SettingsPage } from './pages/SettingsPage';
import { SplashPage } from './pages/SplashPage';
import { ThemeProvider } from './shared/lib/context/ThemeContext';

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
  {
    path: 'blocked',
    element: <BlockerPage />,
  },
  {
    path: '/oauth-callback',
    element: <OAuthCallback />,
  },
  {
    path: '*',
    element: <Navigate to="/" replace />,
  },
]);

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

  const [getBlockInfo] = useGetBlockInfoMutation();

  // Get userId directly from Redux store
  const userId = useSelector(getUserId);

  // Load userId from localStorage on initial load
  useEffect(() => {
    if (!userId) {
      const storedUserId = localStorage.getItem('userId');
      if (storedUserId) {
        dispatch(setUserId(storedUserId));
      }
    }
  }, [userId, dispatch]);

  useEffect(() => {
    const fetchBlockInfo = async () => {
      if (localStorage.getItem('token')) {
        try {
          const blockInfo = await getBlockInfo({}).unwrap();
          if (blockInfo && blockInfo.blockId) {
            dispatch(
              setBlocking({
                isBlocked: true,
                reasons: blockInfo.reasons,
                message: blockInfo.message,
                createdAt: blockInfo.createdAt,
              }),
            );
          } else {
            dispatch(
              setBlocking({
                isBlocked: false,
                reasons: [],
                message: '',
                createdAt: null,
              }),
            );
          }
        } catch (error) {
          console.error('Error checking block status:', error);
          dispatch(
            setBlocking({
              isBlocked: false,
              reasons: [],
              message: '',
              createdAt: null,
            }),
          );
        }
      }
    };

    fetchBlockInfo();
  }, [getBlockInfo, dispatch]);

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

  useEffect(() => {
    Config.init({
      app: 53582290,
      redirectUrl: 'https://testing.couply.ru',
      scope: 'email',
      responseMode: ConfigResponseMode.Callback,
    });
  }, []);

  return (
    <ThemeProvider>
      {isLandscape && <OrientationMessage />}
      <RouterProvider router={router} />
    </ThemeProvider>
  );
}

export default App;
