import { useCallback } from 'react';

import { useLoginMutation } from '../api/authApi';
import { LoginParams, LoginResponse } from '../api/types';

export const useLogin = () => {
  const [login] = useLoginMutation();

  const loginUser = useCallback(
    async (loginParams: LoginParams): Promise<{ data?: LoginResponse; error?: string }> => {
      try {
        const loginResponse = await login(loginParams).unwrap();
        localStorage.setItem('token', loginResponse.token);
        return { data: loginResponse };
      } catch (error: any) {
        if (error?.status === 404) {
          return { error: 'Пользователь с таким email не найден' };
        } else if (error?.status === 401) {
          return { error: 'Неверный пароль' };
        } else if (error?.data?.message) {
          return { error: error.data.message };
        } else {
          return { error: 'Произошла ошибка при входе в систему' };
        }
      }
    },
    [login],
  );

  return { loginUser };
};
