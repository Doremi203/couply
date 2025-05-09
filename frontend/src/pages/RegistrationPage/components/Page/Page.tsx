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

  const [errors, setErrors] = useState({
    contactValue: '',
    password: '',
    confirmPassword: '',
  });

  // const [register, { isLoading }] = useRegisterMutation();

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
        const registrationData = {
          password,
          ...{ email: contactValue },
        };

        await registerAndLogin(registrationData);

        // Store the token in localStorage
        // localStorage.setItem('token', result.token);
        navigate('/enterInfo');
      } catch (error) {
        console.error('Registration failed:', error);

        setErrors({
          ...errors,
          contactValue: 'Ошибка регистрации. Пожалуйста, попробуйте снова.',
        });
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
      </div>

      <CustomButton
        onClick={handleSubmit}
        text="Зарегистрироваться"
        // disabled={isLoading}
        className={styles.submitButton}
      />
    </div>
  );
};

export default RegistrationPage;
