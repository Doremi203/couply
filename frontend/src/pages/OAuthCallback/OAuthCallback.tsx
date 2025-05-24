// src/pages/OAuthCallback.tsx
import { useEffect, useState } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';

export default function OAuthCallback() {
  const { search } = useLocation();
  const navigate = useNavigate();
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const params = new URLSearchParams(search);
    const code = params.get('code');
    const state = params.get('state');
    const deviceId = sessionStorage.getItem('device_id');

    if (!code || !state) {
      setError('Отсутствует code или state в URL');
      return;
    }

    // Проверяем state
    const origState = sessionStorage.getItem('oauth_state');
    if (state !== origState) {
      setError('Неверный OAuth-state');
      return;
    }

    // Достаём verifier
    const codeVerifier = sessionStorage.getItem('oauth_code_verifier');

    // Вызываем ваш бэкенд для обмена
    fetch('https://auth.testing.couply.ru/v1/login/oauth', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        provider: 'yandex',
        code,
        state,
        device_id: deviceId,
        code_verifier: codeVerifier,
      }),
    })
      .then(res => {
        if (!res.ok) throw new Error(`Server error ${res.status}`);
        return res.json();
      })
      .then(result => {
        // здесь точно такой же код, как в handleOAuthSuccess
        localStorage.setItem('token', result.accessToken.token);
        localStorage.setItem('refreshToken', result.refreshToken.token);

        if (result.firstLogin) {
          navigate('/enterInfo');
        } else {
          navigate('/home');
        }
      })
      .catch(err => {
        console.error('OAuth callback error', err);
        setError(err.message);
      });
  }, [search, navigate]);

  if (error) {
    return <div>Ошибка авторизации: {error}</div>;
  }

  return <div>Обрабатываем авторизацию… Пожалуйста, подождите.</div>;
}
