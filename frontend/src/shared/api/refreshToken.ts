import { getRefreshToken, getToken, setTokens } from '../lib/services/TokenService';

export const refreshToken = async () => {
  try {
    const token = getToken();
    const refreshToken = getRefreshToken();

    if (!token || !refreshToken) {
      // console.warn('No token or refresh token found');
      return false;
    }

    const response = await fetch('https://auth.testing.couply.ru/v1/token/refresh', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({ refreshToken }),
    });

    if (!response.ok) {
      // console.warn('Token refresh failed - server returned', response.status);
      return false;
    }

    const data = await response.json();

    if (data.accessToken && data.refreshToken) {
      setTokens(data.accessToken.token, data.refreshToken.token, data.accessToken.expiresIn);
      return true;
    } else {
      console.warn('Token refresh failed - invalid response format');
      return false;
    }
  } catch (error) {
    console.error('Token refresh error:', error);
    return false;
  }
};
