// /**
//  * ProfileView Component
//  *
//  * This component displays a detailed view of a user profile with:
//  * - Full-screen photo view initially
//  * - Ability to scroll up to see detailed profile information
//  * - Common interest tags highlighted
//  * - Touch gestures support for navigation
//  */
// import React, { useState, useRef, useEffect } from 'react';

// import { BackButton } from '../../shared/components/BackButton';
// import { DislikeButton } from '../../shared/components/DislikeButton';
// import { LikeButton } from '../../shared/components/LikeButton';
// import { ProfileData } from '../../shared/components/ProfileCard';

// import ProfileInfo from './components/ProfileInfo';
// import styles from './profileView.module.css';

// export interface ProfileViewProps {
//   profile: ProfileData;
//   onClose: () => void;
//   onLike: (id: number) => void;
// }

// export const ProfileView: React.FC<ProfileViewProps> = ({ profile, onClose, onLike }) => {
//   // State to track the menu position - initially expanded to show full photo
//   const [menuPosition, setMenuPosition] = useState<'collapsed' | 'expanded'>('expanded');

//   // Refs for containers to handle scrolling and touch events
//   const profileInfoRef = useRef<HTMLDivElement>(null);
//   const containerRef = useRef<HTMLDivElement>(null);

//   // Sample photos for the profile
//   const samplePhotos = [
//     profile.user.imageUrl, // Use the profile image as the first photo
//     'woman1.jpg',
//     'man1.jpg',
//     'photo1.png',
//     'woman1.jpg',
//     'man1.jpg',
//   ];

//   // Mock data for the profile details
//   const profileDetails = {
//     bio: profile.user.bio || 'Hello, I am a fashion designer based in Florida.',
//     location: profile.user.location || 'Miami Beach, Florida',
//     lifestyle: profile.user.lifestyle || {
//       kids: "I don't have kids",
//     },
//     passion: profile.user.passion ||
//       profile.user.interests || [
//         'Music',
//         'Travel',
//         'Tea',
//         'Photography',
//         'Fashion',
//         'House Parties',
//       ],
//     photos: profile.user.photos || samplePhotos,
//   };

//   // Simulate common interests (in a real app, this would be compared with the user's interests)
//   const commonInterests = ['Music', 'Travel', 'Photography'];

//   // Check if an interest is common
//   const isCommonInterest = (interest: string) => {
//     return commonInterests.includes(interest);
//   };

//   // State for touch gesture handling
//   const [touchStart, setTouchStart] = useState<number | null>(null);
//   const [touchEnd, setTouchEnd] = useState<number | null>(null);

//   /**
//    * Toggle between expanded (full photo) and collapsed (showing details) states
//    */
//   const toggleMenuPosition = () => {
//     setMenuPosition(menuPosition === 'collapsed' ? 'expanded' : 'collapsed');
//   };

//   // No initial scroll - let the user scroll manually to see details

//   /**
//    * Detect scrolling to automatically collapse the menu when user scrolls down
//    * and expand it (hide info) when user scrolls back up
//    * This creates a natural transition between photo view and details view
//    */
//   useEffect(() => {
//     const handleScroll = () => {
//       if (containerRef.current) {
//         const scrollTop = containerRef.current.scrollTop;

//         // If user scrolls down when menu is expanded, collapse the menu to show info
//         if (menuPosition === 'expanded' && scrollTop > 50) {
//           setMenuPosition('collapsed');
//         }

//         // If user scrolls back up when menu is collapsed, expand the menu to hide info
//         if (menuPosition === 'collapsed' && scrollTop < 20) {
//           setMenuPosition('expanded');
//         }
//       }
//     };

//     const container = containerRef.current;
//     if (container) {
//       container.addEventListener('scroll', handleScroll);
//     }

//     return () => {
//       if (container) {
//         container.removeEventListener('scroll', handleScroll);
//       }
//     };
//   }, [menuPosition]);

//   /**
//    * Handle click on the scroll indicator button
//    * Toggles between expanded and collapsed states with smooth scrolling
//    */
//   const handleToggleClick = () => {
//     toggleMenuPosition();

