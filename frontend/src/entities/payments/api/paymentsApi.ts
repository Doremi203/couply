import { paymentsApi } from '../../../shared/api/baseApi';
import { blockerApiExtended } from '../../blocker';

import {
  CreatePaymentResponse,
  CreatePaymentRequest,
  GetPaymentResponse,
  GetPaymentRequest,
} from './types';

export const paymentsApiExtended = paymentsApi.injectEndpoints({
  endpoints: builder => ({
    createPayment: builder.mutation<CreatePaymentResponse, CreatePaymentRequest>({
      query: userData => ({
        url: '/v1/payments/create',
        method: 'POST',
        body: { ...userData },
      }),
      // invalidatesTags: [{ type: 'User', id: 'LIST' }],
    }),

    getPaymentStatus: builder.mutation<GetPaymentResponse, GetPaymentRequest>({
      query: userData => ({
        url: '/v1/payments/status/by-id',
        method: 'POST',
        body: { ...userData },
      }),
      // invalidatesTags: [{ type: 'User', id: 'LIST' }],
    }),
  }),
});

export const { useCreatePaymentMutation, useGetPaymentStatusMutation } = paymentsApiExtended;
