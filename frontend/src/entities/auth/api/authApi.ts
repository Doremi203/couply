import { v4 as uuidv4 } from 'uuid';

import { baseApi } from '../../../shared/api/baseApi';

import {
  LoginParams,
  LoginResponse,
  RefreshRequest,
  RefreshResponse,
  RegisterParams,
} from './types';

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

    refreshToken: builder.mutation<RefreshResponse, RefreshRequest>({
      query: credentials => ({
        url: '/v1/token/refresh',
        method: 'POST',
        body: credentials,
      }),
    }),
  }),
});

export const { useRegisterMutation, useLoginMutation, useRefreshTokenMutation } = authApi;