//     // Scroll to appropriate position based on new menu state
//     if (menuPosition === 'collapsed') {
//       // If currently collapsed, will expand, so scroll down to show full photo and hide info
//       setTimeout(() => {
//         if (containerRef.current) {
//           containerRef.current.scrollTo({
//             top: containerRef.current.scrollHeight,
//             behavior: 'smooth',
//           });
//         }
//       }, 100);
//     } else {
//       // If currently expanded, will collapse, so scroll up to show details
//       setTimeout(() => {
//         if (containerRef.current) {
//           containerRef.current.scrollTo({
//             top: 0,
//             behavior: 'smooth',
//           });
//         }
//       }, 100);
//     }
//   };

//   /**
//    * Touch gesture handling for swipe navigation
//    * These functions work together to detect and respond to swipe gestures
//    */
//   // Handle touch start - record initial position
//   const handleTouchStart = (e: React.TouchEvent) => {
//     // Store the initial touch position
//     setTouchStart(e.targetTouches[0].clientY);
//     setTouchEnd(null);
//   };

//   // Handle touch move - track finger movement
//   const handleTouchMove = (e: React.TouchEvent) => {
//     // Update the current touch position
//     setTouchEnd(e.targetTouches[0].clientY);

//     // We don't call preventDefault() here to avoid issues with passive event listeners
//   };

//   // Handle touch end - determine if it was a swipe and in which direction
//   const handleTouchEnd = () => {
//     if (!touchStart || !touchEnd) return;

//     const distance = touchStart - touchEnd;
//     const isDownSwipe = distance < -30; // More sensitive - negative means swiping down
//     const isUpSwipe = distance > 30; // More sensitive - positive means swiping up

//     // If we're in collapsed state and swiping down, expand to show full photo and hide info
//     if (isDownSwipe && menuPosition === 'collapsed') {
//       // Expand the menu
//       setMenuPosition('expanded');

//       // Scroll to show the full photo
//       setTimeout(() => {
//         if (containerRef.current) {
//           containerRef.current.scrollTo({
//             top: containerRef.current.scrollHeight,
//             behavior: 'smooth',
//           });
//         }
//       }, 50); // Faster timeout
//     }
//     // If we're in expanded state and swiping up, collapse to show details
//     else if (isUpSwipe && menuPosition === 'expanded') {
//       // Collapse the menu
//       setMenuPosition('collapsed');

//       // Scroll back to the top
//       setTimeout(() => {
//         if (containerRef.current) {
//           containerRef.current.scrollTo({
//             top: 0,
//             behavior: 'smooth',
//           });
//         }
//       }, 50); // Faster timeout
//     }

//     // Reset touch values
//     setTouchStart(null);
//     setTouchEnd(null);
//   };

//   return (
//     <div className={styles.profileViewContainer} ref={containerRef}>
//       {/* Full-screen photo view with overlay content */}
//       <div
//         className={`${styles.profileImageContainer} ${
//           menuPosition === 'expanded' ? styles.expanded : ''
//         }`}
//       >
//         <BackButton onClose={onClose} />
//         {/* Main profile photo */}
//         <img src={profile.user.imageUrl} alt={profile.user.name} className={styles.profileImage} />
//         {/* Gradient overlay for better text visibility */}
//         <div className={styles.profileGradient} />

//         {/* Content displayed on top of the photo */}
//         <div className={styles.photoContent}>
//           {/* User name and basic info */}
//           <h2 className={styles.photoName}>{profile.user.name}</h2>
//           <p className={styles.photoInfo}>
//             {profile.user.age} | {profileDetails.location}
//           </p>

//           {/* Interest tags on photo */}
//           <div className={styles.photoTags}>
//             {profileDetails.passion.slice(0, 5).map((interest, index) => (
//               <span key={index} className={styles.photoTag}>
//                 {interest}
//               </span>
//             ))}
//           </div>

