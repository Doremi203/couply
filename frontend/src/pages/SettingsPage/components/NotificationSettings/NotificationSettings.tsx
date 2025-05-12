// import ArticleIcon from '@mui/icons-material/Article';
// import BedtimeIcon from '@mui/icons-material/Bedtime';
// import EmailIcon from '@mui/icons-material/Email';
// import NotificationsIcon from '@mui/icons-material/Notifications';
// import SmsIcon from '@mui/icons-material/Sms';
// import { Switch } from '@mui/material';
// import React, { useState } from 'react';

// // import { useTheme } from '../../lib/context/ThemeContext';

// import { useTheme } from '../../../../shared/lib/context/ThemeContext';

// import styles from './notificationSettings.module.css';

// interface NotificationOption {
//   id: string;
//   label: string;
//   icon: React.ReactNode;
//   enabled: boolean;
// }

// interface NotificationSettingsProps {
//   className?: string;
// }

// export const NotificationSettings: React.FC<NotificationSettingsProps> = ({ className }) => {
//   const { theme, toggleTheme } = useTheme();

//   const [notifications, setNotifications] = useState<NotificationOption[]>([
//     {
//       id: 'push',
//       label: 'Push-уведомления',
//       icon: <NotificationsIcon />,
//       enabled: true,
//     },
//     {
//       id: 'theme',
//       label: 'Темная тема',
//       icon: <BedtimeIcon />,
//       enabled: false,
//     },
//   ]);

//   const handleToggle = (id: string) => {
//     if (id === 'theme') {
//       toggleTheme();
//     }
//     // setNotifications(
//     //   notifications.map(notification =>
//     //     notification.id === id ? { ...notification, enabled: !notification.enabled } : notification,
//     //   ),
//     // );
//   };

//   return (
//     <div className={`${styles.container} ${className || ''}`}>
//       {notifications.map(notification => (
//         <div key={notification.id} className={styles.notificationItem}>
//           <div className={styles.iconAndLabel}>
//             <div className={styles.icon}>{notification.icon}</div>
//             <span className={styles.label}>{notification.label}</span>
//           </div>
//           <Switch
//             checked={notification.enabled}
//             onChange={() => handleToggle(notification.id)}
//             color="primary"
//             className={styles.switch}
//           />
//         </div>
//       ))}
//     </div>
//   );
// };

// export default NotificationSettings;

import BedtimeIcon from '@mui/icons-material/Bedtime';
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

  // Инициализация состояния с учетом текущей темы
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
        //@ts-ignore
        enabled: theme.isDark,
      },
    ]);
    // ts-ignore
  }, []);

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
