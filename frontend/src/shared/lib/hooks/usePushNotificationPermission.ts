import { useCallback, useEffect, useState } from 'react';

export function usePushNotificationPermission() {
  const [permission, setPermission] = useState<NotificationPermission>(Notification.permission);

  useEffect(() => {
    const handler = () => setPermission(Notification.permission);
    document.addEventListener('visibilitychange', handler);
    return () => {
      document.removeEventListener('visibilitychange', handler);
    };
  }, []);

  const requestPermission = useCallback(async () => {
    if (permission !== 'granted') {
      const result = await Notification.requestPermission();
      setPermission(result);
      return result;
    }
    return permission;
  }, [permission]);

  return { permission, requestPermission };
}
