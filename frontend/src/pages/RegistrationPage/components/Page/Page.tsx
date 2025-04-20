import KeyboardBackspaceIcon from '@mui/icons-material/KeyboardBackspace';
import { useState } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';

import { useRegisterMutation } from '../../../../entities/auth';
import { CustomButton } from '../../../../shared/components/CustomButton';
import { CustomInput } from '../../../../shared/components/CustomInput';

import styles from './registrationPage.module.css';

interface LocationState {
  method: 'phone' | 'email';
}

export const RegistrationPage = () => {
  const navigate = useNavigate();
  const location = useLocation();
  const state = location.state as LocationState;

  // Default to phone if method is not specified
  const method = state?.method || 'phone';

  // State for form values
  const [contactValue, setContactValue] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');

  // State for form validation
  const [errors, setErrors] = useState({
    contactValue: '',
    password: '',
    confirmPassword: '',
  });

  // RTK Query mutation hook for registration
  const [register, { isLoading }] = useRegisterMutation();

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

    // Validate contact value (phone or email)
    if (!contactValue) {
      newErrors.contactValue =
        method === 'phone' ? 'Пожалуйста, введите номер телефона' : 'Пожалуйста, введите email';
      isValid = false;
    } else if (method === 'email' && !/\S+@\S+\.\S+/.test(contactValue)) {
      newErrors.contactValue = 'Пожалуйста, введите корректный email';
      isValid = false;
    } else if (method === 'phone' && !/^\+?[0-9]{10,15}$/.test(contactValue)) {
      newErrors.contactValue = 'Пожалуйста, введите корректный номер телефона';
      isValid = false;
    }

    // Validate password
    if (!password) {
      newErrors.password = 'Пожалуйста, введите пароль';
      isValid = false;
    } else if (password.length < 6) {
      newErrors.password = 'Пароль должен содержать не менее 6 символов';
      isValid = false;
    }

    // Validate confirm password
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
        // Prepare registration data based on method
        const registrationData = {
          password,
          ...(method === 'phone' ? { phone: contactValue } : { email: contactValue }),
        };

        // Call the register mutation
        const result = await register(registrationData).unwrap();

        // Store the token in localStorage
        localStorage.setItem('token', result.token);

        // Navigate to the enter info page
        navigate('/enterInfo');
      } catch (error) {
        // Handle registration errors
        console.error('Registration failed:', error);

        // You could set specific error messages based on the error response
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
          <label>{method === 'phone' ? 'Номер телефона' : 'Email'}</label>
          <CustomInput
            type={method === 'phone' ? 'tel' : 'email'}
            placeholder={method === 'phone' ? '+7 (999) 123-45-67' : 'example@email.com'}
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
        text={isLoading ? 'Регистрация...' : 'Зарегистрироваться'}
        disabled={isLoading}
        className={styles.submitButton}
      />
    </div>
  );
};

export default RegistrationPage;
