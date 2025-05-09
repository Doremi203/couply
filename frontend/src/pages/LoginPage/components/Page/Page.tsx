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

  const [showRegistrationModal, setShowRegistrationModal] = useState(false);

  const [errors, setErrors] = useState({
    email: '',
    password: '',
    confirmPassword: '',
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

  const handleCloseModal = () => {
    setShowRegistrationModal(false);
  };

  const handleNavigateToRegistration = () => {
    navigate('/registration');
    setShowRegistrationModal(false);
  };

  const handleSubmit = async () => {
    if (validateForm()) {
      try {
        const loginData = {
          password,
          email,
        };

        const result = await loginUser(loginData).unwrap();

        // Store the token and its expiration time using TokenService
        // setToken(result.token, result.expiresIn);

        if (result.token) {
          navigate('/home');
        } else {
          navigate('/registration');
        }
      } catch (error) {
        setShowRegistrationModal(true);
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

      <CustomButton onClick={handleSubmit} text="Войти" className={styles.submitButton} />
      <Link onClick={onRegister} className={styles.regLink}>
        Зарегистрироваться
      </Link>

      {showRegistrationModal && (
        <Dialog
          open={showRegistrationModal}
          onClose={handleCloseModal}
          aria-labelledby="alert-dialog-title"
          aria-describedby="alert-dialog-description"
        >
          <div className={styles.notificationPrompt}>
            <h3>Аккаунт не найден</h3>
            <p>Аккаунт с указанным email не существует. Хотите зарегистрироваться?</p>
            <div className={styles.promptButtons}>
              <CustomButton
                onClick={onRegister}
                text="Зарегистрироваться"
                className={styles.allowButton}
              />
            </div>
          </div>
        </Dialog>
      )}
    </div>
  );
};

export default LoginPage;
