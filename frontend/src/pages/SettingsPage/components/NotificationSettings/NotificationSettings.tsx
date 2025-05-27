import BedtimeIcon from '@mui/icons-material/Bedtime';
import CreditCardOffOutlinedIcon from '@mui/icons-material/CreditCardOffOutlined';
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
import ConfirmModal from '../../../../shared/components/ConfirmModal';

import styles from './notificationSettings.module.css';
import {
  useGetSubscriptionMutation,
  useCancelSubscriptionMutation,
} from '../../../../entities/subscription/api/subscriptionApi.ts';

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
  const [getSubscription] = useGetSubscriptionMutation();
  const [cancelSubscription] = useCancelSubscriptionMutation();
  const [isPremium, setIsPremium] = useState(false);
  const [showUnsubscribeModal, setShowUnsubscribeModal] = useState(false);
  const [subscriptionId, setSubscriptionId] = useState<string>('');

  useEffect(() => {
    const fetchSubscription = async () => {
      try {
        const subscription = await getSubscription({}).unwrap();
        setIsPremium(subscription.status === 'SUBSCRIPTION_STATUS_ACTIVE');
        setSubscriptionId(subscription.subscriptionId);
      } catch (error) {
        console.error(error);
      }
    };

    fetchSubscription();
  }, [getSubscription]);

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
      {
        id: 'unsubscribe',
        label: 'Отменить подписку',
        icon: <CreditCardOffOutlinedIcon />,
        enabled: isPremium,
      },
    ]);
  }, [theme, isPushEnabled, isPremium]);

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
      } else if (id === 'unsubscribe') {
        setShowUnsubscribeModal(true);
      }
    } catch (error) {
      console.error('Error toggling:', error);
      alert('Произошла ошибка при изменении настроек');
    } finally {
      setIsLoading(false);
      setActiveId(null);
    }
  };

  const handleUnsubscribe = async () => {
    try {
      await cancelSubscription({ subscriptionId }).unwrap();
      setIsPremium(false);
      setShowUnsubscribeModal(false);
      setNotifications(prev =>
        prev.map(notification =>
          notification.id === 'unsubscribe' ? { ...notification, enabled: false } : notification,
        ),
      );
    } catch (error) {
      console.error('Error canceling subscription:', error);
      alert('Произошла ошибка при отмене подписки');
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

      <ConfirmModal
        isOpen={showUnsubscribeModal}
        onClose={() => setShowUnsubscribeModal(false)}
        onConfirm={handleUnsubscribe}
        title="Отмена подписки"
        message="Вы уверены, что хотите отменить подписку? После отмены вы потеряете доступ к премиум-функциям."
        confirmText="Отменить подписку"
        cancelText="Вернуться"
      />
    </div>
  );
};

export default NotificationSettings;
