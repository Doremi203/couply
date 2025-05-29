import { useState, useEffect, useCallback, useRef } from 'react';

import { useGetTelegramMutation } from '../../../../entities/telegram/api/telegramApi';
import { useGetUserMutation } from '../../../../entities/user';
import { LikesSection, useLikesAndMatches } from '../../../../features/matches';
import { NavBar } from '../../../../shared/components/NavBar';
import TabsSection from '../../../../shared/components/TabsSection';
import { ProfileView } from '../../../../widgets/ProfileView';
import MatchesSection from '../MatchesSection';
import TelegramModal from '../TelegramModal';

import styles from './likesPage.module.css';

export const LikesPage = () => {
  const [activeTab, setActiveTab] = useState<'лайки' | 'мэтчи'>('лайки');
  const [selectedProfile, setSelectedProfile] = useState<any>(null);
  const [showTelegramModal, setShowTelegramModal] = useState<boolean>(false);
  const [getTelegram] = useGetTelegramMutation();
  const [getUser] = useGetUserMutation();

  const [telegram, setTelegram] = useState<string>('');
  const [isLoading, setIsLoading] = useState<boolean>(true);

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
    const fetchTg = async () => {
      setIsLoading(true);
      try {
        const profile = await getUser({}).unwrap();

        console.log('profile', profile);
        //@ts-ignore
        console.log('profile.id', profile.user.id);

        //@ts-ignore
        const tg = await getTelegram(profile.user.id).unwrap();

        console.log('tg', tg);

        //@ts-ignore
        setTelegram(tg.telegramUrl);
      } catch (err) {
        console.error('Error fetching users:', err);
      } finally {
        setIsLoading(false);
      }
    };

    fetchTg();
  }, [getTelegram, getUser]);

  useEffect(() => {
    // If telegram has a value, immediately close the modal
    if (telegram) {
      setShowTelegramModal(false);
      return;
    }

    // Only show the modal if we're done loading and there's no telegram value
    if (!isLoading && telegram === '') {
      setShowTelegramModal(true);
    }
  }, [telegram, isLoading]);

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

      <TelegramModal
        isOpen={showTelegramModal}
        onClose={() => {
          setShowTelegramModal(false);
          // setIsTelegram(true);
        }}
      />
    </div>
  );
};

export default LikesPage;
