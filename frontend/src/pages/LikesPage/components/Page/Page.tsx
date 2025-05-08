import { useState, useEffect, useCallback, useRef } from 'react';

import useFetchUsers from '../../../../entities/user/hooks/useFetchUsers';
import { LikesSection, useMatches } from '../../../../features/matches';
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
    name: 'Константин',
    age: 25,
    imageUrl: 'woman1.jpg',
    liked: false,
    hasLikedYou: true, // This profile has already liked the user
    location: 'Москва, Россия',
    interests: ['Музыка', 'Путешествия', 'Фотография', 'Мода', 'Искусство'],
  },
  {
    id: 2,
    name: 'Иван',
    age: 30,
    imageUrl: 'man1.jpg',
    liked: false,
    hasLikedYou: true, // This profile has already liked the user
    location: 'Санкт-Петербург, Россия',
    interests: ['Спорт', 'Кино', 'Технологии', 'Путешествия'],
  },
  {
    id: 3,
    name: 'Ольга',
    age: 28,
    imageUrl: 'photo1.png',
    liked: false,
    hasLikedYou: false,
    location: 'Казань, Россия',
    interests: ['Книги', 'Йога', 'Кулинария', 'Природа'],
  },
  {
    id: 4,
    name: 'Алексей',
    age: 32,
    imageUrl: 'man1.jpg',
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
    instagram: '@maria_insta',
  },
  {
    id: 102, // Using different ID range for matches
    name: 'Дмитрий',
    age: 31,
    imageUrl: 'man1.jpg',
    telegram: '@dmitry_31',
    instagram: '@dmitry_insta',
  },
];

export const LikesPage = () => {
  const [activeTab, setActiveTab] = useState<'likes' | 'matches'>('likes');

  const {
    //   matches,
    showMatchModal,
    matchedProfile,
    showChatMessage,
    handleLike,
    handleSendMessage,
    handleKeepSwiping,
    handleSocialClick,
    //   incomingMatches,
  } = useMatches();

  // @ts-ignore
  //const likes = useFetchUsers(incomingMatches);
  // @ts-ignore
  //const matchesUsers = useFetchUsers(matches);

  const { selectedProfile, handleProfileClick, handleMatchClick, handleCloseProfile } =
    useProfileView();

  // Handle tab change - memoize to prevent unnecessary re-renders
  const handleTabChange = useCallback((tab: 'likes' | 'matches') => {
    setActiveTab(tab);
  }, []);

  // Handle send message (switch to matches tab) - memoize to prevent unnecessary re-renders
  const handleSendMessageAndSwitchTab = useCallback(() => {
    handleSendMessage();
    setActiveTab('matches');
  }, [handleSendMessage]);

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
      <div className={styles.header}>likes & matches</div>

      <TabsSection
        tabs={['likes', 'matches'] as const}
        activeTab={activeTab}
        onTabChange={handleTabChange}
      />

      {activeTab === 'likes' && (
        // @ts-ignore
        <LikesSection likes={likes} onProfileClick={handleProfileClick} onLike={handleLike} />
      )}

      {activeTab === 'matches' && (
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
