import { useCallback, useEffect, useRef, useState } from 'react';

import { refreshToken } from '../../../shared/api/refreshToken';
import { isTokenExpired, getTokenExpiryTime } from '../../../shared/lib/services/TokenService';

export const useTokenRefresh = () => {
  const [isRefreshing, setIsRefreshing] = useState(false);
  const timerRef = useRef<number | null>(null);

  const checkAndRefreshToken = useCallback(async () => {
    if (isTokenExpired()) {
      try {
        setIsRefreshing(true);
        const success = await refreshToken();
        return success;
      } catch (error) {
        console.error('Failed to refresh token:', error);
        return false;
      } finally {
        setIsRefreshing(false);
      }
    }

    return false;
  }, []);

  useEffect(() => {
    const setupExpiryTimer = () => {
      if (timerRef.current) {
        window.clearTimeout(timerRef.current);
      }

      const expiryTime = getTokenExpiryTime();
      if (!expiryTime) {
        checkAndRefreshToken();
        return;
      }

      const timeUntilExpiry = Math.max(0, expiryTime - Date.now());

      timerRef.current = window.setTimeout(() => {
        checkAndRefreshToken().then(() => {
          setupExpiryTimer();
        });
      }, timeUntilExpiry);
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
    isRefreshing,
  };
};
