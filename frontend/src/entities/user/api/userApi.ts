import { matcherApi } from '../../../shared/api/baseApi';

import { Alcohol, Children, Education, Goal, Smoking, Zodiac } from './constants';
import {
  UserResponse,
  UserRequest,
  UpdateUserParams,
  PhotoParams,
  UsersResponse,
  GetUsersRequest,
} from './types';

//TODO
const basicData = {
  location: 'Москва, Россия', // todo
  bio: 'bio',
  goal: Goal.dating, // УБРАТЬ
  education: Education.unspecified,
  children: Children.unspecified,
  alcohol: Alcohol.neutrally,
  smoking: Smoking.unspecified,
  zodiac: Zodiac.unspecified,
  hidden: false,
  verified: false,
  interest: null, //TODO
  //photos: null, //сллетаь
};

export const userApiExtended = matcherApi.injectEndpoints({
  endpoints: builder => ({
    createUser: builder.mutation<UserResponse, Partial<UserRequest>>({
      query: userData => ({
        url: '/v1/CreateUserV1',
        method: 'POST',
        body: { ...userData, ...basicData },
      }),
      // invalidatesTags: [{ type: 'User', id: 'LIST' }],
    }),

    getUser: builder.mutation<UserResponse, object>({
      query: id => ({
        url: '/v1/GetUserV1',
        method: 'POST',
        body: id,
      }),
      // invalidatesTags: (_result, _error, { id }) => [{ type: 'User', id }],
    }),

    getUsers: builder.mutation<UsersResponse, GetUsersRequest>({
      query: userIds => ({
        url: '/v1/GetUsersV1',
        method: 'POST',
        body: { userIds: userIds },
      }),
      // invalidatesTags: (_result, _error, { id }) => [{ type: 'User', id }],
    }),

    updateUser: builder.mutation<UserResponse, UpdateUserParams>({
      query: ({ data }) => ({
        url: 'v1/UpdateUserV1',
        method: 'POST',
        body: data,
      }),
      // invalidatesTags: (_result, _error, { id }) => [{ type: 'User', id }],
    }),

    deleteUser: builder.mutation<object, object>({
      query: () => ({
        url: '/v1/DeleteUserV1',
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
        url: 'v1/ConfirmPhotoUploadV1',
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
