import React, { useRef, useState } from 'react';

import { PhotoGalleryEdit } from '../../features/photoGallery/components/PhotoGalleryEdit';
import { ProfileData } from '../../features/profileEdit';
import { ProfilePhotoEdit } from '../../features/profileEdit/components/ProfilePhotoEdit';
import { ProfileVisibilitySection } from '../../features/profileVisibility/components/ProfileVisibilitySection';
import AboutMeSection from '../../shared/components/AboutMeSection';
import { InterestsSection } from '../../shared/components/InterestsSection';
import PageHeader from '../../shared/components/PageHeader';
import { SaveButtonSection } from '../../shared/components/SaveButtonSection';

import styles from './editProfile.module.css';

export interface EditProfileProps {
  profileData: ProfileData;
  onBack: () => void;
  onSave: () => void;
  onInputChange: (field: string, value: string) => void;
  onArrayInputChange: (field: string, value: string) => void;
  onPhotoAdd: (file?: File, isAvatar?: boolean) => void;
  onPhotoRemove: (index: number) => void;
}

export const EditProfile: React.FC<EditProfileProps> = ({
  profileData,
  onBack,
  onSave,
  onInputChange,
  onArrayInputChange,
  onPhotoAdd,
  onPhotoRemove,
}) => {
  const fileInputRef = useRef<HTMLInputElement>(null);
  const [isAvatarUpload, setIsAvatarUpload] = useState(false);

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

  return (
    <div className={styles.editContent}>
      <input
        type="file"
        ref={fileInputRef}
        style={{ display: 'none' }}
        accept="image/*"
        onChange={handleFileChange}
      />

      <PageHeader onBack={onBack} title="edit profile" />

      <ProfilePhotoEdit profilePhoto={profileData.photos[0]} onCameraClick={handleCameraClick} />

      <PhotoGalleryEdit
        photos={profileData.photos}
        onPhotoRemove={onPhotoRemove}
        onAddPhotoClick={() => handleCameraClick(false)}
      />

      {/* <BasicInfoForm profileData={profileData} onInputChange={onInputChange} /> */}

      <AboutMeSection about={profileData.about} onInputChange={onInputChange} />

      <InterestsSection
        title="Interests"
        placeholder="Interests (comma separated)"
        values={profileData.interests}
        fieldName="interests"
        onArrayInputChange={onArrayInputChange}
      />

      <InterestsSection
        title="Music"
        placeholder="Favorite music (comma separated)"
        values={profileData.music}
        fieldName="music"
        onArrayInputChange={onArrayInputChange}
      />

      <InterestsSection
        title="Movies"
        placeholder="Favorite movies (comma separated)"
        values={profileData.movies}
        fieldName="movies"
        onArrayInputChange={onArrayInputChange}
      />

      <InterestsSection
        title="Books"
        placeholder="Favorite books (comma separated)"
        values={profileData.books}
        fieldName="books"
        onArrayInputChange={onArrayInputChange}
      />

      <InterestsSection
        title="Hobbies"
        placeholder="Hobbies (comma separated)"
        values={profileData.hobbies}
        fieldName="hobbies"
        onArrayInputChange={onArrayInputChange}
      />

      <ProfileVisibilitySection isHidden={profileData.isHidden} onInputChange={onInputChange} />

      <SaveButtonSection onSave={onSave} />
    </div>
  );
};

export default EditProfile;
