import { paymentsApi } from '../../../shared/api/baseApi';
import { CancelRequest, CreateSubRequst, GetSubscriptionResponse } from '../types';

export const subscriptionApi = paymentsApi.injectEndpoints({
  endpoints: builder => ({
    createSubscription: builder.mutation<GetSubscriptionResponse, CreateSubRequst>({
      query: data => ({
        url: '/v1/subscriptions/create',
        method: 'POST',
        body: data,
      }),
      // invalidatesTags: [{ type: 'User', id: 'LIST' }],
    }),

    cancelSubscription: builder.mutation<object, CancelRequest>({
      query: userData => ({
        url: '/v1/subscriptions/cancel/by-id',
        method: 'POST',
        body: { ...userData },
      }),
      // invalidatesTags: [{ type: 'User', id: 'LIST' }],
    }),

    getSubscription: builder.mutation<GetSubscriptionResponse, object>({
      query: userData => ({
        url: '/v1/subscriptions/active/get',
        method: 'POST',
        body: { ...userData },
      }),
      // invalidatesTags: [{ type: 'User', id: 'LIST' }],
    }),
  }),
});

export const {
  useCreateSubscriptionMutation,
  useCancelSubscriptionMutation,
  useGetSubscriptionMutation,
} = subscriptionApi;
