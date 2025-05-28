import React, { useRef, useState } from 'react';
import { useDispatch } from 'react-redux';

import { useUploadFileToS3Mutation } from '../../entities/photo/api/photoApi';
import { setProfileData } from '../../entities/profile/model/profileSlice';
import { useUpdateUserMutation } from '../../entities/user';
import { useConfirmPhotoMutation } from '../../entities/user/api/userApi';
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
  orderNumber: number;
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
  const [confirmPhoto] = useConfirmPhotoMutation();

  const fileInputRef = useRef<HTMLInputElement>(null);
  const [isAvatarUpload, setIsAvatarUpload] = useState(false);

  const [selectedEducation, setSelectedEducation] = useState<string[]>([
    //@ts-ignore
    educationFromApi[profileData.education],
  ]);
  const [selectedChildren, setSelectedChildren] = useState<string[]>([
    //@ts-ignore
    childrenFromApi[profileData.children],
  ]);
  const [selectedAlcohol, setSelectedAlcohol] = useState<string[]>([
    //@ts-ignore
    alcoholFromApi[profileData.alcohol],
  ]);
  const [selectedSmoking, setSelectedSmoking] = useState<string[]>([
    //@ts-ignore
    smokingFromApi[profileData.smoking],
  ]);
  //@ts-ignore
  const [selectedGoal, setSelectedGoal] = useState<string[]>([goalFromApi[profileData.goal]]);

  //@ts-ignore
  const interest = mapInterestsFromApiFormat(profileData.interest);
  const [selectedInterests, setSelectedInterests] = useState<string[]>(interest);
  const [bio, setBio] = useState(profileData.bio || '');
  const [isHidden, setIsHidden] = useState(profileData.isHidden || false);

  const [photoFiles, setPhotoFiles] = useState<PhotoItem[]>([]);

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

      const orderNumber = photoFiles.length;
      const newPhotoItem = {
        file,
        url: fileUrl,
        orderNumber,
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
      const photoUploadRequests =
        photoFiles.length > 0
          ? photoFiles.map(photo => ({
              orderNumber: photo.orderNumber,
              mimeType: photo.file.type,
            }))
          : undefined;

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
        photoUploadRequests,
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
      const response: any = await updateUser(userData).unwrap();

      //@ts-ignore
      dispatch(setProfileData(userData));

      if (photoFiles.length > 0 && response && response.photoUploadResponses) {
        await Promise.all(
          response.photoUploadResponses.map(
            async (resp: { orderNumber: number; uploadUrl: string }) => {
              const photo = photoFiles.find(p => p.orderNumber === resp.orderNumber);
              if (!photo?.file) return;

              try {
                await uploadFile({
                  url: resp.uploadUrl,
                  file: photo.file,
                }).unwrap();
              } catch (error) {
                console.error('Error uploading file:', error);
              }
            },
          ),
        );

        const orderNumbers = photoFiles.map(photo => photo.orderNumber);
        await confirmPhoto({ orderNumbers }).unwrap();
      }

      onBack();
    } catch (error) {
      console.error('Failed to save profile data:', error);
    }
  };

  return (
    <div>
      <input
        type="file"
        ref={fileInputRef}
        style={{ display: 'none' }}
        accept="image/*"
        onChange={handleFileChange}
      />

      <PageHeader onBack={onBack} title="Редактирование" />

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
    </div>
  );
};

export default EditProfile;

// import React, { useRef, useState } from 'react';
// import { useDispatch } from 'react-redux';

// import { useUploadFileToS3Mutation } from '../../entities/photo/api/photoApi';
// import { setProfileData } from '../../entities/profile/model/profileSlice';
// import { getIsVerified, useUpdateUserMutation } from '../../entities/user';
// import { useConfirmPhotoMutation } from '../../entities/user/api/userApi';
// import {
//   alcoholFromApi,
//   alcoholOptions,
//   alcoholToApi,
//   childrenFromApi,
//   childrenOptions,
//   childrenToApi,
//   educationFromApi,
//   educationOptions,
//   educationToApi,
//   goalFromApi,
//   goalOptions,
//   goalToApi,
//   smokingFromApi,
//   smokingOptions,
//   smokingToApi,
// } from '../../features/filters/components/constants';
// import { mapInterestsFromApiFormat } from '../../features/filters/helpers/mapInterestsFromApiFormat';
// import { mapInterestsToApiFormat } from '../../features/filters/helpers/mapInterestsToApiFormat';
// import { PhotoGalleryEdit } from '../../features/photoGallery/components/PhotoGalleryEdit';
// import { ProfileData } from '../../features/profileEdit';
// import { ProfilePhotoEdit } from '../../features/profileEdit/components/ProfilePhotoEdit';
// import { ProfileVisibilitySection } from '../../features/profileVisibility/components/ProfileVisibilitySection';
// import AboutMeSection from '../../shared/components/AboutMeSection';
// import PageHeader from '../../shared/components/PageHeader';
// import { SaveButtonSection } from '../../shared/components/SaveButtonSection';

