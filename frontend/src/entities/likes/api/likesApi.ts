import { baseApi } from '../../../shared/api/baseApi';

// Типы данных для лайков
export interface Like {
  id: string;
  userId: string;
  targetUserId: string;
  createdAt: string;
}

// Типы данных для матчей
export interface Match {
  id: string;
  users: [string, string]; // ID двух пользователей, у которых совпадение
  createdAt: string;
  lastMessageAt: string | null;
}

// Параметры для создания лайка
export interface CreateLikeParams {
  targetUserId: string;
}

// Расширение базового API для работы с лайками и матчами
export const likesApi = baseApi.injectEndpoints({
  endpoints: (builder) => ({
    // Получение лайков пользователя
    getUserLikes: builder.query<Like[], void>({
      query: () => '/likes/me',
      providesTags: (result) => 
        result 
          ? [
              ...result.map(({ id }) => ({ type: 'Likes' as const, id })),
              { type: 'Likes', id: 'LIST' },
            ]
          : [{ type: 'Likes', id: 'LIST' }],
    }),
    
    // Получение лайков, полученных пользователем
    getReceivedLikes: builder.query<Like[], void>({
      query: () => '/likes/received',
      providesTags: (result) => 
        result 
          ? [
              ...result.map(({ id }) => ({ type: 'Likes' as const, id })),
              { type: 'Likes', id: 'RECEIVED' },
            ]
          : [{ type: 'Likes', id: 'RECEIVED' }],
    }),
    
    // Создание лайка
    createLike: builder.mutation<Like, CreateLikeParams>({
      query: (data) => ({
        url: '/likes',
        method: 'POST',
        body: data,
      }),
      invalidatesTags: [{ type: 'Likes', id: 'LIST' }],
    }),
    
    // Удаление лайка
    deleteLike: builder.mutation<void, string>({
      query: (id) => ({
        url: `/likes/${id}`,
        method: 'DELETE',
      }),
      invalidatesTags: (result, error, id) => [{ type: 'Likes', id }],
    }),
    
    // Получение матчей пользователя
    getUserMatches: builder.query<Match[], void>({
      query: () => '/matches',
      providesTags: (result) => 
        result 
          ? [
              ...result.map(({ id }) => ({ type: 'Matches' as const, id })),
              { type: 'Matches', id: 'LIST' },
            ]
          : [{ type: 'Matches', id: 'LIST' }],
    }),
    
    // Получение конкретного матча
    getMatch: builder.query<Match, string>({
      query: (id) => `/matches/${id}`,
      providesTags: (result, error, id) => [{ type: 'Matches', id }],
    }),
  }),
});

// Экспорт хуков для использования в компонентах
export const {
  useGetUserLikesQuery,
  useGetReceivedLikesQuery,
  useCreateLikeMutation,
  useDeleteLikeMutation,
  useGetUserMatchesQuery,
  useGetMatchQuery,
} = likesApi;