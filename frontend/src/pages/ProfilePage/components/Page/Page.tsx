import React, { useState } from 'react';

import { NavBar } from '../../../../shared/components/NavBar';
import ActivityHistory from '../../../../widgets/ActivityHistory';
import { EditProfile } from '../../../../widgets/EditProfile';
// import { ProfileView } from '../../../LikesPage/components/ProfileView';
import { ProfileView } from '../../../../widgets/ProfileView';
import { ProfileData, ActivityItem } from '../../types';
import { Profile } from '../Profile';
// import { ProfilePreview } from '../ProfilePreview';

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
  const [isEditMode, setIsEditMode] = useState(initialEditMode);
  const [activeTab, setActiveTab] = useState(initialTab);
  const [isProfileHidden, setIsProfileHidden] = useState(false);
  const [isVerified, setIsVerified] = useState(initialVerified);

  const [profileData, setProfileData] = useState<ProfileData>({
    name: 'Jenny',
    age: 22,
    phone: '+91 9876543210',
    dateOfBirth: '1997-05-02',
    email: 'abcqwertyu@gmail.com',
    gender: 'female',
    interests: ['Travel', 'Fashion', 'Music'],
    about: 'I love traveling and exploring new places.',
    music: ['Pop', 'Rock', 'Jazz'],
    movies: ['Comedy', 'Action', 'Drama'],
    books: ['Fiction', 'Biography'],
    hobbies: ['Photography', 'Cooking', 'Hiking'],
    isHidden: false,
    photos: ['/photo1.png', '/woman1.jpg'],
    imageUrl: '/photo1.png',
  });

  const [activityHistory] = useState<ActivityItem[]>([
    { type: 'view', user: 'Alex', date: '2025-03-28T14:30:00' },
    { type: 'like', user: 'Michael', date: '2025-03-27T10:15:00' },
    { type: 'message', user: 'David', date: '2025-03-26T18:45:00' },
    { type: 'view', user: 'Sarah', date: '2025-03-25T09:20:00' },
  ]);

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
    // In a real app, this would trigger a verification process
    setIsVerified(true);
  };

  const handleInputChange = (field: string, value: string) => {
    setProfileData({
      ...profileData,
      [field]: value,
    });
  };

  const handleArrayInputChange = (field: string, value: string) => {
    const values = value.split(',').map(item => item.trim());
    setProfileData({
      ...profileData,
      [field]: values,
    });
  };

  const handlePhotoAdd = (file?: File, isAvatar: boolean = false) => {
    if (file) {
      // Create a URL for the selected file
      const fileUrl = URL.createObjectURL(file);

      if (isAvatar) {
        // If this is an avatar upload, set it as the first photo
        const updatedPhotos = [...profileData.photos];
        updatedPhotos.unshift(fileUrl); // Add to the beginning of the array
        setProfileData({
          ...profileData,
          photos: updatedPhotos,
        });
      } else {
        // Otherwise, add it to the end of the photos array
        setProfileData({
          ...profileData,
          photos: [...profileData.photos, fileUrl],
        });
      }
    } else {
      // Fallback to placeholder if no file is provided
      const placeholderUrl = '/man1.jpg';

      if (isAvatar) {
        // If this is an avatar upload, set it as the first photo
        const updatedPhotos = [...profileData.photos];
        updatedPhotos.unshift(placeholderUrl); // Add to the beginning of the array
        setProfileData({
          ...profileData,
          photos: updatedPhotos,
        });
      } else {
        // Otherwise, add it to the end of the photos array
        setProfileData({
          ...profileData,
          photos: [...profileData.photos, placeholderUrl],
        });
      }
    }
  };

  const handlePhotoRemove = (index: number) => {
    const updatedPhotos = [...profileData.photos];
    updatedPhotos.splice(index, 1);
    setProfileData({
      ...profileData,
      photos: updatedPhotos,
    });
  };

  const handleSaveChanges = () => {
    // In a real app, this would save to a backend
    setIsEditMode(false);
    setActiveTab('profile');
  };

  const formatDate = (dateString: string) => {
    const date = new Date(dateString);
    return date.toLocaleDateString();
  };

  const renderContent = () => {
    switch (activeTab) {
      case 'edit':
        return (
          <EditProfile
            profileData={profileData}
            onBack={() => setActiveTab('profile')}
            onSave={handleSaveChanges}
            onInputChange={handleInputChange}
            onArrayInputChange={handleArrayInputChange}
            onPhotoAdd={handlePhotoAdd}
            onPhotoRemove={handlePhotoRemove}
          />
        );
      case 'activity':
        return (
          <ActivityHistory
            activityHistory={activityHistory}
            onBack={() => setActiveTab('profile')}
            formatDate={formatDate}
          />
        );
      case 'preview':
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
