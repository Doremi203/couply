const SUBSCRIPTION_ENDPOINT = 'https://notificator.testing.couply.ru/v1/push/subscribe';
const UNSUBSCRIPTION_ENDPOINT = 'https://notificator.testing.couply.ru/v1/push/unsubscribe';

export const isPushNotificationSupported = (): boolean => {
  return 'serviceWorker' in navigator && 'PushManager' in window;
};

export const sendSubscriptionToServer = async (subscription: PushSubscription) => {
  const token = localStorage.getItem('token');
  if (!token) {
    throw new Error('User token not found');
  }

  const response = await fetch(SUBSCRIPTION_ENDPOINT, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'User-Token': token,
    },
    body: JSON.stringify({
      endpoint: subscription.endpoint,
      p256dh: btoa(
        String.fromCharCode(...new Uint8Array(subscription.getKey('p256dh') as ArrayBuffer)),
      ),
      authKey: btoa(
        String.fromCharCode(...new Uint8Array(subscription.getKey('auth') as ArrayBuffer)),
      ),
    }),
  });
  if (!response.ok) {
    throw new Error('Failed to send subscription to server');
  }
  console.log('Successfully subscribed to notifications');
};

export const unsubscribeFromPushNotifications = async (subscription: PushSubscription) => {
  const token = localStorage.getItem('token');
  if (!token) {
    throw new Error('User token not found');
  }

  const response = await fetch(UNSUBSCRIPTION_ENDPOINT, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'User-Token': token,
    },
    body: JSON.stringify({
      endpoint: subscription.endpoint,
    }),
  });
  if (!response.ok) {
    throw new Error('Failed to unsubscribe from server');
  }

  console.log('Successfully unsubscribed from push notifications');
};
