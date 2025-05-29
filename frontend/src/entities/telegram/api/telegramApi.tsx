import { baseApi } from '../../../shared/api/baseApi';

export const telegramApi = baseApi.injectEndpoints({
  endpoints: builder => ({
    setTelegram: builder.mutation<object, object>({
      query: user => ({
        url: '/v1/telegram/data/set',
        method: 'POST',
        body: user,
      }),
      // invalidatesTags: [{ type: 'User', id: 'LIST' }],
    }),

    getTelegram: builder.mutation<object, object>({
      query: userId => ({
        url: '/v1/telegram/data/get',
        method: 'POST',
        body: { user_id: userId },
      }),
      // invalidatesTags: [{ type: 'User', id: 'LIST' }],
    }),
  }),
});

export const { useSetTelegramMutation, useGetTelegramMutation } = telegramApi;
