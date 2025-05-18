// Service Worker для обработки push-уведомлений

self.addEventListener('push', evt => {
  let data = {};
  try {
    data = evt.data.json();
  } catch (error) {
    console.error('Error parsing push notification data:', error);
  }
  const title = data.title || 'Новое уведомление';
  const opts = {
    body: data.body,
    icon: '/icon512_rounded.png',
    badge: '/icon512_maskable.png',
    vibrate: [100, 50, 100],
    data: data, // передадим метаданные (например, chatId)
  };
  evt.waitUntil(self.registration.showNotification(title, opts));
});

self.addEventListener('notificationclick', evt => {
  evt.notification.close();
  const url = evt.notification.data?.url || '/';
  evt.waitUntil(clients.openWindow(url));
});
