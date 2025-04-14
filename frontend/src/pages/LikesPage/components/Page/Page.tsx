import { useState } from 'react';

import { LikesSection, useLikesAndMatches } from '../../../../features/likes';
import { NavBar } from '../../../../shared/components/NavBar';
import TabsSection from '../../../../shared/components/TabsSection';
// import { ProfileView, useProfileView } from '../../../../widgets/ProfileView';
import { useProfileView } from '../../hooks/useProfileView';
import MatchesSection from '../MatchesSection';
import { MatchModal } from '../MatchModal';
import { ProfileView } from '../ProfileView';

import styles from './likesPage.module.css';

export const LikesPage = () => {
  const [activeTab, setActiveTab] = useState<'likes' | 'matches'>('likes');

  // Use custom hooks for state management
  const {
    likes,
    matches,
    showMatchModal,
    matchedProfile,
    showChatMessage,
    handleLike,
    handleSendMessage,
    handleKeepSwiping,
    handleSocialClick,
  } = useLikesAndMatches();

  const { selectedProfile, handleProfileClick, handleMatchClick, handleCloseProfile } =
    useProfileView();

  // Handle tab change
  const handleTabChange = (tab: 'likes' | 'matches') => {
    setActiveTab(tab);
  };

  // Handle send message (switch to matches tab)
  const handleSendMessageAndSwitchTab = () => {
    handleSendMessage();
    setActiveTab('matches');
  };

  return (
    <div className={styles.container}>
      <div className={styles.header}>likes & matches</div>

      {/* Tabs section */}
      <TabsSection
        tabs={['likes', 'matches'] as const}
        activeTab={activeTab}
        onTabChange={handleTabChange}
      />

      {/* Content based on active tab */}
      {activeTab === 'likes' && (
        <LikesSection likes={likes} onProfileClick={handleProfileClick} onLike={handleLike} />
      )}

      {activeTab === 'matches' && (
        <MatchesSection
          matches={matches}
          onMatchClick={handleMatchClick}
          onSocialClick={handleSocialClick}
          showChatMessage={showChatMessage}
        />
      )}

      {/* Match modal */}
      {showMatchModal && matchedProfile && (
        <MatchModal
          userImage="man1.jpg" // Placeholder for the current user's image
          matchImage={matchedProfile.imageUrl}
          matchName={matchedProfile.name}
          onSendMessage={handleSendMessageAndSwitchTab}
          onKeepSwiping={handleKeepSwiping}
        />
      )}

      {/* Profile view */}
      {selectedProfile && (
        <ProfileView profile={selectedProfile} onClose={handleCloseProfile} onLike={handleLike} />
      )}

      <NavBar />
    </div>
  );
};

export default LikesPage;
