import { useState, useEffect, useCallback, useMemo, useRef } from 'react';

import {
  useFetchIncomingLikesMutation,
  useFetchMatchesUserIdsMutation,
  useLikeUserMutation,
} from '../../../entities/matches/api/matchesApi';
import { useGetUsersMutation } from '../../../entities/user';
import { LikeProfile, MatchProfile } from '../types';

export const useLikesAndMatches = () => {
  const [fetchMatchesUserIds] = useFetchMatchesUserIdsMutation();
  const [fetchIncomingLikes, { isLoading: isLoadingIncoming }] = useFetchIncomingLikesMutation();
  //@ts-ignore
  const [likeUser] = useLikeUserMutation();
  const [getUsers] = useGetUsersMutation();
  //const [getUser] = useGetUserMutation();

  const [matchesUsers, setMatchesUsers] = useState([]);
  const [likesUsers, setLikesUsers] = useState([]);

  const isInitialized = useRef(false);

  const [likes, setIncomingMatches] = useState<LikeProfile[]>([]);

  useEffect(() => {
    const loadData = async () => {
      try {
        const matchesIds = await fetchMatchesUserIds({
          limit: 10,
          offset: 0,
        }).unwrap();

        const matchesUsers = await getUsers(matchesIds).unwrap();

        setMatchesUsers(matchesUsers);

        // @ts-ignore
        // setMatches(matchesResult.match.map(el => el.mainUserId));

        const incomingResult = await fetchIncomingLikes({
          limit: 10,
          offset: 0,
        }).unwrap();

        const likesIds = incomingResult.likes.map(el => el.senderId);
        console.log('IDS', likesIds);

        const likesUsers = await getUsers(likesIds).unwrap();

        console.log('1', likesUsers);

        setLikesUsers(likesUsers);

        // @ts-ignore
        // const likes = incomingResult.match.map(el => el.mainUserId);

        // setIncomingMatches(likes);
      } catch (error) {
        console.error('Error loading matches data:', error);
      }
    };

    if (!isInitialized.current) {
      loadData();
      isInitialized.current = true;
    }
  }, [fetchIncomingLikes, fetchMatchesUserIds, getUsers]);

  const [matches, setMatches] = useState<MatchProfile[]>([]);
  const [showMatchModal, setShowMatchModal] = useState(false);
  const [matchedProfile, setMatchedProfile] = useState<LikeProfile | null>(null);
  const [showChatMessage, setShowChatMessage] = useState<number | null>(null);

  const handleLike = useCallback(
    async (id: number) => {
      // @ts-ignore
      const likedProfile = likes.find(like => like === id);

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
    [likes, matches],
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
      matches,
      showMatchModal,
      matchedProfile,
      showChatMessage,
      handleLike,
      handleSendMessage,
      handleKeepSwiping,
      handleSocialClick,
      isLoading: isLoadingIncoming || false,
      likes,
      matchesUsers,
      likesUsers,
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
      likes,
      matchesUsers,
      likesUsers,
    ],
  );
};