// import { EditSection } from './components/EditSection/EditSection';
// import { InterestSection } from './components/InterestsSection/InterestsSection';
// import styles from './editProfile.module.css';

// export interface EditProfileProps {
//   profileData: ProfileData;
//   onBack: () => void;
//   onPhotoAdd: (file?: File, isAvatar?: boolean) => void;
//   onPhotoRemove: (index: number) => void;
// }

// interface PhotoItem {
//   file?: File;
//   url: string;
//   orderNumber: number;
//   isExisting?: boolean;
// }

// export const EditProfile: React.FC<EditProfileProps> = ({
//   profileData,
//   onBack,
//   onPhotoAdd,
//   onPhotoRemove,
// }) => {
//   const dispatch = useDispatch();

//   const [updateUser] = useUpdateUserMutation();
//   const [uploadFile] = useUploadFileToS3Mutation();
//   const [confirmPhoto] = useConfirmPhotoMutation();

//   const fileInputRef = useRef<HTMLInputElement>(null);
//   const [isAvatarUpload, setIsAvatarUpload] = useState(false);

//   // Инициализация состояний с существующими данными профиля
//   const [selectedEducation, setSelectedEducation] = useState<string[]>([
//     educationFromApi[profileData.education] || '',
//   ]);
//   const [selectedChildren, setSelectedChildren] = useState<string[]>([
//     childrenFromApi[profileData.children] || '',
//   ]);
//   const [selectedAlcohol, setSelectedAlcohol] = useState<string[]>([
//     alcoholFromApi[profileData.alcohol] || '',
//   ]);
//   const [selectedSmoking, setSelectedSmoking] = useState<string[]>([
//     smokingFromApi[profileData.smoking] || '',
//   ]);
//   const [selectedGoal, setSelectedGoal] = useState<string[]>([goalFromApi[profileData.goal] || '']);

//   const [selectedInterests, setSelectedInterests] = useState<string[]>(
//     mapInterestsFromApiFormat(profileData.interest),
//   );
//   const [bio, setBio] = useState(profileData.bio || '');
//   const [isHidden, setIsHidden] = useState(profileData.isHidden || false);

//   const handleBioChange = (value: string) => {
//     setBio(value);
//   };

//   const handleVisibilityChange = () => {
//     setIsHidden(!isHidden);
//   };

//   // Инициализация фото с существующими изображениями
//   const [photoFiles, setPhotoFiles] = useState<PhotoItem[]>(() => {
//     return (profileData.photos || []).map((photo, index) => ({
//       url: typeof photo === 'string' ? photo : photo.url,
//       orderNumber: index,
//       isExisting: true,
//     }));
//   });

//   const MAX_PHOTOS = 6;

//   const handleCameraClick = (isAvatar: boolean = false) => {
//     const currentCount = photoFiles.filter(p => !p.isExisting).length;
//     if (!isAvatar && currentCount >= MAX_PHOTOS) {
//       alert(`Максимальное количество новых фото: ${MAX_PHOTOS}`);
//       return;
//     }

//     setIsAvatarUpload(isAvatar);
//     if (fileInputRef.current) {
//       fileInputRef.current.click();
//     }
//   };

//   const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
//     const files = event.target.files;
//     if (files && files.length > 0) {
//       const currentCount = photoFiles.filter(p => !p.isExisting).length;
//       if (currentCount >= MAX_PHOTOS) {
//         alert(`Максимальное количество новых фото: ${MAX_PHOTOS}`);
//         event.target.value = '';
//         return;
//       }

//       const file = files[0];
//       const fileUrl = URL.createObjectURL(file);

//       setPhotoFiles(prev => [
//         ...prev,
//         {
//           file,
//           url: fileUrl,
//           orderNumber: prev.length,
//           isExisting: false,
//         },
//       ]);

//       onPhotoAdd(file, isAvatarUpload);
//       event.target.value = '';
//     }
//   };

