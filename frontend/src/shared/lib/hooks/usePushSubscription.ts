import { useCallback, useEffect, useState } from 'react';

const PUBLIC_VAPID_KEY =
  'BAX081_ZDseLaav3K0xcZN00ISk518x0F93ZOGuYSrsE_hhMGwoyQPUnI9rJ8-O27q2LCvWXKwzTdDwGORJwPGM';

export function usePushSubscription() {
  const [subscription, setSubscription] = useState<PushSubscription | null>(null);

  const getExistingSubscription = useCallback(async () => {
    const reg = await navigator.serviceWorker.ready;
    const sub = await reg.pushManager.getSubscription();
    setSubscription(sub);
    return sub;
  }, []);

  const subscribe = useCallback(async (): Promise<PushSubscription> => {
    if (Notification.permission !== 'granted') {
      throw new Error('Нет разрешения на уведомления');
    }
    const reg = await navigator.serviceWorker.ready;
    const sub = await reg.pushManager.subscribe({
      userVisibleOnly: true,
      applicationServerKey: urlBase64ToUint8Array(PUBLIC_VAPID_KEY),
    });
    setSubscription(sub);
    return sub;
  }, []);

  const unsubscribe = useCallback(async () => {
    if (!subscription) return false;
    const success = await subscription.unsubscribe();
    if (success) setSubscription(null);
    return success;
  }, [subscription]);

  useEffect(() => {
    getExistingSubscription().catch(err => {
      console.error(err);
    });
  }, [getExistingSubscription]);

  return { subscription, subscribe, unsubscribe };
}

function urlBase64ToUint8Array(base64String: string): Uint8Array {
  const padding = '='.repeat((4 - (base64String.length % 4)) % 4);
  const base64 = (base64String + padding).replace(/-/g, '+').replace(/_/g, '/');

  const rawData = window.atob(base64);
  const outputArray = new Uint8Array(rawData.length);

  for (let i = 0; i < rawData.length; ++i) {
    outputArray[i] = rawData.charCodeAt(i);
  }

  return outputArray;
}
