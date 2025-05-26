import { useCallback, useEffect, useRef } from 'react';

import { getToken, getRefreshToken, isTokenExpired, getTokenExpiryTime } from '../../../shared/lib/services/TokenService';
import { useRefreshTokenMutation } from '../api/authApi';


export const useTokenRefresh = () => {
  const [refreshToken, { isLoading }] = useRefreshTokenMutation();
  const timerRef = useRef<number | null>(null);

  const checkAndRefreshToken = useCallback(async () => {
    if (isTokenExpired()) {
      const token = getToken();
      const refreshTokenValue = getRefreshToken();
      
      if (token && refreshTokenValue) {
        try {
          await refreshToken({
            token,
            refreshToken: refreshTokenValue,
          }).unwrap();
          console.log('Token refreshed successfully');
        } catch (error) {
          console.error('Failed to refresh token:', error);
        }
      }
    }
  }, [refreshToken]);

  useEffect(() => {
    const setupExpiryTimer = () => {
      if (timerRef.current) {
        window.clearTimeout(timerRef.current);
      }

      const expiryTime = getTokenExpiryTime();
      if (!expiryTime) return;

      const timeUntilRefresh = Math.max(0, expiryTime - Date.now() - 5 * 60 * 1000);
      
      timerRef.current = window.setTimeout(() => {
        checkAndRefreshToken();
        setupExpiryTimer();
      }, timeUntilRefresh);
    };

    checkAndRefreshToken();
    setupExpiryTimer();

    return () => {
      if (timerRef.current) {
        window.clearTimeout(timerRef.current);
      }
    };
  }, [checkAndRefreshToken]);

  return {
    refreshToken: checkAndRefreshToken,
    isRefreshing: isLoading,
  };
};