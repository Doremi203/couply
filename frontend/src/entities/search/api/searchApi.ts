import { matcherApi } from '../../../shared/api/baseApi';

import {
  AddViewRequest,
  CreateFilterRequest,
  FilterResponse,
  SearchRequest,
  SearchResponse,
} from './types';

export const searchApi = matcherApi.injectEndpoints({
  endpoints: builder => ({
    createFilter: builder.mutation<FilterResponse, CreateFilterRequest>({
      query: data => ({
        url: '/v1/search/filters/create',
        method: 'POST',
        body: data,
      }),
      // invalidatesTags: [{ type: 'filter' }],
    }),

    getFilter: builder.query<FilterResponse, object>({
      query: data => ({
        url: '/v1/search/filters/me/get',
        method: 'POST',
        body: data,
      }),
    }),

    searchUsers: builder.mutation<SearchResponse, SearchRequest>({
      query: data => ({
        url: '/v1/search/users/search',
        method: 'POST',
        body: data,
      }),
      // invalidatesTags: [{ type: 'search' }],
    }),

    updateFilter: builder.mutation<FilterResponse, CreateFilterRequest>({
      query: data => ({
        url: '/v1/search/filters/update/me',
        method: 'POST',
        body: data,
      }),
      // invalidatesTags: [{ type: 'filter' }],
    }),

    addView: builder.mutation<object, AddViewRequest>({
      query: data => ({
        url: '/v1/search/views/add',
        method: 'POST',
        body: data,
      }),
      // invalidatesTags: [{ type: 'filter' }],
    }),
  }),
});

export const {
  useCreateFilterMutation,
  useGetFilterQuery,
  useSearchUsersMutation,
  useUpdateFilterMutation,
} = searchApi;
