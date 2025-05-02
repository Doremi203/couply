import { matcherApi } from '../../../shared/api/baseApi';
import {
  Alcohol,
  Children,
  Education,
  GetUserRequest,
  Goal,
  Smoking,
  UpdateUserParams,
  UserRequest,
  UserResponse,
  Zodiac,
} from '../types';

const basicData = {
  location: 'Москва, Россия',
  bio: 'dfg',
  goal: Goal.dating,
  height: 170,
  education: Education.unspecified,
  children: Children.unspecified,
  alcohol: Alcohol.neutrally,
  smoking: Smoking.unspecified,
  zodiac: Zodiac.unspecified,
  hidden: false,
  verified: false,
  interest: null, //TODO
  //photos: null,
};

export const userApiExtended = matcherApi.injectEndpoints({
  endpoints: builder => ({
    createUser: builder.mutation<UserResponse, Partial<UserRequest>>({
      query: userData => ({
        url: '/v1/CreateUserV1',
        method: 'POST',
        body: { ...userData, ...basicData },
      }),
      invalidatesTags: [{ type: 'User', id: 'LIST' }],
    }),

    getUser: builder.mutation<UserResponse, GetUserRequest>({
      query: id => ({
        url: '/v1/GetUserV1',
        method: 'POST',
        body: id,
      }),
      invalidatesTags: (_result, _error, { id }) => [{ type: 'User', id }],
    }),

    updateUser: builder.mutation<UserResponse, UpdateUserParams>({
      query: ({ data }) => ({
        url: 'v1/UpdateUserV1',
        method: 'POST',
        body: data,
      }),
      invalidatesTags: (_result, _error, { id }) => [{ type: 'User', id }],
    }),

    deleteUser: builder.mutation<void, string>({
      query: id => ({
        url: '/v1/DeleteUserV1',
        method: 'POST',
        body: { id },
      }),
      invalidatesTags: (_result, _error, id) => [
        { type: 'User', id },
        { type: 'User', id: 'LIST' },
      ],
    }),
  }),
});

export const {
  useCreateUserMutation,
  useGetUserMutation,
  useUpdateUserMutation,
  useDeleteUserMutation,
} = userApiExtended;
