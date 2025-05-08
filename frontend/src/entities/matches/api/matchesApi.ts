import { matcherApi } from '../../../shared/api/baseApi';

import {
  CreateMatchRequest,
  FetchInMatchesRequest,
  FetchMatchesRequest,
  Match,
  MatchRequest,
  MatchesResponse,
} from './types';

export const matchesApi = matcherApi.injectEndpoints({
  endpoints: builder => ({
    createMatch: builder.mutation<Match, CreateMatchRequest>({
      query: data => ({
        url: '/v1/CreateMatchV1',
        method: 'POST',
        body: data,
      }),
      invalidatesTags: [{ type: 'Matches', id: 'LIST' }],
    }),

    deleteMatch: builder.mutation<object, CreateMatchRequest>({
      query: data => ({
        url: '/v1/DeleteMatchV1',
        method: 'POST',
        body: data,
      }),
      invalidatesTags: [{ type: 'Matches', id: 'LIST' }],
    }),

    updateMatch: builder.mutation<Match, MatchRequest>({
      query: data => ({
        url: '/v1/UpdateMatchV1',
        method: 'POST',
        body: data,
      }),
      invalidatesTags: [{ type: 'Matches', id: 'LIST' }],
    }),

    fetchIncomingMatches: builder.mutation<Match[], FetchInMatchesRequest>({
      query: data => ({
        url: '/v1/FetchIncomingMatchesV1',
        method: 'POST',
        body: data,
      }),
      invalidatesTags: [{ type: 'Matches', id: 'LIST' }],
    }),

    fetchMatches: builder.mutation<MatchesResponse, FetchMatchesRequest>({
      query: data => ({
        url: '/v1/FetchMatchesV1',
        method: 'POST',
        body: data,
      }),
      invalidatesTags: [{ type: 'Matches', id: 'LIST' }],
    }),

    fetchOutgoingMatches: builder.mutation<Match[], FetchMatchesRequest>({
      query: data => ({
        url: '/v1/FetchOutgoingMatchesV1',
        method: 'POST',
        body: data,
      }),
      invalidatesTags: [{ type: 'Matches', id: 'LIST' }],
    }),
  }),
});

export const {
  useCreateMatchMutation,
  useDeleteMatchMutation,
  useUpdateMatchMutation,
  useFetchIncomingMatchesMutation,
  useFetchMatchesMutation,
  useFetchOutgoingMatchesMutation,
} = matchesApi;
