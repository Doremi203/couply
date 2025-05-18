import { useEffect, useRef } from 'react';
import { useNavigate } from 'react-router-dom';

import { ButtonWithIcon } from '../../../../shared/components/ButtonWithIcon';

import styles from './authPage.module.css';

interface YaAuthResponse {
  handler: () => Promise<any>;
}

export const AuthPage = () => {
  const navigate = useNavigate();
  const yandexButtonContainerRef = useRef<HTMLDivElement>(null);

  const onLogin = () => {
    navigate('/login');
  };

  useEffect(() => {
    // Получаем ширину другой кнопки для согласованности
    const emailButton = document.querySelector('[class*="ButtonWithIcon"]');
    const buttonWidth = emailButton ? emailButton.clientWidth : 280;

    // Load Yandex Auth SDK
    const script = document.createElement('script');
    script.src =
      'https://yastatic.net/s3/passport-sdk/autofill/v1/sdk-suggest-with-polyfills-latest.js';
    script.async = true;

    script.onload = () => {
      if (window.YaAuthSuggest && yandexButtonContainerRef.current) {
        // Установим фиксированную ширину для контейнера
        yandexButtonContainerRef.current.style.width = `${buttonWidth}px`;
        yandexButtonContainerRef.current.style.maxWidth = `${buttonWidth}px`;

        window.YaAuthSuggest.init(
          {
            client_id: '1d2e86281ff2444fa0d11e63b09afae0',
            response_type: 'token',
            redirect_uri: 'https://auth.testing.couply.ru/v1/login/yandex',
          },
          'https://auth.testing.couply.ru',
          {
            view: 'button',
            parentId: 'yandex-button-container',
            buttonView: 'main',
            buttonTheme: 'light',
            buttonSize: 'm',
            buttonBorderRadius: 25,
            // Попробуем другие возможные параметры ширины
            buttonWidth: buttonWidth,
          },
        )
          .then(({ handler }: YaAuthResponse) => {
            // После инициализации найдем кнопку и принудительно установим ширину
            setTimeout(() => {
              const yaButton = yandexButtonContainerRef.current?.querySelector('button');
              if (yaButton) {
                yaButton.style.width = `${buttonWidth}px`;
                yaButton.style.maxWidth = `${buttonWidth}px`;
              }
            }, 100);

            return handler();
          })
          .then((data: any) => console.log('Сообщение с токеном', data))
          .catch((error: any) => console.log('Обработка ошибки', error));
      }
    };

    document.head.appendChild(script);

    return () => {
      if (document.head.contains(script)) {
        document.head.removeChild(script);
      }
    };
  }, []);

  return (
    <div className={styles.page}>
      <img src="pv3.jpg" width="200px" height="150px" alt="Logo" className={styles.logo} />
      <span className={styles.text}>Найди того, кто будет похож на тебя, как капля воды.</span>

      <div className={styles.buttons}>
        {/* <ButtonWithIcon
          onClick={onLogin}
          icon={<img src="vk.png" width="20px" height="20px" alt="VK" />}
          text="Войти с VK ID"
          iconClassName={styles.vkIcon}
        /> */}

        <ButtonWithIcon
          onClick={onLogin}
          icon={<img src="email2.png" width="26px" height="26px" alt="email" />}
          text="Войти по почте"
        />

        {/* Обертка с фиксированной шириной для контроля кнопки Яндекса */}
        <div className={styles.yandexButtonWrapper}>
          <div
            id="yandex-button-container"
            ref={yandexButtonContainerRef}
            className={styles.yandexButton}
          />
        </div>
      </div>
    </div>
  );
};

// Add TypeScript declaration for the YaAuthSuggest global variable
declare global {
  interface Window {
    YaAuthSuggest?: any;
  }
}

export default AuthPage;
