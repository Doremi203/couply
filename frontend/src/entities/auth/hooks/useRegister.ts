import { useCallback } from 'react';

import { setTokens } from '../../../shared/lib/services/TokenService';
import { useRegisterMutation, useLoginMutation } from '../api/authApi';
import { LoginParams, LoginResponse, RegisterParams } from '../api/types';

export const useRegister = () => {
  const [register] = useRegisterMutation();
  const [login] = useLoginMutation();

  const registerAndLogin = useCallback(
    async (registerParams: RegisterParams): Promise<{ data?: LoginResponse; error?: string }> => {
      try {
        await register(registerParams).unwrap();

        const loginParams: LoginParams = {
          email: registerParams.email,
          password: registerParams.password,
        };

        const loginResponse = await login(loginParams).unwrap();
        setTokens(loginResponse.token, loginResponse.refreshToken.token, loginResponse.expiresIn);
        return { data: loginResponse };
      } catch (error: any) {
        // Provide specific error messages based on error response
        if (error?.status === 409) {
          return { error: 'Пользователь с таким email уже существует' };
        } else if (error?.status === 400) {
          // Handle validation errors from the API
          if (error?.data?.message?.includes('email')) {
            return { error: 'Некорректный формат email' };
          } else if (
            error?.data?.message?.includes('password must be at least 6 characters long')
          ) {
            return { error: 'Пароль должен содержать не менее 6 символов' };
          } else if (
            error?.data?.message?.includes('password must be at most 16 characters long')
          ) {
            return { error: 'Пароль должен содержать не более 16 символов' };
          } else if (
            error?.data?.message?.includes('password must contain at least one special character')
          ) {
            return { error: 'Пароль должен содержать хотя бы один специальный символ (_-!@#?)' };
          } else if (
            error?.data?.message?.includes('password must contain at least one uppercase letter')
          ) {
            return { error: 'Пароль должен содержать хотя бы одну заглавную букву' };
          } else if (error?.data?.message?.includes('password')) {
            return { error: 'Ненадежный пароль' };
          } else if (error?.data?.message) {
            return { error: error.data.message };
          }
        }
        return { error: 'Произошла ошибка при регистрации' };
      }
    },
    [register, login],
  );

  return { registerAndLogin };
};
