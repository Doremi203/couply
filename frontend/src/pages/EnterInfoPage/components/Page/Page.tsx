import KeyboardBackspaceIcon from '@mui/icons-material/KeyboardBackspace';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

import { CustomButton } from '../../../../shared/components/CustomButton';
import { CustomInput } from '../../../../shared/components/CustomInput';
import { ToggleButtons } from '../../../../shared/components/ToggleButtons';
import {
  isPushNotificationSupported,
  askUserPermission,
  registerServiceWorker,
  createPushSubscription,
  sendSubscriptionToServer,
} from '../../../../shared/lib/services/PushNotificationService';

import styles from './enterInfo.module.css';

export const EnterInfoPage = () => {
  const [currentStep, setCurrentStep] = useState(0);
  const navigate = useNavigate();

  // State for form values
  const [name, setName] = useState('');
  const [birthDate, setBirthDate] = useState('');
  const [userGender, setUserGender] = useState('');
  const [preferredGender, setPreferredGender] = useState('');

  // State for notification permission
  const [showNotificationPrompt, setShowNotificationPrompt] = useState(false);
  const [notificationPermissionRequested, setNotificationPermissionRequested] = useState(false);

  const nextStep = () => {
    if (currentStep === sections.length - 1) {
      // If we're on the last step, check if we should show notification prompt
      if (
        isPushNotificationSupported() &&
        !notificationPermissionRequested &&
        Notification.permission !== 'granted' &&
        Notification.permission !== 'denied'
      ) {
        setShowNotificationPrompt(true);
      } else {
        // If notifications are not supported or permission already requested, navigate to home
        navigate('/home');
      }
    } else {
      // Otherwise, go to the next step
      setCurrentStep(prevStep => prevStep + 1);
    }
  };

  const prevStep = () => {
    if (currentStep > 0) {
      setCurrentStep(prevStep => prevStep - 1);
    }
  };

  const handleUserGenderSelect = (value: string) => {
    setUserGender(value);
  };

  const handlePreferredGenderSelect = (value: string) => {
    setPreferredGender(value);
  };

  // Handle notification permission request
  const handleRequestPermission = async () => {
    setNotificationPermissionRequested(true);
    setShowNotificationPrompt(false);

    try {
      const permission = await askUserPermission();

      if (permission === 'granted') {
        const registration = await registerServiceWorker();

        if (registration) {
          const subscription = await createPushSubscription(registration);

          if (subscription) {
            // In a real app, you would get the userId from authentication
            await sendSubscriptionToServer(subscription, 'user123');
          }
        }
      }

      // Navigate to home page regardless of permission result
      navigate('/home');
    } catch (error) {
      console.error('Error requesting notification permission:', error);
      navigate('/home');
    }
  };

  const handleSkipPermission = () => {
    setNotificationPermissionRequested(true);
    setShowNotificationPrompt(false);
    navigate('/home');
  };

  // Check if the current step's form is valid
  const isCurrentStepValid = () => {
    switch (currentStep) {
      case 0:
        return name.trim() !== '';
      case 1:
        return birthDate !== '';
      case 2:
        return userGender !== '' && preferredGender !== '';
      default:
        return false;
    }
  };

  const sections = [
    <div key="nameSection">
      <h2>Как вас зовут?</h2>
      <CustomInput
        placeholder="Введите имя"
        type="text"
        value={name}
        onChange={e => setName(e.target.value)}
      />
    </div>,
    <div key="birthDateSection">
      <h2>Дата рождения</h2>
      <CustomInput
        placeholder="Выберите дату рождения"
        type="date"
        value={birthDate}
        onChange={e => setBirthDate(e.target.value)}
      />
    </div>,
    <div key="datingSettingsSection">
      <h2>Настройки дейтинга</h2>
      <div>
        <label>Ваш пол:</label>
        <ToggleButtons
          options={[
            { label: 'Женский', value: 'female' },
            { label: 'Мужской', value: 'male' },
          ]}
          onSelect={handleUserGenderSelect}
          value={userGender}
        />
      </div>
      <div>
        <label>Кого вам показывать:</label>
        <ToggleButtons
          options={[
            { label: 'Женщин', value: 'female' },
            { label: 'Мужчин', value: 'male' },
            { label: 'Всех', value: 'other' },
          ]}
          onSelect={handlePreferredGenderSelect}
          value={preferredGender}
        />
      </div>
    </div>,
  ];

  return (
    <div className={styles.container}>
      <div className={styles.backIcon} onClick={prevStep}>
        <KeyboardBackspaceIcon />
      </div>
      {sections[currentStep]}

      {showNotificationPrompt ? (
        <div className={styles.notificationPrompt}>
          <h3>Разрешить уведомления?</h3>
          <p>Получайте уведомления о новых совпадениях и важных событиях в приложении</p>
          <div className={styles.promptButtons}>
            <CustomButton
              onClick={handleRequestPermission}
              text="Разрешить"
              className={styles.allowButton}
            />
            <CustomButton
              onClick={handleSkipPermission}
              text="Не сейчас"
              className={`${styles.skipButton} ${styles.outlinedButton}`}
            />
          </div>
        </div>
      ) : (
        <CustomButton
          onClick={nextStep}
          text="Дальше"
          disabled={!isCurrentStepValid()}
          className={styles.nextButton}
        />
      )}
    </div>
  );
};

export default EnterInfoPage;
