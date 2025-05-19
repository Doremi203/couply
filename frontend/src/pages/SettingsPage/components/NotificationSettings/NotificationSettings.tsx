import BedtimeIcon from '@mui/icons-material/Bedtime';
import CreditCardOffOutlinedIcon from '@mui/icons-material/CreditCardOffOutlined';
import NotificationsIcon from '@mui/icons-material/Notifications';
import { Switch, CircularProgress } from '@mui/material';
import React, { useState, useEffect, useCallback } from 'react';

import { useTheme } from '../../../../shared/lib/context/ThemeContext';
import { usePushNotifications } from '../../../../shared/lib/hooks/usePushNotifications';

import styles from './notificationSettings.module.css';

interface NotificationOption {
  id: string;
  label: string;
  icon: React.ReactNode;
  enabled: boolean;
}

interface NotificationSettingsProps {
  className?: string;
}

export const NotificationSettings: React.FC<NotificationSettingsProps> = ({ className }) => {
  const { theme, toggleTheme } = useTheme();
  const { isSupported, isSubscribed, subscribe, unsubscribe } = usePushNotifications();
  const [notifications, setNotifications] = useState<NotificationOption[]>([]);
  const [isLoading, setIsLoading] = useState(false);
  const [activeId, setActiveId] = useState<string | null>(null);

  useEffect(() => {
    setNotifications([
      {
        id: 'push',
        label: 'Push-уведомления',
        icon: <NotificationsIcon />,
        enabled: isSubscribed,
      },
      {
        id: 'theme',
        label: 'Темная тема',
        icon: <BedtimeIcon />,
        enabled: theme === 'dark',
      },
      {
        id: 'unsubscribe',
        label: 'Отменить подписку',
        icon: <CreditCardOffOutlinedIcon />,
        enabled: false,
      },
    ]);
  }, [theme, isSubscribed]);

  const handleToggle = useCallback(
    async (id: string) => {
      if (isLoading) return;

      setActiveId(id);
      setIsLoading(true);

      try {
        if (id === 'theme') {
          toggleTheme();
          // Для темы не нужен await, сразу заканчиваем обработку
          setIsLoading(false);
          setActiveId(null);
          return;
        } else if (id === 'push') {
          // if (!isSupported) {
          //   alert('Ваш браузер не поддерживает push-уведомления');
          //   setIsLoading(false);
          //   setActiveId(null);
          //   return;
          // }

          if (isSubscribed) {
            await unsubscribe();
          } else {
            await subscribe();
          }
          // await refreshState();
        } else if (id === 'unsubscribe') {
          // Простое переключение состояния без реальной логики
          setNotifications(prev =>
            prev.map(notification =>
              notification.id === id
                ? { ...notification, enabled: !notification.enabled }
                : notification,
            ),
          );
        }
      } catch (error) {
        console.error('Error toggling:', error);
        alert('Произошла ошибка при изменении настроек');
      } finally {
        setIsLoading(false);
        setActiveId(null);
      }
    },
    [isSupported, isSubscribed, subscribe, unsubscribe, toggleTheme],
  );

  return (
    <div className={`${styles.container} ${className || ''}`}>
      {notifications.map(notification => (
        <div key={notification.id} className={styles.notificationItem}>
          <div className={styles.iconAndLabel}>
            <div className={styles.icon}>{notification.icon}</div>
            <span className={styles.label}>{notification.label}</span>
          </div>
          <div style={{ position: 'relative', display: 'flex', alignItems: 'center' }}>
            {isLoading && activeId === notification.id && (
              <CircularProgress size={16} style={{ marginRight: 8 }} />
            )}
            <Switch
              checked={
                notification.id === 'theme'
                  ? theme === 'dark'
                  : notification.id === 'push'
                    ? isSubscribed
                    : notification.enabled
              }
              onChange={() => handleToggle(notification.id)}
              disabled={isLoading && activeId !== notification.id}
              color="primary"
              className={styles.switch}
            />
          </div>
        </div>
      ))}
    </div>
  );
};

export default NotificationSettings;
