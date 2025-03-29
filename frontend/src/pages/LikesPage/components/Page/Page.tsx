import React, { useState } from "react";
import { NavBar } from "../../../../shared/components/NavBar";
import { MatchModal } from "../MatchModal";
import { ProfileView } from "../ProfileView";
import { useNavigate } from "react-router-dom";
import styles from "./likesPage.module.css";
import { ProfileCard } from "../ProfileCard";
import { MatchCard, MatchProfile } from "../MatchCard";

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
        likes & matches
      </div>

      <div className={styles.tabs}>
        <div
          className={`${styles.tab} ${activeTab === "likes" ? styles.activeTab : ""}`}
          onClick={() => setActiveTab("likes")}
        >
          likes
        </div>
        <div
          className={`${styles.tab} ${activeTab === "matches" ? styles.activeTab : ""}`}
          onClick={() => setActiveTab("matches")}
        >
          matches
        </div>
      </div>

      {activeTab === "likes" && (
        <div className={styles.section}>
          {likes.length > 0 ? (
            <div className={styles.profilesGrid}>
              {likes.map((profile) => (
                <ProfileCard
                  key={profile.id}
                  profile={profile}
                  onClick={() => handleProfileClick(profile)}
                  onLike={handleLike}
                  className={styles.profileCard}
                />
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
                <MatchCard
                  key={match.id}
                  match={match}
                  onClick={handleMatchClick}
                  onSocialClick={handleSocialClick}
                  showChatMessage={showChatMessage}
                />
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
