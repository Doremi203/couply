export interface ProfileData {
  education: any;
  latitude: any;
  longitude: any;
  children: any;
  zodiac: any;
  isVerified: any;
  isPremium: any;
  isBlocked: any;
  bio: string;
  height: any;
  name: string;
  age: number;
  phone: string;
  dateOfBirth: string;
  email: string;
  gender: string;
  isHidden: boolean;
  photos: string[];
}

export interface EditProfileProps {
  profileData: ProfileData;
  onBack: () => void;
  onSave: () => void;
  onInputChange: (field: string, value: string) => void;
  onArrayInputChange: (field: string, value: string) => void;
  onPhotoAdd: (file?: File, isAvatar?: boolean) => void;
  onPhotoRemove: (index: number) => void;
}
