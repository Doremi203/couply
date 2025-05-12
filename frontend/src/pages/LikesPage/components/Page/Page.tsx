import { useState, useEffect, useCallback, useRef } from 'react';

import { LikesSection, useLikesAndMatches } from '../../../../features/matches';
import { NavBar } from '../../../../shared/components/NavBar';
import TabsSection from '../../../../shared/components/TabsSection';
import { ProfileView } from '../../../../widgets/ProfileView';
import { useProfileView } from '../../hooks/useProfileView';
import MatchesSection from '../MatchesSection';
import { MatchModal } from '../MatchModal';

import styles from './likesPage.module.css';

const likes = [
  {
    id: 1,
    name: 'Анна',
    age: 25,
    imageUrl: 'girl.jpeg',
    liked: false,
    hasLikedYou: true, // This profile has already liked the user
    location: 'Москва, Россия',
    interests: ['Музыка', 'Путешествия', 'Фотография', 'Мода', 'Искусство'],
  },
  {
    id: 2,
    name: 'Иван',
    age: 30,
    imageUrl: 'boy.jpeg',
    liked: false,
    hasLikedYou: true, // This profile has already liked the user
    location: 'Санкт-Петербург, Россия',
    interests: ['Спорт', 'Кино', 'Технологии', 'Путешествия'],
  },
  {
    id: 3,
    name: 'Ольга',
    age: 28,
    imageUrl: 'woman3.jpeg',
    liked: false,
    hasLikedYou: false,
    location: 'Казань, Россия',
    interests: ['Книги', 'Йога', 'Кулинария', 'Природа'],
  },
  {
    id: 4,
    name: 'Алексей',
    age: 32,
    imageUrl: 'miio.jpeg',
    liked: false,
    hasLikedYou: false,
    location: 'Екатеринбург, Россия',
    interests: ['Музыка', 'Горы', 'Фотография', 'Путешествия'],
  },
];

const matchesUsers = [
  {
    id: 101, // Using different ID range for matches
    name: 'Мария',
    age: 27,
    imageUrl: 'woman1.jpg',
    telegram: '@maria_27',
  },
  {
    id: 102, // Using different ID range for matches
    name: 'Дмитрий',
    age: 31,
    imageUrl: 'boy1.jpeg',
    telegram: '@dmitry_31',
  },
];

export const LikesPage = () => {
  const [activeTab, setActiveTab] = useState<'лайки' | 'мэтчи'>('лайки');

  const {
    //   matches,
    showMatchModal,
    matchedProfile,
    showChatMessage,
    handleLike,
    // handleSendMessage,
    handleKeepSwiping,
    handleSocialClick,
    //   incomingMatches,
  } = useLikesAndMatches();

  const { selectedProfile, handleProfileClick, handleMatchClick, handleCloseProfile } =
    useProfileView();

  // Handle tab change - memoize to prevent unnecessary re-renders
  const handleTabChange = useCallback((tab: 'лайки' | 'мэтчи') => {
    setActiveTab(tab);
  }, []);

  // Handle send message (switch to matches tab) - memoize to prevent unnecessary re-renders
  // const handleSendMessageAndSwitchTab = useCallback(() => {
  //   handleSendMessage();
  //   setActiveTab('matches');
  // }, [handleSendMessage]);

  // Effect to clean up when component unmounts - use a ref to prevent dependency on selectedProfile
  const firstRender = useRef(true);

  useEffect(() => {
    // Skip the first render to avoid unnecessary calls
    if (firstRender.current) {
      firstRender.current = false;
      return;
    }

    return () => {
      // Clean up any resources when navigating away
      if (selectedProfile) {
        handleCloseProfile();
      }
    };
  }, [handleCloseProfile, selectedProfile]);

  return (
    <div className={styles.container} id="likes-page-container">
      <div className={styles.header}>Лайки и мэтчи</div>

      <TabsSection
        tabs={['лайки', 'мэтчи'] as const}
        activeTab={activeTab}
        onTabChange={handleTabChange}
      />

      {activeTab === 'лайки' && (
        // @ts-ignore
        <LikesSection likes={likes} onProfileClick={handleProfileClick} onLike={handleLike} />
      )}

      {activeTab === 'мэтчи' && (
        <MatchesSection
          matches={matchesUsers}
          onMatchClick={handleMatchClick}
          onSocialClick={handleSocialClick}
          showChatMessage={showChatMessage}
        />
      )}

      {showMatchModal && matchedProfile && (
        <MatchModal
          userImage="man1.jpg"
          // matchImage={matchedProfile.user.imageUrl}
          matchImage="djdj"
          matchName={matchedProfile.user.name}
          onKeepSwiping={handleKeepSwiping}
        />
      )}

      {selectedProfile && (
        <ProfileView profile={selectedProfile} onClose={handleCloseProfile} onLike={handleLike} />
      )}

      <div style={{ position: 'relative', zIndex: 1010 }}>
        <NavBar />
      </div>
    </div>
  );
};

export default LikesPage;
