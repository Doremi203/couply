import EmailIcon from '@mui/icons-material/Email';
import { Stack } from '@mui/material';
import { useNavigate } from 'react-router-dom';

import { ButtonWithIcon } from '../../../../shared/components/ButtonWithIcon';

import styles from './authPage.module.css';

export const AuthPage = () => {
  const navigate = useNavigate();

  const onLogin = () => {
    navigate('/login', { state: { method: 'email' } });
  };

  return (
    <div className={styles.page}>
      <Stack
        direction="column"
        spacing={2}
        sx={{
          justifyContent: 'center',
          alignItems: 'center',
        }}
      >
        <img src="logo.png" width="200px" height="150px" alt="Logo" />

        <span className={styles.text}>Найди того, кто будет похож на тебя, как капля воды.</span>

        <Stack
          direction="column"
          spacing={2}
          sx={{
            justifyContent: 'center',
            alignItems: 'center',
          }}
        >
          {/* <ButtonWithIcon
            onClick={onLogin}
            icon={<img src="image.png" width="20px" height="20px" alt="Google" />}
            text="login with google"
            disabled={isLoading}
          /> */}

          <ButtonWithIcon
            onClick={onLogin}
            icon={<img src="vk.png" width="20px" height="20px" alt="VK" />}
            text="login with vk"
          />

          <ButtonWithIcon
            onClick={onLogin}
            icon={<EmailIcon style={{ width: '20px', height: '20px' }} />}
            text="login with email"
          />
        </Stack>
      </Stack>
    </div>
  );
};

export default AuthPage;
