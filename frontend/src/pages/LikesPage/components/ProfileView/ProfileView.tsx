import React, { useState, useRef } from "react";
import styles from "./profileView.module.css";

interface ProfileViewProps {
  profile: {
    id: number;
    name: string;
    age: number;
    imageUrl: string;
    hasLikedYou?: boolean;
    bio?: string;
    location?: string;
    interests?: string[];
    lifestyle?: { [key: string]: string };
    passion?: string[];
    photos?: string[];
  };
  onClose: () => void;
  onLike: (id: number) => void;
}

export const ProfileView: React.FC<ProfileViewProps> = ({
  profile,
  onClose,
  onLike,
}) => {
  // State to track the menu position
  const [menuPosition, setMenuPosition] = useState<'collapsed' | 'expanded'>('collapsed');
  
  // Refs for containers
  const profileInfoRef = useRef<HTMLDivElement>(null);
  
  // Sample photos for the profile
  const samplePhotos = [
    profile.imageUrl, // Use the profile image as the first photo
    "woman1.jpg",
    "man1.jpg",
    "photo1.png",
    "woman1.jpg",
    "man1.jpg",
  ];

  // Mock data for the profile details
  const profileDetails = {
    bio: profile.bio || "Hello, I am a fashion designer based in Florida.",
    location: profile.location || "Miami Beach, Florida",
    lifestyle: profile.lifestyle || {
      kids: "I don't have kids",
    },
    passion: profile.passion || profile.interests || ["Music", "Travel", "Tea", "Photography", "Fashion", "House Parties"],
    photos: profile.photos || samplePhotos,
  };

  // Toggle menu position
  const toggleMenuPosition = () => {
    setMenuPosition(menuPosition === 'collapsed' ? 'expanded' : 'collapsed');
  };

  return (
    <div className={styles.profileViewContainer}>
      <div className={styles.header}>
        <button className={styles.backButton} onClick={onClose}>
          <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M15 18L9 12L15 6" stroke="white" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"/>
          </svg>
        </button>
      </div>

      <div
        className={`${styles.profileImageContainer} ${menuPosition === 'expanded' ? styles.expanded : ''}`}
      >
        <img
          src={profile.imageUrl}
          alt={profile.name}
          className={styles.profileImage}
        />
        <div className={styles.profileGradient}></div>
      </div>

      <div
        ref={profileInfoRef}
        className={`${styles.profileInfo} ${menuPosition === 'expanded' ? styles.expanded : ''}`}
      >
        <div className={styles.scrollIndicator} onClick={toggleMenuPosition}>
          <span>{menuPosition === 'collapsed' ? 'Scroll down to view more' : 'Scroll up to collapse'}</span>
          <svg
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
            style={{ transform: menuPosition === 'expanded' ? 'rotate(180deg)' : 'none' }}
          >
            <path d="M7 10L12 15L17 10" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"/>
          </svg>
        </div>
        <div className={styles.profileNameContainer}>
          <div>
            <h2 className={styles.profileName}>{profile.name}</h2>
            <p className={styles.profileAge}>
              {profile.age} | {profileDetails.location}
            </p>
          </div>
          <div className={styles.actionButtons}>
            <button className={styles.closeButton} onClick={onClose}>
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M18 6L6 18M6 6L18 18" stroke="black" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"/>
              </svg>
            </button>
            <button 
              className={styles.likeButton} 
              onClick={() => onLike(profile.id)}
            >
              <svg width="20" height="18" viewBox="0 0 26 22" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path
                  d="M23.4175 2.3871C22.8188 1.78818 22.108 1.31307 21.3257 0.988918C20.5434 0.664766 19.7049 0.497925 18.8581 0.497925C18.0113 0.497925 17.1728 0.664766 16.3905 0.988918C15.6082 1.31307 14.8974 1.78818 14.2987 2.3871L13.0563 3.6295L11.8139 2.3871C10.6047 1.17788 8.96466 0.498552 7.25456 0.498552C5.54447 0.498552 3.90441 1.17788 2.69519 2.3871C1.48597 3.59632 0.806641 5.23638 0.806641 6.94647C0.806641 8.65657 1.48597 10.2966 2.69519 11.5058L3.93759 12.7482L13.0563 21.867L22.1751 12.7482L23.4175 11.5058C24.0164 10.9072 24.4915 10.1964 24.8156 9.4141C25.1398 8.63179 25.3066 7.79328 25.3066 6.94647C25.3066 6.09966 25.1398 5.26115 24.8156 4.47884C24.4915 3.69653 24.0164 2.98575 23.4175 2.3871Z"
                  fill="#FF4B91"
                />
              </svg>
            </button>
          </div>
        </div>

        <div className={styles.interestTags}>
          {profileDetails.passion.slice(0, 5).map((interest, index) => (
            <span key={index} className={styles.interestTag}>
              {interest}
            </span>
          ))}
        </div>

        <div className={styles.detailedInfo}>
          <section className={styles.infoSection}>
            <h3>Introduce</h3>
            <p>{profileDetails.bio}</p>
          </section>

          <section className={styles.infoSection}>
            <h3>Lifestyle</h3>
            {Object.values(profileDetails.lifestyle).map((value, index) => (
              <p key={index}>{value}</p>
            ))}
          </section>

          <section className={styles.infoSection}>
            <h3>Passion</h3>
            <div className={styles.interestTags}>
              {profileDetails.passion.slice(0, 5).map((interest, index) => (
                <span key={index} className={styles.interestTag}>
                  {interest}
                </span>
              ))}
            </div>
          </section>

          <section className={styles.infoSection}>
            <h3>Photos</h3>
            <div className={styles.photosGrid}>
              {samplePhotos.map((photo, index) => (
                <div key={index} className={styles.photoItem}>
                  <img src={photo} alt={`Photo ${index + 1}`} />
                </div>
              ))}
            </div>
          </section>
        </div>
      </div>
    </div>
  );
};

export default ProfileView;