# Архитектура Redux Toolkit и RTK Query

В этом документе описана архитектура Redux Toolkit и RTK Query в проекте, включая организацию файлов, основные концепции и примеры использования.

## Структура файлов

Архитектура организована в соответствии с принципами Feature-Sliced Design (FSD):

```
src/
├── app/
│   └── store/
│       └── index.ts         # Корневой store Redux
├── entities/
│   ├── profile/
│   │   └── api/
│   │       └── profileApi.ts # API для работы с профилями
│   └── likes/
│       └── api/
│           └── likesApi.ts   # API для работы с лайками и матчами
├── features/
│   └── filters/
│       └── model/
│           └── filtersSlice.ts # Слайс для фильтров
└── shared/
    ├── api/
    │   └── baseApi.ts        # Базовый API для RTK Query
    └── lib/
        └── hooks/
            └── redux.ts      # Типизированные хуки для Redux
```

## Основные компоненты

### 1. Корневой Store (src/app/store/index.ts)

Корневой store объединяет все reducers и middleware:

```typescript
import { configureStore } from '@reduxjs/toolkit';
import { setupListeners } from '@reduxjs/toolkit/query';
import { baseApi } from '../../shared/api/baseApi';
import filtersReducer from '../../features/filters/model/filtersSlice';

export const store = configureStore({
  reducer: {
    // API reducers
    [baseApi.reducerPath]: baseApi.reducer,
    // Feature reducers
    filters: filtersReducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(baseApi.middleware),
});

// Необходимо для refetchOnFocus/refetchOnReconnect
setupListeners(store.dispatch);

// Типы для использования в приложении
export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
```

### 2. Базовый API (src/shared/api/baseApi.ts)

Базовый API для RTK Query, который будет расширяться в других файлах:

```typescript
import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';

export const baseApi = createApi({
  reducerPath: 'api',
  baseQuery: fetchBaseQuery({ 
    baseUrl: API_BASE_URL,
    prepareHeaders: (headers) => {
      // Добавление авторизационных заголовков
      const token = localStorage.getItem('token');
      if (token) {
        headers.set('Authorization', `Bearer ${token}`);
      }
      return headers;
    },
  }),
  endpoints: () => ({}), // Пустые endpoints, будут расширяться в других файлах
  tagTypes: ['Profile', 'Likes', 'Matches'], // Типы тегов для кэширования
});
```

### 3. Типизированные хуки (src/shared/lib/hooks/redux.ts)

Типизированные хуки для использования Redux в компонентах:

```typescript
import { TypedUseSelectorHook, useDispatch, useSelector } from 'react-redux';
import type { RootState, AppDispatch } from '../../../app/store';

// Типизированные хуки для использования в компонентах
export const useAppDispatch = () => useDispatch<AppDispatch>();
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector;
```

### 4. API сервисы (src/entities/*/api/*.ts)

API сервисы расширяют базовый API и предоставляют endpoints для работы с конкретными сущностями:

```typescript
// Пример API для профилей
export const profileApi = baseApi.injectEndpoints({
  endpoints: (builder) => ({
    getProfile: builder.query<Profile, string>({
      query: (id) => `/profiles/${id}`,
      providesTags: (result, error, id) => [{ type: 'Profile', id }],
    }),
    // Другие endpoints...
  }),
});

// Экспорт хуков для использования в компонентах
export const {
  useGetProfileQuery,
  useGetProfilesQuery,
  useUpdateProfileMutation,
} = profileApi;
```

### 5. Слайсы (src/features/*/model/*.ts)

Слайсы для управления состоянием:

```typescript
// Пример слайса для фильтров
export const filtersSlice = createSlice({
  name: 'filters',
  initialState,
  reducers: {
    setAgeRange: (state, action: PayloadAction<[number, number]>) => {
      state.ageRange = action.payload;
    },
    // Другие reducers...
  },
});

// Экспорт actions
export const { setAgeRange, setDistance, /* ... */ } = filtersSlice.actions;

// Экспорт reducer
export default filtersSlice.reducer;
```

## Использование в компонентах

### Использование RTK Query

```tsx
// Пример использования RTK Query в компоненте
const { data: profiles, isLoading, error } = useGetProfilesQuery();

// Пример использования мутации
const [createLike, { isLoading: isLikeLoading }] = useCreateLikeMutation();

// Вызов мутации
const handleLike = async (profileId: string) => {
  try {
    await createLike({ targetUserId: profileId }).unwrap();
    // Обработка успешного лайка
  } catch (error) {
    // Обработка ошибки
    console.error('Failed to like profile:', error);
  }
};
```

### Использование Redux State

```tsx
// Пример использования Redux State в компоненте
const dispatch = useAppDispatch();
const filters = useAppSelector(selectFilters);

// Пример диспатча действия
const handleResetFilters = () => {
  dispatch(resetFilters());
};
```

## Добавление новых API сервисов

Для добавления нового API сервиса:

1. Создайте новый файл в соответствующей папке entities или features
2. Расширьте baseApi с помощью injectEndpoints
3. Определите endpoints с помощью builder.query и builder.mutation
4. Экспортируйте сгенерированные хуки

```typescript
// Пример нового API сервиса
export const newEntityApi = baseApi.injectEndpoints({
  endpoints: (builder) => ({
    // Определите endpoints...
  }),
});

// Экспортируйте хуки
export const { useNewQueryHook, useNewMutationHook } = newEntityApi;
```

## Добавление новых слайсов

Для добавления нового слайса:

1. Создайте новый файл в соответствующей папке features
2. Определите интерфейс состояния и начальное состояние
3. Создайте слайс с помощью createSlice
4. Экспортируйте actions и reducer
5. Добавьте reducer в корневой store

```typescript
// Пример нового слайса
export const newSlice = createSlice({
  name: 'newFeature',
  initialState,
  reducers: {
    // Определите reducers...
  },
});

// Экспортируйте actions и reducer
export const { action1, action2 } = newSlice.actions;
export default newSlice.reducer;
```

Затем добавьте reducer в корневой store:

```typescript
// В src/app/store/index.ts
import newReducer from '../../features/newFeature/model/newSlice';

export const store = configureStore({
  reducer: {
    // Существующие reducers...
    newFeature: newReducer,
  },
  // ...
});
```

## Лучшие практики

1. **Организация по фичам**: Размещайте API и слайсы в соответствующих папках entities и features
2. **Типизация**: Всегда определяйте типы для данных API и состояния
3. **Кэширование**: Используйте теги (tagTypes) для правильного кэширования и инвалидации данных
4. **Селекторы**: Создавайте селекторы для доступа к состоянию
5. **Типизированные хуки**: Используйте типизированные хуки useAppDispatch и useAppSelector вместо обычных useDispatch и useSelector