# Линтинг и форматирование кода

В этом проекте настроены инструменты для обеспечения качества кода и единого стиля форматирования.

## Установленные инструменты

- **ESLint** - анализатор кода для выявления проблемных паттернов
- **Prettier** - форматирование кода для поддержания единого стиля
- **TypeScript** - статическая типизация для JavaScript

## Доступные команды

```bash
# Проверка кода на ошибки и стиль
npm run lint         # Запуск ESLint для проверки кода
npm run format:check # Проверка форматирования кода с помощью Prettier
npm run check        # Запуск всех проверок (lint + format:check + tsc)

# Автоматическое исправление проблем
npm run lint:fix     # Автоматическое исправление проблем ESLint
npm run format       # Форматирование кода с помощью Prettier
npm run fix          # Запуск всех исправлений (lint:fix + format)
```

## Правила линтинга

В проекте настроены следующие группы правил:

1. **Базовые правила JavaScript/TypeScript**
   - Запрет использования `console.log` (разрешены `warn`, `error`, `info`)
   - Обязательное использование `const` и `let` вместо `var`
   - Единый стиль кавычек, точек с запятой и т.д.

2. **Правила React**
   - Проверка корректного использования JSX
   - Проверка ключей в списках
   - Запрет дублирования props
   - Правила для хуков React

3. **Правила TypeScript**
   - Проверка неиспользуемых переменных
   - Соглашения по именованию (интерфейсы с префиксом 'I', PascalCase для типов)
   - Предупреждения о использовании `any`

4. **Правила импортов**
   - Сортировка импортов по группам
   - Запрет циклических зависимостей
   - Запрет дублирования импортов

5. **Правила доступности (a11y)**
   - Проверка атрибутов alt для изображений
   - Проверка корректного использования ARIA-атрибутов
   - Проверка доступности интерактивных элементов

## Настройка IDE

### VS Code

1. Установите расширения:
   - ESLint
   - Prettier

2. Настройте автоматическое форматирование при сохранении, добавив в settings.json:

```json
{
  "editor.formatOnSave": true,
  "editor.defaultFormatter": "esbenp.prettier-vscode",
  "editor.codeActionsOnSave": {
    "source.fixAll.eslint": true
  }
}
```

### WebStorm/IntelliJ IDEA

1. Включите ESLint:
   - Preferences → Languages & Frameworks → JavaScript → Code Quality Tools → ESLint
   - Выберите "Automatic ESLint configuration"

2. Включите Prettier:
   - Preferences → Languages & Frameworks → JavaScript → Prettier
   - Выберите "Automatic Prettier configuration"
   - Включите "Run on save"

## Игнорирование правил

### Временное отключение правил ESLint

```javascript
// eslint-disable-next-line no-console
console.log('Debugging information');

/* eslint-disable */
// Код, для которого отключены все правила
/* eslint-enable */
```

### Игнорирование файлов

- Для ESLint: добавьте пути в `.eslintignore` или в секцию `ignores` в `eslint.config.js`
- Для Prettier: добавьте пути в `.prettierignore`