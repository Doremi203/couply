import React, { useState } from "react";
import styles from "./profilePage.module.css";
import { NavBar } from "../../../../shared/components/NavBar";
import { CustomButton } from "../../../../shared/components/CustomButton";
import VisibilityIcon from "@mui/icons-material/Visibility";
import VisibilityOffIcon from "@mui/icons-material/VisibilityOff";
import EditIcon from "@mui/icons-material/Edit";
import VerifiedIcon from "@mui/icons-material/Verified";
import HistoryIcon from "@mui/icons-material/History";
import { IconButton } from "../../../../shared/components/IconButton";
import { EditProfile } from "../EditProfile";
import { ActivityHistory } from "../ActivityHistory";
import { ProfileView } from "../../../../pages/LikesPage/components/ProfileView/ProfileView";

interface ProfilePageProps {
  initialTab?: string;
  initialEditMode?: boolean;
  initialVerified?: boolean;
}

export const ProfilePage: React.FC<ProfilePageProps> = ({
  initialTab = "profile",
  initialEditMode = false,
  initialVerified = false
}) => {
  // State for edit mode
  const [isEditMode, setIsEditMode] = useState(initialEditMode);
  const [activeTab, setActiveTab] = useState(initialTab);
  const [isProfileHidden, setIsProfileHidden] = useState(false);
  const [isVerified, setIsVerified] = useState(initialVerified);

  // User profile data
  const [profileData, setProfileData] = useState({
    name: "Jenny",
    age: 22,
    phone: "+91 9876543210",
    dateOfBirth: "1997-05-02",
    email: "abcqwertyu@gmail.com",
    gender: "female",
    interests: ["Travel", "Fashion", "Music"],
    about: "I love traveling and exploring new places.",
    music: ["Pop", "Rock", "Jazz"],
    movies: ["Comedy", "Action", "Drama"],
    books: ["Fiction", "Biography"],
    hobbies: ["Photography", "Cooking", "Hiking"],
    isHidden: false,
    photos: [
      "/photo1.png",
      "/woman1.jpg"
    ]
  });

  // Activity history
  const [activityHistory] = useState([
    { type: "view", user: "Alex", date: "2025-03-28T14:30:00" },
    { type: "like", user: "Michael", date: "2025-03-27T10:15:00" },
    { type: "message", user: "David", date: "2025-03-26T18:45:00" },
    { type: "view", user: "Sarah", date: "2025-03-25T09:20:00" },
  ]);

  const handleEditToggle = () => {
    setIsEditMode(!isEditMode);
    setActiveTab(isEditMode ? "profile" : "edit");
  };

  const handleProfileVisibilityToggle = () => {
    setIsProfileHidden(!isProfileHidden);
    setProfileData({
      ...profileData,
      isHidden: !isProfileHidden
    });
  };

  const handleVerificationRequest = () => {
    // In a real app, this would trigger a verification process
    setIsVerified(true);
  };

  const handleInputChange = (field: string, value: string) => {
    setProfileData({
      ...profileData,
      [field]: value
    });
  };

  const handleArrayInputChange = (field: string, value: string) => {
    // Split by commas and trim whitespace
    const values = value.split(',').map(item => item.trim());
    setProfileData({
      ...profileData,
      [field]: values
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
          photos: updatedPhotos
        });
      } else {
        // Otherwise, add it to the end of the photos array
        setProfileData({
          ...profileData,
          photos: [...profileData.photos, fileUrl]
        });
      }
    } else {
      // Fallback to placeholder if no file is provided
      const placeholderUrl = "/man1.jpg";
      
      if (isAvatar) {
        // If this is an avatar upload, set it as the first photo
        const updatedPhotos = [...profileData.photos];
        updatedPhotos.unshift(placeholderUrl); // Add to the beginning of the array
        setProfileData({
          ...profileData,
          photos: updatedPhotos
        });
      } else {
        // Otherwise, add it to the end of the photos array
        setProfileData({
          ...profileData,
          photos: [...profileData.photos, placeholderUrl]
        });
      }
    }
  };

  const handlePhotoRemove = (index: number) => {
    const updatedPhotos = [...profileData.photos];
    updatedPhotos.splice(index, 1);
    setProfileData({
      ...profileData,
      photos: updatedPhotos
    });
  };

  const handleSaveChanges = () => {
    // In a real app, this would save to a backend
    setIsEditMode(false);
    setActiveTab("profile");
  };

  const formatDate = (dateString: string) => {
    const date = new Date(dateString);
    return date.toLocaleDateString();
  };

  const renderProfileView = () => {
    return (
      <div className={styles.profileContent}>
        <div className={styles.profileHeader}>
          <div className={styles.header}> profile</div>
         
          <div className={styles.profileActions}>
            <IconButton onClick={handleEditToggle} touchFriendly={true}>
              <EditIcon />
            </IconButton>
            <IconButton onClick={handleProfileVisibilityToggle} >
              {isProfileHidden ? <VisibilityOffIcon /> : <VisibilityIcon />}
            </IconButton>
            <IconButton onClick={() => setActiveTab("activity")} touchFriendly={true}>
              <HistoryIcon />
            </IconButton>
          </div>
        </div>

        <div className={styles.profileInfo}>
          <div className={styles.profileImageContainer}>
            <img
              src={profileData.photos[0] || "/photo1.png"}
              alt="Profile"
              className={styles.profilePic}
            />
            {isVerified && (
              <div className={styles.verificationBadge}>
                <VerifiedIcon />
              </div>
            )}
          </div>
          <h2>{profileData.name}, {profileData.age}</h2>
          {!isVerified && (
            <CustomButton
              text="Verify Profile"
              onClick={handleVerificationRequest}
              className={styles.verifyButton}
            />
          )}
        </div>

        <div className={styles.photoGallery}>
          <div className={styles.sectionHeader}>
            <h3>Photos</h3>
            {/* <span className={styles.editLink} onClick={() => setActiveTab("edit")}>Edit</span> */}
          </div>
          <div className={styles.photosGrid}>
            {profileData.photos.map((photo, index) => (
              <div key={index} className={styles.photoItem}>
                <img src={photo} alt={`Photo ${index + 1}`} />
              </div>
            ))}
          </div>
        </div>

        <div className={styles.profileSection}>
          <div className={styles.sectionHeader}>
            <h3>About Me</h3>
          </div>
          <p>{profileData.about}</p>
        </div>

        <div className={styles.profileSection}>
          <div className={styles.sectionHeader}>
            <h3>Basic Information</h3>
            {/* <span className={styles.editLink} onClick={() => setActiveTab("edit")}>Edit</span> */}
          </div>
          <div className={styles.infoGrid}>
            <div className={styles.infoItem}>
              <span className={styles.infoLabel}>Name:</span>
              <span>{profileData.name}</span>
            </div>
            <div className={styles.infoItem}>
              <span className={styles.infoLabel}>Age:</span>
              <span>{profileData.age}</span>
            </div>
            <div className={styles.infoItem}>
              <span className={styles.infoLabel}>Gender:</span>
              <span>{profileData.gender === "female" ? "Female" : "Male"}</span>
            </div>
            <div className={styles.infoItem}>
              <span className={styles.infoLabel}>Email:</span>
              <span>{profileData.email}</span>
            </div>
            <div className={styles.infoItem}>
              <span className={styles.infoLabel}>Phone:</span>
              <span>{profileData.phone}</span>
            </div>
          </div>
        </div>

        <div className={styles.profileSection}>
          <div className={styles.sectionHeader}>
            <h3>Interests</h3>
            {/* <span className={styles.editLink} onClick={() => setActiveTab("edit")}>Edit</span> */}
          </div>
          <div className={styles.tagsList}>
            {profileData.interests.map((interest, index) => (
              <div key={index} className={styles.tag}>
                {interest}
              </div>
            ))}
          </div>
        </div>

        <div className={styles.profileSection}>
          <div className={styles.sectionHeader}>
            <h3>Music</h3>
            {/* <span className={styles.editLink} onClick={() => setActiveTab("edit")}>Edit</span> */}
          </div>
          <div className={styles.tagsList}>
            {profileData.music.map((item, index) => (
              <div key={index} className={styles.tag}>
                {item}
              </div>
            ))}
          </div>
        </div>

        <div className={styles.profileSection}>
          <div className={styles.sectionHeader}>
            <h3>Movies</h3>
            {/* <span className={styles.editLink} onClick={() => setActiveTab("edit")}>Edit</span> */}
          </div>
          <div className={styles.tagsList}>
            {profileData.movies.map((item, index) => (
              <div key={index} className={styles.tag}>
                {item}
              </div>
            ))}
          </div>
        </div>

        <div className={styles.profileSection}>
          <div className={styles.sectionHeader}>
            <h3>Books</h3>
            {/* <span className={styles.editLink} onClick={() => setActiveTab("edit")}>Edit</span> */}
          </div>
          <div className={styles.tagsList}>
            {profileData.books.map((item, index) => (
              <div key={index} className={styles.tag}>
                {item}
              </div>
            ))}
          </div>
        </div>

        <div className={styles.profileSection}>
          <div className={styles.sectionHeader}>
            <h3>Hobbies</h3>
            {/* <span className={styles.editLink} onClick={() => setActiveTab("edit")}>Edit</span> */}
          </div>
          <div className={styles.tagsList}>
            {profileData.hobbies.map((item, index) => (
              <div key={index} className={styles.tag}>
                {item}
              </div>
            ))}
          </div>
        </div>

        <div className={styles.profileActions}>
          <CustomButton
            text="preview"
            onClick={() => setActiveTab("preview")}
            className={styles.previewButton}
          />
        </div>
      </div>
    );
  };

  const renderEditMode = () => {
    return (
      <EditProfile
        profileData={profileData}
        onBack={() => setActiveTab("profile")}
        onSave={handleSaveChanges}
        onInputChange={handleInputChange}
        onArrayInputChange={handleArrayInputChange}
        onPhotoAdd={handlePhotoAdd}
        onPhotoRemove={handlePhotoRemove}
      />
    );
  };

  const renderActivityHistory = () => {
    return (
      <ActivityHistory
        activityHistory={activityHistory}
        onBack={() => setActiveTab("profile")}
        formatDate={formatDate}
      />
    );
  };

  const renderProfilePreview = () => {
    // Create a profile object that matches the ProfileView component's expected props
    const profile = {
      id: 1, // Dummy ID
      name: profileData.name,
      age: profileData.age,
      imageUrl: profileData.photos[0] || "/photo1.png",
      bio: profileData.about,
      location: "Your Location", // You can add location to profileData if needed
      interests: profileData.interests,
      passion: [...profileData.interests, ...profileData.hobbies],
      photos: profileData.photos,
      lifestyle: {
        kids: "I don't have kids", // Example lifestyle data
      }
    };

    // Custom styles for the preview badge
    const previewBadgeStyle: React.CSSProperties = {
      position: 'absolute',
      top: '20px',
      right: '15px',
      backgroundColor: '#ff9f43',
      color: 'white',
      padding: '5px 10px',
      borderRadius: '20px',
      fontSize: '12px',
      zIndex: 1001, // Higher than the ProfileView's z-index
      fontWeight: 'bold'
    };

    return (
      <div style={{ position: 'relative' as const }}>
        {/* Preview badge overlay */}
        <div style={previewBadgeStyle}>
          <span>Preview Mode</span>
        </div>
        
        {/* Use the ProfileView component from LikesPage */}
        <ProfileView
          profile={profile}
          onClose={() => setActiveTab("profile")}
          onLike={() => {}} // Empty function since we don't need like functionality in preview
        />
      </div>
    );
  };

  const renderContent = () => {
    switch (activeTab) {
      case "edit":
        return renderEditMode();
      case "activity":
        return renderActivityHistory();
      case "preview":
        return renderProfilePreview();
      default:
        return renderProfileView();
    }
  };

  return (
    <div className={styles.pageContainer}>
      {renderContent()}
      <NavBar />
    </div>
  );
};

export default ProfilePage;
