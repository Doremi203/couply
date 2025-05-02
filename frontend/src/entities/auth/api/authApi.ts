import { v4 as uuidv4 } from 'uuid';

import { baseApi } from '../../../shared/api/baseApi';
import { AuthResponse, LoginParams, RegisterParams } from '../types';

export const authApi = baseApi.injectEndpoints({
  endpoints: builder => ({
    register: builder.mutation<AuthResponse, RegisterParams>({
      query: credentials => ({
        url: '/v1/register/basic',
        method: 'POST',
        body: credentials,
        headers: {
          'Idempotency-Key': uuidv4(),
        },
      }),
    }),

    login: builder.mutation<AuthResponse, LoginParams>({
      query: credentials => ({
        url: '/auth/login',
        method: 'POST',
        body: credentials,
      }),
    }),

    // logout: builder.mutation<void, void>({
    //   query: () => ({
    //     url: '/auth/logout',
    //     method: 'POST',
    //   }),
    // }),

    // // Получение текущего пользователя
    // getCurrentUser: builder.query<User, void>({
    //   query: () => '/auth/me',
    //   providesTags: ['User'],
    // }),

    // // Обновление профиля пользователя после регистрации
    // updateUserProfile: builder.mutation<User, Partial<User>>({
    //   query: userData => ({
    //     url: '/auth/profile',
    //     method: 'PATCH',
    //     body: userData,
    //   }),
    //   invalidatesTags: ['User'],
    // }),

    // // Запрос на сброс пароля
    // requestPasswordReset: builder.mutation<void, { email: string } | { phone: string }>({
    //   query: data => ({
    //     url: '/auth/reset-password',
    //     method: 'POST',
    //     body: data,
    //   }),
    // }),

    // // Подтверждение сброса пароля
    // confirmPasswordReset: builder.mutation<void, { token: string; password: string }>({
    //   query: data => ({
    //     url: '/auth/reset-password/confirm',
    //     method: 'POST',
    //     body: data,
    //   }),
    // }),
  }),
});

// Экспорт хуков для использования в компонентах
export const {
  useRegisterMutation,
  useLoginMutation,
  // useLogoutMutation,
  // useGetCurrentUserQuery,
  // useUpdateUserProfileMutation,
  // useRequestPasswordResetMutation,
  // useConfirmPasswordResetMutation,
} = authApi;
