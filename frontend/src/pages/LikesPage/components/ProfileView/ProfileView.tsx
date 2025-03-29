/**
 * ProfileView Component
 *
 * This component displays a detailed view of a user profile with:
 * - Full-screen photo view initially
 * - Ability to scroll up to see detailed profile information
 * - Common interest tags highlighted
 * - Touch gestures support for navigation
 */
import React, { useState, useRef, useEffect } from "react";
import styles from "./profileView.module.css";
import { DislikeButton } from "../../../../shared/components/DislikeButton";
import { LikeButton } from "../../../../shared/components/LikeButton";

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
  // State to track the menu position - initially expanded to show full photo
  const [menuPosition, setMenuPosition] = useState<'collapsed' | 'expanded'>('expanded');
  
  // Refs for containers to handle scrolling and touch events
  const profileInfoRef = useRef<HTMLDivElement>(null);
  const containerRef = useRef<HTMLDivElement>(null);
  
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

  // Simulate common interests (in a real app, this would be compared with the user's interests)
  const commonInterests = ["Music", "Travel", "Photography"];
  
  // Check if an interest is common
  const isCommonInterest = (interest: string) => {
    return commonInterests.includes(interest);
  };

  // State for touch gesture handling
  const [touchStart, setTouchStart] = useState<number | null>(null);
  const [touchEnd, setTouchEnd] = useState<number | null>(null);
  
  /**
   * Toggle between expanded (full photo) and collapsed (showing details) states
   */
  const toggleMenuPosition = () => {
    setMenuPosition(menuPosition === 'collapsed' ? 'expanded' : 'collapsed');
  };
  
  // No initial scroll - let the user scroll manually to see details

  /**
   * Detect scrolling to automatically collapse the menu when user scrolls up
   * This creates a natural transition from photo view to details view
   */
  useEffect(() => {
    const handleScroll = () => {
      if (containerRef.current) {
        // If user scrolls up when menu is expanded, collapse the menu
        if (menuPosition === 'expanded' && containerRef.current.scrollTop < 50) {
          setMenuPosition('collapsed');
        }
      }
    };
    
    const container = containerRef.current;
    if (container) {
      container.addEventListener('scroll', handleScroll);
    }
    
    return () => {
      if (container) {
        container.removeEventListener('scroll', handleScroll);
      }
    };
  }, [menuPosition]);
  
  /**
   * Handle click on the scroll indicator button
   * Toggles between expanded and collapsed states with smooth scrolling
   */
  const handleToggleClick = () => {
    toggleMenuPosition();
    
    // Scroll to appropriate position based on new menu state
    if (menuPosition === 'collapsed') {
      // If currently collapsed, will expand, so scroll down to show full photo
      setTimeout(() => {
        if (containerRef.current) {
          containerRef.current.scrollTo({
            top: containerRef.current.scrollHeight,
            behavior: 'smooth'
          });
        }
      }, 100);
    } else {
      // If currently expanded, will collapse, so scroll up to show details
      setTimeout(() => {
        if (containerRef.current) {
          containerRef.current.scrollTo({
            top: 0,
            behavior: 'smooth'
          });
        }
      }, 100);
    }
  };
  
  /**
   * Touch gesture handling for swipe navigation
   * These functions work together to detect and respond to swipe gestures
   */
  // Handle touch start - record initial position
  const handleTouchStart = (e: React.TouchEvent) => {
    // Store the initial touch position
    setTouchStart(e.targetTouches[0].clientY);
    setTouchEnd(null);
  };
  
  // Handle touch move - track finger movement
  const handleTouchMove = (e: React.TouchEvent) => {
    // Update the current touch position
    setTouchEnd(e.targetTouches[0].clientY);
    
    // We don't call preventDefault() here to avoid issues with passive event listeners
  };
  
  // Handle touch end - determine if it was a swipe and in which direction
  const handleTouchEnd = () => {
    if (!touchStart || !touchEnd) return;
    
    const distance = touchStart - touchEnd;
    const isDownSwipe = distance < -30; // More sensitive - negative means swiping down
    const isUpSwipe = distance > 30; // More sensitive - positive means swiping up
    
    // If we're in collapsed state and swiping down, expand to show full photo
    if (isDownSwipe && menuPosition === 'collapsed') {
      // Expand the menu
      setMenuPosition('expanded');
      
      // Scroll to show the full photo
      setTimeout(() => {
        if (containerRef.current) {
          containerRef.current.scrollTo({
            top: containerRef.current.scrollHeight,
            behavior: 'smooth'
          });
        }
      }, 50); // Faster timeout
    }
    // If we're in expanded state and swiping up, collapse to show details
    else if (isUpSwipe && menuPosition === 'expanded') {
      // Collapse the menu
      setMenuPosition('collapsed');
      
      // Scroll back to the top
      setTimeout(() => {
        if (containerRef.current) {
          containerRef.current.scrollTo({
            top: 0,
            behavior: 'smooth'
          });
        }
      }, 50); // Faster timeout
    }
    
    // Reset touch values
    setTouchStart(null);
    setTouchEnd(null);
  };

  return (
    <div className={styles.profileViewContainer} ref={containerRef}>
      {/* Full-screen photo view with overlay content */}
      <div
        className={`${styles.profileImageContainer} ${menuPosition === 'expanded' ? styles.expanded : ''}`}
      >
        {/* Back button in top-left corner */}
        <button className={styles.photoBackButton} onClick={onClose}>
        <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
<path d="M19 12H5" stroke="white" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
<path d="M12 19L5 12L12 5" stroke="white" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
</svg>

          {/* <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M15 18L9 12L15 6" stroke="white" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"/>
          </svg> */}
        </button>
        
        {/* Main profile photo */}
        <img
          src={profile.imageUrl}
          alt={profile.name}
          className={styles.profileImage}
        />
        {/* Gradient overlay for better text visibility */}
        <div className={styles.profileGradient}></div>
        
        {/* Content displayed on top of the photo */}
        <div className={styles.photoContent}>
          {/* User name and basic info */}
          <h2 className={styles.photoName}>{profile.name}</h2>
          <p className={styles.photoInfo}>
            {profile.age} | {profileDetails.location}
          </p>
          
          {/* Interest tags on photo */}
          <div className={styles.photoTags}>
            {profileDetails.passion.slice(0, 5).map((interest, index) => (
              <span key={index} className={styles.photoTag}>
                {interest}
              </span>
            ))}
          </div>
          
          {/* Scroll indicator at bottom of photo */}
          <div className={styles.photoScrollIndicator} onClick={handleToggleClick}>
              Scroll up to view details
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" style={{ transform: 'rotate(180deg)' }}>
                <path d="M7 10L12 15L17 10" stroke="white" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"/>
              </svg>
            </div>
        </div>
        
        {/* Action buttons (close and like) */}
        {/* <div className={styles.photoButtons}>
          <button className={styles.photoButton} onClick={onClose}>
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M18 6L6 18M6 6L18 18" stroke="black" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"/>
            </svg>
          </button>
          <button className={styles.photoButton} onClick={() => onLike(profile.id)}>
            <svg width="20" height="18" viewBox="0 0 26 22" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path
                d="M23.4175 2.3871C22.8188 1.78818 22.108 1.31307 21.3257 0.988918C20.5434 0.664766 19.7049 0.497925 18.8581 0.497925C18.0113 0.497925 17.1728 0.664766 16.3905 0.988918C15.6082 1.31307 14.8974 1.78818 14.2987 2.3871L13.0563 3.6295L11.8139 2.3871C10.6047 1.17788 8.96466 0.498552 7.25456 0.498552C5.54447 0.498552 3.90441 1.17788 2.69519 2.3871C1.48597 3.59632 0.806641 5.23638 0.806641 6.94647C0.806641 8.65657 1.48597 10.2966 2.69519 11.5058L3.93759 12.7482L13.0563 21.867L22.1751 12.7482L23.4175 11.5058C24.0164 10.9072 24.4915 10.1964 24.8156 9.4141C25.1398 8.63179 25.3066 7.79328 25.3066 6.94647C25.3066 6.09966 25.1398 5.26115 24.8156 4.47884C24.4915 3.69653 24.0164 2.98575 23.4175 2.3871Z"
                fill="#FF4B91"
              />
            </svg>
          </button>
        </div> */}
        {/* <LikeButton onClick={() => onLike(profile.id)} />
        <DislikeButton onClick={() => onDislike(profile.id)} /> */}
        <LikeButton onClick={() => onLike(profile.id)} className={styles.likeButton} />
        <DislikeButton onClick={() => onLike(profile.id)} className={styles.dislikeButton} />
      </div>

      {/* Detailed profile information section */}
      <div
        ref={profileInfoRef}
        className={`${styles.profileInfo} ${menuPosition === 'expanded' ? styles.expanded : styles.collapsed}`}
        onTouchStart={handleTouchStart}
        onTouchMove={handleTouchMove}
        onTouchEnd={handleTouchEnd}
      >
        {/* Scroll indicator button */}
        <div className={styles.scrollIndicator} onClick={handleToggleClick}>
          <span>{menuPosition === 'collapsed' ? 'Scroll down to view more' : 'Scroll up to view details'}</span>
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
        {/* Profile header with name and action buttons */}
        <div className={styles.profileNameContainer}>
          <div>
            <h2 className={styles.profileName}>{profile.name}</h2>
            <p className={styles.profileAge}>
              {profile.age} | {profileDetails.location}
            </p>
          </div>
          {/* Action buttons in the details view */}
          {/* <div className={styles.actionButtons}>
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
          </div> */}
        </div>

        {/* Interest tags with common interests highlighted */}
        <div className={styles.interestTags}>
          {profileDetails.passion.slice(0, 5).map((interest, index) => (
            <span
              key={index}
              className={`${styles.interestTag} ${isCommonInterest(interest) ? styles.commonInterest : ''}`}
            >
              {/* "Common" badge appears at the top of common interest tags */}
              {isCommonInterest(interest) && <span className={styles.commonBadge}>Common</span>}
              {interest}
            </span>
          ))}
        </div>

        {/* Detailed profile information sections */}
        <div className={styles.detailedInfo}>
          {/* Bio section */}
          <section className={styles.infoSection}>
            <h3>Introduce</h3>
            <p>{profileDetails.bio}</p>
          </section>

          {/* Lifestyle section */}
          <section className={styles.infoSection}>
            <h3>Lifestyle</h3>
            {Object.values(profileDetails.lifestyle).map((value, index) => (
              <p key={index}>{value}</p>
            ))}
          </section>

          {/* Passion/interests section with common interests highlighted */}
          <section className={styles.infoSection}>
            <h3>Passion</h3>
            <div className={styles.interestTags}>
              {profileDetails.passion.slice(0, 5).map((interest, index) => (
                <span
                  key={index}
                  className={`${styles.interestTag} ${isCommonInterest(interest) ? styles.commonInterest : ''}`}
                >
                  {isCommonInterest(interest) && <span className={styles.commonBadge}>Common</span>}
                  {interest}
                </span>
              ))}
            </div>
          </section>

          {/* Photo gallery section */}
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