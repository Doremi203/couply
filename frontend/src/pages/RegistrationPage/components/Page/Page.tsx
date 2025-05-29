import KeyboardBackspaceIcon from '@mui/icons-material/KeyboardBackspace';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

import { useRegister } from '../../../../entities/auth';
import { CustomButton } from '../../../../shared/components/CustomButton';
import { CustomInput } from '../../../../shared/components/CustomInput';

import styles from './registrationPage.module.css';

export const RegistrationPage = () => {
  const navigate = useNavigate();

  const [contactValue, setContactValue] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [isLoading, setIsLoading] = useState(false);

  const [errors, setErrors] = useState({
    contactValue: '',
    password: '',
    confirmPassword: '',
    general: '',
  });

  const { registerAndLogin } = useRegister();

  const goBack = () => {
    navigate('/auth');
  };

  const validateForm = () => {
    let isValid = true;
    const newErrors = {
      contactValue: '',
      password: '',
      confirmPassword: '',
      general: '',
    };

    if (!contactValue) {
      newErrors.contactValue = 'Пожалуйста, введите email';
      isValid = false;
    } else if (!/\S+@\S+\.\S+/.test(contactValue)) {
      newErrors.contactValue = 'Пожалуйста, введите корректный email';
      isValid = false;
    }

    if (!password) {
      newErrors.password = 'Пожалуйста, введите пароль';
      isValid = false;
    } else if (password.length < 6) {
      newErrors.password = 'Пароль должен содержать не менее 6 символов';
      isValid = false;
    }

    if (!confirmPassword) {
      newErrors.confirmPassword = 'Пожалуйста, подтвердите пароль';
      isValid = false;
    } else if (confirmPassword !== password) {
      newErrors.confirmPassword = 'Пароли не совпадают';
      isValid = false;
    }

    setErrors(newErrors);
    return isValid;
  };

  const handleSubmit = async () => {
    if (validateForm()) {
      try {
        setIsLoading(true);
        setErrors({ contactValue: '', password: '', confirmPassword: '', general: '' });

        const registrationData = {
          password,
          ...{ email: contactValue },
        };

        const result = await registerAndLogin(registrationData);

        if (result.data) {
          navigate('/enterPhone');
        } else if (result.error) {
          // Handle specific error types
          if (result.error.includes('email') || result.error.includes('существует')) {
            setErrors({
              ...errors,
              contactValue: result.error,
            });
          } else if (result.error.includes('пароль') || result.error.includes('password')) {
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
          general: 'Произошла ошибка при регистрации. Пожалуйста, попробуйте позже.',
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

      <h2 className={styles.title}>Регистрация</h2>

      <div className={styles.form}>
        <div className={styles.inputGroup}>
          <label>Email</label>
          <CustomInput
            type="email"
            placeholder="example@email.com"
            value={contactValue}
            onChange={e => setContactValue(e.target.value)}
          />
          {errors.contactValue && <span className={styles.errorText}>{errors.contactValue}</span>}
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

        <div className={styles.inputGroup}>
          <label>Подтверждение пароля</label>
          <CustomInput
            type="password"
            placeholder="Повторите пароль"
            value={confirmPassword}
            onChange={e => setConfirmPassword(e.target.value)}
          />
          {errors.confirmPassword && (
            <span className={styles.errorText}>{errors.confirmPassword}</span>
          )}
        </div>

        {errors.general && <div className={styles.generalError}>{errors.general}</div>}
      </div>

      <div data-testid="submit-button">
        <CustomButton
          onClick={handleSubmit}
          text={isLoading ? 'Регистрация...' : 'Зарегистрироваться'}
          className={styles.submitButton}
          disabled={isLoading}
        />
      </div>
    </div>
  );
};

export default RegistrationPage;
