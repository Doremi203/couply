import KeyboardBackspaceIcon from '@mui/icons-material/KeyboardBackspace';
import { Link } from '@mui/material';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

import { useLoginMutation } from '../../../../entities/auth';
import { CustomButton } from '../../../../shared/components/CustomButton';
import { CustomInput } from '../../../../shared/components/CustomInput';
// import { setToken } from '../../../../shared/lib/services/TokenService';

import styles from './loginPage.module.css';

export const LoginPage = () => {
  const navigate = useNavigate();

  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const [errors, setErrors] = useState({
    email: '',
    password: '',
    confirmPassword: '',
  });

  const [login, { isLoading }] = useLoginMutation();

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
      confirmPassword: '',
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

  const handleSubmit = async () => {
    if (validateForm()) {
      try {
        const registrationData = {
          password,
          ...{ email: email },
        };

        const result = await login(registrationData).unwrap();

        // Store the token and its expiration time using TokenService
        // setToken(result.token, result.expiresIn);

        if (result.token) {
          navigate('/home');
        } else {
          navigate('/registration');
        }
      } catch (error) {
        //TODO
        // navigate('/registration');
        console.error('Registration failed:', error);
        setErrors({
          ...errors,
          email: 'Такого аккаунта не существует',
        });
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
      </div>

      <CustomButton
        onClick={handleSubmit}
        text="Войти"
        disabled={isLoading}
        className={styles.submitButton}
      />
      <Link onClick={onRegister} className={styles.regLink}>
        Зарегистрироваться
      </Link>
    </div>
  );
};

export default LoginPage;
