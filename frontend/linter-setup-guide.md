# Настройка линтера и прекоммитных хуков в React TypeScript проекте

## Введение

Линтинг кода и автоматическое форматирование - важные инструменты для поддержания качества кода в проекте. Они помогают:

- Обнаруживать потенциальные ошибки и проблемы
- Поддерживать единый стиль кода
- Следовать лучшим практикам
- Улучшать доступность и производительность

В этом руководстве описан процесс настройки ESLint, Prettier и прекоммитных хуков с использованием Husky и lint-staged.

## Шаг 1: Установка необходимых пакетов

Сначала необходимо установить все требуемые пакеты:

```bash
# Установка ESLint и базовых плагинов
npm install --save-dev eslint @eslint/js typescript-eslint

# Установка плагинов для React
npm install --save-dev eslint-plugin-react eslint-plugin-react-hooks eslint-plugin-react-refresh

# Установка плагинов для импортов и доступности
npm install --save-dev eslint-plugin-import eslint-plugin-jsx-a11y

# Установка Prettier и интеграции с ESLint
npm install --save-dev prettier eslint-plugin-prettier

# Установка Husky и lint-staged для прекоммитных хуков
npm install --save-dev husky lint-staged
```

## Шаг 2: Настройка ESLint

ESLint настраивается через файл `eslint.config.js` в корне проекта. Современный ESLint использует "flat config" формат:

```javascript
import js from '@eslint/js';
import importPlugin from 'eslint-plugin-import';
import jsxA11y from 'eslint-plugin-jsx-a11y';
import prettierPlugin from 'eslint-plugin-prettier';
import react from 'eslint-plugin-react';
import reactHooks from 'eslint-plugin-react-hooks';
import reactRefresh from 'eslint-plugin-react-refresh';
import globals from 'globals';
import tseslint from 'typescript-eslint';

export default tseslint.config(
  // Игнорируемые директории
  { ignores: ['dist', 'node_modules', '.git', 'build'] },

  // Основная конфигурация
  {
    extends: [js.configs.recommended, ...tseslint.configs.recommended],
    files: ['**/*.{ts,tsx,js,jsx}'],

    // Настройки языка
    languageOptions: {
      ecmaVersion: 2020,
      globals: {
        ...globals.browser,
        ...globals.node,
        ...globals.es2020,
      },
      parserOptions: {
        ecmaFeatures: {
          jsx: true,
        },
      },
    },

    // Настройки React
    settings: {
      react: {
        version: 'detect', // Автоматически определять версию React
      },
    },

    // Подключаемые плагины
    plugins: {
      'react-hooks': reactHooks,
      'react-refresh': reactRefresh,
      react: react,
      import: importPlugin,
      'jsx-a11y': jsxA11y,
      prettier: prettierPlugin,
      '@typescript-eslint': tseslint.plugin,
    },

    // Правила линтинга
    rules: {
      // Базовые правила JavaScript
      'no-console': ['warn', { allow: ['warn', 'error', 'info'] }],
      'no-unused-vars': 'off', // Отключено в пользу версии TypeScript
      'no-duplicate-imports': 'error',
      'no-var': 'error',
      'prefer-const': 'error',

      // Правила React
      'react/jsx-uses-react': 'error',
      'react/jsx-uses-vars': 'error',
      'react/jsx-key': 'error',
      'react/jsx-no-duplicate-props': 'error',

      // Правила React Hooks
      ...reactHooks.configs.recommended.rules,

      // Правила TypeScript
      '@typescript-eslint/no-unused-vars': [
        'error',
        { argsIgnorePattern: '^_', varsIgnorePattern: '^_' },
      ],
      '@typescript-eslint/no-explicit-any': 'warn',

      // Правила импортов
      'import/order': [
        'warn',
        {
          groups: ['builtin', 'external', 'internal', 'parent', 'sibling', 'index'],
          'newlines-between': 'always',
          alphabetize: { order: 'asc', caseInsensitive: true },
        },
      ],

      // Правила доступности
      'jsx-a11y/alt-text': 'warn',
      'jsx-a11y/click-events-have-key-events': 'warn',
    },
  },

  // Специальные настройки для тестов
  {
    files: ['**/*.test.{ts,tsx,js,jsx}', '**/*.spec.{ts,tsx,js,jsx}'],
    rules: {
      'max-len': 'off',
      '@typescript-eslint/no-explicit-any': 'off',
    },
  },

  // Специальные настройки для Storybook
  {
    files: ['**/*.stories.{ts,tsx,js,jsx}'],
    rules: {
      'no-console': 'off', // Разрешаем console.log в файлах историй
    },
  },
);
```

### Объяснение ключевых частей конфигурации ESLint:

1. **Импорты плагинов**: Импортируем все необходимые плагины для ESLint.
2. **Игнорируемые директории**: Указываем, какие директории не нужно проверять.
3. **Extends**: Расширяем базовые конфигурации рекомендуемыми правилами.
4. **Files**: Указываем, какие файлы нужно проверять.
5. **LanguageOptions**: Настраиваем параметры языка, включая глобальные переменные.
6. **Settings**: Дополнительные настройки для плагинов, например версия React.
7. **Plugins**: Подключаем все необходимые плагины.
8. **Rules**: Настраиваем правила линтинга.
9. **Специальные настройки**: Создаем отдельные конфигурации для тестов и Storybook.

