// import {
//   createApi,
//   fetchBaseQuery,
//   retry,
//   BaseQueryFn,
//   FetchArgs,
//   FetchBaseQueryError,
//   FetchBaseQueryMeta,
// } from '@reduxjs/toolkit/query/react';

// import { LoginResponse } from '../../entities/auth/types';
// import { isTokenExpired, getToken, setToken } from '../lib/services/TokenService';

// // Base URLs for APIs
// const AUTH_BASE_URL = 'https://auth.testing.couply.ru';
// const MATCHER_API_URL = 'https://matcher.testing.couply.ru';

// // Create a custom base query with token refresh logic for auth API
// const baseQueryWithReauth: BaseQueryFn<
//   string | FetchArgs,
//   unknown,
//   FetchBaseQueryError,
//   object,
//   FetchBaseQueryMeta
// > = async (args, api, extraOptions) => {
//   // Create our base query
//   const baseQuery = fetchBaseQuery({
//     baseUrl: AUTH_BASE_URL,
//     prepareHeaders: headers => {
//       // Check if token is about to expire
//       if (isTokenExpired()) {
//         console.log('Token is expired or about to expire');
//       }

//       // Add auth headers
//       const token = getToken();
//       if (token) {
//         headers.set('Authorization', `Bearer ${token}`);
//         headers.set('user-token', token);
//       }
//       return headers;
//     },
//   });

//   // Execute the initial query
//   let result = await baseQuery(args, api, extraOptions);

//   // If we get a 401 Unauthorized response, try to refresh the token
//   if (result.error && result.error.status === 401) {
//     console.log('Attempting to refresh token');

//     // Try to get a new token
//     const refreshQuery = fetchBaseQuery({ baseUrl: AUTH_BASE_URL });
//     const refreshResult = await refreshQuery(
//       { url: '/v1/refresh', method: 'POST', headers: { Authorization: `Bearer ${getToken()}` } },
//       api,
//       extraOptions,
//     );

//     // If we got a new token
//     if (refreshResult.data) {
//       const refreshData = refreshResult.data as LoginResponse;

//       // Store the new token
//       setToken(refreshData.token, refreshData.expiresIn);

//       // Retry the original query with new token
//       result = await baseQuery(args, api, extraOptions);
//     } else {
//       // If refresh failed, redirect to login
//       console.log('Token refresh failed, redirecting to login');
//       window.location.href = '/login';
//     }
//   }

//   return result;
// };

// // Create a custom base query with token refresh logic for matcher API
// const matcherQueryWithReauth: BaseQueryFn<
//   string | FetchArgs,
//   unknown,
//   FetchBaseQueryError,
//   object,
//   FetchBaseQueryMeta
// > = async (args, api, extraOptions) => {
//   // Create our base query
//   const baseQuery = fetchBaseQuery({
//     baseUrl: MATCHER_API_URL,
//     prepareHeaders: headers => {
//       // Check if token is about to expire
//       if (isTokenExpired()) {
//         console.log('Token is expired or about to expire');
//       }

//       // Add auth headers
//       const token = getToken();
//       if (token) {
//         headers.set('Authorization', `Bearer ${token}`);
//         headers.set('user-token', token);
//       }
//       return headers;
//     },
//   });

//   // Execute the initial query
//   let result = await baseQuery(args, api, extraOptions);

//   // If we get a 401 Unauthorized response, try to refresh the token
//   if (result.error && result.error.status === 401) {
//     console.log('Attempting to refresh token');

//     // Try to get a new token
//     const refreshQuery = fetchBaseQuery({ baseUrl: AUTH_BASE_URL });
//     const refreshResult = await refreshQuery(
//       { url: '/v1/refresh', method: 'POST', headers: { Authorization: `Bearer ${getToken()}` } },
//       api,
//       extraOptions,
//     );

//     // If we got a new token
//     if (refreshResult.data) {
//       const refreshData = refreshResult.data as LoginResponse;

//       // Store the new token
//       setToken(refreshData.token, refreshData.expiresIn);

//       // Retry the original query with new token
//       result = await baseQuery(args, api, extraOptions);
//     } else {
//       // If refresh failed, redirect to login
//       console.log('Token refresh failed, redirecting to login');
//       window.location.href = '/login';
//     }
//   }

//   return result;
// };

// export const baseApi = createApi({
//   reducerPath: 'api',
//   baseQuery: retry(baseQueryWithReauth, { maxRetries: 1 }),
//   endpoints: () => ({}),
//   tagTypes: ['Profile', 'Matches', 'User'],
// });

// export const matcherApi = createApi({
//   reducerPath: 'matherApi',
//   baseQuery: retry(matcherQueryWithReauth, { maxRetries: 1 }),
//   endpoints: () => ({}),
//   tagTypes: ['Profile', 'Matches', 'User'],
// });

import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';

const AUTH_BASE_URL = 'https://auth.testing.couply.ru';
const MATCHER_API_URL = 'https://matcher.testing.couply.ru';

export const baseApi = createApi({
  reducerPath: 'api',
  baseQuery: fetchBaseQuery({
    baseUrl: AUTH_BASE_URL,
    prepareHeaders: headers => {
      const token = localStorage.getItem('token');
      if (token) {
        headers.set('Authorization', `Bearer ${token}`);
      }
      return headers;
    },
  }),
  endpoints: () => ({}),
  tagTypes: ['Profile', 'Matches', 'User'], // TODO
});

export const matcherApi = createApi({
  reducerPath: 'matherApi',
  baseQuery: fetchBaseQuery({
    baseUrl: MATCHER_API_URL,
    prepareHeaders: headers => {
      const token = localStorage.getItem('token');
      if (token) {
        headers.set('Authorization', `Bearer ${token}`);
      }
      return headers;
    },
  }),
  endpoints: () => ({}),
  tagTypes: ['Profile', 'Matches', 'User'], // TODO
});
