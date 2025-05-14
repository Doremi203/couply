import { useCallback, useState } from 'react';

import { useLoginMutation } from '../api/authApi';
import { LoginParams, LoginResponse } from '../api/types';

export const useLogin = () => {
  const [login] = useLoginMutation();

  const loginUser = useCallback(
    async (loginParams: LoginParams): Promise<LoginResponse> => {
      try {
        await login(loginParams).unwrap();

        const loginResponse = await login(loginParams).unwrap();

        localStorage.setItem('token', loginResponse.token);

        return loginResponse;
      } catch (error) {
        // TODO
        //@ts-ignore
        throw new Error(error?.data?.message || 'An error occurred during registration or login');
      }
    },
    [login],
  );

  return { loginUser };
};