//   const handleSave = async () => {
//     try {
//       // Фильтрация новых фото для загрузки
//       const newPhotos = photoFiles.filter(p => !p.isExisting);

//       console.log(profileData);

//       const photoUploadRequests = newPhotos.map(photo => ({
//         orderNumber: photo.orderNumber,
//         mimeType: photo.file!.type,
//       }));

//       const userData = {
//         // ...profileData,
//         name: profileData.name,
//         age: profileData.age,
//         gender: profileData.gender,
//         latitude: profileData.latitude,
//         longitude: profileData.longitude,
//         zodiac: profileData.zodiac,
//         height: profileData.height,
//         bio,
//         isHidden,
//         isVerified: profileData.isVerified,
//         isPremium: profileData.isPremium,
//         isBlocked: profileData.isBlocked,
//         children: childrenToApi[selectedChildren[0]],
//         alcohol: alcoholToApi[selectedAlcohol[0]],
//         smoking: smokingToApi[selectedSmoking[0]],
//         goal: goalToApi[selectedGoal[0]],
//         interest: mapInterestsToApiFormat(selectedInterests),
//         education: educationToApi[selectedEducation[0]],
//         photoUploadRequests,
//         // photos: photoFiles.map(p => p.url),
//       };

//       const response = await updateUser(userData).unwrap();

//       // Обновление данных профиля в хранилище
//       dispatch(setProfileData(userData));

//       // Загрузка новых фото
//       if (newPhotos.length > 0 && response?.photoUploadResponses) {
//         await Promise.all(
//           response.photoUploadResponses.map(async (resp: any) => {
//             const photo = newPhotos.find(p => p.orderNumber === resp.orderNumber);
//             if (!photo?.file) return;

//             await uploadFile({
//               url: resp.uploadUrl,
//               file: photo.file,
//             }).unwrap();
//           }),
//         );

//         // Подтверждение загрузки фото
//         const orderNumbers = newPhotos.map(p => p.orderNumber);
//         await confirmPhoto({ orderNumbers }).unwrap();
//       }

//       onBack();
//     } catch (error) {
//       console.error('Ошибка сохранения профиля:', error);
//     }
//   };

//   const handlePhotoRemove = (index: number) => {
//     setPhotoFiles(prev => prev.filter((_, i) => i !== index));
//     onPhotoRemove(index);
//   };

//   return (
//     <div>
//       <input
//         type="file"
//         ref={fileInputRef}
//         style={{ display: 'none' }}
//         accept="image/*"
//         onChange={handleFileChange}
//       />

//       <PageHeader onBack={onBack} title="Редактирование" />

//       <div className={styles.editContent}>
//         <ProfilePhotoEdit
//           profilePhoto={photoFiles[0]?.url || ''}
//           onCameraClick={() => handleCameraClick(true)}
//         />

//         <PhotoGalleryEdit
//           photos={photoFiles.map(p => p.url)}
//           onPhotoRemove={handlePhotoRemove}
//           onAddPhotoClick={() => handleCameraClick(false)}
//           maxPhotos={MAX_PHOTOS}
//         />

//         <AboutMeSection about={bio} onInputChange={handleBioChange} />

//         <EditSection
//           title="Цель"
//           options={Object.values(goalOptions)}
//           selectedOptions={selectedGoal}
//           onToggle={value => setSelectedGoal([value])}
//         />

//         <InterestSection selectedOptions={selectedInterests} onSelect={setSelectedInterests} />

//         <EditSection
//           title="Образование"
//           options={Object.values(educationOptions)}
//           selectedOptions={selectedEducation}
//           onToggle={value => setSelectedEducation([value])}
//         />

//         <EditSection
//           title="Курение"
//           options={Object.values(smokingOptions)}
//           selectedOptions={selectedSmoking}
//           onToggle={value => setSelectedSmoking([value])}
//         />

//         <EditSection
//           title="Алкоголь"
//           options={Object.values(alcoholOptions)}
//           selectedOptions={selectedAlcohol}
//           onToggle={value => setSelectedAlcohol([value])}
//         />

//         <EditSection
//           title="Дети"
//           options={Object.values(childrenOptions)}
//           selectedOptions={selectedChildren}
//           onToggle={value => setSelectedChildren([value])}
//         />

//         <ProfileVisibilitySection isHidden={isHidden} onInputChange={handleVisibilityChange} />

//         <SaveButtonSection onSave={handleSave} />

//         <div className={styles.bottom} />
//       </div>
//     </div>
//   );
// };

// export default EditProfile;
