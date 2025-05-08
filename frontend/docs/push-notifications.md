# Документация по Push-уведомлениям в Couply

## Содержание

1. [Введение](#введение)
2. [Архитектура системы уведомлений](#архитектура-системы-уведомлений)
3. [Настройка клиентской части](#настройка-клиентской-части)
   - [Service Worker](#service-worker)
   - [Web App Manifest](#web-app-manifest)
   - [Регистрация Service Worker](#регистрация-service-worker)
   - [Запрос разрешений](#запрос-разрешений)
   - [Подписка на уведомления](#подписка-на-уведомления)
4. [Настройка серверной части](#настройка-серверной-части)
   - [Генерация VAPID ключей](#генерация-vapid-ключей)
   - [Хранение подписок](#хранение-подписок)
   - [Отправка уведомлений](#отправка-уведомлений)
5. [Интеграция в приложение](#интеграция-в-приложение)
   - [Инициализация при запуске](#инициализация-при-запуске)
   - [Отправка уведомлений о совпадениях](#отправка-уведомлений-о-совпадениях)
6. [Тестирование](#тестирование)
7. [Устранение неполадок](#устранение-неполадок)
8. [Ограничения и совместимость](#ограничения-и-совместимость)

## Введение

Push-уведомления позволяют приложению Couply отправлять пользователям уведомления даже когда они не находятся в приложении. Это особенно полезно для оповещения о новых совпадениях (взаимных лайках), сообщениях и других важных событиях.

Система push-уведомлений в Couply использует Web Push API, который поддерживается большинством современных браузеров. Для работы push-уведомлений требуется:

1. Service Worker для обработки уведомлений
2. Разрешение пользователя на отправку уведомлений
3. Серверная часть для отправки уведомлений

## Архитектура системы уведомлений

Система push-уведомлений в Couply состоит из следующих компонентов:

1. **Service Worker** (`public/service-worker.js`) - обрабатывает входящие push-события и отображает уведомления
2. **PushNotificationService** (`src/shared/lib/services/PushNotificationService.ts`) - управляет регистрацией service worker и подпиской на уведомления
3. **usePushNotifications** (`src/shared/lib/hooks/usePushNotifications.ts`) - React-хук для удобного использования push-уведомлений в компонентах
4. **MatchNotificationService** (`src/shared/lib/services/MatchNotificationService.ts`) - сервис для отправки уведомлений о совпадениях
5. **Серверная часть** (не реализована в текущей версии) - отправляет push-уведомления подписанным пользователям

Процесс работы push-уведомлений:

1. Пользователь разрешает отправку уведомлений
2. Браузер регистрирует service worker и создает подписку
3. Подписка отправляется на сервер и сохраняется в базе данных
4. При возникновении события (например, взаимного лайка) сервер отправляет push-сообщение
5. Service worker получает push-сообщение и отображает уведомление
6. Пользователь может взаимодействовать с уведомлением (открыть приложение, закрыть уведомление)

## Настройка клиентской части

### Service Worker

Service Worker - это JavaScript-файл, который работает в фоновом режиме и обрабатывает push-уведомления. В Couply используется файл `public/service-worker.js`.

Основные функции Service Worker:

1. Обработка push-событий:

```javascript
self.addEventListener('push', function (event) {
  const data = event.data.json();

  const options = {
    body: data.body,
    icon: data.icon || '/icon512_rounded.png',
    badge: '/icon512_maskable.png',
    data: {
      url: data.url || '/',
      dateOfArrival: Date.now(),
      primaryKey: 1,
    },
  };

  event.waitUntil(self.registration.showNotification(data.title, options));
});
```

2. Обработка кликов по уведомлениям:

```javascript
self.addEventListener('notificationclick', function (event) {
  const notification = event.notification;
  const action = event.action;
  const url = notification.data.url;

  if (action === 'close') {
    notification.close();
  } else {
    event.waitUntil(
      clients.matchAll({ type: 'window' }).then(function (clientList) {
        // Если есть открытое окно, фокусируемся на нем
        for (let i = 0; i < clientList.length; i++) {
          const client = clientList[i];
          if (client.url === url && 'focus' in client) {
            return client.focus();
          }
        }
        // Если нет открытого окна, открываем новое
        if (clients.openWindow) {
          return clients.openWindow(url);
        }
      }),
    );
    notification.close();
  }
});
```

### Web App Manifest

Web App Manifest (`public/manifest.json`) - это JSON-файл, который содержит информацию о веб-приложении и позволяет установить его на устройство пользователя. Для push-уведомлений важно указать `gcm_sender_id`:

```json
{
  "name": "Couply",
  "short_name": "Couply",
  "description": "Приложение для знакомств",
  "start_url": "/",
  "display": "standalone",
  "background_color": "#202C83",
  "theme_color": "#202C83",
  "icons": [
    {
      "src": "icon512_rounded.png",
      "sizes": "512x512",
      "type": "image/png",
      "purpose": "any"
    },
    {
      "src": "icon512_maskable.png",
      "sizes": "512x512",
      "type": "image/png",
      "purpose": "maskable"
    }
  ],
  "gcm_sender_id": "103953800507"
}
```

### Регистрация Service Worker

Для регистрации Service Worker используется функция `registerServiceWorker` из `PushNotificationService`:

```typescript
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
```

### Запрос разрешений

Для отправки push-уведомлений необходимо получить разрешение пользователя. Это делается с помощью функции `askUserPermission`:

```typescript
export const askUserPermission = async (): Promise<NotificationPermission> => {
  return await Notification.requestPermission();
};
```

Возможные значения разрешения:

- `granted` - разрешено
- `denied` - запрещено
- `default` - пользователь не принял решение

### Подписка на уведомления

После получения разрешения и регистрации Service Worker, необходимо создать подписку на push-уведомления:

```typescript
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
```

Подписка содержит информацию, необходимую для отправки уведомлений конкретному пользователю, и должна быть отправлена на сервер.

## Настройка серверной части

### Генерация VAPID ключей

VAPID (Voluntary Application Server Identification) ключи используются для идентификации сервера при отправке push-уведомлений. Для генерации ключей можно использовать библиотеку `web-push`:

```bash
npx web-push generate-vapid-keys
```

Это создаст пару ключей:

- Публичный ключ - используется на клиенте
- Приватный ключ - используется на сервере

### Хранение подписок

Сервер должен хранить подписки пользователей в базе данных. Каждая подписка связывается с конкретным пользователем:

```typescript
interface PushSubscriptionRecord {
  userId: string;
  subscription: PushSubscription;
  createdAt: Date;
  updatedAt: Date;
}
```

### Отправка уведомлений

Для отправки push-уведомлений на сервере можно использовать библиотеку `web-push`:

```javascript
const webpush = require('web-push');

// Настройка VAPID ключей
webpush.setVapidDetails(
  'mailto:example@couply.com',
  process.env.PUBLIC_VAPID_KEY,
  process.env.PRIVATE_VAPID_KEY,
);

// Отправка уведомления
async function sendPushNotification(subscription, payload) {
  try {
    await webpush.sendNotification(subscription, JSON.stringify(payload));
    return true;
  } catch (error) {
    console.error('Error sending push notification:', error);
    return false;
  }
}
```

Пример payload для уведомления о совпадении:

```javascript
const payload = {
  title: `Новое совпадение с ${matchName}!`,
  body: 'У вас взаимная симпатия!',
  icon: matchImage,
  url: '/likes',
};
```

## Интеграция в приложение

### Инициализация при запуске

В Couply инициализация push-уведомлений происходит при запуске приложения в компоненте `App`:

```tsx
function App() {
  const [userId] = useState('user123');
  const { isSupported, permission, initialize, isInitializing } = usePushNotifications();
  const [showPermissionPrompt, setShowPermissionPrompt] = useState(false);

  // Инициализация push-уведомлений при загрузке приложения
  useEffect(() => {
    if (isSupported && permission === 'granted' && !isInitializing) {
      initialize(userId).then(success => {
        console.log('Push notifications initialized:', success);
      });
    } else if (isSupported && permission !== 'granted' && !isInitializing) {
      // Показываем запрос на разрешение уведомлений через некоторое время после загрузки
      const timer = setTimeout(() => {
        setShowPermissionPrompt(true);
      }, 5000);

      return () => clearTimeout(timer);
    }
  }, [isSupported, permission, isInitializing, initialize, userId]);

  // ...
}
```

### Отправка уведомлений о совпадениях

Отправка уведомлений о совпадениях происходит в компоненте `LikesPage` при возникновении взаимного лайка:

```tsx
// Check if this profile has already liked the user
if (likedProfile.hasLikedYou) {
  // It's a match! Show the match modal
  setMatchedProfile(likedProfile);
  setShowMatchModal(true);

  // Show in-app notification
  setNotificationProfile(likedProfile);
  setShowNotification(true);

  // Send push notification
  sendMatchNotification({
    userId: 'user123',
    matchId: likedProfile.id,
    matchName: likedProfile.name,
    matchImage: likedProfile.imageUrl,
  });

  // Don't open the profile view when it's a match
  setSelectedProfile(null);
}
```

## Тестирование

Для тестирования push-уведомлений можно использовать следующие методы:

1. **Локальное тестирование**:

   - Запустите приложение с помощью `npm run dev`
   - Разрешите отправку уведомлений
   - Создайте взаимный лайк
   - Проверьте, что уведомление отображается

2. **Тестирование с помощью DevTools**:

   - Откройте DevTools в Chrome
   - Перейдите на вкладку Application > Service Workers
   - Нажмите на кнопку "Push" для отправки тестового push-события

3. **Тестирование с помощью curl**:
   - Получите подписку пользователя
   - Используйте curl для отправки push-уведомления:
   ```bash
   curl -X POST -H "Content-Type: application/json" -H "Authorization: vapid t=<token>" -d '{"subscription": <subscription>, "payload": <payload>}' https://web-push-codelab.glitch.me/api/send-push-msg
   ```

## Устранение неполадок

### Уведомления не отображаются

1. Проверьте, что пользователь разрешил отправку уведомлений:

```javascript
console.log(Notification.permission);
```

2. Проверьте, что Service Worker зарегистрирован:

```javascript
navigator.serviceWorker.getRegistrations().then(registrations => {
  console.log(registrations);
});
```

3. Проверьте, что подписка создана:

```javascript
navigator.serviceWorker.ready.then(registration => {
  registration.pushManager.getSubscription().then(subscription => {
    console.log(subscription);
  });
});
```

### Ошибки при создании подписки

1. Проверьте, что VAPID ключи корректны
2. Проверьте, что сайт использует HTTPS (кроме localhost)
3. Проверьте, что manifest.json правильно настроен

## Ограничения и совместимость

### Поддержка браузерами

Push API поддерживается следующими браузерами:

- Chrome (Desktop и Android)
- Firefox (Desktop и Android)
- Edge
- Opera
- Samsung Internet

Не поддерживается:

- Safari (iOS и macOS) - использует Apple Push Notification Service (APNs)
- Internet Explorer

### Ограничения

1. **Частота отправки**: браузеры могут ограничивать количество уведомлений, которые можно отправить за определенный период времени
2. **Размер payload**: максимальный размер payload обычно ограничен 4KB
3. **Требование HTTPS**: для работы Push API требуется HTTPS (кроме localhost)
4. **Взаимодействие с пользователем**: уведомления должны быть видимыми для пользователя (userVisibleOnly: true)

### Альтернативы для iOS

Для iOS можно использовать следующие альтернативы:

1. **Web Push для Safari на macOS**: начиная с Safari 16.4 на macOS поддерживается Web Push API
2. **PWA на домашнем экране**: если пользователь добавил приложение на домашний экран, можно использовать локальные уведомления
3. **Нативное приложение**: создать нативное приложение с использованием React Native или другой технологии
