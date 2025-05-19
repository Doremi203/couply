import React from 'react';

import { usePushNotifications } from '../../../../shared/lib/hooks/usePushNotifications';

import styles from './pushToggle.module.css';

interface PushToggleProps {
  className?: string;
}

export const PushToggle: React.FC<PushToggleProps> = ({ className }) => {
  const { isSupported, isSubscribed, subscribe, unsubscribe } = usePushNotifications();

  const handleToggle = async () => {
    if (!isSupported) {
      alert('Ваш браузер не поддерживает push-уведомления');
      return;
    }

    try {
      if (isSubscribed) {
        await unsubscribe();
      } else {
        await subscribe();
      }
    } catch (error) {
      console.error('Error toggling push notifications:', error);
      alert('Произошла ошибка при работе с уведомлениями');
    }
  };

  if (!isSupported) {
    return null;
  }

  return (
    <button onClick={handleToggle} className={`${styles.toggleButton} ${className || ''}`}>
      {isSubscribed ? 'Отключить уведомления' : 'Включить уведомления'}
    </button>
  );
};

export default PushToggle;
