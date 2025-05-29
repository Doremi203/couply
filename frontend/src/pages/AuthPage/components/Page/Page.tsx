import { useNavigate } from 'react-router-dom';

import { ButtonWithIcon } from '../../../../shared/components/ButtonWithIcon';
import VkOneTapButton, {
  OAuthLoginV1Response,
} from '../../../../shared/components/VKIDButton/VkOneTapButton.tsx';
import {
  generateCodeChallenge,
  generateCodeVerifier,
  generateDeviceId,
  generateState,
} from '../../../../shared/lib/services/OAuthService.ts';

import styles from './authPage.module.css';

export const AuthPage = () => {
  const navigate = useNavigate();

  const onLogin = () => {
    navigate('/login');
  };

  const onYandexClick = async () => {
    try {
      const codeVerifier = generateCodeVerifier();
      sessionStorage.setItem('oauth_code_verifier', codeVerifier);
      const state = generateState();
      sessionStorage.setItem('oauth_state', state);
      const deviceId = generateDeviceId();
      sessionStorage.setItem('device_id', deviceId);
      const codeChallenge = await generateCodeChallenge(codeVerifier);

      const params = new URLSearchParams({
        response_type: 'code',
        client_id: '1d2e86281ff2444fa0d11e63b09afae0',
        redirect_uri: 'https://testing.couply.ru/oauth-callback',
        device_id: deviceId,
        state: state,
        code_challenge: codeChallenge,
        code_challenge_method: 'S256',
        force_confirm: 'yes',
      });

      window.location.href = `https://oauth.yandex.ru/authorize?${params.toString()}`;
    } catch (err) {
      console.error('Yandex OAuth error:', err);
    }
  };

  const handleVKIDError = (error: any) => {
    console.error('VK ID login error:', error);
  };

  const handleOAuthSuccess = (result: OAuthLoginV1Response) => {
    localStorage.setItem('token', result.accessToken.token);
    localStorage.setItem('refreshToken', result.refreshToken.token);
    if (result.firstLogin) {
      navigate('/enterInfo');
    } else {
      navigate('/home');
    }
  };

  const handleVKIDAuthError = (error: Error) => {
    console.error('VK ID authentication error:', error);
  };

  const openTermsOfService = () => {
    navigate('/terms');
  };

  const openPrivacyPolicy = () => {
    navigate('/privacy');
  };

  return (
    <div className={styles.page}>
      <img src="pv3.jpg" width="200px" height="150px" alt="Logo" className={styles.logo} />
      <span className={styles.text}>Найди того, кто будет похож на тебя, как капля воды.</span>

      <div className={styles.buttons}>
        <div data-testid="email-button">
          <ButtonWithIcon
            onClick={onLogin}
            icon={<img src="email2.png" width="26px" height="26px" alt="email" />}
            text="Войти по почте"
          />
        </div>

        <ButtonWithIcon
          onClick={onYandexClick}
          icon={<img src="yandex.png" width="26px" height="26px" alt="email" />}
          text="Войти c Яндекс ID"
          className={styles.yandex}
        />

        <VkOneTapButton
          onError={handleVKIDError}
          onSuccess={handleOAuthSuccess}
          onAuthError={handleVKIDAuthError}
        />
      </div>

      <div className={styles.terms}>
        <p>
          Продолжая, вы принимаете{' '}
          <button onClick={openTermsOfService} className={styles.link}>
            пользовательское соглашение
          </button>{' '}
          и{' '}
          <button onClick={openPrivacyPolicy} className={styles.link}>
            политику конфиденциальности
          </button>
        </p>
      </div>
    </div>
  );
};

export default AuthPage;
