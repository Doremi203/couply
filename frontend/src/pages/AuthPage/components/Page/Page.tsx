import { useNavigate } from 'react-router-dom';

import { ButtonWithIcon } from '../../../../shared/components/ButtonWithIcon';

import styles from './authPage.module.css';

export const AuthPage = () => {
  const navigate = useNavigate();

  const onLogin = () => {
    navigate('/login');
  };

  return (
    <div className={styles.page}>
      <img src="logo.png" width="200px" height="150px" alt="Logo" className={styles.logo} />
      <span className={styles.text}>Найди того, кто будет похож на тебя, как капля воды.</span>

      <div className={styles.buttons}>
        <ButtonWithIcon
          onClick={onLogin}
          icon={<img src="vk.png" width="20px" height="20px" alt="VK" />}
          text="Войти с VK ID"
          iconClassName={styles.vkIcon}
        />

        <ButtonWithIcon
          onClick={onLogin}
          icon={<img src="email2.png" width="26px" height="26px" alt="email" />}
          text="Войти по почте"
        />
      </div>
    </div>
  );
};

export default AuthPage;
