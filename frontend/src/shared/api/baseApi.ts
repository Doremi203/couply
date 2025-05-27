import {
  createApi,
  fetchBaseQuery,
  BaseQueryFn,
  FetchArgs,
  FetchBaseQueryError,
  FetchBaseQueryMeta,
  QueryReturnValue,
} from '@reduxjs/toolkit/query/react';

import {
  getRefreshToken,
  getToken,
  clearTokens,
  isTokenExpired,
} from '../lib/services/TokenService';

import { refreshToken as refreshTokenFunction } from './refreshToken';

const AUTH_BASE_URL = 'https://auth.testing.couply.ru';
const MATCHER_API_URL = 'https://matcher.testing.couply.ru';
const BLOCKER_API_URL = 'https://blocker.testing.couply.ru';
const PAYMENTS_API_URL = 'https://payments.testing.couply.ru';
const NOTIFICATOR_API_URL = 'https://notificator.testing.couply.ru';

type BaseQueryType = BaseQueryFn<
  string | FetchArgs,
  unknown,
  FetchBaseQueryError,
  object,
  FetchBaseQueryMeta
>;

const baseQueryWithReauth = (baseQuery: BaseQueryType): BaseQueryType => {
  let isRefreshing = false;
  let pendingRequests: Array<{
    resolve: (value: QueryReturnValue<unknown, FetchBaseQueryError, FetchBaseQueryMeta>) => void;
    args: string | FetchArgs;
    api: any;
    extraOptions: object;
  }> = [];

  // We'll use a simple flag to prevent concurrent refresh attempts
  let isRefreshInProgress = false;

  const refreshTokenFn = async (_api: any, _extraOptions: object) => {
    try {
      // If a refresh is already in progress, don't start another one
      if (isRefreshInProgress) {
        console.log('Token refresh already in progress, skipping');
        return false;
      }

      // Set the flag to prevent concurrent refreshes
      isRefreshInProgress = true;

      const token = getToken();
      const refreshToken = getRefreshToken();

      if (!token || !refreshToken) {
        console.warn('No token or refresh token found');
        clearTokens();
        return false;
      }

      // Use our standalone refresh token function
      const success = await refreshTokenFunction();

      if (!success) {
        console.warn('Token refresh failed');
        // Only redirect if we're not already on the auth page
        if (!window.location.pathname.includes('/auth')) {
          clearTokens();
          window.location.href = '/auth';
        }
      }

      // Always reset the flag
      isRefreshInProgress = false;
      return success;
    } catch (error) {
      console.error('Error in refreshTokenFn:', error);
      isRefreshInProgress = false;
      return false;
    }
  };

  return async (args: string | FetchArgs, api: any, extraOptions: object) => {
    // Check if token is about to expire and we're not already refreshing
    if (isTokenExpired() && !isRefreshing) {
      // Prevent concurrent refresh attempts
      isRefreshing = true;

      try {
        const refreshSuccess = await refreshTokenFn(api, extraOptions);

        if (!refreshSuccess) {
          console.warn('Proactive token refresh failed, proceeding with original request');
        } else {
          console.log('Proactive token refresh succeeded');
        }
      } catch (error) {
        console.error('Error during proactive token refresh:', error);
      } finally {
        // Always reset the refreshing flag
        isRefreshing = false;
      }
    }

    const result = await baseQuery(args, api, extraOptions);

    if (result.error?.status === 401) {
      if (!isRefreshing) {
        isRefreshing = true;

        const refreshSuccess = await refreshTokenFn(api, extraOptions);

        if (refreshSuccess) {
          const promises = pendingRequests.map(async request => {
            const result = await baseQuery(request.args, request.api, request.extraOptions);
            request.resolve(result);
            return result;
          });

          await Promise.all(promises);
        }

        isRefreshing = false;
        pendingRequests = [];
      }

      if (isRefreshing) {
        return new Promise<QueryReturnValue<unknown, FetchBaseQueryError, FetchBaseQueryMeta>>(
          resolve => {
            pendingRequests.push({ resolve, args, api, extraOptions });
          },
        );
      }
    }

    return result;
  };
};

interface ApiOptions {
  reducerPath: string;
  baseQueryOptions: {
    baseUrl: string;
    prepareHeaders: (headers: Headers) => Headers;
  };
  endpoints: (builder: any) => Record<string, any>;
}

const createApiWithReauth = (options: ApiOptions) =>
  createApi({
    ...options,
    baseQuery: baseQueryWithReauth(fetchBaseQuery(options.baseQueryOptions)),
  });

export const baseApi = createApiWithReauth({
  reducerPath: 'api',
  baseQueryOptions: {
    baseUrl: AUTH_BASE_URL,
    prepareHeaders: (headers: Headers) => {
      const token = localStorage.getItem('token');
      if (token) {
        headers.set('Authorization', `Bearer ${token}`);
        headers.set('user-token', token);
      }
      return headers;
    },
  },
  endpoints: () => ({}),
});

export const matcherApi = createApiWithReauth({
  reducerPath: 'matcherApi',
  baseQueryOptions: {
    baseUrl: MATCHER_API_URL,
    prepareHeaders: (headers: Headers) => {
      const token = localStorage.getItem('token');
      if (token) {
        headers.set('Authorization', `Bearer ${token}`);
        headers.set('user-token', token);
      }
      return headers;
    },
  },
  endpoints: () => ({}),
});

export const blockerApi = createApiWithReauth({
  reducerPath: 'blockerApi',
  baseQueryOptions: {
    baseUrl: BLOCKER_API_URL,
    prepareHeaders: (headers: Headers) => {
      const token = localStorage.getItem('token');
      if (token) {
        headers.set('Authorization', `Bearer ${token}`);
        headers.set('user-token', token);
      }
      return headers;
    },
  },
  endpoints: () => ({}),
});

export const paymentsApi = createApiWithReauth({
  reducerPath: 'paymentsApi',
  baseQueryOptions: {
    baseUrl: PAYMENTS_API_URL,
    prepareHeaders: (headers: Headers) => {
      const token = localStorage.getItem('token');
      if (token) {
        headers.set('Authorization', `Bearer ${token}`);
        headers.set('user-token', token);
      }
      return headers;
    },
  },
  endpoints: () => ({}),
});

export const notificatorApi = createApiWithReauth({
  reducerPath: 'notificatorApi',
  baseQueryOptions: {
    baseUrl: NOTIFICATOR_API_URL,
    prepareHeaders: (headers: Headers) => {
      const token = localStorage.getItem('token');
      if (token) {
        headers.set('user-token', token);
      }
      return headers;
    },
  },
  endpoints: () => ({}),
});
