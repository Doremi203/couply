// import React, { useEffect } from 'react';
// import { useGetProfilesQuery } from '../../../../entities/profile/api/profileApi';
// import { useCreateLikeMutation } from '../../../../entities/likes/api/likesApi';
// import { useAppSelector, useAppDispatch } from '../../../../shared/lib/hooks/redux';
// import { 
//   selectFilters, 
//   setAgeRange, 
//   setDistance,
//   resetFilters
// } from '../../../../features/filters/model/filtersSlice';
// import styles from './homePage.module.css';

// const HomePage: React.FC = () => {
//   const dispatch = useAppDispatch();
//   const filters = useAppSelector(selectFilters);
  
//   // Получение профилей с использованием RTK Query
//   const { data: profiles, isLoading, error } = useGetProfilesQuery();
  
//   // Мутация для создания лайка
//   const [createLike, { isLoading: isLikeLoading }] = useCreateLikeMutation();
  
//   // Пример использования действий Redux
//   const handleResetFilters = () => {
//     dispatch(resetFilters());
//   };
  
//   const handleAgeRangeChange = (min: number, max: number) => {
//     dispatch(setAgeRange([min, max]));
//   };
  
//   const handleDistanceChange = (distance: number) => {
//     dispatch(setDistance(distance));
//   };
  
//   // Пример обработчика лайка
//   const handleLike = async (profileId: string) => {
//     try {
//       await createLike({ targetUserId: profileId }).unwrap();
//       // Обработка успешного лайка
//     } catch (error) {
//       // Обработка ошибки
//       console.error('Failed to like profile:', error);
//     }
//   };
  
//   // Фильтрация профилей на основе состояния фильтров
//   const filteredProfiles = React.useMemo(() => {
//     if (!profiles) return [];
    
//     return profiles.filter(profile => {
//       const age = profile.age;
//       const [minAge, maxAge] = filters.ageRange;
      
//       // Фильтрация по возрасту
//       if (age < minAge || age > maxAge) {
//         return false;
//       }
      
//       // Фильтрация по интересам
//       if (filters.interests.length > 0) {
//         const hasMatchingInterests = profile.interests.some(interest => 
//           filters.interests.includes(interest)
//         );
//         if (!hasMatchingInterests) {
//           return false;
//         }
//       }
      
//       // Фильтрация по наличию фото
//       if (filters.showOnlyWithPhoto && profile.photos.length === 0) {
//         return false;
//       }
      
//       return true;
//     });
//   }, [profiles, filters]);
  
//   // Пример использования useEffect с Redux
//   useEffect(() => {
//     // Можно выполнить какие-то действия при изменении фильтров
//     console.log('Filters changed:', filters);
//   }, [filters]);
  
//   if (isLoading) {
//     return <div>Loading profiles...</div>;
//   }
  
//   if (error) {
//     return <div>Error loading profiles</div>;
//   }
  
//   return (
//     <div className={styles.homePage}>
//       <div className={styles.filtersSection}>
//         <h2>Filters</h2>
//         <div className={styles.filterControls}>
//           <div>
//             <label>Age Range: {filters.ageRange[0]} - {filters.ageRange[1]}</label>
//             <div>
//               <input
//                 type="range"
//                 min="18"
//                 max="80"
//                 value={filters.ageRange[0]}
//                 onChange={(e) => handleAgeRangeChange(parseInt(e.target.value), filters.ageRange[1])}
//               />
//               <input
//                 type="range"
//                 min="18"
//                 max="80"
//                 value={filters.ageRange[1]}
//                 onChange={(e) => handleAgeRangeChange(filters.ageRange[0], parseInt(e.target.value))}
//               />
//             </div>
//           </div>
          
//           <div>
//             <label>Distance: {filters.distance} km</label>
//             <input
//               type="range"
//               min="1"
//               max="100"
//               value={filters.distance}
//               onChange={(e) => handleDistanceChange(parseInt(e.target.value))}
//             />
//           </div>
          
//           <div>
//             <label>
//               <input 
//                 type="checkbox" 
//                 checked={filters.showOnlyWithPhoto}
//                 onChange={() => dispatch({ type: 'filters/toggleShowOnlyWithPhoto' })}
//               />
//               Show only with photo
//             </label>
//           </div>
          
//           <button onClick={handleResetFilters}>Reset Filters</button>
//         </div>
//       </div>
      
//       <div className={styles.profilesSection}>
//         <h2>Profiles ({filteredProfiles.length})</h2>
//         {filteredProfiles.length === 0 ? (
//           <div>No profiles match your filters</div>
//         ) : (
//           <div className={styles.profilesList}>
//             {filteredProfiles.map(profile => (
//               <div key={profile.id} className={styles.profileCard}>
//                 <h3>{profile.name}, {profile.age}</h3>
//                 {profile.photos.length > 0 && (
//                   <img 
//                     src={profile.photos[0]} 
//                     alt={`${profile.name}'s photo`} 
//                     className={styles.profilePhoto}
//                   />
//                 )}
//                 <p>{profile.bio}</p>
//                 <div className={styles.interests}>
//                   {profile.interests.map(interest => (
//                     <span key={interest} className={styles.interestTag}>
//                       {interest}
//                     </span>
//                   ))}
//                 </div>
//                 <div className={styles.actions}>
//                   <button 
//                     onClick={() => handleLike(profile.id)}
//                     disabled={isLikeLoading}
//                   >
//                     Like
//                   </button>
//                   {/* Другие действия */}
//                 </div>
//               </div>
//             ))}
//           </div>
//         )}
//       </div>
//     </div>
//   );
// };

// export default HomePage;