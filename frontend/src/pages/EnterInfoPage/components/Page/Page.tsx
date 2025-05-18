import KeyboardBackspaceIcon from '@mui/icons-material/KeyboardBackspace';
import PhotoCameraIcon from '@mui/icons-material/PhotoCamera';
import { useState, useRef } from 'react';
import { useNavigate } from 'react-router-dom';

import { useUploadFileToS3Mutation } from '../../../../entities/photo/api/photoApi';
import { Gender, Goal } from '../../../../entities/user/api/constants';
import {
  useConfirmPhotoMutation,
  useCreateUserMutation,
} from '../../../../entities/user/api/userApi';
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
import getAge from '../../helpers/getAge';
import { GeoLocationRequest } from '../GeoLocationRequest';
import { FixedPhotoGallery } from '../PhotoGallery/PhotoGallery';

import styles from './enterInfo.module.css';

export const EnterInfoPage = () => {
  const [currentStep, setCurrentStep] = useState(0);
  const navigate = useNavigate();

  const [createUser, { isLoading }] = useCreateUserMutation();
  const [confirmPhoto] = useConfirmPhotoMutation();

  const [name, setName] = useState('');
  const [telegram, setTelegram] = useState('');
  const [birthDate, setBirthDate] = useState('');
  const [height, setHeight] = useState('');
  const [userGender, setUserGender] = useState('');
  const [userGoal, setUserGoal] = useState('');
  const [profilePhoto, setProfilePhoto] = useState<string | null>(null);
  const fileInputRef = useRef<HTMLInputElement>(null);

  const [showNotificationPrompt, setShowNotificationPrompt] = useState(false);
  const [notificationPermissionRequested, setNotificationPermissionRequested] = useState(false);

  const [coords, setCoords] = useState<{ lat: number; lng: number } | null>(null);
  const [useManualLocation, setUseManualLocation] = useState(false);
  const [manualLocation, setManualLocation] = useState('');

  // Add proper interface for the photo object
  interface PhotoItem {
    file: File;
    url: string;
  }

  // Update the userPhotos state with the proper type
  const [userPhotos, setUserPhotos] = useState<PhotoItem[]>([]);

  const handleLocationReceived = (coordinates: { lat: number; lng: number }) => {
    setCoords(coordinates);
  };

  const [uploadFile] = useUploadFileToS3Mutation();

  const goalOptions = [
    { label: 'Отношения', value: 'relationship', enum: Goal.relationship },
    { label: 'Дружба', value: 'friendship', enum: Goal.friendship },
    { label: 'Общение', value: 'justChatting', enum: Goal.justChatting },
    { label: 'Свидания', value: 'dating', enum: Goal.dating },
  ];

  const nextStep = async () => {
    if (currentStep === sections.length - 1) {
      try {
        const calculatedAge = getAge(birthDate);
        const ageValue = typeof calculatedAge === 'number' ? calculatedAge : 0;

        let genderEnum: Gender;
        switch (userGender) {
          case 'male':
            genderEnum = Gender.male;
            break;
          case 'female':
            genderEnum = Gender.female;
            break;
          default:
            genderEnum = Gender.unspecified;
        }

        // Find the selected goal enum directly from our mapping
        const selectedGoalOption = goalOptions.find(option => option.value === userGoal);
        const goalEnum = selectedGoalOption ? selectedGoalOption.enum : Goal.unspecified;

        const locationString = coords
          ? `${coords.lat.toFixed(6)},${coords.lng.toFixed(6)}`
          : manualLocation;

        if (profilePhoto) {
          localStorage.setItem('profilePhotoUrl', profilePhoto);
        }

        const photoUploadRequests = userPhotos.map((photo, index) => ({
          orderNumber: index,
          mimeType: photo.file.type,
        }));

        const orderNumbers = userPhotos.map((_photo, index) => index);

        const userData = {
          name,
          telegram,
          age: ageValue,
          gender: genderEnum,
          height: Number(height),
          goal: goalEnum,
          ...(locationString ? { location: locationString } : {}),
          photoUploadRequests,
          latitude: coords?.lat,
          longitude: coords?.lng,
        };

        const response = await createUser(userData).unwrap();

        // @ts-ignore
        if (response.photoUploadResponses) {
          await Promise.all(
            //@ts-ignore
            response.photoUploadResponses.map(async resp => {
              const photo = userPhotos[resp.orderNumber];
              if (!photo?.file) return;

              try {
                await uploadFile({
                  url: resp.uploadUrl,
                  file: photo.file,
                }).unwrap();
              } catch (error) {
                console.error(`Ошибка загрузки файла ${photo.file.name}:`, error);
              }
            }),
          );
        }

        try {
          await confirmPhoto({ orderNumbers }).unwrap();
        } catch (error) {
          console.error('Photo confirmation failed:', error);
          throw error;
        }

        if (
          isPushNotificationSupported() &&
          !notificationPermissionRequested
          // Notification.permission !== 'granted' &&
          // Notification.permission !== 'denied'
        ) {
          setShowNotificationPrompt(true);
        } else {
          navigate('/home');
        }
      } catch (error) {
        console.error('Error creating user:', error);
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

  const handleUserGoalSelect = (value: string) => {
    setUserGoal(value);
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

      setUserPhotos(prevPhotos => {
        const newPhotos = [...prevPhotos];
        if (newPhotos.length > 0) {
          newPhotos[0] = { file, url: fileUrl };
        } else {
          newPhotos.push({ file, url: fileUrl });
        }
        return newPhotos;
      });

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
        return name.trim() !== '' && telegram.trim() !== '';
      case 1: {
        const calculatedAge = getAge(birthDate);
        return (
          birthDate !== '' &&
          typeof calculatedAge === 'number' &&
          calculatedAge >= 18 &&
          userGender !== '' &&
          height.trim() !== ''
        );
      }
      case 2:
        return userGoal !== '';
      case 3:
        return userPhotos.length > 0;
      case 4:
        return coords !== null || (useManualLocation && manualLocation.trim() !== '');
      default:
        return false;
    }
  };

  const handleAddPhoto = () => {
    if (userPhotos.length >= 6) return;

    const input = document.createElement('input');
    input.type = 'file';
    input.accept = 'image/*';
    input.onchange = e => {
      const target = e.target as HTMLInputElement;
      const files = target.files;

      if (files && files.length > 0) {
        const file = files[0];
        const fileUrl = URL.createObjectURL(file);

        setUserPhotos(prevPhotos => {
          // If this is the first photo, also set it as the profile photo
          if (prevPhotos.length === 0) {
            setProfilePhoto(fileUrl);
          }
          return [...prevPhotos, { file, url: fileUrl }];
        });
      }
    };
    input.click();
  };

  const handleRemovePhoto = (index: number) => {
    setUserPhotos(prevPhotos => {
      const newPhotos = prevPhotos.filter((_, i) => i !== index);

      // Update profile photo if we're removing the first photo
      if (index === 0) {
        // If there are still photos left, set the new first photo as profile
        if (newPhotos.length > 0) {
          setProfilePhoto(newPhotos[0].url);
        } else {
          // Otherwise clear the profile photo
          setProfilePhoto(null);
        }
      }

      return newPhotos;
    });
  };

  const sections = [
    <div key="nameSection">
      <h2>Как вас зовут?</h2>
      <CustomInput
        placeholder="Введите имя"
        type="text"
        value={name}
        onChange={e => setName(e.target.value)}
        className={styles.input}
      />
      <h2 className={styles.genderLabel}>Ваш Telegram</h2>
      <CustomInput
        placeholder="@username"
        type="text"
        value={telegram}
        onChange={e => setTelegram(e.target.value)}
        className={styles.input}
      />
    </div>,
    <div key="birthDateSection">
      <h2>Дата рождения</h2>
      <CustomInput
        placeholder="Выберите дату рождения"
        type="date"
        value={birthDate}
        onChange={e => setBirthDate(e.target.value)}
        className={styles.input}
      />

      {birthDate &&
        (() => {
          const age = getAge(birthDate);
          return typeof age === 'number' && age < 18;
        })() && <div className={styles.error}>Для регистрации необходимо быть старше 18 лет</div>}
      <div>
        <h2 className={styles.genderLabel}>Ваш пол:</h2>
        <ToggleButtons
          options={[
            { label: 'Женский', value: 'female' },
            { label: 'Мужской', value: 'male' },
          ]}
          onSelect={handleUserGenderSelect}
          value={userGender}
          className={styles.toggleButtons}
        />
      </div>
      <h2 className={styles.genderLabel}>Ваш рост</h2>
      <CustomInput
        placeholder="180 см"
        type="number"
        value={height}
        onChange={e => setHeight(e.target.value)}
        className={styles.input}
      />
    </div>,
    <div key="goalSection">
      <h2>Какова ваша цель?</h2>
      <ToggleButtons
        options={goalOptions.map(({ label, value }) => ({ label, value }))}
        onSelect={handleUserGoalSelect}
        value={userGoal}
        className={styles.toggleButtons}
      />
    </div>,
    <div key="datingSettingsSection">
      <h2>Загрузите ваше фото</h2>
      <div>
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
        <FixedPhotoGallery
          photos={userPhotos.map(photo => photo.url)}
          onPhotoRemove={handleRemovePhoto}
          onAddPhotoClick={handleAddPhoto}
          title="Мои фото"
        />
      </div>
    </div>,
    <div key="locationSection">
      <h2>Включите геопозицию</h2>
      <div>
        <div>
          <div className={styles.geoText}>
            Чтобы мы подобрали вам людей не только близких по духу, но и по расположению
          </div>
          {!useManualLocation ? (
            <>
              <GeoLocationRequest onLocationReceived={handleLocationReceived} />
              {!coords && (
                <div className={styles.manualLocationOption}>
                  <p>Не хотите разрешать доступ к геопозиции?</p>
                  <CustomButton
                    onClick={() => setUseManualLocation(true)}
                    text="Ввести вручную"
                    className={styles.outlinedButton}
                  />
                </div>
              )}
            </>
          ) : (
            <div className={styles.manualLocationInput}>
              <h3>Укажите ваше местоположение</h3>
              <CustomInput
                placeholder="Например: Москва"
                type="text"
                value={manualLocation}
                onChange={e => setManualLocation(e.target.value)}
                className={styles.input}
              />
              <CustomButton
                onClick={() => setUseManualLocation(false)}
                text="Использовать геопозицию"
                className={`${styles.outlinedButton} ${styles.smallButton}`}
              />
            </div>
          )}
        </div>
      </div>
    </div>,
  ];

  console.log(birthDate);

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
