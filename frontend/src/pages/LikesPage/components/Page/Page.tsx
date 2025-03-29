import React, { useState } from "react";
import { NavBar } from "../../../../shared/NavBar";
import { MatchModal } from "../MatchModal";
import { ProfileView } from "../ProfileView";
import { useNavigate } from "react-router-dom";
import styles from "./likesPage.module.css";

// Define types for our data
interface LikeProfile {
  id: number;
  name: string;
  age: number;
  imageUrl: string;
  liked?: boolean;
  hasLikedYou?: boolean;
  bio?: string;
  location?: string;
  interests?: string[];
  lifestyle?: { [key: string]: string };
  passion?: string[];
}

interface MatchProfile {
  id: number;
  name: string;
  age: number;
  imageUrl: string;
  telegram: string;
  instagram: string;
}

// SVG icons for social media
const TelegramIcon = () => (
  <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
    <path d="M9.78 18.65L10.06 14.65L17.76 7.65C18.1 7.34 17.71 7.18 17.28 7.45L7.86 13.35L3.95 12.05C3.05 11.79 3.04 11.16 4.15 10.7L19.92 4.55C20.69 4.24 21.41 4.76 21.15 5.98L18.31 18.65C18.13 19.65 17.58 19.87 16.8 19.45L12.6 16.32L10.6 18.25C10.38 18.47 10.2 18.65 9.78 18.65Z" fill="#2AABEE"/>
  </svg>
);

const InstagramIcon = () => (
  <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
    <path d="M12 2.16C15.2 2.16 15.58 2.17 16.82 2.22C18.63 2.31 19.69 2.78 20.35 3.09C21.2 3.49 21.81 3.98 22.44 4.6C23.07 5.23 23.56 5.84 23.96 6.69C24.27 7.35 24.74 8.41 24.83 10.22C24.88 11.46 24.89 11.84 24.89 15.04C24.89 18.24 24.88 18.62 24.83 19.86C24.74 21.67 24.27 22.73 23.96 23.39C23.56 24.24 23.07 24.85 22.44 25.48C21.81 26.11 21.2 26.6 20.35 27C19.69 27.31 18.63 27.78 16.82 27.87C15.58 27.92 15.2 27.93 12 27.93C8.8 27.93 8.42 27.92 7.18 27.87C5.37 27.78 4.31 27.31 3.65 27C2.8 26.6 2.19 26.11 1.56 25.48C0.93 24.85 0.44 24.24 0.04 23.39C-0.27 22.73 -0.74 21.67 -0.83 19.86C-0.88 18.62 -0.89 18.24 -0.89 15.04C-0.89 11.84 -0.88 11.46 -0.83 10.22C-0.74 8.41 -0.27 7.35 0.04 6.69C0.44 5.84 0.93 5.23 1.56 4.6C2.19 3.97 2.8 3.48 3.65 3.08C4.31 2.77 5.37 2.3 7.18 2.21C8.42 2.16 8.8 2.15 12 2.15V2.16ZM12 0C8.74 0 8.33 0.01 7.07 0.06C5.26 0.15 4.46 0.39 3.77 0.69C3.05 1.01 2.46 1.42 1.87 2.01C1.28 2.6 0.87 3.19 0.55 3.91C0.25 4.6 0.01 5.4 -0.08 7.21C-0.13 8.47 -0.14 8.88 -0.14 12.14C-0.14 15.4 -0.13 15.81 -0.08 17.07C0.01 18.88 0.25 19.68 0.55 20.37C0.87 21.09 1.28 21.68 1.87 22.27C2.46 22.86 3.05 23.27 3.77 23.59C4.46 23.89 5.26 24.13 7.07 24.22C8.33 24.27 8.74 24.28 12 24.28C15.26 24.28 15.67 24.27 16.93 24.22C18.74 24.13 19.54 23.89 20.23 23.59C20.95 23.27 21.54 22.86 22.13 22.27C22.72 21.68 23.13 21.09 23.45 20.37C23.75 19.68 23.99 18.88 24.08 17.07C24.13 15.81 24.14 15.4 24.14 12.14C24.14 8.88 24.13 8.47 24.08 7.21C23.99 5.4 23.75 4.6 23.45 3.91C23.13 3.19 22.72 2.6 22.13 2.01C21.54 1.42 20.95 1.01 20.23 0.69C19.54 0.39 18.74 0.15 16.93 0.06C15.67 0.01 15.26 0 12 0Z" fill="#E4405F"/>
    <path d="M12 5.89C8.6 5.89 5.84 8.65 5.84 12.05C5.84 15.45 8.6 18.21 12 18.21C15.4 18.21 18.16 15.45 18.16 12.05C18.16 8.65 15.4 5.89 12 5.89ZM12 16.05C9.79 16.05 8 14.26 8 12.05C8 9.84 9.79 8.05 12 8.05C14.21 8.05 16 9.84 16 12.05C16 14.26 14.21 16.05 12 16.05Z" fill="#E4405F"/>
    <path d="M19.85 5.65C19.85 6.45 19.2 7.1 18.4 7.1C17.6 7.1 16.95 6.45 16.95 5.65C16.95 4.85 17.6 4.2 18.4 4.2C19.2 4.2 19.85 4.85 19.85 5.65Z" fill="#E4405F"/>
  </svg>
);

