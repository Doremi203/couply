import KeyboardBackspaceIcon from '@mui/icons-material/KeyboardBackspace';
import { Dialog, Link } from '@mui/material';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

import { useLogin } from '../../../../entities/auth/hooks/useLogin';
import { CustomButton } from '../../../../shared/components/CustomButton';
import { CustomInput } from '../../../../shared/components/CustomInput';

import styles from './loginPage.module.css';

export const LoginPage = () => {
  const navigate = useNavigate();

  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [showRegistrationModal, setShowRegistrationModal] = useState(false);

  const [errors, setErrors] = useState({
    email: '',
    password: '',
    general: '',
  });

  const { loginUser } = useLogin();

  const goBack = () => {
    navigate('/auth');
  };

  const onRegister = () => {
    navigate('/registration');
  };

  const validateForm = () => {
    let isValid = true;
    const newErrors = {
      email: '',
      password: '',
      general: '',
    };

    if (!email) {
      newErrors.email = 'Пожалуйста, введите email';
      isValid = false;
    } else if (!/\S+@\S+\.\S+/.test(email)) {
      newErrors.email = 'Пожалуйста, введите корректный email';
      isValid = false;
    }

    if (!password) {
      newErrors.password = 'Пожалуйста, введите пароль';
      isValid = false;
    } else if (password.length < 6) {
      newErrors.password = 'Пароль должен содержать не менее 6 символов';
      isValid = false;
    }

    setErrors(newErrors);
    return isValid;
  };

  const handleCloseModal = () => {
    setShowRegistrationModal(false);
  };

  const handleSubmit = async () => {
    if (validateForm()) {
      try {
        setIsLoading(true);
        setErrors({ email: '', password: '', general: '' });

        const loginData = {
          password,
          email,
        };

        const result = await loginUser(loginData);

        if (result.data?.token) {
          navigate('/home');
        } else if (result.error) {
          console.log(result.error);
          // Handle specific error types
          if (result.error.includes('не найден')) {
            setShowRegistrationModal(true);
            setErrors({
              ...errors,
              email: 'Такого аккаунта не существует',
            });
          } else if (result.error.includes('пароль')) {
            setErrors({
              ...errors,
              password: result.error,
            });
          } else {
            setErrors({
              ...errors,
              general: result.error,
            });
          }
        }
      } catch {
        setErrors({
          ...errors,
          general: 'Произошла ошибка при входе. Пожалуйста, попробуйте позже.',
        });
      } finally {
        setIsLoading(false);
      }
    }
  };

  return (
    <div className={styles.container}>
      <div className={styles.backIcon} onClick={goBack}>
        <KeyboardBackspaceIcon />
      </div>

      <h2 className={styles.title}>Авторизация</h2>

      <div className={styles.form}>
        <div className={styles.inputGroup}>
          <label>Email</label>
          <CustomInput
            type="email"
            placeholder="example@email.com"
            value={email}
            onChange={e => setEmail(e.target.value)}
          />
          {errors.email && <span className={styles.errorText}>{errors.email}</span>}
        </div>

        <div className={styles.inputGroup}>
          <label>Пароль</label>
          <CustomInput
            type="password"
            placeholder="Введите пароль"
            value={password}
            onChange={e => setPassword(e.target.value)}
          />
          {errors.password && <span className={styles.errorText}>{errors.password}</span>}
        </div>

        {errors.general && <div className={styles.generalError}>{errors.general}</div>}
      </div>

      <CustomButton
        onClick={handleSubmit}
        text={isLoading ? 'Вход...' : 'Войти'}
        className={styles.submitButton}
        disabled={isLoading}
      />
      <Link onClick={onRegister} className={styles.regLink}>
        Зарегистрироваться
      </Link>

      <Dialog open={showRegistrationModal} onClose={handleCloseModal}>
        <div className={styles.notificationPrompt}>
          <h3>Аккаунт не найден</h3>
          <p>Такого аккаунта не существует. Хотите зарегистрироваться?</p>
          <div className={styles.promptButtons}>
            <CustomButton
              onClick={onRegister}
              text="Зарегистрироваться"
              className={styles.allowButton}
            />
            <CustomButton onClick={handleCloseModal} text="Отмена" className={styles.skipButton} />
          </div>
        </div>
      </Dialog>
    </div>
  );
};

export default LoginPage;
