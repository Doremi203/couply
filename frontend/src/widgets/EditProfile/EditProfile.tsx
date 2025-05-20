import React, { useRef, useState } from 'react';

import { useUploadFileToS3Mutation } from '../../entities/photo/api/photoApi';
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
import { mapInterestsToBackendFormat } from '../../features/filters/helpers/mapInterestsToApiFormat';
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
  const [selectedInterests, setSelectedInterests] = useState<string[]>([]);
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
        interest: mapInterestsToBackendFormat(selectedInterests),
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
      };

      // @ts-ignore - The API seems to work differently in practice vs type definition
      const response: any = await updateUser(userData).unwrap();

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