// Mock data for likes
const likesData: LikeProfile[] = [
  {
    id: 1,
    name: "Анна",
    age: 25,
    imageUrl: "woman1.jpg",
    liked: false,
    hasLikedYou: true, // This profile has already liked the user
    location: "Москва, Россия",
    interests: ["Музыка", "Путешествия", "Фотография", "Мода", "Искусство"],
  },
  {
    id: 2,
    name: "Иван",
    age: 30,
    imageUrl: "man1.jpg",
    liked: false,
    hasLikedYou: true, // This profile has already liked the user
    location: "Санкт-Петербург, Россия",
    interests: ["Спорт", "Кино", "Технологии", "Путешествия"],
  },
  {
    id: 3,
    name: "Ольга",
    age: 28,
    imageUrl: "photo1.png",
    liked: false,
    hasLikedYou: false,
    location: "Казань, Россия",
    interests: ["Книги", "Йога", "Кулинария", "Природа"],
  },
  {
    id: 4,
    name: "Алексей",
    age: 32,
    imageUrl: "man1.jpg",
    liked: false,
    hasLikedYou: false,
    location: "Екатеринбург, Россия",
    interests: ["Музыка", "Горы", "Фотография", "Путешествия"],
  },
];

// Mock data for matches
const matchesData: MatchProfile[] = [
  {
    id: 101, // Using different ID range for matches
    name: "Мария",
    age: 27,
    imageUrl: "woman1.jpg",
    telegram: "@maria_27",
    instagram: "@maria_insta",
  },
  {
    id: 102, // Using different ID range for matches
    name: "Дмитрий",
    age: 31,
    imageUrl: "man1.jpg",
    telegram: "@dmitry_31",
    instagram: "@dmitry_insta",
  },
];

// Helper function to generate a unique ID for new matches
const generateMatchId = (existingMatches: MatchProfile[]): number => {
  const maxId = existingMatches.reduce((max: number, match: MatchProfile) => 
    Math.max(max, match.id), 100);
  return maxId + 1;
};

