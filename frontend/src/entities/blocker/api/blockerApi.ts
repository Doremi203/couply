import { blockerApi } from '../../../shared/api/baseApi';

import { CreateBlockRequest, GetBlockResponse } from './types';


export const blockerApiExtended = blockerApi.injectEndpoints({
  endpoints: builder => ({

    getBlockInfo: builder.mutation<GetBlockResponse, object>({
        query: () => ({
          url: '/v1/block/get',
          method: 'POST',
          body: {  },
        }),
        // invalidatesTags: [{ type: 'User', id: 'LIST' }],
      }),

      createComplaint: builder.mutation<object, CreateBlockRequest>({
        query: userData => ({
          url: '/v1/block/reports/create',
          method: 'POST',
          body: { ...userData },
        }),
        // invalidatesTags: [{ type: 'User', id: 'LIST' }],
      }),
    }),
});

export const {
  useGetBlockInfoMutation,
  useCreateComplaintMutation,
} = blockerApiExtended;
