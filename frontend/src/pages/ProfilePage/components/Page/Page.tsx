import React, { useState } from "react";
import styles from "./profilePage.module.css";
import { NavBar } from "../../../../shared/components/NavBar";
import { CustomInput } from "../../../../shared/components/CustomInput";
import { CustomButton } from "../../../../shared/components/CustomButton";
import { ToggleButtons } from "../../../../shared/components/ToggleButtons";
import KeyboardBackspaceIcon from "@mui/icons-material/KeyboardBackspace";
import VisibilityIcon from "@mui/icons-material/Visibility";
import VisibilityOffIcon from "@mui/icons-material/VisibilityOff";
import EditIcon from "@mui/icons-material/Edit";
import AddIcon from "@mui/icons-material/Add";
import CloseIcon from "@mui/icons-material/Close";
import VerifiedIcon from "@mui/icons-material/Verified";
import HistoryIcon from "@mui/icons-material/History";
import PhotoCameraIcon from "@mui/icons-material/PhotoCamera";

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

  const handlePhotoAdd = () => {
    // In a real app, this would open a file picker
    // For now, we'll just add a placeholder
    setProfileData({
      ...profileData,
      photos: [...profileData.photos, "/man1.jpg"]
    });
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
          <div className={styles.backButton} onClick={() => setActiveTab("profile")}>
            <KeyboardBackspaceIcon />
          </div>
          <h5>Profile</h5>
          <div className={styles.profileActions}>
            <div className={styles.actionIcon} onClick={handleEditToggle}>
              <EditIcon />
            </div>
            <div className={styles.actionIcon} onClick={handleProfileVisibilityToggle}>
              {isProfileHidden ? <VisibilityOffIcon /> : <VisibilityIcon />}
            </div>
            <div className={styles.actionIcon} onClick={() => setActiveTab("activity")}>
              <HistoryIcon />
            </div>
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
          <div className={styles.editProfileButton}>
            <CustomButton
              text="Edit Profile"
              onClick={handleEditToggle}
              className={styles.editButton}
            />
          </div>
        </div>

        <div className={styles.photoGallery}>
          <div className={styles.sectionHeader}>
            <h3>Photos</h3>
            <span className={styles.editLink} onClick={() => setActiveTab("edit")}>Edit</span>
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
            <span className={styles.editLink} onClick={() => setActiveTab("edit")}>Edit</span>
          </div>
          <p>{profileData.about}</p>
        </div>

        <div className={styles.profileSection}>
          <div className={styles.sectionHeader}>
            <h3>Basic Information</h3>
            <span className={styles.editLink} onClick={() => setActiveTab("edit")}>Edit</span>
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
            <span className={styles.editLink} onClick={() => setActiveTab("edit")}>Edit</span>
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
            <span className={styles.editLink} onClick={() => setActiveTab("edit")}>Edit</span>
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
            <span className={styles.editLink} onClick={() => setActiveTab("edit")}>Edit</span>
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
            <span className={styles.editLink} onClick={() => setActiveTab("edit")}>Edit</span>
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
            <span className={styles.editLink} onClick={() => setActiveTab("edit")}>Edit</span>
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
            text="View as Others See"
            onClick={() => setActiveTab("preview")}
            className={styles.previewButton}
          />
        </div>
      </div>
    );
  };

  const renderEditMode = () => {
    return (
      <div className={styles.editContent}>
        <div className={styles.profileHeader}>
          <div className={styles.backButton} onClick={() => setActiveTab("profile")}>
            <KeyboardBackspaceIcon />
          </div>
          <h5>Edit Profile</h5>
        </div>

        <div className={styles.photoEditSection}>
          <h3>Profile Photo</h3>
          <div className={styles.profileImageEdit}>
            <img
              src={profileData.photos[0] || "/photo1.png"}
              alt="Profile"
              className={styles.profilePic}
            />
            <div className={styles.photoEditIcon}>
              <PhotoCameraIcon />
            </div>
          </div>
        </div>

        <div className={styles.photoGalleryEdit}>
          <h3>Photos</h3>
          <div className={styles.photosGrid}>
            {profileData.photos.map((photo, index) => (
              <div key={index} className={styles.photoItemEdit}>
                <img src={photo} alt={`Photo ${index + 1}`} />
                <div 
                  className={styles.removePhotoIcon}
                  onClick={() => handlePhotoRemove(index)}
                >
                  <CloseIcon />
                </div>
              </div>
            ))}
            <div className={styles.addPhotoItem} onClick={handlePhotoAdd}>
              <AddIcon />
            </div>
          </div>
        </div>

        <div className={styles.editSection}>
          <h3>Basic Information</h3>
          <div className={styles.editField}>
            <label>Name</label>
            <CustomInput
              type="text"
              placeholder="Your name"
              value={profileData.name}
              onChange={(e) => handleInputChange("name", e.target.value)}
            />
          </div>
          <div className={styles.editField}>
            <label>Age</label>
            <CustomInput
              type="number"
              placeholder="Your age"
              value={profileData.age.toString()}
              onChange={(e) => handleInputChange("age", e.target.value)}
            />
          </div>
          <div className={styles.editField}>
            <label>Date of Birth</label>
            <CustomInput
              type="date"
              placeholder="Date of birth"
              value={profileData.dateOfBirth}
              onChange={(e) => handleInputChange("dateOfBirth", e.target.value)}
            />
          </div>
          <div className={styles.editField}>
            <label>Phone</label>
            <CustomInput
              type="tel"
              placeholder="Your phone number"
              value={profileData.phone}
              onChange={(e) => handleInputChange("phone", e.target.value)}
            />
          </div>
          <div className={styles.editField}>
            <label>Email</label>
            <CustomInput
              type="email"
              placeholder="Your email"
              value={profileData.email}
              onChange={(e) => handleInputChange("email", e.target.value)}
            />
          </div>
          <div className={styles.editField}>
            <label>Gender</label>
            <ToggleButtons
              options={[
                { label: "Female", value: "female" },
                { label: "Male", value: "male" },
              ]}
              value={profileData.gender}
              onSelect={(value) => handleInputChange("gender", value)}
            />
          </div>
        </div>

        <div className={styles.editSection}>
          <h3>About Me</h3>
          <textarea
            className={styles.textareaInput}
            placeholder="Tell something about yourself"
            value={profileData.about}
            onChange={(e) => handleInputChange("about", e.target.value)}
          />
        </div>

        <div className={styles.editSection}>
          <h3>Interests</h3>
          <CustomInput
            type="text"
            placeholder="Interests (comma separated)"
            value={profileData.interests.join(", ")}
            onChange={(e) => handleArrayInputChange("interests", e.target.value)}
          />
        </div>

        <div className={styles.editSection}>
          <h3>Music</h3>
          <CustomInput
            type="text"
            placeholder="Favorite music (comma separated)"
            value={profileData.music.join(", ")}
            onChange={(e) => handleArrayInputChange("music", e.target.value)}
          />
        </div>

        <div className={styles.editSection}>
          <h3>Movies</h3>
          <CustomInput
            type="text"
            placeholder="Favorite movies (comma separated)"
            value={profileData.movies.join(", ")}
            onChange={(e) => handleArrayInputChange("movies", e.target.value)}
          />
        </div>

        <div className={styles.editSection}>
          <h3>Books</h3>
          <CustomInput
            type="text"
            placeholder="Favorite books (comma separated)"
            value={profileData.books.join(", ")}
            onChange={(e) => handleArrayInputChange("books", e.target.value)}
          />
        </div>

        <div className={styles.editSection}>
          <h3>Hobbies</h3>
          <CustomInput
            type="text"
            placeholder="Hobbies (comma separated)"
            value={profileData.hobbies.join(", ")}
            onChange={(e) => handleArrayInputChange("hobbies", e.target.value)}
          />
        </div>

        <div className={styles.editSection}>
          <h3>Profile Visibility</h3>
          <div className={styles.toggleOption}>
            <span>Hide my profile</span>
            <ToggleButtons
              options={[
                { label: "No", value: "visible" },
                { label: "Yes", value: "hidden" },
              ]}
              value={profileData.isHidden ? "hidden" : "visible"}
              onSelect={(value) => handleInputChange("isHidden", value === "hidden" ? "true" : "false")}
            />
          </div>
        </div>

        <div className={styles.saveButtonContainer}>
          <CustomButton
            text="Save Changes"
            onClick={handleSaveChanges}
            className={styles.saveButton}
          />
        </div>
      </div>
    );
  };

  const renderActivityHistory = () => {
    return (
      <div className={styles.activityContent}>
        <div className={styles.profileHeader}>
          <div className={styles.backButton} onClick={() => setActiveTab("profile")}>
            <KeyboardBackspaceIcon />
          </div>
          <h5>Activity History</h5>
        </div>

        <div className={styles.activityList}>
          {activityHistory.map((activity, index) => (
            <div key={index} className={styles.activityItem}>
              <div className={styles.activityIcon}>
                {activity.type === "view" && <VisibilityIcon />}
                {activity.type === "like" && <span>❤️</span>}
                {activity.type === "message" && <span>💬</span>}
              </div>
              <div className={styles.activityDetails}>
                <span className={styles.activityUser}>{activity.user}</span>
                <span className={styles.activityType}>
                  {activity.type === "view" && "viewed your profile"}
                  {activity.type === "like" && "liked your profile"}
                  {activity.type === "message" && "sent you a message"}
                </span>
                <span className={styles.activityDate}>{formatDate(activity.date)}</span>
              </div>
            </div>
          ))}
        </div>
      </div>
    );
  };

  const renderProfilePreview = () => {
    // This shows how others see your profile
    return (
      <div className={styles.previewContent}>
        <div className={styles.profileHeader}>
          <div className={styles.backButton} onClick={() => setActiveTab("profile")}>
            <KeyboardBackspaceIcon />
          </div>
          <h5>Profile Preview</h5>
          <div className={styles.previewBadge}>
            <span>Preview Mode</span>
          </div>
        </div>

        <div className={styles.profileInfo}>
          <img
            src={profileData.photos[0] || "/photo1.png"}
            alt="Profile"
            className={styles.profilePic}
          />
          <h2>{profileData.name}, {profileData.age}</h2>
          {isVerified && (
            <div className={styles.verificationBadge}>
              <VerifiedIcon />
            </div>
          )}
        </div>

        <div className={styles.photoGallery}>
          <h3>Photos</h3>
          <div className={styles.photosGrid}>
            {profileData.photos.map((photo, index) => (
              <div key={index} className={styles.photoItem}>
                <img src={photo} alt={`Photo ${index + 1}`} />
              </div>
            ))}
          </div>
        </div>

        <div className={styles.profileSection}>
          <h3>About</h3>
          <p>{profileData.about}</p>
        </div>

        <div className={styles.profileSection}>
          <h3>Interests</h3>
          <div className={styles.tagsList}>
            {profileData.interests.map((interest, index) => (
              <div key={index} className={`${styles.tag} ${index % 2 === 0 ? styles.commonInterest : ''}`}>
                {interest}
                {index % 2 === 0 && <span className={styles.commonBadge}>Common</span>}
              </div>
            ))}
          </div>
        </div>

        <div className={styles.profileSection}>
          <h3>Music</h3>
          <div className={styles.tagsList}>
            {profileData.music.map((item, index) => (
              <div key={index} className={`${styles.tag} ${index === 0 ? styles.commonInterest : ''}`}>
                {item}
                {index === 0 && <span className={styles.commonBadge}>Common</span>}
              </div>
            ))}
          </div>
        </div>

        <div className={styles.profileSection}>
          <h3>Movies</h3>
          <div className={styles.tagsList}>
            {profileData.movies.map((item, index) => (
              <div key={index} className={styles.tag}>
                {item}
              </div>
            ))}
          </div>
        </div>

        <div className={styles.profileSection}>
          <h3>Books</h3>
          <div className={styles.tagsList}>
            {profileData.books.map((item, index) => (
              <div key={index} className={`${styles.tag} ${index === 0 ? styles.commonInterest : ''}`}>
                {item}
                {index === 0 && <span className={styles.commonBadge}>Common</span>}
              </div>
            ))}
          </div>
        </div>

        <div className={styles.profileSection}>
          <h3>Hobbies</h3>
          <div className={styles.tagsList}>
            {profileData.hobbies.map((item, index) => (
              <div key={index} className={styles.tag}>
                {item}
              </div>
            ))}
          </div>
        </div>
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