export const LikesPage = () => {
  const navigate = useNavigate();
  const [activeTab, setActiveTab] = useState<"likes" | "matches">("likes");
  const [likes, setLikes] = useState<LikeProfile[]>(likesData);
  const [matches, setMatches] = useState<MatchProfile[]>(matchesData);
  const [showMatchModal, setShowMatchModal] = useState(false);
  const [matchedProfile, setMatchedProfile] = useState<LikeProfile | null>(null);
  const [selectedProfile, setSelectedProfile] = useState<LikeProfile | null>(null);
  const [showChatMessage, setShowChatMessage] = useState<number | null>(null);

  const handleLike = (id: number) => {
    // Find the profile that was liked
    const likedProfile = likes.find((like) => like.id === id);
    
    if (likedProfile) {
      // Check if this profile has already liked the user
      if (likedProfile.hasLikedYou) {
        // It's a match! Show the match modal
        setMatchedProfile(likedProfile);
        setShowMatchModal(true);
      }
      
      // Update the liked status
      const updatedLikes = likes.map((like) =>
        like.id === id ? { ...like, liked: true } : like
      );
      
      // Create a new match object with a unique ID
      const newMatch: MatchProfile = {
        id: generateMatchId(matches), // Generate a unique ID
        name: likedProfile.name,
        age: likedProfile.age,
        imageUrl: likedProfile.imageUrl,
        telegram: `@${likedProfile.name.toLowerCase()}_${likedProfile.age}`,
        instagram: `@${likedProfile.name.toLowerCase()}_insta`,
      };
      
      // Add to matches and remove from likes if it was a match
      if (likedProfile.hasLikedYou) {
        setMatches([...matches, newMatch]);
        setLikes(updatedLikes.filter((like) => like.id !== id));
      } else {
        setLikes(updatedLikes);
      }
    }
  };

  const handleProfileClick = (profile: LikeProfile) => {
    setSelectedProfile(profile);
  };

  const handleMatchClick = (match: MatchProfile) => {
    // Convert match to a profile format that ProfileView can use
    const matchAsProfile = {
      id: match.id,
      name: match.name,
      age: match.age,
      imageUrl: match.imageUrl,
      hasLikedYou: true,
      bio: "This is a match! You can contact them via social media.",
      location: "",
      interests: [],
      lifestyle: {
        "contact": `Telegram: ${match.telegram}`,
        "social": `Instagram: ${match.instagram}`
      },
    };
    setSelectedProfile(matchAsProfile);
  };

  const handleCloseProfile = () => {
    setSelectedProfile(null);
  };

  const handleSendMessage = () => {
    setShowMatchModal(false);
    // Navigate to chat page
    navigate("/chat");
  };

  const handleKeepSwiping = () => {
    setShowMatchModal(false);
  };

  const handleSocialClick = (matchId: number, type: 'telegram' | 'instagram') => {
    // Show a message when clicking on social media buttons
    setShowChatMessage(matchId);
    
    // Log which social media was clicked
    console.log(`Opening ${type} for match ID ${matchId}`);
    
    // Hide the message after 2 seconds
    setTimeout(() => {
      setShowChatMessage(null);
    }, 2000);
  };

  return (
    <div className={styles.container}>
      <div className={styles.header}>
        <h1 className={styles.title}>Лайки и Мэтчи</h1>
      </div>

      <div className={styles.tabs}>
        <div
          className={`${styles.tab} ${activeTab === "likes" ? styles.activeTab : ""}`}
          onClick={() => setActiveTab("likes")}
        >
          Лайки
        </div>
        <div
          className={`${styles.tab} ${activeTab === "matches" ? styles.activeTab : ""}`}
          onClick={() => setActiveTab("matches")}
        >
          Мэтчи
        </div>
      </div>

      {activeTab === "likes" && (
        <div className={styles.section}>
          {likes.length > 0 ? (
            <div className={styles.profilesGrid}>
              {likes.map((profile) => (
                <div 
                  key={profile.id} 
                  className={styles.profileCard}
                  onClick={() => handleProfileClick(profile)}
                >
                  <img
                    src={profile.imageUrl}
                    alt={profile.name}
                    className={styles.profileImage}
                  />
                  <div className={styles.profileInfo}>
                    <div className={styles.profileName}>
                      {profile.name}, {profile.age}
                    </div>
                  </div>
                  <button
                    className={styles.likeButton}
                    onClick={(e) => {
                      e.stopPropagation(); // Prevent profile click when clicking the like button
                      handleLike(profile.id);
                    }}
                  >
                    <svg
                      width="20"
                      height="18"
                      viewBox="0 0 26 22"
                      fill="none"
                      xmlns="http://www.w3.org/2000/svg"
                    >
                      <path
                        d="M23.4175 2.3871C22.8188 1.78818 22.108 1.31307 21.3257 0.988918C20.5434 0.664766 19.7049 0.497925 18.8581 0.497925C18.0113 0.497925 17.1728 0.664766 16.3905 0.988918C15.6082 1.31307 14.8974 1.78818 14.2987 2.3871L13.0563 3.6295L11.8139 2.3871C10.6047 1.17788 8.96466 0.498552 7.25456 0.498552C5.54447 0.498552 3.90441 1.17788 2.69519 2.3871C1.48597 3.59632 0.806641 5.23638 0.806641 6.94647C0.806641 8.65657 1.48597 10.2966 2.69519 11.5058L3.93759 12.7482L13.0563 21.867L22.1751 12.7482L23.4175 11.5058C24.0164 10.9072 24.4915 10.1964 24.8156 9.4141C25.1398 8.63179 25.3066 7.79328 25.3066 6.94647C25.3066 6.09966 25.1398 5.26115 24.8156 4.47884C24.4915 3.69653 24.0164 2.98575 23.4175 2.3871Z"
                        fill="white"
                      />
                    </svg>
                  </button>
                </div>
              ))}
            </div>
          ) : (
            <div className={styles.emptyState}>
              <div className={styles.emptyStateIcon}>
                <svg width="64" height="64" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <path d="M12 21.35L10.55 20.03C5.4 15.36 2 12.28 2 8.5C2 5.42 4.42 3 7.5 3C9.24 3 10.91 3.81 12 5.09C13.09 3.81 14.76 3 16.5 3C19.58 3 22 5.42 22 8.5C22 12.28 18.6 15.36 13.45 20.04L12 21.35Z" fill="#E0E0E0"/>
                </svg>
              </div>
              <p className={styles.emptyStateText}>У вас пока нет лайков</p>
              <p className={styles.emptyStateSubtext}>Продолжайте искать новые знакомства</p>
            </div>
          )}
        </div>
      )}

      {activeTab === "matches" && (
        <div className={styles.section}>
          {matches.length > 0 ? (
            <div className={styles.matchesContainer}>
              {matches.map((match) => (
                <div
                  key={match.id}
                  className={styles.matchCard}
                  onClick={() => handleMatchClick(match)}
                >
                  <img
                    src={match.imageUrl}
                    alt={match.name}
                    className={styles.matchImage}
                  />
                  <div className={styles.matchInfo}>
                    <div className={styles.matchName}>
                      {match.name}, {match.age}
                    </div>
                    {showChatMessage === match.id && (
                      <div className={styles.chatMessage}>
                        Открыто в новой вкладке
                      </div>
                    )}
                  </div>
                  <div className={styles.socialButtons} onClick={(e) => e.stopPropagation()}>
                    <button
                      className={styles.socialButton}
                      title={match.telegram}
                      onClick={() => handleSocialClick(match.id, 'telegram')}
                    >
                      <TelegramIcon />
                    </button>
                    <button
                      className={styles.socialButton}
                      title={match.instagram}
                      onClick={() => handleSocialClick(match.id, 'instagram')}
                    >
                      <InstagramIcon />
                    </button>
                  </div>
                </div>
              ))}
            </div>
          ) : (
            <div className={styles.emptyState}>
              <div className={styles.emptyStateIcon}>
                <svg width="64" height="64" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <path d="M12 21.35L10.55 20.03C5.4 15.36 2 12.28 2 8.5C2 5.42 4.42 3 7.5 3C9.24 3 10.91 3.81 12 5.09C13.09 3.81 14.76 3 16.5 3C19.58 3 22 5.42 22 8.5C22 12.28 18.6 15.36 13.45 20.04L12 21.35Z" fill="#E0E0E0"/>
                </svg>
              </div>
              <p className={styles.emptyStateText}>У вас пока нет мэтчей</p>
              <p className={styles.emptyStateSubtext}>Лайкайте профили, чтобы найти мэтчи</p>
            </div>
          )}
        </div>
      )}

      {showMatchModal && matchedProfile && (
        <MatchModal
          userImage="man1.jpg" // Placeholder for the current user's image
          matchImage={matchedProfile.imageUrl}
          matchName={matchedProfile.name}
          onSendMessage={handleSendMessage}
          onKeepSwiping={handleKeepSwiping}
        />
      )}

      {selectedProfile && (
        <ProfileView
          profile={selectedProfile}
          onClose={handleCloseProfile}
          onLike={handleLike}
        />
      )}

      <NavBar />
    </div>
  );
};

export default LikesPage;
