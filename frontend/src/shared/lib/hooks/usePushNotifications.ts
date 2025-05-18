import { useState, useEffect } from 'react';

import {
  isPushNotificationSupported,
  askUserPermission,
  registerServiceWorker,
  createPushSubscription,
  sendSubscriptionToServer,
  unsubscribeFromPushNotifications,
} from '../services/PushNotificationService';

interface PushNotificationState {
  isSupported: boolean;
  permission: NotificationPermission | null;
  subscription: PushSubscription | null;
  isSubscribed: boolean;
  isInitializing: boolean;
  error: Error | null;
}

interface UsePushNotificationsReturn extends PushNotificationState {
  initialize: (userId: string) => Promise<boolean>;
  subscribe: () => Promise<boolean>;
  unsubscribe: () => Promise<boolean>;
}

/**
 * Хук для работы с push-уведомлениями
 */
export const usePushNotifications = (): UsePushNotificationsReturn => {
  const [state, setState] = useState<PushNotificationState>({
    isSupported: false,
    permission: null,
    subscription: null,
    isSubscribed: false,
    isInitializing: true,
    error: null,
  });

  // Проверяем поддержку push-уведомлений при монтировании компонента
  useEffect(() => {
    const checkSupport = async () => {
      try {
        const supported = isPushNotificationSupported();

        if (supported) {
          const permission = Notification.permission;
          setState(prev => ({
            ...prev,
            isSupported: supported,
            permission,
            isInitializing: false,
          }));
        } else {
          setState(prev => ({
            ...prev,
            isSupported: false,
            isInitializing: false,
          }));
        }
      } catch (error) {
        setState(prev => ({
          ...prev,
          error: error as Error,
          isInitializing: false,
        }));
      }
    };

    checkSupport();
  }, []);

  // Функция для инициализации push-уведомлений
  const initialize = async (userId: string): Promise<boolean> => {
    if (!state.isSupported) {
      return false;
    }

    setState(prev => ({ ...prev, isInitializing: true }));

    try {
      // Запрашиваем разрешение на отправку уведомлений
      const permission = await askUserPermission();

      if (permission !== 'granted') {
        setState(prev => ({
          ...prev,
          permission,
          isInitializing: false,
        }));
        return false;
      }

      // Регистрируем сервис-воркер
      const registration = await registerServiceWorker();

      if (!registration) {
        throw new Error('Failed to register service worker');
      }

      // Создаем подписку на push-уведомления
      const subscription = await createPushSubscription(registration);

      if (!subscription) {
        throw new Error('Failed to create push subscription');
      }

      // Отправляем подписку на сервер
      const success = await sendSubscriptionToServer(subscription, userId);

      if (!success) {
        throw new Error('Failed to send subscription to server');
      }

      setState(prev => ({
        ...prev,
        permission,
        subscription,
        isSubscribed: true,
        isInitializing: false,
      }));

      return true;
    } catch (error) {
      setState(prev => ({
        ...prev,
        error: error as Error,
        isInitializing: false,
      }));
      return false;
    }
  };

  // Функция для подписки на push-уведомления
  const subscribe = async (): Promise<boolean> => {
    if (!state.isSupported) {
      return false;
    }

    setState(prev => ({ ...prev, isInitializing: true }));

    try {
      // Запрашиваем разрешение на отправку уведомлений
      const permission = await askUserPermission();

      if (permission !== 'granted') {
        setState(prev => ({
          ...prev,
          permission,
          isInitializing: false,
        }));
        return false;
      }

      // Регистрируем сервис-воркер
      const registration = await registerServiceWorker();

      if (!registration) {
        throw new Error('Failed to register service worker');
      }

      // Создаем подписку на push-уведомления
      const subscription = await createPushSubscription(registration);

      if (!subscription) {
        throw new Error('Failed to create push subscription');
      }

      // Отправляем подписку на сервер
      const userId = localStorage.getItem('userId') || 'user123';
      const success = await sendSubscriptionToServer(subscription, userId);

      if (!success) {
        throw new Error('Failed to send subscription to server');
      }

      setState(prev => ({
        ...prev,
        permission,
        subscription,
        isSubscribed: true,
        isInitializing: false,
      }));

      return true;
    } catch (error) {
      setState(prev => ({
        ...prev,
        error: error as Error,
        isInitializing: false,
      }));
      return false;
    }
  };

  // Функция для отписки от push-уведомлений
  const unsubscribe = async (): Promise<boolean> => {
    if (!state.isSupported || !state.subscription) {
      return false;
    }

    setState(prev => ({ ...prev, isInitializing: true }));

    try {
      // Отписываемся от push-уведомлений
      const success = await unsubscribeFromPushNotifications(state.subscription);

      if (!success) {
        throw new Error('Failed to unsubscribe from push notifications');
      }

      setState(prev => ({
        ...prev,
        subscription: null,
        isSubscribed: false,
        isInitializing: false,
      }));

      return true;
    } catch (error) {
      setState(prev => ({
        ...prev,
        error: error as Error,
        isInitializing: false,
      }));
      return false;
    }
  };

  return {
    ...state,
    initialize,
    subscribe,
    unsubscribe,
  };
};
