import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Button,
  Typography,
} from '@mui/material';
import React, { useRef, useState, useEffect } from 'react';
import { useDispatch } from 'react-redux';

import { useUploadFileToS3Mutation } from '../../entities/photo/api/photoApi';
import { setProfileData } from '../../entities/profile/model/profileSlice';
import { useUpdateUserMutation } from '../../entities/user';
import {
  alcoholFromApi,
  alcoholOptions,
  alcoholToApi,
  childrenFromApi,
  childrenOptions,
  childrenToApi,
  educationFromApi,
  educationOptions,
  educationToApi,
  goalFromApi,
  goalOptions,
  goalToApi,
  smokingFromApi,
  smokingOptions,
  smokingToApi,
} from '../../features/filters/components/constants';
import { mapInterestsFromApiFormat } from '../../features/filters/helpers/mapInterestsFromApiFormat';
import { mapInterestsToApiFormat } from '../../features/filters/helpers/mapInterestsToApiFormat';
import { PhotoGalleryEdit } from '../../features/photoGallery/components/PhotoGalleryEdit';
import { ProfileData } from '../../features/profileEdit';
import { ProfilePhotoEdit } from '../../features/profileEdit/components/ProfilePhotoEdit';
import { ProfileVisibilitySection } from '../../features/profileVisibility/components/ProfileVisibilitySection';
import AboutMeSection from '../../shared/components/AboutMeSection';
import PageHeader from '../../shared/components/PageHeader';
import { SaveButtonSection } from '../../shared/components/SaveButtonSection';

import { EditSection } from './components/EditSection/EditSection';
import { InterestSection } from './components/InterestsSection/InterestsSection';
import styles from './editProfile.module.css';

export interface EditProfileProps {
  profileData: ProfileData;
  onBack: () => void;
  onPhotoAdd: (file?: File, isAvatar?: boolean) => void;
  onPhotoRemove: (index: number) => void;
}

interface PhotoItem {
  file: File;
  url: string;
  order_number: number;
}

