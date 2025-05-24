// VkOneTapButton.tsx
import * as VKID from '@vkid/sdk';
import React, { useEffect, useRef } from 'react';

import { generateCodeVerifier } from '../../lib/services/OAuthService';

export interface VkOneTapButtonProps {
  /** Тема кнопки: светлая или тёмная */
  scheme?: VKID.Scheme;
  /** Язык виджета */
  lang?: VKID.Languages;
  /** Обработчик ошибок */
  onError?: (error: Error) => void;
  /** Успешная авторизация через VK OneTap */
  onSuccess?: (result: OAuthLoginV1Response) => void;
  /** Ошибка при обмене кода авторизации */
  onAuthError?: (error: Error) => void;
}

export interface Token {
  /** Поле token (string) */
  token: string;
  /** Поле expires_in (int32) */
  expiresIn: number;
}

export interface OAuthLoginV1Response {
  /** Поле access_token (Token) */
  accessToken: Token;
  /** Поле refresh_token (Token) */
  refreshToken: Token;
  /** Поле first_login (bool) */
  firstLogin: boolean;
}

const VkOneTapButton: React.FC<VkOneTapButtonProps> = ({
  scheme = VKID.Scheme.LIGHT,
  lang = VKID.Languages.RUS,
  onError = err => console.error('VK OneTap error:', err),
  onSuccess = res => console.log('VK login success:', res),
  onAuthError = err => console.error('VK auth exchange error:', err),
}) => {
  const containerRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (!containerRef.current) return;

    const codeVerifier = generateCodeVerifier();
    sessionStorage.getItem('vk_code_verifier');

    VKID.Config.update({ codeVerifier: codeVerifier });

    const oneTap = new VKID.OneTap();

    // Рендерим в наш контейнер
    oneTap
      .render({
        container: containerRef.current,
        scheme,
        lang,
      })
      .on(VKID.WidgetEvents.ERROR, onError);
    oneTap.on(VKID.OneTapInternalEvents.LOGIN_SUCCESS, (payload: VKID.AuthResponse) => {
      const { code, device_id, state } = payload;
      const codeVerifier = sessionStorage.getItem('vk_code_verifier');
      fetch('https://auth.testing.couply.ru/v1/login/oauth', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          provider: 'vk',
          state,
          device_id,
          code,
          code_verifier: codeVerifier,
        }),
      })
        .then(res => res.json())
        .then(onSuccess)
        .catch(onAuthError);
    });

    return () => {
      oneTap.close();
    };
  }, [scheme, lang, onError, onSuccess, onAuthError]);

  return <div id="VkIdSdkOneTap" ref={containerRef} style={{ display: 'inline-block' }} />;
};

export default VkOneTapButton;
