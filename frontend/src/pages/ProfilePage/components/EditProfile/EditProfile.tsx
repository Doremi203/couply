import React, { useRef } from "react";
import styles from "./editProfile.module.css";
import { CustomInput } from "../../../../shared/components/CustomInput";
import { CustomButton } from "../../../../shared/components/CustomButton";
import { ToggleButtons } from "../../../../shared/components/ToggleButtons";
import KeyboardBackspaceIcon from "@mui/icons-material/KeyboardBackspace";
import PhotoCameraIcon from "@mui/icons-material/PhotoCamera";
import AddIcon from "@mui/icons-material/Add";
import CloseIcon from "@mui/icons-material/Close";

interface ProfileData {
  name: string;
  age: number;
  phone: string;
  dateOfBirth: string;
  email: string;
  gender: string;
  interests: string[];
  about: string;
  music: string[];
  movies: string[];
  books: string[];
  hobbies: string[];
  isHidden: boolean;
  photos: string[];
}

interface EditProfileProps {
  profileData: ProfileData;
  onBack: () => void;
  onSave: () => void;
  onInputChange: (field: string, value: string) => void;
  onArrayInputChange: (field: string, value: string) => void;
  onPhotoAdd: (file?: File) => void;
  onPhotoRemove: (index: number) => void;
}

export const EditProfile: React.FC<EditProfileProps> = ({
  profileData,
  onBack,
  onSave,
  onInputChange,
  onArrayInputChange,
  onPhotoAdd,
  onPhotoRemove
}) => {
  const fileInputRef = useRef<HTMLInputElement>(null);

  const handleCameraClick = () => {
    // Trigger the hidden file input click
    if (fileInputRef.current) {
      fileInputRef.current.click();
    }
  };

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const files = event.target.files;
    if (files && files.length > 0) {
      // Call the onPhotoAdd function with the selected file
      onPhotoAdd(files[0]);
      // Reset the file input value so the same file can be selected again
      event.target.value = '';
    }
  };

  return (
    <div className={styles.editContent}>
      {/* Hidden file input for photo selection */}
      <input
        type="file"
        ref={fileInputRef}
        style={{ display: 'none' }}
        accept="image/*"
        onChange={handleFileChange}
      />
      <div className={styles.profileHeader}>
        <div className={styles.backButton} onClick={onBack}>
          <KeyboardBackspaceIcon />
        </div>
        {/* <h5>Edit Profile</h5> */}
        <div className={styles.header}> edit profile</div>
      </div>

      <div className={styles.photoEditSection}>
        <h3>Profile Photo</h3>
        <div className={styles.profileImageEdit}>
          <img
            src={profileData.photos[0] || "/photo1.png"}
            alt="Profile"
            className={styles.profilePic}
          />
          <div className={styles.photoEditIcon} onClick={handleCameraClick}>
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
                onClick={() => onPhotoRemove(index)}
              >
                <CloseIcon />
              </div>
            </div>
          ))}
          <div className={styles.addPhotoItem} onClick={handleCameraClick}>
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
            onChange={(e) => onInputChange("name", e.target.value)}
          />
        </div>
        <div className={styles.editField}>
          <label>Age</label>
          <CustomInput
            type="number"
            placeholder="Your age"
            value={profileData.age.toString()}
            onChange={(e) => onInputChange("age", e.target.value)}
          />
        </div>
        <div className={styles.editField}>
          <label>Date of Birth</label>
          <CustomInput
            type="date"
            placeholder="Date of birth"
            value={profileData.dateOfBirth}
            onChange={(e) => onInputChange("dateOfBirth", e.target.value)}
          />
        </div>
        <div className={styles.editField}>
          <label>Phone</label>
          <CustomInput
            type="tel"
            placeholder="Your phone number"
            value={profileData.phone}
            onChange={(e) => onInputChange("phone", e.target.value)}
          />
        </div>
        <div className={styles.editField}>
          <label>Email</label>
          <CustomInput
            type="email"
            placeholder="Your email"
            value={profileData.email}
            onChange={(e) => onInputChange("email", e.target.value)}
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
            onSelect={(value) => onInputChange("gender", value)}
          />
        </div>
      </div>

      <div className={styles.editSection}>
        <h3>About Me</h3>
        <div className={styles.textareaContainer}>
          <textarea
            className={styles.textareaInput}
            placeholder="Tell something about yourself"
            value={profileData.about}
            onChange={(e) => onInputChange("about", e.target.value)}
            maxLength={500}
          />
          <div className={styles.characterCount}>
            {profileData.about.length}/500
          </div>
        </div>
      </div>

      <div className={styles.editSection}>
        <h3>Interests</h3>
        <CustomInput
          type="text"
          placeholder="Interests (comma separated)"
          value={profileData.interests.join(", ")}
          onChange={(e) => onArrayInputChange("interests", e.target.value)}
        />
      </div>

      <div className={styles.editSection}>
        <h3>Music</h3>
        <CustomInput
          type="text"
          placeholder="Favorite music (comma separated)"
          value={profileData.music.join(", ")}
          onChange={(e) => onArrayInputChange("music", e.target.value)}
        />
      </div>

      <div className={styles.editSection}>
        <h3>Movies</h3>
        <CustomInput
          type="text"
          placeholder="Favorite movies (comma separated)"
          value={profileData.movies.join(", ")}
          onChange={(e) => onArrayInputChange("movies", e.target.value)}
        />
      </div>

      <div className={styles.editSection}>
        <h3>Books</h3>
        <CustomInput
          type="text"
          placeholder="Favorite books (comma separated)"
          value={profileData.books.join(", ")}
          onChange={(e) => onArrayInputChange("books", e.target.value)}
        />
      </div>

      <div className={styles.editSection}>
        <h3>Hobbies</h3>
        <CustomInput
          type="text"
          placeholder="Hobbies (comma separated)"
          value={profileData.hobbies.join(", ")}
          onChange={(e) => onArrayInputChange("hobbies", e.target.value)}
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
            onSelect={(value) => onInputChange("isHidden", value === "hidden" ? "true" : "false")}
          />
        </div>
      </div>

      <div className={styles.saveButtonContainer}>
        <CustomButton
          text="Save Changes"
          onClick={onSave}
          className={styles.saveButton}
        />
      </div>
    </div>
  );
};

export default EditProfile;