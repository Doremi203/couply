import { KeyboardBackspace } from '@mui/icons-material';
import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';

import { useSendCodeMutation, useConfirmPhoneMutation } from '../../../../entities/phone';
import CodeInput from '../../../../shared/components/CodeInput/CodeInput';
import { CustomButton } from '../../../../shared/components/CustomButton';
import { PhoneInput } from '../../../../shared/components/PhoneInput/PhoneInput';
import { getFormattedPhone } from '../../helpers/getFormattedPhone';

import styles from './enterPhone.module.css';

export const EnterPhonePage = () => {
  const [currentStep, setCurrentStep] = useState(0);
  const navigate = useNavigate();

  const [phone, setPhone] = useState('');
  const [code, setCode] = useState('');
  const [error, setError] = useState('');
  const [isLoading, setIsLoading] = useState(false);

  const [sendAgainIn, setSendAgainIn] = useState<number | null>(null);

  const [sendCode] = useSendCodeMutation();
  const [confirm] = useConfirmPhoneMutation();

  useEffect(() => {
    if (sendAgainIn === null || sendAgainIn <= 0) {
      setSendAgainIn(null);
      return;
    }

    const timer = setInterval(() => {
      setSendAgainIn(prev => (prev !== null ? prev - 1 : null));
    }, 1000);

    return () => clearInterval(timer);
  }, [sendAgainIn]);

  const handleBack = () => {
    if (currentStep > 0) {
      setCurrentStep(prev => prev - 1);
    } else {
      navigate(-1);
    }
  };

  const validatePhone = () => {
    // Очищаем телефон от всех символов, кроме цифр
    const cleanPhone = phone.replace(/\D/g, '');

    if (!phone || cleanPhone.length < 11) {
      setError('Пожалуйста, введите корректный номер телефона');
      return false;
    }

    return true;
  };

  const handleSendCode = async () => {
    if (!validatePhone()) {
      return;
    }

    setIsLoading(true);
    setError('');

    try {
      const formattedPhone = getFormattedPhone(phone);
      const timeout = await sendCode({ phone: formattedPhone }).unwrap();
      setSendAgainIn(timeout.sendAgainIn);
      setCurrentStep(1);
    } catch (err: any) {
      // Handle specific API errors
      if (err?.status === 429) {
        setError('Слишком много попыток. Пожалуйста, попробуйте позже.');
      } else if (err?.status === 400) {
        setError('Некорректный формат номера телефона');
      } else if (err?.data?.message) {
        setError(err.data.message);
      } else {
        setError('Ошибка отправки кода. Попробуйте снова.');
      }
    } finally {
      setIsLoading(false);
    }
  };

  const getResendButtonText = () => {
    if (sendAgainIn === null || sendAgainIn <= 0) {
      return 'Отправить снова';
    }
    return `Отправить снова (${sendAgainIn} сек)`;
  };

  const validateCode = () => {
    if (!code || code.length !== 6 || !/^\d+$/.test(code)) {
      setError('Пожалуйста, введите правильный код подтверждения (6 цифр)');
      return false;
    }

    return true;
  };

  const handleConfirmCode = async () => {
    if (!validateCode()) {
      return;
    }

    setIsLoading(true);
    setError('');

    try {
      const formattedPhone = getFormattedPhone(phone);
      await confirm({ phone: formattedPhone, code }).unwrap();
      navigate('/enterInfo');
    } catch (err: any) {
      // Handle specific API errors
      if (err?.status === 400) {
        setError('Неверный код подтверждения');
      } else if (err?.status === 404) {
        setError('Код подтверждения истек. Пожалуйста, запросите новый код.');
      } else if (err?.data?.message) {
        setError(err.data.message);
      } else {
        setError('Ошибка при подтверждении номера. Пожалуйста, попробуйте снова.');
      }
    } finally {
      setIsLoading(false);
    }
  };

  const sections = [
    <div key="phoneSection" className={styles.section}>
      <h2>Введите ваш телефон</h2>
      <div className={styles.phoneInput}>
        <PhoneInput value={phone} onChange={setPhone} />
      </div>
      {error && <div className={styles.error}>{error}</div>}
    </div>,

    <div key="codeSection" className={styles.section}>
      <h2>Введите код из SMS</h2>
      <div className={styles.code}>
        <CodeInput onCodeChange={setCode} />
        {error && <div className={styles.error}>{error}</div>}
        <div className={styles.resend}>
          Не получили код?{' '}
          <button
            onClick={handleSendCode}
            className={sendAgainIn === null || sendAgainIn <= 0 ? styles.activeResend : ''}
            disabled={sendAgainIn !== null && sendAgainIn > 0}
          >
            {getResendButtonText()}
          </button>
        </div>
      </div>
    </div>,
  ];

  return (
    <div className={styles.container}>
      <div className={styles.backIcon} onClick={handleBack}>
        <KeyboardBackspace />
      </div>

      {sections[currentStep]}

      <CustomButton
        onClick={currentStep === 0 ? handleSendCode : handleConfirmCode}
        text={
          isLoading
            ? currentStep === 0
              ? 'Отправка...'
              : 'Проверка...'
            : currentStep === 0
              ? 'Получить код'
              : 'Подтвердить'
        }
        disabled={isLoading || (currentStep === 0 ? !phone : code.length < 6)}
        className={styles.nextButton}
      />
    </div>
  );
};

export default EnterPhonePage;
