import BedtimeIcon from '@mui/icons-material/Bedtime';
import NotificationsIcon from '@mui/icons-material/Notifications';
import { CircularProgress, Switch } from '@mui/material';
import React, { useEffect, useState } from 'react';

import { useTheme } from '../../../../shared/lib/context/ThemeContext';
import { usePushNotificationPermission } from '../../../../shared/lib/hooks/usePushNotificationPermission';
import { usePushSubscription } from '../../../../shared/lib/hooks/usePushSubscription';
import {
  sendSubscriptionToServer,
  unsubscribeFromPushNotifications,
} from '../../../../shared/lib/services/PushNotificationService.ts';

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
  const [notifications, setNotifications] = useState<NotificationOption[]>([]);
  const [isLoading, setIsLoading] = useState(false);
  const [activeId, setActiveId] = useState<string | null>(null);
  const { subscription, subscribe, unsubscribe } = usePushSubscription();
  const { requestPermission } = usePushNotificationPermission();
  const isPushEnabled = Boolean(subscription);

  useEffect(() => {
    setNotifications([
      {
        id: 'push',
        label: 'Push-уведомления',
        icon: <NotificationsIcon />,
        enabled: isPushEnabled,
      },
      {
        id: 'theme',
        label: 'Темная тема',
        icon: <BedtimeIcon />,
        enabled: theme === 'dark',
      },
    ]);
  }, [theme, isPushEnabled]);

  const handleToggle = async (id: string) => {
    if (isLoading) return;

    setActiveId(id);
    setIsLoading(true);

    try {
      if (id === 'theme') {
        toggleTheme();
      } else if (id === 'push') {
        if (!isPushEnabled) {
          const permissionResult = await requestPermission();
          if (permissionResult === 'granted') {
            const newSubscription = await subscribe();
            await sendSubscriptionToServer(newSubscription);
          }
        } else if (subscription) {
          await unsubscribeFromPushNotifications(subscription);
          await unsubscribe();
        }
      }
    } catch (error) {
      console.error('Error toggling:', error);
      alert('Произошла ошибка при изменении настроек');
    } finally {
      setIsLoading(false);
      setActiveId(null);
    }
  };

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
                    ? isPushEnabled
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
