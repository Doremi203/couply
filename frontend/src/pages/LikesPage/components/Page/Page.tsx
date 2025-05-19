import { useState, useEffect, useCallback, useRef } from 'react';

import { useDislikeUserMutation } from '../../../../entities/matches';
import { LikesSection, useLikesAndMatches } from '../../../../features/matches';
import { NavBar } from '../../../../shared/components/NavBar';
import TabsSection from '../../../../shared/components/TabsSection';
import { ProfileView } from '../../../../widgets/ProfileView';
import { useProfileView } from '../../hooks/useProfileView';
import MatchesSection from '../MatchesSection';

import styles from './likesPage.module.css';

export const LikesPage = () => {
  const [activeTab, setActiveTab] = useState<'лайки' | 'мэтчи'>('лайки');

  const {
    showChatMessage,
    handleLike,
    handleSocialClick,
    matchesUsers,
    likesUsers,
    likes,
    handleDislike,
  } = useLikesAndMatches();

  const { selectedProfile, handleProfileClick, handleMatchClick, handleCloseProfile } =
    useProfileView();

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
          likesUsers={likesUsers}
          likes={likes}
          onProfileClick={handleProfileClick}
          onLike={handleLike}
        />
      )}

      {activeTab === 'мэтчи' && (
        <MatchesSection
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
          onLike={handleLike}
          onDislike={handleDislike}
        />
      )}

      <div style={{ position: 'relative', zIndex: 1010 }}>
        <NavBar />
      </div>
    </div>
  );
};

export default LikesPage;
