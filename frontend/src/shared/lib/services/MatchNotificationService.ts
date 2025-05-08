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
    // В реальном приложении здесь будет запрос к серверу
    // Сервер будет отправлять push-уведомление на устройство пользователя

    // Имитация запроса к серверу
    console.log('Sending match notification to server:', data);

    // Для демонстрации, создаем локальное уведомление, если push API недоступен
    if ('Notification' in window && Notification.permission === 'granted') {
      // Если сервис-воркер не зарегистрирован, показываем обычное уведомление
      if (!navigator.serviceWorker.controller) {
        new Notification(`Новое совпадение с ${data.matchName}!`, {
          body: 'У вас взаимная симпатия!',
          icon: data.matchImage,
        });
      } else {
        // Если сервис-воркер зарегистрирован, отправляем через него
        // В реальном приложении это будет делать сервер
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
            // actions не поддерживается в типе NotificationOptions
          });
        });
      }
    }

    // В реальном приложении здесь будет проверка ответа от сервера
    return true;
  } catch (error) {
    console.error('Error sending match notification:', error);
    return false;
  }
};
