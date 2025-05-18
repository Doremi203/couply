import React, { useRef, useState } from 'react';

import { useUpdateUserMutation } from '../../entities/user';
import {
  alcoholOptions,
  alcoholToApi,
  childrenOptions,
  childrenToApi,
  educationOptions,
  goalOptions,
  goalToApi,
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

export const EditProfile: React.FC<EditProfileProps> = ({
  profileData,
  onBack,
  onPhotoAdd,
  onPhotoRemove,
}) => {
  const [updateUser] = useUpdateUserMutation();
  const fileInputRef = useRef<HTMLInputElement>(null);
  const [isAvatarUpload, setIsAvatarUpload] = useState(false);

  const [selectedEducation, setSelectedEducation] = useState<string[]>([]);
  const [selectedChildren, setSelectedChildren] = useState<string[]>([]);
  const [selectedAlcohol, setSelectedAlcohol] = useState<string[]>([]);
  const [selectedSmoking, setSelectedSmoking] = useState<string[]>([]);
  const [selectedGoal, setSelectedGoal] = useState<string[]>([]);
  const [selectedInterests, setSelectedInterests] = useState<string[]>([]);
  const [bio, setBio] = useState(profileData.bio || '');
  const [isHidden, setIsHidden] = useState(profileData.isHidden || false);

  const handleCameraClick = (isAvatar: boolean = false) => {
    setIsAvatarUpload(isAvatar);
    if (fileInputRef.current) {
      fileInputRef.current.click();
    }
  };

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const files = event.target.files;
    if (files && files.length > 0) {
      onPhotoAdd(files[0], isAvatarUpload);
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
          photos={profileData.photos}
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

        <SaveButtonSection
          onSave={async () => {
            try {
              const userData = {
                name: profileData.name,
                age: profileData.age,
                bio: bio,
                isHidden: isHidden,
                //@ts-ignore
                children: childrenToApi[selectedChildren],
                //@ts-ignore
                alcohol: alcoholToApi[selectedAlcohol],
                //@ts-ignore
                smoking: smokingToApi[selectedSmoking],
                //@ts-ignore
                goal: goalToApi[selectedGoal],
                interests: mapInterestsToBackendFormat(selectedInterests),
                height: profileData.height,
              };

              //@ts-ignore
              await updateUser(userData).unwrap();

              onBack();
            } catch (error) {
              console.error('Failed to save profile data:', error);
            }
          }}
        />
      </div>
    </div>
  );
};

export default EditProfile;
