import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';

// Базовый URL API
const API_BASE_URL = 'https://api.example.com'; // Замените на ваш реальный API URL

// Базовый API для RTK Query
export const baseApi = createApi({
  reducerPath: 'api',
  baseQuery: fetchBaseQuery({ 
    baseUrl: API_BASE_URL,
    prepareHeaders: (headers) => {
      // Здесь можно добавить авторизационные заголовки
      const token = localStorage.getItem('token');
      if (token) {
        headers.set('Authorization', `Bearer ${token}`);
      }
      return headers;
    },
  }),
  endpoints: () => ({}), // Пустые endpoints, будут расширяться в других файлах
  tagTypes: ['Profile', 'Likes', 'Matches'], // Добавьте здесь все типы тегов для кэширования
});