// import { useState } from 'react';

// import { sendMatchNotification } from '../../../shared/lib/services/MatchNotificationService';
// import { likesData, matchesData, generateMatchId } from '../mockData';
// import { LikeProfile, MatchProfile } from '../types';

// export const useLikesAndMatches = () => {
//   const [likes, setLikes] = useState<LikeProfile[]>(likesData);
//   const [matches, setMatches] = useState<MatchProfile[]>(matchesData);
//   const [showMatchModal, setShowMatchModal] = useState(false);
//   const [matchedProfile, setMatchedProfile] = useState<LikeProfile | null>(null);
//   const [showChatMessage, setShowChatMessage] = useState<number | null>(null);

//   const handleLike = (id: number) => {
//     // Check if this is a like from the likes tab (not from matches tab)
//     // Match IDs are in a different range (101+) than like IDs
//     const isFromLikesTab = id < 100;

//     if (isFromLikesTab) {
//       // Find the profile that was liked
//       const likedProfile = likes.find(like => like.id === id);

//       if (likedProfile) {
//         // Check if this profile has already liked the user
//         if (likedProfile.hasLikedYou) {
//           // It's a match! Show the match modal
//           setMatchedProfile(likedProfile);
//           setShowMatchModal(true);

//           // Send push notification
//           sendMatchNotification({
//             userId: 'user123', // В реальном приложении это должно быть получено из аутентификации
//             matchId: likedProfile.id,
//             matchName: likedProfile.name,
//             matchImage: likedProfile.imageUrl,
//           });
//         }

//         // Update the liked status
//         const updatedLikes = likes.map(like => (like.id === id ? { ...like, liked: true } : like));

//         // Create a new match object with a unique ID
//         const newMatch: MatchProfile = {
//           id: generateMatchId(matches), // Generate a unique ID
//           name: likedProfile.name,
//           age: likedProfile.age,
//           imageUrl: likedProfile.imageUrl,
//           telegram: `@${likedProfile.name.toLowerCase()}_${likedProfile.age}`,
//           instagram: `@${likedProfile.name.toLowerCase()}_insta`,
//         };

//         // Add to matches and remove from likes if it was a match
//         if (likedProfile.hasLikedYou) {
//           setMatches([...matches, newMatch]);
//           setLikes(updatedLikes.filter(like => like.id !== id));
//         } else {
//           setLikes(updatedLikes);
//         }
//       }
//     }
//     // If it's from matches tab (id >= 100), do nothing
//   };

//   const handleSendMessage = () => {
//     setShowMatchModal(false);
//   };

//   const handleKeepSwiping = () => {
//     setShowMatchModal(false);
//   };

//   const handleSocialClick = (matchId: number, type: 'telegram' | 'instagram') => {
//     // Show a message when clicking on social media buttons
//     setShowChatMessage(matchId);

//     // Log which social media was clicked
//     console.log(`Opening ${type} for match ID ${matchId}`);

//     // Hide the message after 2 seconds
//     setTimeout(() => {
//       setShowChatMessage(null);
//     }, 2000);
//   };

//   return {
//     likes,
//     matches,
//     showMatchModal,
//     matchedProfile,
//     showChatMessage,
//     handleLike,
//     handleSendMessage,
//     handleKeepSwiping,
//     handleSocialClick,
//   };
// };

// TODO delete
