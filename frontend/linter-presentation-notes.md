# Шпаргалка по настройке линтера и прекоммитных хуков

## Что такое линтер и зачем он нужен?

**Линтер** - это инструмент, который анализирует код на соответствие определенным правилам и стандартам.

**Преимущества использования линтера:**

- Обнаружение ошибок до запуска кода
- Поддержание единого стиля кода в команде
- Следование лучшим практикам
- Улучшение качества и поддерживаемости кода

## Основные инструменты

1. **ESLint** - анализатор кода для JavaScript/TypeScript
2. **Prettier** - форматирование кода
3. **Husky** - управление Git-хуками
4. **lint-staged** - запуск линтера только для измененных файлов

## Процесс настройки (пошагово)

### 1. Установка пакетов

```bash
# Основные пакеты
npm install --save-dev eslint prettier

# Плагины для React и TypeScript
npm install --save-dev eslint-plugin-react typescript-eslint

# Интеграция ESLint и Prettier
npm install --save-dev eslint-plugin-prettier

# Прекоммитные хуки
npm install --save-dev husky lint-staged
```

### 2. Настройка ESLint (eslint.config.js)

```javascript
// Импорт плагинов
import js from '@eslint/js';
import react from 'eslint-plugin-react';
import tseslint from 'typescript-eslint';

// Конфигурация
export default tseslint.config(
  { ignores: ['dist', 'node_modules'] },
  {
    extends: [js.configs.recommended],
    files: ['**/*.{ts,tsx,js,jsx}'],
    plugins: {
      react: react,
      '@typescript-eslint': tseslint.plugin,
    },
    rules: {
      // Правила для JavaScript
      'no-console': 'warn',
      'no-unused-vars': 'off',

      // Правила для React
      'react/jsx-key': 'error',

      // Правила для TypeScript
      '@typescript-eslint/no-unused-vars': 'error',
    },
  },
);
```

### 3. Настройка Prettier (.prettierrc)

```json
{
  "semi": true,
  "singleQuote": true,
  "tabWidth": 2,
  "trailingComma": "all",
  "printWidth": 100
}
```

### 4. Настройка скриптов в package.json

```json
{
  "scripts": {
    "lint": "eslint .",
    "lint:fix": "eslint . --fix",
    "format": "prettier --write \"src/**/*.{ts,tsx,js,jsx}\"",
    "prepare": "husky"
  },
  "lint-staged": {
    "*.{js,jsx,ts,tsx}": ["eslint --fix", "prettier --write"]
  }
}
```

### 5. Настройка Husky

```bash
# Создание директории
mkdir -p .husky

# Создание pre-commit хука
echo '#!/usr/bin/env sh
. "$(dirname -- "$0")/_/husky.sh"

npx lint-staged
' > .husky/pre-commit

# Установка прав на выполнение
chmod +x .husky/pre-commit
```

## Как это работает на практике

2. При сохранении файла VS Code автоматически форматирует код (если настроен)
3. Разработчик делает `git add` измененных файлов
4. При выполнении `git commit`:
   - Husky перехватывает команду
   - Запускает lint-staged
   - lint-staged запускает ESLint и Prettier для измененных файлов
   - Если есть ошибки, коммит отменяется
   - Если ошибок нет, коммит выполняется

## Демонстрация

1. **Показать конфигурацию ESLint**

   - Объяснить основные правила
   - Показать, как настраиваются правила для разных типов файлов

2. **Показать работу линтера**

   - Запустить `npm run lint`
   - Показать найденные предупреждения/ошибки

3. **Показать автоматическое исправление**

   - Запустить `npm run lint:fix`
   - Показать, как изменился код

4. **Показать работу прекоммитных хуков**
   - Внести изменение с ошибкой
   - Попытаться сделать коммит
   - Показать, как хук блокирует коммит с ошибкой

## Ключевые выводы

1. Линтеры и форматирование кода - важная часть процесса разработки
2. Автоматизация проверки кода экономит время и повышает качество
3. Прекоммитные хуки гарантируют, что в репозиторий попадает только качественный код
4. Единый стиль кода улучшает читаемость и поддерживаемость проекта
