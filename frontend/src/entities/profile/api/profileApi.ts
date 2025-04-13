import { baseApi } from '../../../shared/api/baseApi';

// Типы данных для профиля
export interface Profile {
  id: string;
  name: string;
  age: number;
  bio: string;
  photos: string[];
  interests: string[];
  // Добавьте другие поля профиля по необходимости
}

// Параметры для обновления профиля
export interface UpdateProfileParams {
  id: string;
  data: Partial<Profile>;
}

// Расширение базового API для работы с профилями
export const profileApi = baseApi.injectEndpoints({
  endpoints: builder => ({
    // Получение профиля пользователя
    getProfile: builder.query<Profile, string>({
      query: id => `/profiles/${id}`,
      providesTags: (_result, _error, id) => [{ type: 'Profile', id }],
    }),

    // Получение списка профилей
    getProfiles: builder.query<Profile[], void>({
      query: () => '/profiles',
      providesTags: result =>
        result
          ? [
              ...result.map(({ id }) => ({ type: 'Profile' as const, id })),
              { type: 'Profile', id: 'LIST' },
            ]
          : [{ type: 'Profile', id: 'LIST' }],
    }),

    // Обновление профиля
    updateProfile: builder.mutation<Profile, UpdateProfileParams>({
      query: ({ id, data }) => ({
        url: `/profiles/${id}`,
        method: 'PATCH',
        body: data,
      }),
      invalidatesTags: (_result, _error, { id }) => [{ type: 'Profile', id }],
    }),
  }),
});

// Экспорт хуков для использования в компонентах
export const { useGetProfileQuery, useGetProfilesQuery, useUpdateProfileMutation } = profileApi;
