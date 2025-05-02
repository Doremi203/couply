import KeyboardBackspaceIcon from '@mui/icons-material/KeyboardBackspace';
import PhotoCameraIcon from '@mui/icons-material/PhotoCamera';
import { useState, useRef } from 'react';
import { useDispatch } from 'react-redux';
import { useNavigate } from 'react-router-dom';

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

// function generateS3Url(fileName, bucketName, regionName) {
//   // Формируем базовый URL
//   return `https://${bucketName}.s3.${regionName}.amazonaws.com/${fileName}`;
// }

export const EnterInfoPage = () => {
  const [currentStep, setCurrentStep] = useState(0);
  const navigate = useNavigate();
  const dispatch = useDispatch();

  // Initialize the createUser mutation
  const [createUser, { isLoading }] = useCreateUserMutation();

  // State for form values
  const [name, setName] = useState('');
  const [birthDate, setBirthDate] = useState('');
  const [userGender, setUserGender] = useState('');
  const [profilePhoto, setProfilePhoto] = useState<string | null>(null);
  const fileInputRef = useRef<HTMLInputElement>(null);

  // State for notification permission
  const [showNotificationPrompt, setShowNotificationPrompt] = useState(false);
  const [notificationPermissionRequested, setNotificationPermissionRequested] = useState(false);

  const nextStep = async () => {
    if (currentStep === sections.length - 1) {
      // If we're on the last step, submit user data
      try {
        // Create user data object with default values according to UserRequest interface
        const userData = {
          name,
          age: 20,
          gender: userGender,
          birthDate, // This is not in the UserRequest interface but might be used elsewhere
          photos: [{ url: profilePhoto }], // Set to null as expected by the API
        };

        // Store the profile photo URL in localStorage for later use
        if (profilePhoto) {
          localStorage.setItem('profilePhotoUrl', profilePhoto);
        }

        // Send data to the API
        const response = await createUser(userData).unwrap();
        console.log('Response:', response);
        console.log('Response type:', typeof response);
        console.log('Response keys:', Object.keys(response));

        // Store the user ID if available in the response
        if (response && response.user && response.user.id) {
          console.log('User ID:', response.user.id);
          // Save user ID to Redux store
          dispatch(setUserId(response.user.id));
          // Also save to localStorage for persistence across page reloads
          localStorage.setItem('userId', response.user.id);
        }

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
      // Otherwise, go to the next step
      setCurrentStep(prevStep => prevStep + 1);
    }
  };

  const prevStep = () => {
    if (currentStep > 0) {
      setCurrentStep(prevStep => prevStep - 1);
    } else {
      navigate('/');
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

      // Create a URL for the file
      const fileUrl = URL.createObjectURL(file);
      setProfilePhoto(fileUrl);

      event.target.value = '';
    }
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
            // Use the userId from localStorage, or fallback to a default
            const userIdFromStorage = localStorage.getItem('userId');
            const userIdToUse = userIdFromStorage || 'user123';
            await sendSubscriptionToServer(subscription, userIdToUse);
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
