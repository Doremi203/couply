import { matcherApi } from '../../../shared/api/baseApi';

import { CreateFilterRequest, FilterResponse, SearchRequest, SearchResponse } from './types';

export const searchApi = matcherApi.injectEndpoints({
  endpoints: builder => ({
    createFilter: builder.mutation<FilterResponse, CreateFilterRequest>({
      query: data => ({
        url: '/v1/CreateFilterV1',
        method: 'POST',
        body: data,
      }),
      invalidatesTags: [{ type: 'filter' }],
    }),

    getFilter: builder.query<FilterResponse, object>({
      query: data => ({
        url: '/v1/GetFilterV1',
        method: 'POST',
        body: data,
      }),
    }),

    searchUsers: builder.mutation<SearchResponse, SearchRequest>({
      query: data => ({
        url: '/v1/SearchUsersV1',
        method: 'POST',
        body: data,
      }),
      invalidatesTags: [{ type: 'search' }],
    }),

    updateFilter: builder.mutation<FilterResponse, CreateFilterRequest>({
      query: data => ({
        url: '/v1/UpdateFilterV1',
        method: 'POST',
        body: data,
      }),
      invalidatesTags: [{ type: 'filter' }],
    }),
  }),
});

export const {
  useCreateFilterMutation,
  useGetFilterQuery,
  useSearchUsersMutation,
  useUpdateFilterMutation,
} = searchApi;
