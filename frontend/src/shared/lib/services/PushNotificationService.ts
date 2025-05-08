/**
 * Сервис для работы с push-уведомлениями
 */

// Публичный VAPID ключ (в реальном приложении должен быть получен с сервера)
const PUBLIC_VAPID_KEY =
  'BLBz-HVJLnXooDAeahFWQK8qQwJgm8Tw0HOTW_-qTINeW5-3CRe7NKfXWGNzCj_y7A6qYSHJZsYYCwL1Vd4kSFo';

// URL для отправки подписки на сервер
const SUBSCRIPTION_ENDPOINT = 'https://api.example.com/push-subscriptions';

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
  userId: string,
): Promise<boolean> => {
  try {
    const response = await fetch(SUBSCRIPTION_ENDPOINT, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        subscription,
        userId,
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
