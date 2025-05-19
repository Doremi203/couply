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
  const [getUsers] = useGetUsersMutation();
  const [likeUser] = useLikeUserMutation();

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

        const matchesUsers = await getUsers(matchesIds.userIds).unwrap();

        setMatchesUsers(matchesUsers);

        const incomingResult = await fetchIncomingLikes({
          limit: 10,
          offset: 0,
        }).unwrap();

        const likesIds = incomingResult.likes.map(el => el.senderId);

        const likesUsers = await getUsers(likesIds).unwrap();

        setLikesUsers(likesUsers);

        setIncomingMatches(incomingResult);
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
      const likedProfile = likes.find(like => like === id);

      if (likedProfile) {
        try {
          await likeUser({
            targetUserId: likedProfile,
            message: '',
          });

          setShowMatchModal(true);

          setMatchedProfile(await getUser({ id: likedProfile }).unwrap());

          setMatches(matches.concat(likedProfile));

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

  const handleSocialClick = useCallback((matchId: number) => {
    setShowChatMessage(matchId);

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
