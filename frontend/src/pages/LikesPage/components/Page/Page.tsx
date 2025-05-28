import { useState, useEffect, useCallback, useRef } from 'react';

import { LikesSection, useLikesAndMatches } from '../../../../features/matches';
import { NavBar } from '../../../../shared/components/NavBar';
import TabsSection from '../../../../shared/components/TabsSection';
import { ProfileView } from '../../../../widgets/ProfileView';
import MatchesSection from '../MatchesSection';

import styles from './likesPage.module.css';

export const LikesPage = () => {
  const [activeTab, setActiveTab] = useState<'лайки' | 'мэтчи'>('лайки');
  const [selectedProfile, setSelectedProfile] = useState<any>(null);

  const {
    showChatMessage,
    handleLike,
    handleSocialClick,
    matchesUsers,
    likesUsers,
    likes,
    handleDislike,
  } = useLikesAndMatches();

  const handleProfileClick = useCallback((profile: any) => {
    setSelectedProfile(profile);
  }, []);

  const handleMatchClick = useCallback((match: any) => {
    setSelectedProfile(match);
  }, []);

  const handleCloseProfile = useCallback(() => {
    setSelectedProfile(null);
  }, []);

  const handleTabChange = useCallback((tab: 'лайки' | 'мэтчи') => {
    setActiveTab(tab);
  }, []);

  const firstRender = useRef(true);

  useEffect(() => {
    if (firstRender.current) {
      firstRender.current = false;
      return;
    }

    return () => {
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
        <LikesSection
          //@ts-ignore
          likesUsers={likesUsers}
          //@ts-ignore
          likes={likes}
          onProfileClick={handleProfileClick}
          //@ts-ignore
          onLike={handleLike}
        />
      )}

      {activeTab === 'мэтчи' && (
        <MatchesSection
          //@ts-ignore
          matches={matchesUsers}
          onMatchClick={handleMatchClick}
          onSocialClick={handleSocialClick}
          showChatMessage={showChatMessage}
        />
      )}

      {selectedProfile && (
        <ProfileView
          profile={selectedProfile}
          onClose={handleCloseProfile}
          //@ts-ignore
          onLike={handleLike}
          //@ts-ignore
          onDislike={handleDislike}
          isMatchView={activeTab === 'мэтчи'}
        />
      )}

      <div style={{ position: 'relative', zIndex: 1010 }}>
        <NavBar />
      </div>
    </div>
  );
};

export default LikesPage;
