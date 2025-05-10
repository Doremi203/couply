import { useCallback } from 'react';

import { useRegisterMutation, useLoginMutation } from '../api/authApi';
import { LoginParams, LoginResponse, RegisterParams } from '../api/types';

export const useRegister = () => {
  const [register] = useRegisterMutation();
  const [login] = useLoginMutation();

  const registerAndLogin = useCallback(
    async (registerParams: RegisterParams): Promise<LoginResponse> => {
      try {
        await register(registerParams).unwrap();

        const loginParams: LoginParams = {
          email: registerParams.email,
          password: registerParams.password,
        };

        const loginResponse = await login(loginParams).unwrap();

        localStorage.setItem('token', loginResponse.token);
        return loginResponse;
      } catch (error) {
        // TODO
        //@ts-ignore
        throw new Error(error?.data?.message || 'An error occurred during registration or login');
      }
    },
    [register, login],
  );

  return { registerAndLogin };
};
