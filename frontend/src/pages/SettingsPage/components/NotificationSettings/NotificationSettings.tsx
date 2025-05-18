import BedtimeIcon from '@mui/icons-material/Bedtime';
import CreditCardOffOutlinedIcon from '@mui/icons-material/CreditCardOffOutlined';
import NotificationsIcon from '@mui/icons-material/Notifications';
import { Switch } from '@mui/material';
import React, { useState, useEffect } from 'react';

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
        id: 'theme',
        label: 'Отменить подписку',
        icon: <CreditCardOffOutlinedIcon />,
        //@ts-ignore
        enabled: theme.isDark, //TODO
      },
    ]);
  }, [theme, isSubscribed]);

  const handleToggle = async (id: string) => {
    if (id === 'theme') {
      toggleTheme();
    } else if (id === 'push') {
      if (isSupported) {
        if (isSubscribed) {
          await unsubscribe();
        } else {
          await subscribe();
        }
      } else {
        alert('Ваш браузер не поддерживает push-уведомления');
      }
    } else {
      setNotifications(prev =>
        prev.map(notification =>
          notification.id === id
            ? { ...notification, enabled: !notification.enabled }
            : notification,
        ),
      );
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
          <Switch
            checked={
              notification.id === 'theme'
                ? theme === 'dark'
                : notification.id === 'push'
                  ? isSubscribed
                  : notification.enabled
            }
            onChange={() => handleToggle(notification.id)}
            color="primary"
            className={styles.switch}
          />
        </div>
      ))}
    </div>
  );
};

export default NotificationSettings;
