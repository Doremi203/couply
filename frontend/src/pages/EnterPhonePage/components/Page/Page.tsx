import { KeyboardBackspace } from '@mui/icons-material';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

import CodeInput from '../../../../shared/components/CodeInput/CodeInput';
import { CustomButton } from '../../../../shared/components/CustomButton';
import { PhoneInput } from '../../../../shared/components/PhoneInput/PhoneInput';

import styles from './enterPhone.module.css';

export const EnterPhonePage = () => {
  const [currentStep, setCurrentStep] = useState(0);
  const navigate = useNavigate();
  // const dispatch = useDispatch();

  const [phone, setPhone] = useState('');
  const [code, setCode] = useState('');
  const [error, setError] = useState('');
  const [isLoading, setIsLoading] = useState(false);

  //   const [verifyPhone] = useVerifyPhoneMutation();
  //   const [confirmCode] = useConfirmCodeMutation();

  // const phoneInputRef = useRef<HTMLInputElement>(null);
  // const codeInputRef = useRef<HTMLInputElement>(null);

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
      // await verifyPhone({ phone }).unwrap();
      setCurrentStep(1);
      setError('');
    } catch (err) {
      setError('Ошибка отправки кода. Попробуйте снова.');
    } finally {
      setIsLoading(false);
    }
  };

  const handleConfirmCode = async () => {
    if (!/^\d{6}$/.test(code)) {
      setError('Код должен содержать 6 цифр');
      return;
    }

    setIsLoading(true);
    try {
      // const result = await confirmCode({ phone, code }).unwrap();
      //   dispatch(setVerifiedPhone(phone));
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
          Не получили код? <button onClick={handleSendCode}>Отправить снова</button>
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
