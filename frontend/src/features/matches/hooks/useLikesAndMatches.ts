import { useState, useEffect, useCallback, useMemo, useRef } from 'react';

import {
  useFetchIncomingLikesMutation,
  useFetchMatchesMutation,
  useLikeUserMutation,
} from '../../../entities/matches/api/matchesApi';
import { LikeProfile, MatchProfile } from '../types';

export const useLikesAndMatches = () => {
  const [fetchMatches] = useFetchMatchesMutation();
  const [fetchIncomingLikes, { isLoading: isLoadingIncoming }] = useFetchIncomingLikesMutation();
  const [likeUser] = useLikeUserMutation();
  //const [getUser] = useGetUserMutation();

  const isInitialized = useRef(false);

  const [incomingMatches, setIncomingMatches] = useState<LikeProfile[]>([]);

  useEffect(() => {
    const loadData = async () => {
      try {
        const matchesResult = await fetchMatches({
          limit: 10,
          offset: 0,
        }).unwrap();

        // @ts-ignore
        setMatches(matchesResult.match.map(el => el.mainUserId));

        const incomingResult = await fetchIncomingLikes({
          limit: 10,
          offset: 0,
        }).unwrap();

        // @ts-ignore
        const incomingMatches = incomingResult.match.map(el => el.mainUserId);

        setIncomingMatches(incomingMatches);
      } catch (error) {
        console.error('Error loading matches data:', error);
      }
    };

    if (!isInitialized.current) {
      loadData();
      isInitialized.current = true;
    }
  }, [fetchMatches, fetchIncomingLikes]);

  const [matches, setMatches] = useState<MatchProfile[]>([]);
  const [showMatchModal, setShowMatchModal] = useState(false);
  const [matchedProfile, setMatchedProfile] = useState<LikeProfile | null>(null);
  const [showChatMessage, setShowChatMessage] = useState<number | null>(null);

  const handleLike = useCallback(
    async (id: number) => {
      // @ts-ignore
      const likedProfile = incomingMatches.find(like => like === id);

      if (likedProfile) {
        try {
          // await updateMatch({
          //   mainUserId: likedProfile,
          //   // @ts-ignore
          //   chosenUserId: userId,
          //   approved: true,
          // });

          // await likeUser({
          //   targetUserId: likedProfile,
          //   // @ts-ignore
          //   chosenUserId: userId,
          // });

          setShowMatchModal(true);

          // @ts-ignore
          setMatchedProfile(await getUser({ id: likedProfile }).unwrap());

          // TODO NOTIFICATION
          // if (likedProfile.hasLikedYou) {
          //   setMatchedProfile(likedProfile);
          //   setShowMatchModal(true);
          //   sendMatchNotification({
          //     userId: userId,
          //     matchId: likedProfile.id,
          //     matchName: likedProfile.name,
          //     matchImage: likedProfile.imageUrl,
          //   });
          // }

          setMatches(matches.concat(likedProfile));

          // @ts-ignore
          setIncomingMatches(incomingMatches.filter(like => like !== id));
        } catch (error) {
          console.error('Error creating match:', error);
        }
      }
    },
    [incomingMatches, matches],
  );

  const handleSendMessage = useCallback(() => {
    setShowMatchModal(false);
  }, []);

  const handleKeepSwiping = useCallback(() => {
    setShowMatchModal(false);
  }, []);

  const handleSocialClick = useCallback((matchId: number, type: 'telegram' | 'instagram') => {
    // Показываем сообщение при клике на кнопки социальных сетей
    setShowChatMessage(matchId);

    // Логируем, какая социальная сеть была нажата
    console.log(`Opening ${type} for match ID ${matchId}`);

    // Скрываем сообщение через 2 секунды
    setTimeout(() => {
      setShowChatMessage(null);
    }, 2000);
  }, []);

  return useMemo(
    () => ({
      // matches,
      showMatchModal,
      matchedProfile,
      showChatMessage,
      handleLike,
      handleSendMessage,
      handleKeepSwiping,
      handleSocialClick,
      isLoading: isLoadingIncoming || false,
      // incomingMatches,
    }),
    [
      showMatchModal,
      matchedProfile,
      showChatMessage,
      handleLike,
      handleSendMessage,
      handleKeepSwiping,
      handleSocialClick,
      isLoadingIncoming,
    ],
  );
};
