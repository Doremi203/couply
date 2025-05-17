import BedtimeIcon from '@mui/icons-material/Bedtime';
import CreditCardOffOutlinedIcon from '@mui/icons-material/CreditCardOffOutlined';
import NotificationsIcon from '@mui/icons-material/Notifications';
import { Switch } from '@mui/material';
import React, { useState, useEffect } from 'react';

import { useTheme } from '../../../../shared/lib/context/ThemeContext';

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

  useEffect(() => {
    setNotifications([
      {
        id: 'push',
        label: 'Push-уведомления',
        icon: <NotificationsIcon />,
        enabled: true,
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
  }, [theme]);

  const handleToggle = (id: string) => {
    if (id === 'theme') {
      toggleTheme();
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
            //@ts-ignore
            checked={notification.id === 'theme' ? theme.isDark : notification.enabled}
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