//           {/* Scroll indicator at bottom of photo - always visible */}
//           <div
//             className={`${styles.photoScrollIndicator} ${
//               menuPosition === 'collapsed' ? styles.showScrollBack : ''
//             }`}
//             onClick={handleToggleClick}
//           >
//             {menuPosition === 'collapsed' ? 'Scroll back to photo' : 'Scroll up to view details'}
//             <svg
//               width="16"
//               height="16"
//               viewBox="0 0 24 24"
//               fill="none"
//               xmlns="http://www.w3.org/2000/svg"
//               style={{
//                 transform: menuPosition === 'collapsed' ? 'rotate(180deg)' : 'none',
//               }}
//             >
//               <path
//                 d="M7 10L12 15L17 10"
//                 stroke="white"
//                 strokeWidth="2"
//                 strokeLinecap="round"
//                 strokeLinejoin="round"
//               />
//             </svg>
//           </div>
//         </div>
//         <div onClick={e => e.stopPropagation()}>
//           <LikeButton
//             onClick={() => {
//               // Call the onLike function to trigger the match modal
//               onLike(profile.user.id);
//             }}
//             className={styles.likeButton}
//           />
//         </div>
//         <div onClick={e => e.stopPropagation()}>
//           <DislikeButton
//             onClick={() => {
//               // Close the ProfileView and return to the likes page
//               onClose();
//             }}
//             className={styles.dislikeButton}
//           />
//         </div>
//       </div>

//       {/* Detailed profile information section */}
//       <ProfileInfo
//         profile={profile}
//         profileDetails={profileDetails}
//         menuPosition={menuPosition}
//         handleToggleClick={handleToggleClick}
//         handleTouchStart={handleTouchStart}
//         handleTouchMove={handleTouchMove}
//         handleTouchEnd={handleTouchEnd}
//         isCommonInterest={isCommonInterest}
//         profileInfoRef={profileInfoRef}
//       />
//     </div>
//   );
// };

// export default ProfileView;

/**
 * ProfileView Component
 *
 * This component displays a detailed view of a user profile with:
 * - Full-screen photo view initially
 * - Ability to scroll up to see detailed profile information
 * - Common interest tags highlighted
 * - Touch gestures support for navigation
 */
import VerifiedIcon from '@mui/icons-material/Verified';
import React, { useState, useRef, useEffect } from 'react';

// import { BackButton } from '../../../../shared/components/BackButton';
// import { DislikeButton } from '../../../../shared/components/DislikeButton';
// import { LikeButton } from '../../../../shared/components/LikeButton';

// import { ProfileInfo } from './components/ProfileInfo';
import { BackButton } from '../../shared/components/BackButton';
import { DislikeButton } from '../../shared/components/DislikeButton';
import { LikeButton } from '../../shared/components/LikeButton';

import ProfileInfo from './components/ProfileInfo';
import styles from './profileView.module.css';

