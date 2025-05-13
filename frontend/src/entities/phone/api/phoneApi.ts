import { baseApi } from '../../../shared/api/baseApi';

import { CodeResponse, ConfirmParams, PhoneParams } from './types';

export const phoneApi = baseApi.injectEndpoints({
  endpoints: builder => ({
    confirmPhone: builder.mutation<object, ConfirmParams>({
      query: credentials => ({
        url: '/v1/phone/confirm',
        method: 'POST',
        body: credentials,
      }),
    }),

    sendCode: builder.mutation<CodeResponse, PhoneParams>({
      query: credentials => ({
        url: '/v1/phone/send-code',
        method: 'POST',
        body: credentials,
      }),
    }),
  }),
});

export const { useConfirmPhoneMutation, useSendCodeMutation } = phoneApi;
