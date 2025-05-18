/**
 * Сервис для работы с push-уведомлениями
 */

// Публичный VAPID ключ (в реальном приложении должен быть получен с сервера)
const PUBLIC_VAPID_KEY =
  'BLBz-HVJLnXooDAeahFWQK8qQwJgm8Tw0HOTW_-qTINeW5-3CRe7NKfXWGNzCj_y7A6qYSHJZsYYCwL1Vd4kSFo';

// URL для отправки подписки на сервер
const SUBSCRIPTION_ENDPOINT = '/v1/push/subscribe';
const UNSUBSCRIPTION_ENDPOINT = '/v1/push/unsubscribe';

/**
 * Проверяет поддержку push-уведомлений в браузере
 */
export const isPushNotificationSupported = (): boolean => {
  return 'serviceWorker' in navigator && 'PushManager' in window;
};

/**
 * Запрашивает разрешение на отправку уведомлений
 */
export const askUserPermission = async (): Promise<NotificationPermission> => {
  return await Notification.requestPermission();
};

/**
 * Регистрирует сервис-воркер
 */
export const registerServiceWorker = async (): Promise<ServiceWorkerRegistration | null> => {
  if (!isPushNotificationSupported()) {
    console.log('Push notifications not supported');
    return null;
  }

  try {
    const registration = await navigator.serviceWorker.register('/service-worker.js');
    console.log('Service Worker registered successfully', registration);
    return registration;
  } catch (error) {
    console.error('Service Worker registration failed:', error);
    return null;
  }
};

/**
 * Создает подписку на push-уведомления
 */
export const createPushSubscription = async (
  registration: ServiceWorkerRegistration,
): Promise<PushSubscription | null> => {
  try {
    const subscription = await registration.pushManager.subscribe({
      userVisibleOnly: true,
      applicationServerKey: urlBase64ToUint8Array(PUBLIC_VAPID_KEY),
    });

    console.log('Push subscription created:', subscription);
    return subscription;
  } catch (error) {
    console.error('Error creating push subscription:', error);
    return null;
  }
};

/**
 * Отправляет подписку на сервер
 */
export const sendSubscriptionToServer = async (
  subscription: PushSubscription,
  _userId: string, // Unused parameter but kept for API compatibility
): Promise<boolean> => {
  try {
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
        p256dh: btoa(String.fromCharCode(...new Uint8Array(subscription.getKey('p256dh') as ArrayBuffer))),
        authKey: btoa(String.fromCharCode(...new Uint8Array(subscription.getKey('auth') as ArrayBuffer))),
      }),
    });

    if (!response.ok) {
      throw new Error('Failed to send subscription to server');
    }

    console.log('Subscription sent to server successfully');
    return true;
  } catch (error) {
    console.error('Error sending subscription to server:', error);
    return false;
  }
};

/**
 * Отправляет запрос на отписку от push-уведомлений
 */
export const unsubscribeFromPushNotifications = async (
  subscription: PushSubscription,
): Promise<boolean> => {
  try {
    const token = localStorage.getItem('token');
    if (!token) {
      throw new Error('User token not found');
    }

    // Отправляем запрос на сервер для отписки
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

    // Отписываемся на стороне браузера
    await subscription.unsubscribe();

    console.log('Successfully unsubscribed from push notifications');
    return true;
  } catch (error) {
    console.error('Error unsubscribing from push notifications:', error);
    return false;
  }
};

/**
 * Инициализирует push-уведомления
 */
export const initializePushNotifications = async (userId: string): Promise<boolean> => {
  if (!isPushNotificationSupported()) {
    console.log('Push notifications not supported');
    return false;
  }

  try {
    const permission = await askUserPermission();
    if (permission !== 'granted') {
      console.log('Notification permission not granted');
      return false;
    }

    const registration = await registerServiceWorker();
    if (!registration) {
      return false;
    }

    const subscription = await createPushSubscription(registration);
    if (!subscription) {
      return false;
    }

    const success = await sendSubscriptionToServer(subscription, userId);
    return success;
  } catch (error) {
    console.error('Error initializing push notifications:', error);
    return false;
  }
};

/**
 * Преобразует base64 строку в Uint8Array
 * (необходимо для applicationServerKey)
 */
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
