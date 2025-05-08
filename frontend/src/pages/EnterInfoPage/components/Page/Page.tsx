import KeyboardBackspaceIcon from '@mui/icons-material/KeyboardBackspace';
import PhotoCameraIcon from '@mui/icons-material/PhotoCamera';
import { useState, useRef } from 'react';
import { useDispatch } from 'react-redux';
import { useNavigate } from 'react-router-dom';

import { Gender } from '../../../../entities/user';
import { useCreateUserMutation } from '../../../../entities/user/api/userApi';
import { setUserId } from '../../../../entities/user/model/userSlice';
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
  const dispatch = useDispatch();

  const [createUser, { isLoading }] = useCreateUserMutation();

  const [name, setName] = useState('');
  const [birthDate, setBirthDate] = useState('');
  const [userGender, setUserGender] = useState('');
  const [profilePhoto, setProfilePhoto] = useState<string | null>(null);
  const fileInputRef = useRef<HTMLInputElement>(null);

  const [showNotificationPrompt, setShowNotificationPrompt] = useState(false);
  const [notificationPermissionRequested, setNotificationPermissionRequested] = useState(false);

  const nextStep = async () => {
    if (currentStep === sections.length - 1) {
      try {
        //TODO
        const userData = {
          name,
          age: 20,
          gender: Gender.male,
          birthDate,
          photos: [{ url: profilePhoto }],
        };

        // TODO
        if (profilePhoto) {
          localStorage.setItem('profilePhotoUrl', profilePhoto);
        }

        //ВЕРНУТЬ
        //const response = await createUser(userData).unwrap();

        // надо ли сохранять в локал TODO
        // if (response && response.user && response.user.id) {
        //   dispatch(setUserId(response.user.id));
        //   localStorage.setItem('userId', response.user.id);
        // }

        // After creating user, check if we should show notification prompt
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
      } catch (error) {
        console.error('Error creating user:', error);
        // Still proceed to notification or home page even if there's an error
        if (
          isPushNotificationSupported() &&
          !notificationPermissionRequested &&
          Notification.permission !== 'granted' &&
          Notification.permission !== 'denied'
        ) {
          setShowNotificationPrompt(true);
        } else {
          navigate('/home');
        }
      }
    } else {
      setCurrentStep(prevStep => prevStep + 1);
    }
  };

  const prevStep = () => {
    if (currentStep > 0) {
      setCurrentStep(prevStep => prevStep - 1);
    } else {
      navigate('/auth');
    }
  };

  const handleUserGenderSelect = (value: string) => {
    setUserGender(value);
  };

  const handleCameraClick = () => {
    if (fileInputRef.current) {
      fileInputRef.current.click();
    }
  };

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const files = event.target.files;
    if (files && files.length > 0) {
      const file = files[0];

      const fileUrl = URL.createObjectURL(file);
      setProfilePhoto(fileUrl);

      event.target.value = '';
    }
  };

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
            const userIdFromStorage = localStorage.getItem('userId');
            const userIdToUse = userIdFromStorage || 'user123';
            await sendSubscriptionToServer(subscription, userIdToUse);
          }
        }
      }

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

  const isCurrentStepValid = () => {
    switch (currentStep) {
      case 0:
        return name.trim() !== '';
      case 1:
        return birthDate !== '';
      case 2:
        return userGender !== '' && profilePhoto !== null;
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
      <div>
        <h2>Ваш пол:</h2>
        <ToggleButtons
          options={[
            { label: 'Женский', value: 'female' },
            { label: 'Мужской', value: 'male' },
          ]}
          onSelect={handleUserGenderSelect}
          value={userGender}
        />
      </div>
    </div>,
    <div key="datingSettingsSection">
      <h2>Загрузите ваше фото</h2>
      <div>
        {/* <label>Загрузите ваше фото:</label> */}
        <div className={styles.photoUploadContainer}>
          {profilePhoto ? (
            <div className={styles.photoPreview}>
              <img src={profilePhoto} alt="Profile" className={styles.profilePic} />
              <div className={styles.editIcon} onClick={handleCameraClick}>
                <PhotoCameraIcon />
              </div>
            </div>
          ) : (
            <div className={styles.photoPlaceholder} onClick={handleCameraClick}>
              <PhotoCameraIcon />
              <span>Нажмите, чтобы выбрать фото</span>
            </div>
          )}
        </div>
      </div>
    </div>,
  ];

  return (
    <div className={styles.container}>
      <input
        type="file"
        ref={fileInputRef}
        style={{ display: 'none' }}
        accept="image/*"
        onChange={handleFileChange}
      />
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
          text={isLoading ? 'Загрузка...' : 'Дальше'}
          disabled={!isCurrentStepValid() || isLoading}
          className={styles.nextButton}
        />
      )}
    </div>
  );
};

export default EnterInfoPage;
