// api.ts
import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';

export const fileUploadApi = createApi({
  reducerPath: 'fileUploadApi',
  baseQuery: fetchBaseQuery({ baseUrl: '/' }),
  endpoints: (builder) => ({
    uploadFileToS3: builder.mutation<void, { url: string; file: File }>({
      query: ({ url, file }) => ({
        url: url,
        method: 'PUT',
        body: file,
        headers: {
          'Content-Type': file.type,
        },
      }),
    }),
  }),
});

export const { useUploadFileToS3Mutation } = fileUploadApi;