import { baseApi } from '../../../shared/api/baseApi';

export const telegramApi = baseApi.injectEndpoints({
  endpoints: builder => ({
    setTelegram: builder.mutation<object, object>({
      query: user => ({
        url: '/v1/telegram/data/set',
        method: 'POST',
        body: JSON.stringify(user),
      }),
      // invalidatesTags: [{ type: 'User', id: 'LIST' }],
    }),

    getTelegram: builder.mutation<object, object>({
      query: () => ({
        url: '/v1/telegram/data/get',
        method: 'POST',
        body: {},
      }),
      // invalidatesTags: [{ type: 'User', id: 'LIST' }],
    }),
  }),
});

export const { useSetTelegramMutation, useGetTelegramMutation } = telegramApi;
