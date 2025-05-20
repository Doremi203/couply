import { useState, useEffect, useCallback, useMemo, useRef } from 'react';

import {
  useDislikeUserMutation,
  useFetchIncomingLikesMutation,
  useFetchMatchesUserIdsMutation,
  useLikeUserMutation,
} from '../../../entities/matches/api/matchesApi';
import { Like } from '../../../entities/matches/types';
import { useGetUsersMutation } from '../../../entities/user';
import { UserData } from '../../../entities/user/types';
import { LikeProfile, MatchProfile } from '../types';

const ITEMS_PER_PAGE = 10;

export const useLikesAndMatches = () => {
  const [fetchMatchesUserIds] = useFetchMatchesUserIdsMutation();
  const [fetchIncomingLikes, { isLoading: isLoadingIncoming }] = useFetchIncomingLikesMutation();
  const [getUsers] = useGetUsersMutation();
  const [likeUser] = useLikeUserMutation();

  const [matchesUsers, setMatchesUsers] = useState<UserData[]>([]);
  const [likesUsers, setLikesUsers] = useState<UserData[]>([]);
  const [matchesOffset, setMatchesOffset] = useState(0);
  const [likesOffset, setLikesOffset] = useState(0);
  const [hasMoreMatches, setHasMoreMatches] = useState(true);
  const [hasMoreLikes, setHasMoreLikes] = useState(true);

  const isInitialized = useRef(false);

  const [likes, setIncomingMatches] = useState<Like[]>([]);

  const loadMatches = useCallback(async (offset: number) => {
    try {
      const matchesIds = await fetchMatchesUserIds({
        limit: ITEMS_PER_PAGE,
        offset,
      }).unwrap();

      if (matchesIds.userIds.length === 0) {
        setHasMoreMatches(false);
        return;
      }

          //@ts-ignore
      const matchesUsersResponse = await getUsers(matchesIds.userIds).unwrap();
      // const newMatchesUsers = matchesUsersResponse.users.map(user => user.user);
      
      if (offset === 0) {
            //@ts-ignore
        setMatchesUsers(matchesUsersResponse);
      } else {
            //@ts-ignore
        setMatchesUsers(prev => [...prev, ...matchesUsersResponse]);
      }
    } catch (error) {
      console.error('Error loading matches data:', error);
    }
  }, [fetchMatchesUserIds, getUsers]);

  const loadLikes = useCallback(async (offset: number) => {
    try {
      const incomingResult = await fetchIncomingLikes({
        limit: ITEMS_PER_PAGE,
        offset,
      }).unwrap();

      if (incomingResult.likes.length === 0) {
        setHasMoreLikes(false);
        return;
      }

      const likesIds = incomingResult.likes.map(el => el.senderId);
          //@ts-ignore
      const likesUsersResponse = await getUsers(likesIds).unwrap();
      // const newLikesUsers = likesUsersResponse.users.map(user => user.user);


      // console.log(likesUsersResponse)
  

      if (offset === 0) {
            //@ts-ignore
        setLikesUsers(likesUsersResponse);
        setIncomingMatches(incomingResult.likes);
      } else {
            //@ts-ignore
        setLikesUsers(prev => [...prev, ...likesUsersResponse]);
        setIncomingMatches(prev => [...prev, ...incomingResult.likes]);
      }
    } catch (error) {
      console.error('Error loading likes data:', error);
    }
  }, [fetchIncomingLikes, getUsers]);

  useEffect(() => {
    const loadData = async () => {
      await Promise.all([
        loadMatches(0),
        loadLikes(0),
      ]);
    };

    if (!isInitialized.current) {
      loadData();
      isInitialized.current = true;
    }
  }, [loadMatches, loadLikes]);

  const loadMoreMatches = useCallback(async () => {
    if (!hasMoreMatches || isLoadingIncoming) return;
    
    const newOffset = matchesOffset + ITEMS_PER_PAGE;
    setMatchesOffset(newOffset);
    await loadMatches(newOffset);
  }, [hasMoreMatches, isLoadingIncoming, matchesOffset, loadMatches]);

  const loadMoreLikes = useCallback(async () => {
    if (!hasMoreLikes || isLoadingIncoming) return;
    
    const newOffset = likesOffset + ITEMS_PER_PAGE;
    setLikesOffset(newOffset);
    await loadLikes(newOffset);
  }, [hasMoreLikes, isLoadingIncoming, likesOffset, loadLikes]);

  const [matches, setMatches] = useState<MatchProfile[]>([]);
  const [showMatchModal, setShowMatchModal] = useState(false);
  const [matchedProfile, setMatchedProfile] = useState<LikeProfile | null>(null);
  const [showChatMessage, setShowChatMessage] = useState<number | null>(null);

  const [dislike] = useDislikeUserMutation();

  const handleDislike = useCallback(
    async (id: string) => {
        try {
          await dislike({
            targetUserId: id,
          });

          setIncomingMatches(prev => prev.filter(like => like.senderId !== id));
        } catch (error) {
          console.error('Error creating match:', error);
        }
      },
    [dislike, likes],
  );

  const handleLike = useCallback(
    async (id: string) => {
      const likedProfile = likes.find(like => like.senderId === id);

      if (likedProfile) {
        try {
          await likeUser({
            targetUserId: likedProfile.senderId,
            message: '',
          });

          setShowMatchModal(true);

          const userResponse = await getUsers({ userIds: [likedProfile.senderId] }).unwrap();
          const userData = userResponse.users[0].user;
          
          // Convert UserData to LikeProfile
          const likeProfile = {
            name: userData.name,
            age: userData.age,
            imageUrl: userData.photos?.[0]?.url || '',
            hasLikedYou: true,
            bio: userData.bio,
            location: userData.location,
            interests: [],
            lifestyle: {},
            passion: [],
            photos: userData.photos,
          } as any;
          
          setMatchedProfile(likeProfile);

          setMatches(prev => [...prev, {
            id: parseInt(userData.id),
            name: userData.name,
            age: userData.age,
            imageUrl: userData.photos?.[0]?.url || '',
            telegram: '',
            instagram: '',
          }]);

          setIncomingMatches(prev => prev.filter(like => like.senderId !== id));
        } catch (error) {
          console.error('Error creating match:', error);
        }
      }
    },
    [likeUser, likes, getUsers],
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
      handleDislike,
      loadMoreMatches,
      loadMoreLikes,
      hasMoreMatches,
      hasMoreLikes,
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
      handleDislike,
      loadMoreMatches,
      loadMoreLikes,
      hasMoreMatches,
      hasMoreLikes,
    ],
  );
};
