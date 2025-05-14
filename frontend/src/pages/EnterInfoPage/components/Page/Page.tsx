import KeyboardBackspaceIcon from '@mui/icons-material/KeyboardBackspace';
import PhotoCameraIcon from '@mui/icons-material/PhotoCamera';
import { useState, useRef } from 'react';
import { useDispatch } from 'react-redux';
import { useNavigate } from 'react-router-dom';

import { Gender } from '../../../../entities/user/api/constants';
import {
  useConfirmPhotoMutation,
  useCreateUserMutation,
} from '../../../../entities/user/api/userApi';
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
import getAge from '../../helpers/getAge';
import { GeoLocationRequest } from '../GeoLocationRequest';
import { FixedPhotoGallery } from '../PhotoGallery/PhotoGallery';

import styles from './enterInfo.module.css';

export const EnterInfoPage = () => {
  const [currentStep, setCurrentStep] = useState(0);
  const navigate = useNavigate();
  const dispatch = useDispatch();

  const [createUser, { isLoading }] = useCreateUserMutation();
  const [confirmPhoto] = useConfirmPhotoMutation();

  const [name, setName] = useState('');
  const [birthDate, setBirthDate] = useState('');
  const [height, setHeight] = useState('');
  const [userGender, setUserGender] = useState('');
  const [profilePhoto, setProfilePhoto] = useState<string | null>(null);
  const fileInputRef = useRef<HTMLInputElement>(null);

  const [showNotificationPrompt, setShowNotificationPrompt] = useState(false);
  const [notificationPermissionRequested, setNotificationPermissionRequested] = useState(false);

  const [coords, setCoords] = useState<{ lat: number; lng: number } | null>(null);

  const [userPhotos, setUserPhotos] = useState([]);

  const handleLocationReceived = (coordinates: { lat: number; lng: number }) => {
    setCoords(coordinates);
  };

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

        const locationString = coords ? `${coords.lat.toFixed(6)},${coords.lng.toFixed(6)}` : '';

        if (profilePhoto) {
          localStorage.setItem('profilePhotoUrl', profilePhoto);
        }

        // TODO
        if (profilePhoto) {
          localStorage.setItem('profilePhotoUrl', profilePhoto);
        }

        const photoUploadRequests = userPhotos.map((photo, index) => ({
          orderNumber: index,
          mimeType: photo.file.type,
        }));

        const orderNumbers = userPhotos.map((photo, index) => index);

        const userData = {
          name,
          age: ageValue,
          gender: genderEnum,
          height: String(height),
          ...(locationString ? { location: locationString } : {}),
          photoUploadRequests, // Добавляем массив метаданных
          latitude: coords?.lat,
          longitude: coords?.lng,
        };

        const response = await createUser(userData).unwrap();

        //TODO
        if (response.photoUploadResponses) {
          // Загружаем каждое фото на соответствующий URL
          await Promise.all(
            response.photoUploadResponses.map(async (resp: any) => {
              const photo = userPhotos[resp.orderNumber];
              if (!photo) return;

              await fetch(resp.uploadUrl, {
                method: 'PUT',
                body: photo.file,
                headers: {
                  'Content-Type': photo.file.type,
                },
              });
            }),
          );
        }

        try {
          await confirmPhoto({ orderNumbers }).unwrap();
        } catch (error) {
          console.error('Photo confirmation failed:', error);
          throw error;
        }

        // надо ли сохранять в локал TODO
        // if (response && response.user && response.user.id) {
        //   dispatch(setUserId(response.user.id));
        //   localStorage.setItem('userId', response.user.id);
        // }

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
        return name.trim() !== '';
      case 1: {
        const calculatedAge = getAge(birthDate);
        return birthDate !== '' && typeof calculatedAge === 'number' && calculatedAge >= 18;
      }
      case 2:
        return true;
      case 3:
        return true;
      case 4:
        return true;
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
        setUserPhotos(prevPhotos => [...prevPhotos, { file, url: fileUrl }]);
      }
    };
    input.click();
  };

  const handleRemovePhoto = (index: number) => {
    const newPhotos = userPhotos.filter((_, i) => i !== index);
    setUserPhotos(newPhotos);
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
    <div key="datingSettingsSection">
      <h2>Включите геопозицию</h2>
      <div>
        <div>
          <div className={styles.geoText}>
            Чтобы мы подобрали вам людей не только близких по духу, но и по расположению
          </div>
          <GeoLocationRequest onLocationReceived={handleLocationReceived} />
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
