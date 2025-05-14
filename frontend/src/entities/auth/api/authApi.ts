import { v4 as uuidv4 } from 'uuid';

import { baseApi } from '../../../shared/api/baseApi';

import { LoginParams, LoginResponse, RegisterParams } from './types';

export const authApi = baseApi.injectEndpoints({
  endpoints: builder => ({
    register: builder.mutation<object, RegisterParams>({
      query: credentials => ({
        url: '/v1/register/basic',
        method: 'POST',
        body: credentials,
        headers: {
          'Idempotency-Key': uuidv4(),
        },
      }),
    }),

    login: builder.mutation<LoginResponse, LoginParams>({
      query: credentials => ({
        url: '/v1/login/basic',
        method: 'POST',
        body: credentials,
      }),
    }),

    // refreshToken: builder.mutation<LoginResponse, void>({
    //   query: () => ({
    //     url: '/v1/refresh',
    //     method: 'POST',
    //     headers: {
    //       Authorization: `Bearer ${getToken()}`,
    //     },
    //   }),
    // }),
  }),
});

export const {
  useRegisterMutation,
  useLoginMutation,
  // useRefreshTokenMutation,
} = authApi;
