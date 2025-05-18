import { matcherApi } from '../../../shared/api/baseApi';

import {
  FetchMatchesRequest,
  FetchMatchesResponse,
  FetchMatchesUserIdsResponse,
  LikeRequest,
  LikeResponse,
  MatchRequest,
} from './types';

export const matchesApi = matcherApi.injectEndpoints({
  endpoints: builder => ({
    likeUser: builder.mutation<LikeResponse, LikeRequest>({
      query: data => ({
        url: '/v1/matches/likes/add',
        method: 'POST',
        body: data,
      }),
      // invalidatesTags: [{ type: 'Matches', id: 'LIST' }],
    }),

    dislikeUser: builder.mutation<object, MatchRequest>({
      query: data => ({
        url: '/v1/matches/dislikes/add',
        method: 'POST',
        body: data,
      }),
      // invalidatesTags: [{ type: 'Matches', id: 'LIST' }],
    }),

    deleteMatch: builder.mutation<object, MatchRequest>({
      query: data => ({
        url: '/v1/matches/delete/byId',
        method: 'POST',
        body: data,
      }),
      // invalidatesTags: [{ type: 'Matches', id: 'LIST' }],
    }),

    fetchIncomingLikes: builder.mutation<FetchMatchesResponse, FetchMatchesRequest>({
      query: data => ({
        url: '/v1/matches/likes/incoming/list',
        method: 'POST',
        body: data,
      }),
      // invalidatesTags: [{ type: 'Matches', id: 'LIST' }],
    }),

    fetchMatchesUserIds: builder.mutation<FetchMatchesUserIdsResponse, FetchMatchesRequest>({
      query: data => ({
        url: '/v1/matches/likes/outgoing/list',
        method: 'POST',
        body: data,
      }),
      // invalidatesTags: [{ type: 'Matches', id: 'LIST' }],
    }),

    fetchOutgoingLikes: builder.mutation<FetchMatchesResponse, FetchMatchesRequest>({
      query: data => ({
        url: '/v1/matches/list',
        method: 'POST',
        body: data,
      }),
      // invalidatesTags: [{ type: 'Matches', id: 'LIST' }],
    }),
  }),
});

export const {
  useDeleteMatchMutation,
  useFetchIncomingLikesMutation,
  useFetchOutgoingLikesMutation,
  useLikeUserMutation,
  useDislikeUserMutation,
  useFetchMatchesUserIdsMutation,
} = matchesApi;