interface ProfileViewProps {
  profile: {
    id: number;
    name: string;
    age: number;
    imageUrl: string;
    hasLikedYou?: boolean;
    verified?: boolean;
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

export const ProfileView: React.FC<ProfileViewProps> = ({ profile, onClose, onLike }) => {
  // State to track the menu position - initially expanded to show full photo
  const [menuPosition, setMenuPosition] = useState<'collapsed' | 'expanded'>('expanded');

  // Refs for containers to handle scrolling and touch events
  const profileInfoRef = useRef<HTMLDivElement>(null);
  const containerRef = useRef<HTMLDivElement>(null);

  // Sample photos for the profile
  const samplePhotos = [
    profile.imageUrl, // Use the profile image as the first photo
    'woman1.jpg',
    'man1.jpg',
    'photo1.png',
    'woman1.jpg',
    'man1.jpg',
  ];

  // Mock data for the profile details
  const profileDetails = {
    bio: profile.bio || 'Hello, I am a fashion designer based in Florida.',
    location: profile.location || 'Miami Beach, Florida',
    lifestyle: profile.lifestyle || {
      kids: "I don't have kids",
    },
    passion: profile.passion ||
      profile.interests || ['Music', 'Travel', 'Tea', 'Photography', 'Fashion', 'House Parties'],
    photos: profile.photos || samplePhotos,
  };

  // Simulate common interests (in a real app, this would be compared with the user's interests)
  const commonInterests = ['Music', 'Travel', 'Photography'];

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
   * Detect scrolling to automatically collapse the menu when user scrolls down
   * and expand it (hide info) when user scrolls back up
   * This creates a natural transition between photo view and details view
   */
  useEffect(() => {
    const handleScroll = () => {
      if (containerRef.current) {
        const scrollTop = containerRef.current.scrollTop;

        // If user scrolls down when menu is expanded, collapse the menu to show info
        if (menuPosition === 'expanded' && scrollTop > 50) {
          setMenuPosition('collapsed');
        }

        // If user scrolls back up when menu is collapsed, expand the menu to hide info
        if (menuPosition === 'collapsed' && scrollTop < 20) {
          setMenuPosition('expanded');
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
      // If currently collapsed, will expand, so scroll down to show full photo and hide info
      setTimeout(() => {
        if (containerRef.current) {
          containerRef.current.scrollTo({
            top: containerRef.current.scrollHeight,
            behavior: 'smooth',
          });
        }
      }, 100);
    } else {
      // If currently expanded, will collapse, so scroll up to show details
      setTimeout(() => {
        if (containerRef.current) {
          containerRef.current.scrollTo({
            top: 0,
            behavior: 'smooth',
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

    // If we're in collapsed state and swiping down, expand to show full photo and hide info
    if (isDownSwipe && menuPosition === 'collapsed') {
      // Expand the menu
      setMenuPosition('expanded');

      // Scroll to show the full photo
      setTimeout(() => {
        if (containerRef.current) {
          containerRef.current.scrollTo({
            top: containerRef.current.scrollHeight,
            behavior: 'smooth',
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
            behavior: 'smooth',
          });
        }
      }, 50); // Faster timeout
    }

    // Reset touch values
    setTouchStart(null);
    setTouchEnd(null);
  };

  return (
    <div
      className={styles.profileViewContainer}
      ref={containerRef}
      onClick={e => e.stopPropagation()} // Prevent clicks from propagating
    >
      {/* Full-screen photo view with overlay content */}
      <div
        className={`${styles.profileImageContainer} ${
          menuPosition === 'expanded' ? styles.expanded : ''
        }`}
      >
        <BackButton
          onClose={() => {
            onClose();
            // Ensure the profile view is completely closed
            setTimeout(() => {
              document.body.style.overflow = 'auto'; // Re-enable scrolling
            }, 100);
          }}
        />
        {/* Main profile photo */}
        <img src={profile.imageUrl} alt={profile.name} className={styles.profileImage} />
        {/* Gradient overlay for better text visibility */}
        <div className={styles.profileGradient} />

        {/* Content displayed on top of the photo */}
        <div className={styles.photoContent}>
          {/* User name and basic info */}
          <h2 className={styles.photoName}>
            {profile.name}
            {profile.verified && (
              <div className={styles.verifiedBadge}>
                <VerifiedIcon />
              </div>
            )}
          </h2>
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

          {/* Scroll indicator at bottom of photo - always visible */}
          <div
            className={`${styles.photoScrollIndicator} ${
              menuPosition === 'collapsed' ? styles.showScrollBack : ''
            }`}
            onClick={handleToggleClick}
          >
            {menuPosition === 'collapsed' ? 'Scroll back to photo' : 'Scroll up to view details'}
            <svg
              width="16"
              height="16"
              viewBox="0 0 24 24"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
              style={{
                transform: menuPosition === 'collapsed' ? 'rotate(180deg)' : 'none',
              }}
            >
              <path
                d="M7 10L12 15L17 10"
                stroke="white"
                strokeWidth="2"
                strokeLinecap="round"
                strokeLinejoin="round"
              />
            </svg>
          </div>
        </div>
        <div onClick={e => e.stopPropagation()}>
          <LikeButton
            onClick={() => {
              // Call the onLike function to trigger the match modal
              onLike(profile.id);
            }}
            className={styles.likeButton}
          />
        </div>
        <div onClick={e => e.stopPropagation()}>
          <DislikeButton
            onClick={() => {
              // Close the ProfileView and return to the likes page
              onClose();
            }}
            className={styles.dislikeButton}
          />
        </div>
      </div>

      {/* Detailed profile information section */}
      <ProfileInfo
        profile={profile}
        profileDetails={profileDetails}
        menuPosition={menuPosition}
        handleToggleClick={handleToggleClick}
        handleTouchStart={handleTouchStart}
        handleTouchMove={handleTouchMove}
        handleTouchEnd={handleTouchEnd}
        isCommonInterest={isCommonInterest}
        profileInfoRef={profileInfoRef}
      />
    </div>
  );
};

export default ProfileView;
