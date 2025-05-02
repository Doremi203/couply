import EmailIcon from '@mui/icons-material/Email';
import { Stack } from '@mui/material';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

import { useLoginMutation } from '../../../../entities/auth';
import { ButtonWithIcon } from '../../../../shared/components/ButtonWithIcon';

import styles from './authPage.module.css';

export const AuthPage = () => {
  const navigate = useNavigate();
  const [login, { isLoading }] = useLoginMutation();
  const [error, setError] = useState('');

  const handleSocialLogin = () => {
    // In a real app, this would integrate with social login providers
    // For now, just navigate to the enter info page
    navigate('/enterInfo');
  };

  const handlePhoneLogin = () => {
    navigate('/registration', { state: { method: 'phone' } });
  };

  const handleEmailLogin = () => {
    navigate('/login', { state: { method: 'email' } });
  };

  // This function would be used if we implemented a login form on this page
  // @ts-ignore
  const _handleLogin = async (credentials: LoginParams) => {
    try {
      setError('');
      const result = await login(credentials).unwrap();

      // Store the token in localStorage
      localStorage.setItem('token', result.token);

      // Navigate to the home page after successful login
      navigate('/home');
    } catch (error) {
      console.error('Login failed:', error);
      setError('Неверные учетные данные. Пожалуйста, попробуйте снова.');
    }
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

        {error && <div className={styles.errorMessage}>{error}</div>}

        <Stack
          direction="column"
          spacing={2}
          sx={{
            justifyContent: 'center',
            alignItems: 'center',
          }}
        >
          <ButtonWithIcon
            onClick={handleSocialLogin}
            icon={<img src="image.png" width="20px" height="20px" alt="Google" />}
            text="login with google"
            disabled={isLoading}
          />

          <ButtonWithIcon
            onClick={handleSocialLogin}
            icon={<img src="vk.png" width="20px" height="20px" alt="VK" />}
            text="login with vk"
            disabled={isLoading}
          />

          <ButtonWithIcon
            onClick={handlePhoneLogin}
            icon={<img src="phone.png" width="20px" height="20px" alt="Phone" />}
            text="login with phone"
            disabled={isLoading}
          />

          <ButtonWithIcon
            onClick={handleEmailLogin}
            icon={<EmailIcon style={{ width: '20px', height: '20px' }} />}
            text="login with email"
            disabled={isLoading}
          />
        </Stack>
      </Stack>
    </div>
  );
};

export default AuthPage;
