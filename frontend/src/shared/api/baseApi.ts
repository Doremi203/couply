import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';

// Базовый URL API
const API_BASE_URL = 'https://auth.testing.couply.ru'; // Замените на ваш реальный API URL
const MATCHER_API_URL = 'https://matcher.testing.couply.ru';

// Базовый API для RTK Query
export const baseApi = createApi({
  reducerPath: 'api',
  baseQuery: fetchBaseQuery({
    baseUrl: API_BASE_URL,
    prepareHeaders: headers => {
      // Здесь можно добавить авторизационные заголовки
      const token = localStorage.getItem('token');
      if (token) {
        headers.set('Authorization', `Bearer ${token}`);
      }
      return headers;
    },
  }),
  endpoints: () => ({}), // Пустые endpoints, будут расширяться в других файлах
  tagTypes: ['Profile', 'Matches', 'User'], // Добавлен тег 'User' для аутентификации
});

// Базовый API для RTK Query
export const matcherApi = createApi({
  reducerPath: 'matherApi',
  baseQuery: fetchBaseQuery({
    baseUrl: MATCHER_API_URL,
    prepareHeaders: headers => {
      // Здесь можно добавить авторизационные заголовки
      const token = localStorage.getItem('token');
      if (token) {
        headers.set('Authorization', `Bearer ${token}`);
      }
      return headers;
    },
  }),
  endpoints: () => ({}), // Пустые endpoints, будут расширяться в других файлах
  tagTypes: ['Profile', 'Matches', 'User'], // Добавлен тег 'User' для аутентификации
});
