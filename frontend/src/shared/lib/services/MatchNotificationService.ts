/**
 * Сервис для отправки уведомлений о совпадениях
 */

interface MatchNotificationData {
  userId: string;
  matchId: number;
  matchName: string;
  matchImage: string;
}

/**
 * Отправляет запрос на сервер для отправки push-уведомления о совпадении
 */
export const sendMatchNotification = async (data: MatchNotificationData): Promise<boolean> => {
  try {
    if ('Notification' in window && Notification.permission === 'granted') {
      if (!navigator.serviceWorker.controller) {
        new Notification(`Новое совпадение с ${data.matchName}!`, {
          body: 'У вас взаимная симпатия!',
          icon: data.matchImage,
        });
      } else {
        navigator.serviceWorker.ready.then(registration => {
          registration.showNotification(`Новое совпадение с ${data.matchName}!`, {
            body: 'У вас взаимная симпатия!',
            icon: data.matchImage,
            badge: '/icon512_maskable.png',
            data: {
              url: '/likes',
              dateOfArrival: Date.now(),
              primaryKey: data.matchId,
            },
          });
        });
      }
    }

    return true;
  } catch (error) {
    console.error('Error sending match notification:', error);
    return false;
  }
};