export const EditProfile: React.FC<EditProfileProps> = ({
  profileData,
  onBack,
  onPhotoAdd,
  onPhotoRemove,
}) => {
  const dispatch = useDispatch();

  const [updateUser] = useUpdateUserMutation();
  const [uploadFile] = useUploadFileToS3Mutation();

  const fileInputRef = useRef<HTMLInputElement>(null);
  const [isAvatarUpload, setIsAvatarUpload] = useState(false);

  const [selectedEducation, setSelectedEducation] = useState<string[]>([
    //@ts-ignore
    profileData.education ? educationFromApi[profileData.education] : '',
  ]);
  const [selectedChildren, setSelectedChildren] = useState<string[]>([
    //@ts-ignore
    profileData.children ? childrenFromApi[profileData.children] : '',
  ]);
  const [selectedAlcohol, setSelectedAlcohol] = useState<string[]>([
    //@ts-ignore
    profileData.alcohol ? alcoholFromApi[profileData.alcohol] : '',
  ]);
  const [selectedSmoking, setSelectedSmoking] = useState<string[]>([
    //@ts-ignore
    profileData.smoking ? smokingFromApi[profileData.smoking] : '',
  ]);
  //@ts-ignore
  const [selectedGoal, setSelectedGoal] = useState<string[]>([
    //@ts-ignore
    profileData.goal ? goalFromApi[profileData.goal] : '',
  ]);

  //@ts-ignore
  const interest = profileData.interest ? mapInterestsFromApiFormat(profileData.interest) : [];
  const [selectedInterests, setSelectedInterests] = useState<string[]>(interest);
  const [bio, setBio] = useState(profileData.bio || '');
  const [isHidden, setIsHidden] = useState(profileData.isHidden || false);

  const [photoFiles, setPhotoFiles] = useState<PhotoItem[]>([]);
  const [hasUnsavedChanges, setHasUnsavedChanges] = useState(false);
  const [showConfirmModal, setShowConfirmModal] = useState(false);

  const MAX_PHOTOS = 6;

  const handleCameraClick = (isAvatar: boolean = false) => {
    if (!isAvatar && Array.isArray(profileData.photos) && profileData.photos.length >= MAX_PHOTOS) {
      alert(`Максимальное количество фото: ${MAX_PHOTOS}`);
      return;
    }

    setIsAvatarUpload(isAvatar);
    if (fileInputRef.current) {
      fileInputRef.current.click();
    }
  };

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const files = event.target.files;
    if (files && files.length > 0) {
      if (
        !isAvatarUpload &&
        Array.isArray(profileData.photos) &&
        profileData.photos.length >= MAX_PHOTOS
      ) {
        alert(`Максимальное количество фото: ${MAX_PHOTOS}`);
        event.target.value = '';
        return;
      }

      const file = files[0];
      const fileUrl = URL.createObjectURL(file);

      const order_number = profileData.photos.length;

      const newPhotoItem = {
        file,
        url: fileUrl,
        order_number,
      };
      setPhotoFiles(prev => [...prev, newPhotoItem]);

      onPhotoAdd(file, isAvatarUpload);
      event.target.value = '';
    }
  };

  const handleGoalSelect = (value: string) => {
    //@ts-ignore
    setSelectedGoal(value);
  };

  const handleSmokingSelect = (value: string) => {
    //@ts-ignore
    setSelectedSmoking(value);
  };

  const handleAlcoholSelect = (value: string) => {
    //@ts-ignore
    setSelectedAlcohol(value);
  };

  const handleChildrenSelect = (value: string) => {
    //@ts-ignore
    setSelectedChildren(value);
  };

  const handleEducationSelect = (value: string) => {
    //@ts-ignore
    setSelectedEducation(value);
  };

  const handleBioChange = (value: string) => {
    setBio(value);
  };

  const handleVisibilityChange = () => {
    setIsHidden(!isHidden);
  };

  const handleSave = async () => {
    try {
      const userData = {
        name: profileData.name,
        age: profileData.age,
        bio: bio,
        latitude: profileData.latitude,
        longitude: profileData.longitude,
        gender: profileData.gender,
        isHidden: isHidden,
        //@ts-ignore
        children: childrenToApi[selectedChildren],
        //@ts-ignore
        alcohol: alcoholToApi[selectedAlcohol],
        //@ts-ignore
        smoking: smokingToApi[selectedSmoking],
        //@ts-ignore
        goal: goalToApi[selectedGoal],
        interest: mapInterestsToApiFormat(selectedInterests),
        height: profileData.height,
        // photoUploadRequests,
        zodiac: profileData.zodiac,
        //@ts-ignore
        height: profileData.height,
        //@ts-ignore
        education: educationToApi[selectedEducation],
        isVerified: profileData.isVerified,
        isPremium: profileData.isPremium,
        isBlocked: profileData.isBlocked,
        photos: profileData.photos,
      };

      // @ts-ignore - The API seems to work differently in practice vs type definition
      const _response: any = await updateUser(userData).unwrap();

      //@ts-ignore
      dispatch(setProfileData(userData));

      const token = localStorage.getItem('token');
      const fetchWithRetry = async (url: string, options: RequestInit, maxRetries = 3) => {
        let retries = 0;

        while (retries < maxRetries) {
          try {
            const response = await fetch(url, options);

            if (response.status === 500) {
              retries++;
              console.log(`Получен статус 500, попытка ${retries} из ${maxRetries}`);

              const delay = 1000 * Math.pow(2, retries - 1);
              await new Promise(resolve => setTimeout(resolve, delay));
              continue;
            }

            return response;
          } catch (error) {
            retries++;
            console.error(`Ошибка запроса, попытка ${retries} из ${maxRetries}:`, error);

            if (retries >= maxRetries) {
              throw error;
            }

            const delay = 1000 * Math.pow(2, retries - 1);
            await new Promise(resolve => setTimeout(resolve, delay));
          }
        }

        throw new Error('Превышено максимальное количество попыток');
      };

      if (photoFiles.length > 0) {
        console.log(`Загрузка ${photoFiles.length} фото`);

        await Promise.all(
          photoFiles.map(async (photoFile, index) => {
            console.log(index);
            const order_number = profileData.photos.length + index - 1;

            const body = {
              token,
              bucket: 'testing-couply-profile-photos',
              order_number,
            };

            const getUrlResponse = await fetchWithRetry(
              'https://functions.yandexcloud.net/d4efh4n0sevvo2f928ri',
              {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(body),
              },
            );

            if (!getUrlResponse.ok)
              throw new Error(`Ошибка получения URL загрузки для фото ${index + 1}`);

            const { url } = await getUrlResponse.json();
            console.log(`Получен URL для фото ${index + 1}: ${url}`);

            await uploadFile({
              url,
              file: photoFile.file,
            }).unwrap();

            console.log(`Фото ${index + 1} успешно загружено`);

            return url.split('?')[0];
          }),
        );
      }

      onBack();
    } catch (error) {
      console.error('Failed to save profile data:', error);
    }
  };

  const checkForChanges = () => {
    const initialEducation = profileData.education ? educationFromApi[profileData.education] : '';
    const initialChildren = profileData.children ? childrenFromApi[profileData.children] : '';
    const initialAlcohol = profileData.alcohol ? alcoholFromApi[profileData.alcohol] : '';
    const initialSmoking = profileData.smoking ? smokingFromApi[profileData.smoking] : '';
    const initialGoal = profileData.goal ? goalFromApi[profileData.goal] : '';
    const initialInterests = profileData.interest
      ? mapInterestsFromApiFormat(profileData.interest)
      : [];
    const initialBio = profileData.bio || '';
    const initialIsHidden = profileData.isHidden || false;

    return (
      JSON.stringify(selectedEducation) !== JSON.stringify([initialEducation]) ||
      JSON.stringify(selectedChildren) !== JSON.stringify([initialChildren]) ||
      JSON.stringify(selectedAlcohol) !== JSON.stringify([initialAlcohol]) ||
      JSON.stringify(selectedSmoking) !== JSON.stringify([initialSmoking]) ||
      JSON.stringify(selectedGoal) !== JSON.stringify([initialGoal]) ||
      JSON.stringify(selectedInterests) !== JSON.stringify(initialInterests) ||
      bio !== initialBio ||
      isHidden !== initialIsHidden ||
      photoFiles.length > 0
    );
  };

  const handleBack = () => {
    if (hasUnsavedChanges) {
      setShowConfirmModal(true);
    } else {
      onBack();
    }
  };

  const handleConfirmClose = () => {
    setShowConfirmModal(false);
    onBack();
  };

  const handleSaveAndClose = async () => {
    await handleSave();
    setShowConfirmModal(false);
    onBack();
  };

  useEffect(() => {
    const changes = checkForChanges();
    setHasUnsavedChanges(changes);
  }, [
    selectedEducation,
    selectedChildren,
    selectedAlcohol,
    selectedSmoking,
    selectedGoal,
    selectedInterests,
    bio,
    isHidden,
    photoFiles,
  ]);

  return (
    <div>
      <input
        type="file"
        ref={fileInputRef}
        style={{ display: 'none' }}
        accept="image/*"
        onChange={handleFileChange}
      />

      <PageHeader onBack={handleBack} title="Редактирование" />

      <div className={styles.editContent}>
        <ProfilePhotoEdit
          profilePhoto={
            Array.isArray(profileData.photos) && profileData.photos.length > 0
              ? typeof profileData.photos[0] === 'string'
                ? profileData.photos[0]
                : (profileData.photos[0] as any)?.url || ''
              : ''
          }
          onCameraClick={handleCameraClick}
        />

        <PhotoGalleryEdit
          photos={profileData.photos || []}
          onPhotoRemove={onPhotoRemove}
          onAddPhotoClick={() => handleCameraClick(false)}
        />

        <AboutMeSection about={bio} onInputChange={handleBioChange} />

        <EditSection
          title="Цель"
          options={Object.values(goalOptions)}
          selectedOptions={selectedGoal}
          onToggle={handleGoalSelect}
        />

        <InterestSection
          selectedOptions={selectedInterests}
          onSelect={selected => setSelectedInterests(selected)}
        />

        <EditSection
          title="Образование"
          options={Object.values(educationOptions)}
          selectedOptions={selectedEducation}
          onToggle={handleEducationSelect}
        />

        <EditSection
          title="Курение"
          options={Object.values(smokingOptions)}
          selectedOptions={selectedSmoking}
          onToggle={handleSmokingSelect}
        />

        <EditSection
          title="Алкоголь"
          options={Object.values(alcoholOptions)}
          selectedOptions={selectedAlcohol}
          onToggle={handleAlcoholSelect}
        />

        <EditSection
          title="Дети"
          options={Object.values(childrenOptions)}
          selectedOptions={selectedChildren}
          onToggle={handleChildrenSelect}
        />

        <ProfileVisibilitySection isHidden={isHidden} onInputChange={handleVisibilityChange} />

        <SaveButtonSection onSave={handleSave} />

        <div className={styles.bottom} />
      </div>

      <Dialog
        open={showConfirmModal}
        onClose={() => setShowConfirmModal(false)}
        PaperProps={{
          sx: {
            borderRadius: '16px',
            padding: '24px',
            maxWidth: '400px',
            width: '100%',
            fontFamily: 'Jost',
            boxShadow: '0px 4px 10px var(--shadow-color)',
          },
        }}
      >
        <DialogTitle
          sx={{
            padding: 0,
            marginBottom: '16px',
            '& .MuiTypography-root': {
              fontSize: '20px',
              fontWeight: 600,
              color: '#1A1A1A',
            },
            fontFamily: 'Jost',
          }}
        >
          Есть несохраненные изменения
        </DialogTitle>
        <DialogContent
          sx={{
            padding: 0,
            marginBottom: '24px',
            fontFamily: 'Jost',
          }}
        >
          <Typography
            sx={{
              fontSize: '16px',
              color: '#666666',
              lineHeight: '24px',
              fontFamily: 'Jost',
            }}
          >
            Вы хотите сохранить изменения перед выходом?
          </Typography>
        </DialogContent>
        <DialogActions
          sx={{
            padding: 0,
            gap: '12px',
            justifyContent: 'flex-end',
            fontFamily: 'Jost',
          }}
        >
          <Button
            onClick={handleConfirmClose}
            sx={{
              padding: '12px 24px',
              borderRadius: '12px',
              textTransform: 'none',
              fontSize: '16px',
              fontWeight: 500,
              color: '#666666',
              // backgroundColor: 'grey',
              '&:hover': {
                backgroundColor: '#F5F5F5',
              },
              fontFamily: 'Jost',
            }}
          >
            Выйти
          </Button>
          <Button
            onClick={handleSaveAndClose}
            variant="contained"
            sx={{
              padding: '12px 24px',
              borderRadius: '12px',
              textTransform: 'none',
              fontSize: '16px',
              fontWeight: 500,
              backgroundColor: '#3b5eda',
              '&:hover': {
                backgroundColor: '#3b5eda',
              },
              fontFamily: 'Jost',
            }}
          >
            Сохранить
          </Button>
        </DialogActions>
      </Dialog>
    </div>
  );
};

export default EditProfile;
