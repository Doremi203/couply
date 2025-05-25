import { createSlice, PayloadAction } from '@reduxjs/toolkit';

interface BlockingState {
  isBlocked: boolean;
  reasons: string[];
  message: string;
  createdAt: string | null;
}

const initialState: BlockingState = {
  isBlocked: false,
  reasons: [],
  message: '',
  createdAt: null,
};

const blockingSlice = createSlice({
  name: 'blocking',
  initialState,
  reducers: {
    setBlocking: (
      state,
      action: PayloadAction<{
        isBlocked: boolean;
        reasons: string[];
        message: string;
        createdAt: string | null;
      }>,
    ) => {
      state.isBlocked = action.payload.isBlocked;
      state.reasons = action.payload.reasons;
      state.message = action.payload.message;
      state.createdAt = action.payload.createdAt;
    },
    clearBlocking: state => {
      state.isBlocked = false;
      state.reasons = [];
      state.message = '';
      state.createdAt = null;
    },
  },
});

export const { setBlocking, clearBlocking } = blockingSlice.actions;
export default blockingSlice.reducer;
