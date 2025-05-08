import { matcherApi } from '../../../shared/api/baseApi';

import {
  CreateFilterRequest,
  GetFilterRequest,
  GetFilterResponse,
  SearchRequest,
  SearchResponse,
} from './types';

export const searchApi = matcherApi.injectEndpoints({
  endpoints: builder => ({
    createFilter: builder.mutation<GetFilterResponse, CreateFilterRequest>({
      query: data => ({
        url: '/v1/CreateFilterV1',
        method: 'POST',
        body: data,
      }),
      invalidatesTags: [{ type: 'Matches', id: 'LIST' }], //TODO
    }),

    getFilter: builder.query<GetFilterResponse, GetFilterRequest>({
      query: data => ({
        url: '/v1/GetFilterV1',
        method: 'POST',
        body: data,
      }),
      //   invalidatesTags: [{ type: 'Matches', id: 'LIST' }],
    }),

    searchUsers: builder.mutation<SearchResponse, SearchRequest>({
      query: data => ({
        url: '/v1/SearchUsersV1',
        method: 'POST',
        body: data,
      }),
      invalidatesTags: [{ type: 'Matches', id: 'LIST' }],
    }),

    updateFilter: builder.mutation<GetFilterResponse, CreateFilterRequest>({
      query: data => ({
        url: '/v1/UpdateFilterV1',
        method: 'POST',
        body: data,
      }),
      invalidatesTags: [{ type: 'Matches', id: 'LIST' }],
    }),
  }),
});

export const {
  useCreateFilterMutation,
  useGetFilterQuery,
  useSearchUsersMutation,
  useUpdateFilterMutation,
} = searchApi;
