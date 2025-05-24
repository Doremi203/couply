import { useNavigate } from 'react-router-dom';

import { ButtonWithIcon } from '../../../../shared/components/ButtonWithIcon';

import styles from './authPage.module.css';

export const AuthPage = () => {
  const navigate = useNavigate();

  const onLogin = () => {
    navigate('/login');
  };

  const onYandex = () => {
    window.location.href =
      'https://oauth.yandex.ru/authorize?response_type=code&client_id=1d2e86281ff2444fa0d11e63b09afae0&redirect_uri=https://auth.testing.couply.ru/v1/login/oauth/yandex&force_confirm=yes';
  };

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

        <ButtonWithIcon
          onClick={onYandex}
          icon={<img src="yandex.png" width="26px" height="26px" alt="email" />}
          text="Войти c Яндекс ID"
          className={styles.yandex}
        />
      </div>
    </div>
  );
};

export default AuthPage;
