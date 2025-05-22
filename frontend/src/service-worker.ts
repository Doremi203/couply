/// <reference lib="webworker" />
import { precacheAndRoute } from 'workbox-precaching';

declare const self: ServiceWorkerGlobalScope;

precacheAndRoute(self.__WB_MANIFEST);

self.addEventListener('install', event => {
  event.waitUntil(self.skipWaiting());
});

self.addEventListener('activate', event => {
  // сразу начать контролировать все страницы под своим scope
  event.waitUntil(self.clients.claim());
});

self.addEventListener('push', evt => {
  try {
    if (!evt.data) {
      return;
    }
    const data = evt.data.json();
    if (!data.title || !data.body || !data.icon) {
      console.error('Invalid push notification data:', data);
      return;
    }
    const options = {
      body: data.body,
      icon: data.icon,
      data: data.url || '/',
    };
    evt.waitUntil(self.registration.showNotification(data.title, options));
  } catch (e) {
    console.error('Error parsing push notification data:', e);
  }
});

self.addEventListener('notificationclick', evt => {
  evt.notification.close();
  const url = evt.notification.data.url;

  evt.waitUntil(
    self.clients
      .matchAll({
        type: 'window',
        includeUncontrolled: true,
      })
      .then(list => {
        for (const client of list) {
          if (client.url === url && 'focus' in client) {
            return client.focus();
          }
        }
        if (self.clients.openWindow) {
          return self.clients.openWindow(url);
        }
      }),
  );
});
