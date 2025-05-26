import {
  createApi,
  fetchBaseQuery,
  BaseQueryFn,
  FetchArgs,
  FetchBaseQueryError,
  FetchBaseQueryMeta,
  QueryReturnValue,
} from '@reduxjs/toolkit/query/react';

import { getRefreshToken, getToken, setTokens, clearTokens } from '../lib/services/TokenService';

interface RefreshTokenResponse {
  accessToken: {
    token: string;
    expiresIn: number;
  };
  refreshToken: {
    token: string;
    expiresIn: number;
  };
}

const AUTH_BASE_URL = 'https://auth.testing.couply.ru';
const MATCHER_API_URL = 'https://matcher.testing.couply.ru';
const BLOCKER_API_URL = 'https://blocker.testing.couply.ru';
const PAYMENTS_API_URL = 'https://payments.testing.couply.ru';
const NOTIFICATOR_API_URL = 'https://notificator.testing.couply.ru';

const authBaseQuery = fetchBaseQuery({
  baseUrl: AUTH_BASE_URL,
  prepareHeaders: headers => {
    return headers;
  },
});

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

  return async (args: string | FetchArgs, api: any, extraOptions: object) => {
    const result = await baseQuery(args, api, extraOptions);

    // If 401 Unauthorized, attempt token refresh
    if (result.error?.status === 401) {
      if (!isRefreshing) {
        isRefreshing = true;

        try {
          const token = getToken();
          const refreshToken = getRefreshToken();

          if (!token || !refreshToken) {
            clearTokens();
            return result;
          }

          const refreshResult = await authBaseQuery(
            {
              url: '/v1/token/refresh',
              method: 'POST',
              body: {
                token: token,
                refreshToken: refreshToken,
              },
            },
            api,
            extraOptions,
          );

          if (refreshResult.data) {
            const data = refreshResult.data as RefreshTokenResponse;
            setTokens(data.accessToken.token, data.refreshToken.token, data.accessToken.expiresIn);

            const promises = pendingRequests.map(async request => {
              const result = await baseQuery(request.args, request.api, request.extraOptions);
              request.resolve(result);
              return result;
            });

            await Promise.all(promises);
          } else {
            clearTokens();
            window.location.href = '/auth';
          }
        } catch {
          clearTokens();
          // window.location.href = '/auth';
        } finally {
          isRefreshing = false;
          pendingRequests = [];
        }
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
