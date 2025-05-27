import { useCallback, useEffect, useRef } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import {
  useDislikeUserMutation,
  useFetchIncomingLikesMutation,
  useFetchMatchesUserIdsMutation,
  useLikeUserMutation,
  useDeleteMatchMutation,
} from '../../../entities/matches/api/matchesApi';
import {
  addMatch,
  appendLikes,
  appendMatches,
  removeLike,
  removeMatch,
  selectHasMoreLikes,
  selectHasMoreMatches,
  selectLikes,
  selectLikesOffset,
  selectLikesUsers,
  selectMatches,
  selectMatchesOffset,
  selectMatchedProfile,
  selectShowChatMessage,
  selectShowMatchModal,
  setHasMoreLikes,
  setHasMoreMatches,
  setLikes,
  setLikesOffset,
  setLikesUsers,
  setMatches,
  setMatchesOffset,
  setMatchedProfile,
  setShowChatMessage,
  setShowMatchModal,
} from '../../../entities/matches/model/matchesSlice';
import { useGetUsersMutation } from '../../../entities/user';

const ITEMS_PER_PAGE = 10;

export const useLikesAndMatches = () => {
  const dispatch = useDispatch();

  const likes = useSelector(selectLikes);
  const likesUsers = useSelector(selectLikesUsers);
  const matches = useSelector(selectMatches);
  const matchesOffset = useSelector(selectMatchesOffset);
  const likesOffset = useSelector(selectLikesOffset);
  const hasMoreMatches = useSelector(selectHasMoreMatches);
  const hasMoreLikes = useSelector(selectHasMoreLikes);
  const showMatchModal = useSelector(selectShowMatchModal);
  const matchedProfile = useSelector(selectMatchedProfile);
  const showChatMessage = useSelector(selectShowChatMessage);

  const [fetchMatchesUserIds] = useFetchMatchesUserIdsMutation();
  const [fetchIncomingLikes, { isLoading: isLoadingIncoming }] = useFetchIncomingLikesMutation();
  const [getUsers] = useGetUsersMutation();
  const [likeUser] = useLikeUserMutation();
  const [dislike] = useDislikeUserMutation();
  const [deleteMatch] = useDeleteMatchMutation();

  const isInitialized = useRef(false);

  const loadMatches = useCallback(
    async (offset: number) => {
      try {
        const matchesIds = await fetchMatchesUserIds({
          limit: ITEMS_PER_PAGE,
          offset,
        }).unwrap();

        if (matchesIds.userIds.length === 0) {
          dispatch(setHasMoreMatches(false));
          return;
        }

        console.log('matchesIds', matchesIds);
        //@ts-ignore
        const matchesUsersResponse = await getUsers(matchesIds.userIds).unwrap();

        if (offset === 0) {
          //@ts-ignore
          dispatch(setMatches(matchesUsersResponse));
        } else {
          //@ts-ignore
          dispatch(appendMatches(matchesUsersResponse));
        }
      } catch (error) {
        console.error('Error loading matches data:', error);
      }
    },
    [dispatch, fetchMatchesUserIds, getUsers],
  );

  const loadLikes = useCallback(
    async (offset: number) => {
      try {
        const incomingResult = await fetchIncomingLikes({
          limit: ITEMS_PER_PAGE,
          offset,
        }).unwrap();

        if (incomingResult.likes.length === 0) {
          dispatch(setHasMoreLikes(false));
          return;
        }

        const likesIds = incomingResult.likes.map(el => el.senderId);
        //@ts-ignore
        const likesUsersResponse = await getUsers(likesIds).unwrap();

        if (offset === 0) {
          //@ts-ignore
          dispatch(setLikesUsers(likesUsersResponse));
          dispatch(setLikes(incomingResult.likes));
        } else {
          //@ts-ignore
          dispatch(appendLikes({ likes: incomingResult.likes, users: likesUsersResponse }));
        }
      } catch (error) {
        console.error('Error loading likes data:', error);
      }
    },
    [dispatch, fetchIncomingLikes, getUsers],
  );

  useEffect(() => {
    const loadData = async () => {
      await Promise.all([loadMatches(0), loadLikes(0)]);
    };

    if (!isInitialized.current) {
      loadData();
      isInitialized.current = true;
    }
  }, [loadMatches, loadLikes]);

  const loadMoreMatches = useCallback(async () => {
    if (!hasMoreMatches || isLoadingIncoming) return;

    const newOffset = matchesOffset + ITEMS_PER_PAGE;
    dispatch(setMatchesOffset(newOffset));
    await loadMatches(newOffset);
  }, [dispatch, hasMoreMatches, isLoadingIncoming, matchesOffset, loadMatches]);

  const loadMoreLikes = useCallback(async () => {
    if (!hasMoreLikes || isLoadingIncoming) return;

    const newOffset = likesOffset + ITEMS_PER_PAGE;
    dispatch(setLikesOffset(newOffset));
    await loadLikes(newOffset);
  }, [dispatch, hasMoreLikes, isLoadingIncoming, likesOffset, loadLikes]);

  const handleDislike = useCallback(
    async (id: string) => {
      if (!id) {
        console.error('Invalid like ID provided');
        return;
      }
      try {
        await dislike({
          targetUserId: id,
        });
        dispatch(removeLike(id));
      } catch (error) {
        console.error('Error creating match:', error);
      }
    },
    [dispatch, dislike],
  );

  const handleLike = useCallback(
    async (id: string) => {
      const likedProfile = likes.find(like => like.senderId === id);

      console.log('handleLike called with id:', id);
      console.log('likedProfile found:', likedProfile);

      try {
        // Case 1: This is a profile that has already liked the user
        if (likedProfile) {
          // Immediately remove the like from the UI to provide instant feedback
          dispatch(removeLike(id));

          const response = await likeUser({
            targetUserId: likedProfile.senderId,
            message: '',
          }).unwrap();

          console.log('Like response:', response);

          if (response.isMatch) {
            console.log('Match detected!');
            // Get user data for the match
            const userResponse = await getUsers({ userIds: [likedProfile.senderId] }).unwrap();
            const userData = userResponse.users[0].user;

            console.log('User data for match:', userData);

            // Create the profile object for the match modal
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

            console.log('Created like profile:', likeProfile);

            // First add the match to the matches list
            dispatch(addMatch(userData));

            // Then reload the matches list to ensure it's fully updated
            await loadMatches(0);

            // Finally show the match modal
            dispatch(setMatchedProfile(likeProfile));
            dispatch(setShowMatchModal(true));

            console.log('Match modal state updated');
          }
        }
        // Case 2: This is a new like (not a match yet)
        else {
          console.log('Sending new like to user:', id);

          // Send the like
          const response = await likeUser({
            targetUserId: id,
            message: '',
          }).unwrap();

          console.log('New like response:', response);

          // If it's a match (the other user had already liked this user through another channel)
          if (response.isMatch) {
            console.log('Unexpected match detected!');

            // Get user data for the match
            const userResponse = await getUsers({ userIds: [id] }).unwrap();
            const userData = userResponse.users[0].user;

            console.log('User data for unexpected match:', userData);

            // Create the profile object for the match modal
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

            // Add the match and show the modal
            dispatch(addMatch(userData));
            await loadMatches(0);
            dispatch(setMatchedProfile(likeProfile));
            dispatch(setShowMatchModal(true));

            console.log('Match modal state updated for unexpected match');
          }
        }
      } catch (error) {
        console.error('Error processing like:', error);
        // If there was an error and we removed a like, add it back
        if (likedProfile) {
          await loadLikes(0);
        }
      }
    },
    [dispatch, likeUser, likes, getUsers, loadMatches, loadLikes],
  );

  const handleSendMessage = useCallback(() => {
    dispatch(setShowMatchModal(false));
  }, [dispatch]);

  const handleKeepSwiping = useCallback(async () => {
    dispatch(setShowMatchModal(false));

    // Reload both likes and matches lists to ensure UI is up-to-date
    await Promise.all([loadMatches(0), loadLikes(0)]);
  }, [dispatch, loadMatches, loadLikes]);

  const handleSocialClick = useCallback(
    (matchId: number) => {
      dispatch(setShowChatMessage(matchId));

      setTimeout(() => {
        dispatch(setShowChatMessage(null));
      }, 2000);
    },
    [dispatch],
  );

  const handleRemoveMatch = useCallback(
    async (id: string) => {
      if (!id) {
        console.error('Invalid match ID provided');
        return;
      }
      try {
        await deleteMatch({
          targetUserId: id,
        });
        dispatch(removeMatch(id));
      } catch (error) {
        console.error('Error removing match:', error);
      }
    },
    [dispatch, deleteMatch],
  );

  return {
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
    matchesUsers: matches,
    likesUsers,
    handleDislike,
    handleRemoveMatch,
    loadMoreMatches,
    loadMoreLikes,
    hasMoreMatches,
    hasMoreLikes,
  };
};
