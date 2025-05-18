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

  const handleSendCode = async () => {
    setIsLoading(true);
    try {
      setCurrentStep(1);
      const formattedPhone = getFormattedPhone(phone);
      const timeout = await sendCode({ phone: formattedPhone }).unwrap();
      setSendAgainIn(timeout.sendAgainIn);
      setError('');
    } catch {
      setError('Ошибка отправки кода. Попробуйте снова.');
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

  const handleConfirmCode = async () => {
    const formattedPhone = getFormattedPhone(phone);

    setIsLoading(true);
    try {
      await confirm({ phone: formattedPhone, code }).unwrap();
      navigate('/enterInfo');
    } catch (err) {
      setError('Неверный код подтверждения');
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
        text={isLoading ? 'Отправка...' : currentStep === 0 ? 'Получить код' : 'Подтвердить'}
        disabled={isLoading || (currentStep === 0 ? !phone : code.length < 6)}
        className={styles.nextButton}
      />
    </div>
  );
};

export default EnterPhonePage;
