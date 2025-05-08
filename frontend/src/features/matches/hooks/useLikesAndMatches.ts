import { useState, useEffect, useCallback, useMemo, useRef } from 'react';
import { useSelector } from 'react-redux';

import {
  useFetchIncomingMatchesMutation,
  useFetchMatchesMutation,
  useUpdateMatchMutation,
} from '../../../entities/matches/api/matchesApi';
import { useGetUserMutation } from '../../../entities/user';
import { getUserId } from '../../../entities/user/model/userSlice';
import { LikeProfile, MatchProfile } from '../types';

export const useMatches = () => {
  //const userId = useSelector(getUserId);
  const userId = 1;

  const [fetchMatches] = useFetchMatchesMutation();
  const [fetchIncomingMatches, { isLoading: isLoadingIncoming }] =
    useFetchIncomingMatchesMutation();
  const [getUser] = useGetUserMutation();

  const isInitialized = useRef(false);

  const [incomingMatches, setIncomingMatches] = useState<LikeProfile[]>([]);

  // useEffect(() => {
  //   if (!userId) return;

  //   const loadData = async () => {
  //     try {
  //       const matchesResult = await fetchMatches({
  //         mainUserId: userId,
  //         limit: 10,
  //         offset: 0,
  //       }).unwrap();

  //       // @ts-ignore
  //       setMatches(matchesResult.match.map(el => el.mainUserId));

  //       const incomingResult = await fetchIncomingMatches({
  //         chosenUserId: userId,
  //         limit: 10,
  //         offset: 0,
  //       }).unwrap();

  //       // @ts-ignore
  //       const incomingMatches = incomingResult.match.map(el => el.mainUserId);

  //       setIncomingMatches(incomingMatches);
  //     } catch (error) {
  //       console.error('Error loading matches data:', error);
  //     }
  //   };

  //   if (!isInitialized.current) {
  //     loadData();
  //     isInitialized.current = true;
  //   }
  // }, [userId, fetchMatches, fetchIncomingMatches]);

  const [updateMatch] = useUpdateMatchMutation();

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
          await updateMatch({
            mainUserId: likedProfile,
            // @ts-ignore
            chosenUserId: userId,
            approved: true,
          });

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
    [incomingMatches, updateMatch, userId, getUser, matches],
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
      matches,
      showMatchModal,
      matchedProfile,
      showChatMessage,
      handleLike,
      handleSendMessage,
      handleKeepSwiping,
      handleSocialClick,
      isLoadingIncoming,
      incomingMatches,
    ],
  );
};