## Шаг 3: Настройка Prettier

Prettier настраивается через файл `.prettierrc` в корне проекта:

```json
{
  "semi": true,
  "singleQuote": true,
  "tabWidth": 2,
  "trailingComma": "all",
  "printWidth": 100,
  "bracketSpacing": true,
  "arrowParens": "avoid",
  "endOfLine": "lf",
  "jsxSingleQuote": false,
  "bracketSameLine": false,
  "quoteProps": "as-needed"
}
```

Также создаем файл `.prettierignore` для указания файлов, которые не нужно форматировать:

```
# Игнорировать сборки и зависимости
dist
node_modules
build
coverage

# Игнорировать файлы конфигурации
package-lock.json
yarn.lock

# Игнорировать статические файлы
*.svg
*.png
*.jpg
```

## Шаг 4: Настройка скриптов в package.json

Добавляем скрипты для запуска линтера и форматирования в `package.json`:

```json
{
  "scripts": {
    "lint": "eslint .",
    "lint:fix": "eslint . --fix",
    "format": "prettier --write \"src/**/*.{ts,tsx,js,jsx,json,css,scss,md}\"",
    "format:check": "prettier --check \"src/**/*.{ts,tsx,js,jsx,json,css,scss,md}\"",
    "check": "npm run lint && npm run format:check && tsc --noEmit",
    "fix": "npm run lint:fix && npm run format",
    "prepare": "husky"
  }
}
```

## Шаг 5: Настройка lint-staged

Добавляем конфигурацию lint-staged в `package.json`:

```json
{
  "lint-staged": {
    "*.{js,jsx,ts,tsx}": ["eslint --fix", "prettier --write"],
    "*.{json,css,scss,md}": ["prettier --write"]
  }
}
```

Эта конфигурация указывает, какие команды нужно выполнять для каких файлов при коммите.

## Шаг 6: Настройка Husky

1. Инициализируем Husky:

```bash
npx husky init
```

Если проект не является Git-репозиторием, создаем директорию `.husky` вручную:

```bash
mkdir -p .husky
```

2. Создаем файл `.husky/pre-commit` для запуска lint-staged перед коммитом:

```bash
#!/usr/bin/env sh
. "$(dirname -- "$0")/_/husky.sh"

npx lint-staged
```

3. Делаем файл исполняемым:

```bash
chmod +x .husky/pre-commit
```

4. Создаем вспомогательный файл `.husky/_/husky.sh`:

```bash
#!/usr/bin/env sh
if [ -z "$husky_skip_init" ]; then
  debug () {
    if [ "$HUSKY_DEBUG" = "1" ]; then
      echo "husky (debug) - $1"
    fi
  }

  readonly hook_name="$(basename -- "$0")"
  debug "starting $hook_name..."

  if [ "$HUSKY" = "0" ]; then
    debug "HUSKY env variable is set to 0, skipping hook"
    exit 0
  fi

  if [ -f ~/.huskyrc ]; then
    debug "sourcing ~/.huskyrc"
    . ~/.huskyrc
  fi

  readonly husky_skip_init=1
  export husky_skip_init
  sh -e "$0" "$@"
  exitCode="$?"

  if [ $exitCode != 0 ]; then
    echo "husky - $hook_name hook exited with code $exitCode (error)"
  fi

  exit $exitCode
fi
```

5. Делаем этот файл исполняемым:

```bash
chmod +x .husky/_/husky.sh
```

## Шаг 7: Настройка VS Code

Создаем файл `.vscode/settings.json` для автоматического форматирования при сохранении:

```json
{
  "editor.formatOnSave": true,
  "editor.defaultFormatter": "esbenp.prettier-vscode",
  "editor.codeActionsOnSave": {
    "source.fixAll.eslint": true
  },
  "eslint.validate": ["javascript", "javascriptreact", "typescript", "typescriptreact"],
  "typescript.tsdk": "node_modules/typescript/lib",
  "prettier.requireConfig": true,
  "files.eol": "\n"
}
```

## Как это работает

1. **ESLint** анализирует код на соответствие правилам и находит потенциальные проблемы.
2. **Prettier** форматирует код в соответствии с заданными правилами стиля.
3. **lint-staged** запускает линтер и форматирование только для файлов, которые были изменены и добавлены в коммит.
4. **Husky** запускает lint-staged перед каждым коммитом через pre-commit хук.

## Преимущества такой настройки

1. **Автоматизация**: Проверка и форматирование кода происходят автоматически.
2. **Единый стиль**: Весь код в проекте следует единому стилю.
3. **Раннее обнаружение проблем**: Потенциальные ошибки обнаруживаются до коммита.
4. **Экономия времени**: Не нужно тратить время на ручное форматирование и проверку кода.
5. **Улучшение качества кода**: Следование лучшим практикам повышает качество кода.

## Использование

- Для проверки кода: `npm run lint`
- Для автоматического исправления проблем: `npm run lint:fix`
- Для форматирования кода: `npm run format`
- Для проверки форматирования: `npm run format:check`
- Для запуска всех проверок: `npm run check`
- Для запуска всех исправлений: `npm run fix`

При коммите в Git код будет автоматически проверяться и форматироваться благодаря настройке Husky и lint-staged.
