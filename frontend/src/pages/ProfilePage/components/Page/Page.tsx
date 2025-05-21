import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import {
  selectProfileData,
  selectProfileLoading,
  setProfileData,
  setLoading,
  setError,
} from '../../../../entities/profile/model/profileSlice';
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

interface PhotoItem {
  file: File;
  url: string;
  orderNumber: number;
}

export const ProfilePage: React.FC<ProfilePageProps> = ({
  initialTab = 'profile',
  initialEditMode = false,
  initialVerified = false,
}) => {
  const dispatch = useDispatch();
  const [getUser] = useGetUserMutation();
  const profileData = useSelector(selectProfileData);
  const isLoading = useSelector(selectProfileLoading);

  const [newPhotoFiles, setNewPhotoFiles] = React.useState<PhotoItem[]>([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        dispatch(setLoading(true));
        const data = await getUser({}).unwrap();
        dispatch(setProfileData(data.user));
      } catch (error) {
        console.error('Failed to fetch user:', error);
        dispatch(setError('Failed to fetch user data'));
      } finally {
        dispatch(setLoading(false));
      }
    };

    fetchData();
  }, [getUser, dispatch]);

  const [isEditMode, setIsEditMode] = React.useState(initialEditMode);
  const [activeTab, setActiveTab] = React.useState(initialTab);
  const [isProfileHidden, setIsProfileHidden] = React.useState(false);
  const [isVerified, setIsVerified] = React.useState(initialVerified);

  const MAX_PHOTOS = 6;

  if (isLoading) {
    return <div className={styles.loader} />;
  }

  const handleEditToggle = () => {
    setIsEditMode(!isEditMode);
    setActiveTab(isEditMode ? 'profile' : 'edit');
  };

  const handleProfileVisibilityToggle = () => {
    setIsProfileHidden(!isProfileHidden);
    dispatch(
      setProfileData({
        ...profileData,
        isHidden: !isProfileHidden,
      }),
    );
  };

  const handleVerificationRequest = () => {
    setIsVerified(true);
  };

  const handlePhotoAdd = (file?: File, isAvatar: boolean = false) => {
    if (!file) return;

    // Create object URL from the file
    const fileUrl = URL.createObjectURL(file);
    const currentPhotos = [...(profileData.photos || [])];

    // Get the next order number
    const orderNumber = newPhotoFiles.length;

    // Add to newPhotoFiles for later upload
    const newPhotoItem = {
      file,
      url: fileUrl,
      orderNumber,
    };
    setNewPhotoFiles(prev => [...prev, newPhotoItem]);

    if (isAvatar) {
      // If it's an avatar, place it at the first position
      if (currentPhotos.length > 0) {
        // Replace the first photo with the new avatar
        currentPhotos[0] = { url: fileUrl, isNew: true };
      } else {
        // If there are no photos, add the avatar as the first one
        currentPhotos.push({ url: fileUrl, isNew: true });
      }
    } else {
      // Regular photo addition, limit to MAX_PHOTOS
      if (currentPhotos.length < MAX_PHOTOS) {
        currentPhotos.push({ url: fileUrl, isNew: true });
      } else {
        alert(`Максимальное количество фото: ${MAX_PHOTOS}`);
        return;
      }
    }

    dispatch(
      setProfileData({
        ...profileData,
        photos: currentPhotos,
      }),
    );
  };

  const handlePhotoRemove = (index: number) => {
    //@ts-ignore
    const updatedPhotos = [...profileData.photos];
    if (index >= 0 && index < updatedPhotos.length) {
      // Check if it's a new photo that hasn't been uploaded yet
      const removedPhoto = updatedPhotos[index];
      if (removedPhoto.isNew) {
        // Also remove from newPhotoFiles
        setNewPhotoFiles(prev => prev.filter(p => p.url !== removedPhoto.url));
      } else {
        // For existing photos, we'd need to call an API to remove them
        // This depends on your backend API
      }

      // Remove from UI
      updatedPhotos.splice(index, 1);
      dispatch(
        setProfileData({
          ...profileData,
          photos: updatedPhotos,
        }),
      );
    }
  };

  const renderContent = () => {
    switch (activeTab) {
      case 'edit':
        return (
          <EditProfile
            //@ts-ignore
            profileData={profileData}
            onBack={() => {
              setActiveTab('profile');
            }}
            onPhotoAdd={handlePhotoAdd}
            onPhotoRemove={handlePhotoRemove}
          />
        );
      case 'preview':
        return (
          //@ts-ignore
          <ProfileView profile={profileData} onClose={() => setActiveTab('profile')} isProfile />
        );
      default:
        return (
          <Profile
            //@ts-ignore
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
