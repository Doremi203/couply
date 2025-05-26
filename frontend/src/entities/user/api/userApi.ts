import { matcherApi } from '../../../shared/api/baseApi';

import { Alcohol, Children, Education, Smoking, Zodiac } from './constants';
import {
  UserResponse,
  UserRequest,
  UpdateUserParams,
  PhotoParams,
  UsersResponse,
  GetUsersRequest,
} from './types';


const basicData = {
  bio: 'bio',
  education: Education.unspecified,
  children: Children.unspecified,
  alcohol: Alcohol.unspecified,
  smoking: Smoking.unspecified,
  zodiac: Zodiac.unspecified,
  hidden: false,
  verified: false,
  interest: null,
};

export const userApiExtended = matcherApi.injectEndpoints({
  endpoints: builder => ({
    createUser: builder.mutation<UserResponse, Partial<UserRequest>>({
      query: userData => ({
        url: '/v1/users/create',
        method: 'POST',
        body: { ...userData, ...basicData },
      }),
      // invalidatesTags: [{ type: 'User', id: 'LIST' }],
    }),

    getUser: builder.mutation<UserResponse, object>({
      query: id => ({
        url: '/v1/users/get/me',
        method: 'POST',
        body: id,
      }),
      // invalidatesTags: (_result, _error, { id }) => [{ type: 'User', id }],
    }),

    getUserById: builder.mutation<UserResponse, object>({
      query: () => ({
        url: '/v1/users/get/by-id',
        method: 'POST',
        // body: id,
      }),
      // invalidatesTags: (_result, _error, { id }) => [{ type: 'User', id }],
    }),

    getUsers: builder.mutation<UsersResponse, GetUsersRequest>({
      query: userIds => ({
        url: '/v1/users/batch/get',
        method: 'POST',
        body: { userIds: userIds },
      }),
      // invalidatesTags: (_result, _error, { id }) => [{ type: 'User', id }],
    }),

    updateUser: builder.mutation<UserResponse, UpdateUserParams>({
      query: userData => ({
        url: 'v1/users/update/me',
        method: 'POST',
        body: { ...userData },
      }),
      // invalidatesTags: (_result, _error, { id }) => [{ type: 'User', id }],
    }),

    updateUserById: builder.mutation<UserResponse, UpdateUserParams>({
      query: ({ data }) => ({
        url: 'v1/users/update/by-id',
        method: 'POST',
        body: data,
      }),
      // invalidatesTags: (_result, _error, { id }) => [{ type: 'User', id }],
    }),

    deleteUser: builder.mutation<object, object>({
      query: () => ({
        url: '/v1/users/delete/me',
        method: 'POST',
        body: {},
      }),
      // invalidatesTags: (_result, _error, id) => [
      //   { type: 'User', id },
      //   { type: 'User', id: 'LIST' },
      // ],
    }),
    confirmPhoto: builder.mutation<object, PhotoParams>({
      query: data => ({
        url: 'v1/users/me/photos/confirm',
        method: 'POST',
        body: data,
      }),
      // invalidatesTags: (_result, _error, { id }) => [{ type: 'User', id }],
    }),
  }),
});

export const {
  useCreateUserMutation,
  useGetUserMutation,
  useUpdateUserMutation,
  useDeleteUserMutation,
  useConfirmPhotoMutation,
  useGetUsersMutation,
} = userApiExtended;
