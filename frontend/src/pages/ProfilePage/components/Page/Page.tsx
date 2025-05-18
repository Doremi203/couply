import React, { useEffect, useState } from 'react';

import { useGetUserMutation } from '../../../../entities/user';
import { NavBar } from '../../../../shared/components/NavBar';
import { EditProfile } from '../../../../widgets/EditProfile';
import { ProfileView } from '../../../../widgets/ProfileView';
import { Profile } from '../Profile';

import styles from './profilePage.module.css';

interface ProfilePageProps {
  initialTab?: string;
  initialEditMode?: boolean;
  initialVerified?: boolean;
}

export const ProfilePage: React.FC<ProfilePageProps> = ({
  initialTab = 'profile',
  initialEditMode = false,
  initialVerified = false,
}) => {
  const [getUser] = useGetUserMutation();

  const [profileData, setProfileData] = useState<any>({
    name: '',
    age: 0,
    phone: '',
    dateOfBirth: '',
    email: '',
    gender: '',
    interests: [],
    about: '',
    music: [],
    movies: [],
    books: [],
    hobbies: [],
    isHidden: false,
    photos: [],
    bio: '',
  });

  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const data = await getUser({}).unwrap();
        setProfileData(data.user);
      } catch (error) {
        console.error('Failed to fetch user:', error);
      } finally {
        setIsLoading(false);
      }
    };

    fetchData();
  }, [getUser]);

  const [isEditMode, setIsEditMode] = useState(initialEditMode);
  const [activeTab, setActiveTab] = useState(initialTab);
  const [isProfileHidden, setIsProfileHidden] = useState(false);
  const [isVerified, setIsVerified] = useState(initialVerified);

  if (isLoading) {
    return <div className={styles.loader} />;
  }

  const handleEditToggle = () => {
    setIsEditMode(!isEditMode);
    setActiveTab(isEditMode ? 'profile' : 'edit');
  };

  const handleProfileVisibilityToggle = () => {
    setIsProfileHidden(!isProfileHidden);
    setProfileData({
      ...profileData,
      isHidden: !isProfileHidden,
    });
  };

  const handleVerificationRequest = () => {
    setIsVerified(true);
  };

  const handlePhotoAdd = (file?: File, isAvatar: boolean = false) => {
    if (file) {
      const fileUrl = URL.createObjectURL(file);

      if (isAvatar) {
        const updatedPhotos = ['man1.jpg'];
        updatedPhotos.unshift(fileUrl);
        setProfileData({
          ...profileData,
          photos: updatedPhotos,
        });
      } else {
        setProfileData({
          ...profileData,
          photos: ['man1.jpg', fileUrl],
        });
      }
    } else {
      const placeholderUrl = '/man1.jpg';

      if (isAvatar) {
        const updatedPhotos = ['man1.jpg'];
        updatedPhotos.unshift(placeholderUrl);
        setProfileData({
          ...profileData,
          photos: updatedPhotos,
        });
      } else {
        setProfileData({
          ...profileData,
          photos: ['man1.jpg', placeholderUrl],
        });
      }
    }
  };

  const handlePhotoRemove = (index: number) => {
    const updatedPhotos = ['man1.jpg'];
    updatedPhotos.splice(index, 1);
    setProfileData({
      ...profileData,
      photos: updatedPhotos,
    });
  };

  const renderContent = () => {
    switch (activeTab) {
      case 'edit':
        return (
          <EditProfile
            profileData={profileData}
            onBack={() => setActiveTab('profile')}
            onPhotoAdd={handlePhotoAdd}
            onPhotoRemove={handlePhotoRemove}
          />
        );
      case 'preview':
        //@ts-ignore
        return <ProfileView profile={profileData} onClose={() => setActiveTab('profile')} />;
      default:
        return (
          <Profile
            profileData={profileData}
            isVerified={isVerified}
            isProfileHidden={isProfileHidden}
            onEditToggle={handleEditToggle}
            onVisibilityToggle={handleProfileVisibilityToggle}
            onActivityClick={() => setActiveTab('activity')}
            onPreviewClick={() => setActiveTab('preview')}
            onVerificationRequest={handleVerificationRequest}
          />
        );
    }
  };

  return (
    <div className={styles.pageContainer}>
      {renderContent()}
      <div style={{ position: 'relative', zIndex: 1010 }}>
        <NavBar />
      </div>
    </div>
  );
};

export default ProfilePage;
