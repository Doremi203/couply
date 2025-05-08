/* 1 */
// import VerifiedIcon from '@mui/icons-material/Verified';
// import { useState } from 'react';
// import { useSwipeable } from 'react-swipeable';

// import { DislikeButton } from '../../../../shared/components/DislikeButton';
// import { LikeButton } from '../../../../shared/components/LikeButton';
// import { ProfileView } from '../../../../widgets/ProfileView';

// import styles from './profileSlider.module.css';

// const profiles = [
//   {
//     id: 1,
//     name: 'Анна',
//     age: 25,
//     bio: '',
//     images: ['man1.jpg', 'man2.jpg', 'man3.jpg'], // Изменено на массив фотографий
//     location: 'Москва, Россия',
//     verified: true,
//     interests: ['Музыка', 'Путешествия', 'Фотография', 'Спорт', 'Искусство'],
//     lifestyle: {
//       kids: 'Нет детей',
//       smoking: 'Не курю',
//       alcohol: 'Иногда',
//       education: 'Высшее образование',
//     },
//     passion: ['Музыка', 'Путешествия', 'Фотография', 'Спорт', 'Искусство'],
//   },
//   {
//     id: 2,
//     name: 'Иван',
//     age: 30,
//     bio: 'Пишу музыку и люблю кататься на велосипедецшуатмщышоватмщшоывтамдоытвдаломтыдвоамтлдыоватмдлывоатмдылова.',
//     images: ['photo1.png', 'photo2.png', 'photo3.png'],
//     location: 'Санкт-Петербург, Россия',
//     verified: false,
//     interests: ['Музыка', 'Велоспорт', 'Технологии', 'Кино', 'Путешествия'],
//     lifestyle: {
//       kids: 'Нет детей',
//       smoking: 'Не курю',
//       alcohol: 'Редко',
//       education: 'Высшее образование',
//     },
//     passion: ['Музыка', 'Велоспорт', 'Технологии', 'Кино', 'Путешествия'],
//   },
//   {
//     id: 3,
//     name: 'Ольга',
//     age: 28,
//     bio: 'Увлекаюсь фотографией и кулинарией.',
//     images: ['woman1.jpg', 'woman2.jpg', 'woman3.jpg'],
//     location: 'Казань, Россия',
//     verified: true,
//     interests: ['Фотография', 'Кулинария', 'Книги', 'Йога', 'Природа'],
//     lifestyle: {
//       kids: 'Нет детей',
//       smoking: 'Не курю',
//       alcohol: 'Не пью',
//       education: 'Высшее образование',
//     },
//     passion: ['Фотография', 'Кулинария', 'Книги', 'Йога', 'Природа'],
//   },
// ];

// export const ProfileSlider = () => {
//   const [currentIndex, setCurrentIndex] = useState(0);
//   const [currentPhotoIndex, setCurrentPhotoIndex] = useState(0); // Новое состояние для текущей фотографии
//   const [selectedProfile, setSelectedProfile] = useState<(typeof profiles)[0] | null>(null);

//   const handleNextUser = () => {
//     setCurrentIndex(prevIndex => (prevIndex + 1) % profiles.length);
//     setCurrentPhotoIndex(0); // Сбрасываем индекс фото при переходе к новому пользователю
//   };

//   const handlePrevUser = () => {
//     setCurrentIndex(prevIndex => (prevIndex - 1 + profiles.length) % profiles.length);
//     setCurrentPhotoIndex(0); // Сбрасываем индекс фото при переходе к новому пользователю
//   };

//   const handleNextPhoto = () => {
//     const currentUser = profiles[currentIndex];
//     setCurrentPhotoIndex(prevIndex => (prevIndex + 1) % currentUser.images.length);
//   };

//   const handlePrevPhoto = () => {
//     const currentUser = profiles[currentIndex];
//     setCurrentPhotoIndex(
//       prevIndex => (prevIndex - 1 + currentUser.images.length) % currentUser.images.length,
//     );
//   };

//   const handlers = useSwipeable({
//     onSwipedLeft: handleNextUser,
//     onSwipedRight: handlePrevUser,
//     trackMouse: true,
//   });

//   const currentProfile = profiles[currentIndex];

//   const handleProfileClick = (e: React.MouseEvent<HTMLDivElement>) => {
//     const rect = e.currentTarget.getBoundingClientRect();
//     const clickPosition = e.clientX - rect.left;
//     const width = rect.width;

//     // Если клик в правой четверти изображения - следующее фото
//     if (clickPosition > width * 0.75) {
//       handleNextPhoto();
//     }
//     // Если клик в левой четверти изображения - предыдущее фото
//     else if (clickPosition < width * 0.25) {
//       handlePrevPhoto();
//     }
//     // Если клик по центру - открываем профиль
//     else {
//       setSelectedProfile(currentProfile);
//     }
//   };

//   const handleCloseProfile = () => {
//     setSelectedProfile(null);
//   };

//   const handleLike = () => {
//     handleCloseProfile();
//   };

//   let bioLines = 0;

//   if (currentProfile.bio.length > 0 && currentProfile.bio.length <= 50) {
//     bioLines = 1;
//   } else if (currentProfile.bio.length > 50) {
//     bioLines = 2;
//   }

//   const nameClass = [
//     styles.name,
//     bioLines === 1 && styles.nameWithBioOne,
//     bioLines === 2 && styles.nameWithBioTwo,
//   ]
//     .filter(Boolean)
//     .join(' ');

//   return (
//     <div className={styles.slider}>
//       <div className={styles.profileСard} {...handlers} onClick={handleProfileClick}>
//         <img
//           src={currentProfile.images[currentPhotoIndex]} // Используем текущую фотографию
//           alt={currentProfile.name}
//           className={styles.profileImage}
//           draggable="false" // Запрещаем перетаскивание изображения
//         />
//         <div className={nameClass}>
//           {currentProfile.name}, {currentProfile.age}
//           {currentProfile.verified && (
//             <div className={styles.verifiedBadge}>
//               <VerifiedIcon />
//             </div>
//           )}
//         </div>
//         <div className={styles.bio}>{currentProfile.bio}</div>
//         <div className={styles.photoCounter}>
//           {currentPhotoIndex + 1}/{currentProfile.images.length}
//         </div>
//       </div>
//       <div className={styles.controls}>
//         <DislikeButton onClick={handleNextUser} className={styles.dislikeButton} />
//         <LikeButton onClick={handlePrevUser} className={styles.likeButton} />
//       </div>

//       {selectedProfile && (
//         <ProfileView
//           profile={{
//             ...selectedProfile,
//             imageUrl: selectedProfile.images[currentPhotoIndex],
//           }}
//           onClose={handleCloseProfile}
//           onLike={handleLike}
//         />
//       )}
//     </div>
//   );
// };

// export default ProfileSlider;

/* 2 */

import VerifiedIcon from '@mui/icons-material/Verified';
import { useEffect, useState } from 'react';
import { useSelector } from 'react-redux';
import { useSwipeable } from 'react-swipeable';

import {
  useCreateFilterMutation,
  useSearchUsersMutation,
} from '../../../../entities/search/api/searchApi';
import { GenderPriority } from '../../../../entities/search/api/types';
import {
  Alcohol,
  Art,
  Education,
  Gastronomy,
  getUserId,
  Goal,
  Hobby,
  Selfdevelopment,
  Smoking,
  Social,
  Sport,
  Zodiac,
  Children,
} from '../../../../entities/user';
import { DislikeButton } from '../../../../shared/components/DislikeButton';
import { LikeButton } from '../../../../shared/components/LikeButton';
import { ProfileView } from '../../../../widgets/ProfileView';

import styles from './profileSlider.module.css';

const profiles = [
  {
    id: 1,
    name: 'Анна',
    age: 25,
    bio: '',
    images: ['man1.jpg', 'man1.jpg', 'man1.jpg'],
    location: 'Москва, Россия',
    verified: true,
    interests: ['Музыка', 'Путешествия', 'Фотография', 'Спорт', 'Искусство'],
    lifestyle: {
      kids: 'Нет детей',
      smoking: 'Не курю',
      alcohol: 'Иногда',
      education: 'Высшее образование',
    },
    passion: ['Музыка', 'Путешествия', 'Фотография', 'Спорт', 'Искусство'],
  },
  {
    id: 2,
    name: 'Иван',
    age: 30,
    bio: 'Пишу музыку и люблю кататься на велосипедецшуатмщышоватмщшоывтамдоытвдаломтыдвоамтлдыоватмдлывоатмдылова.',
    images: ['photo1.png', 'photo2.png', 'photo3.png'],
    location: 'Санкт-Петербург, Россия',
    verified: false,
    interests: ['Музыка', 'Велоспорт', 'Технологии', 'Кино', 'Путешествия'],
    lifestyle: {
      kids: 'Нет детей',
      smoking: 'Не курю',
      alcohol: 'Редко',
      education: 'Высшее образование',
    },
    passion: ['Музыка', 'Велоспорт', 'Технологии', 'Кино', 'Путешествия'],
  },
  {
    id: 3,
    name: 'Ольга',
    age: 28,
    bio: 'Увлекаюсь фотографией и кулинарией.',
    images: ['woman1.jpg', 'woman1.jpg', 'woman1.jpg'],
    location: 'Казань, Россия',
    verified: true,
    interests: ['Фотография', 'Кулинария', 'Книги', 'Йога', 'Природа'],
    lifestyle: {
      kids: 'Нет детей',
      smoking: 'Не курю',
      alcohol: 'Не пью',
      education: 'Высшее образование',
    },
    passion: ['Фотография', 'Кулинария', 'Книги', 'Йога', 'Природа'],
  },
];

export const ProfileSlider = () => {
  // const [searchUsers] = useSearchUsersMutation();
  // const [createFilter] = useCreateFilterMutation();

  // const userId = useSelector(getUserId);
  // const [profiles, setProfiles] = useState([]);

  // useEffect(() => {
  //   if (!userId) return; // Защита от пустого userId

  //   const fetchData = async () => {
  //     const defaultFilter = getDefaultFilter(userId);
  //     await createFilter(defaultFilter).unwrap(); // (опционально) если нужно дождаться

  //     const response = await searchUsers({ userId, limit: 10, offset: 0 }).unwrap();
  //     setProfiles(response.users); // response должен быть массивом
  //   };

  //   fetchData();
  // }, [userId]);

  // console.log(profiles);

  // const [currentIndex, setCurrentIndex] = useState(0);
  // const [currentPhotoIndex, setCurrentPhotoIndex] = useState(0);
  // const [selectedProfile, setSelectedProfile] = useState<(typeof profiles)[0] | null>(null);

  // const handleNextUser = () => {
  //   setCurrentIndex(prevIndex => (prevIndex + 1) % profiles.length);
  //   setCurrentPhotoIndex(0);
  // };

  const [searchUsers] = useSearchUsersMutation();
  const [createFilter] = useCreateFilterMutation();
  const userId = useSelector(getUserId);
  // const [loading, setLoading] = useState(true);
  // const [error, setError] = useState<string | null>(null);

  const [currentIndex, setCurrentIndex] = useState(0);
  const [currentPhotoIndex, setCurrentPhotoIndex] = useState(0);
  const [selectedProfile, setSelectedProfile] = useState<(typeof profiles)[0] | null>(null);

  // const handleNextUser = () => {
  //   setCurrentIndex(prevIndex => (prevIndex + 1) % profiles.length);
  //   setCurrentPhotoIndex(0);
  // };

  // const handlePrevUser = () => {
  //   setCurrentIndex(prevIndex => (prevIndex - 1 + profiles.length) % profiles.length);
  //   setCurrentPhotoIndex(0);
  // };

  const handleNextUser = () => {
    if (profiles.length === 0) return;
    setCurrentIndex(prev => (prev + 1) % profiles.length);
    setCurrentPhotoIndex(0);
  };

  const handlePrevUser = () => {
    if (profiles.length === 0) return;
    setCurrentIndex(prev => (prev - 1 + profiles.length) % profiles.length);
    setCurrentPhotoIndex(0);
  };

  const handleNextPhoto = () => {
    const currentUser = profiles[currentIndex];
    setCurrentPhotoIndex(prevIndex => (prevIndex + 1) % currentUser.images.length);
  };

  const handlePrevPhoto = () => {
    const currentUser = profiles[currentIndex];
    setCurrentPhotoIndex(
      prevIndex => (prevIndex - 1 + currentUser.images.length) % currentUser.images.length,
    );
  };

  // const handlers = useSwipeable({
  //   onSwipedLeft: handleNextUser,
  //   onSwipedRight: handlePrevUser,
  //   trackMouse: true,
  // });

  const handlers = useSwipeable({
    onSwipedLeft: () => handleNextUser(), // Свайп влево -> следующий
    onSwipedRight: () => handlePrevUser(), // Свайп вправо -> предыдущий
    trackMouse: true,
  });

  const currentProfile = profiles[currentIndex];

  // if (loading) {
  //   return <div className={styles.loading}>Загрузка...</div>;
  // }

  // if (error) {
  //   return <div className={styles.error}>{error}</div>;
  // }

  if (!profiles.length) {
    return <div className={styles.empty}>Нет профилей для отображения</div>;
  }

  const handleProfileClick = (e: React.MouseEvent<HTMLDivElement>) => {
    const rect = e.currentTarget.getBoundingClientRect();
    const clickPosition = e.clientX - rect.left;
    const width = rect.width;

    if (clickPosition > width * 0.75) {
      handleNextPhoto();
    } else if (clickPosition < width * 0.25) {
      handlePrevPhoto();
    } else {
      setSelectedProfile(currentProfile);
    }
  };

  const handleCloseProfile = () => {
    setSelectedProfile(null);
  };

  const handleLike = () => {
    handleCloseProfile();
  };

  // const renderProfileInfo = () => {
  //   if (!currentProfile) return null;

  //   switch (currentPhotoIndex) {
  //     case 0:
  //       return <div className={styles.bio}>{currentProfile.bio || 'Нет информации'}</div>;
  //     case 1:
  //       return (
  //         <div className={styles.lifestyle}>
  //           <h4>Образ жизни</h4>
  //           <ul>
  //             {currentProfile.lifestyle &&
  //               Object.entries(currentProfile.lifestyle).map(([key, value]) => (
  //                 <li key={key}>
  //                   <strong>{key}:</strong> {value}
  //                 </li>
  //               ))}
  //           </ul>
  //         </div>
  //       );
  //     case 2:
  //       return (
  //         <div className={styles.interests}>
  //           <div className={styles.interestsList}>
  //             {currentProfile.interests?.slice(0, 4).map((interest, index) => (
  //               <span key={index} className={styles.interestTag}>
  //                 {interest}
  //               </span>
  //             ))}
  //           </div>
  //         </div>
  //       );
  //     default:
  //       return null;
  //   }
  // };

  const renderProfileInfo = () => {
    switch (currentPhotoIndex) {
      case 0: // Первая фотография - показываем био
        return <div className={styles.bio}>{currentProfile.bio}</div>;
      case 1: // Вторая фотография - показываем lifestyle
        return (
          <div className={styles.lifestyle}>
            <h4>Образ жизни</h4>
            <ul>
              {Object.entries(currentProfile.lifestyle).map(([key, value]) => (
                <li key={key}>
                  <strong>{key}:</strong> {value}
                </li>
              ))}
            </ul>
          </div>
        );
      case 2: // Третья фотография - показываем интересы
        return (
          <div className={styles.interests}>
            <div className={styles.interestsList}>
              {currentProfile.interests.slice(0, 4).map((interest, index) => (
                <span key={index} className={styles.interestTag}>
                  {interest}
                </span>
              ))}
            </div>
          </div>
        );
      default:
        return null;
    }
  };

  let bioLines = 0;

  if (currentProfile.bio.length > 0 && currentProfile.bio.length <= 50) {
    bioLines = 1;
  } else if (currentProfile.bio.length > 50) {
    bioLines = 2;
  }

  const nameClass = [
    styles.name,
    bioLines === 1 && styles.nameWithBioOne,
    bioLines === 2 && styles.nameWithBioTwo,
  ]
    .filter(Boolean)
    .join(' ');

  return (
    <div className={styles.slider}>
      <div className={styles.profileСard} {...handlers} onClick={handleProfileClick}>
        <img
          src={currentProfile.images[currentPhotoIndex]}
          alt={currentProfile.name}
          className={styles.profileImage}
          draggable="false"
        />
        <div className={nameClass}>
          {currentProfile.name}, {currentProfile.age}
          {currentProfile.verified && (
            <div className={styles.verifiedBadge}>
              <VerifiedIcon />
            </div>
          )}
        </div>

        {renderProfileInfo()}

        <div className={styles.photoCounter}>
          {currentPhotoIndex + 1}/{currentProfile.images.length}
        </div>
      </div>
      <div className={styles.controls}>
        <DislikeButton onClick={handleNextUser} className={styles.dislikeButton} />
        <LikeButton onClick={handlePrevUser} className={styles.likeButton} />
      </div>

      {selectedProfile && (
        <ProfileView
          profile={{
            ...selectedProfile,
            imageUrl: selectedProfile.images[currentPhotoIndex],
          }}
          onClose={handleCloseProfile}
          onLike={handleLike}
        />
      )}
    </div>
  );
};

export default ProfileSlider;

/* 3 */

// import VerifiedIcon from '@mui/icons-material/Verified';
// import { useEffect, useState } from 'react';
// import { useSelector } from 'react-redux';
// import { useSwipeable } from 'react-swipeable';

// import {
//   useCreateFilterMutation,
//   useSearchUsersMutation,
// } from '../../../../entities/search/api/searchApi';
// import { GenderPriority } from '../../../../entities/search/api/types';
// import {
//   Alcohol,
//   Art,
//   Education,
//   Gastronomy,
//   Goal,
//   Hobby,
//   Selfdevelopment,
//   Smoking,
//   Social,
//   Sport,
//   Zodiac,
//   Children,
//   getUserId,
// } from '../../../../entities/user';
// import { DislikeButton } from '../../../../shared/components/DislikeButton';
// import { LikeButton } from '../../../../shared/components/LikeButton';
// import { ProfileView } from '../../../../widgets/ProfileView';

// import styles from './profileSlider.module.css';

// const profiles = [
//   {
//     id: 1,
//     name: 'Анна',
//     age: 25,
//     bio: '',
//     images: ['man1.jpg', 'man2.jpg', 'man3.jpg'], // Изменено на массив фотографий
//     location: 'Москва, Россия',
//     verified: true,
//     interests: ['Музыка', 'Путешествия', 'Фотография', 'Спорт', 'Искусство'],
//     lifestyle: {
//       kids: 'Нет детей',
//       smoking: 'Не курю',
//       alcohol: 'Иногда',
//       education: 'Высшее образование',
//     },
//     passion: ['Музыка', 'Путешествия', 'Фотография', 'Спорт', 'Искусство'],
//   },
//   {
//     id: 2,
//     name: 'Иван',
//     age: 30,
//     bio: 'Пишу музыку и люблю кататься на велосипедецшуатмщышоватмщшоывтамдоытвдаломтыдвоамтлдыоватмдлывоатмдылова.',
//     images: ['photo1.png', 'photo2.png', 'photo3.png'],
//     location: 'Санкт-Петербург, Россия',
//     verified: false,
//     interests: ['Музыка', 'Велоспорт', 'Технологии', 'Кино', 'Путешествия'],
//     lifestyle: {
//       kids: 'Нет детей',
//       smoking: 'Не курю',
//       alcohol: 'Редко',
//       education: 'Высшее образование',
//     },
//     passion: ['Музыка', 'Велоспорт', 'Технологии', 'Кино', 'Путешествия'],
//   },
//   {
//     id: 3,
//     name: 'Ольга',
//     age: 28,
//     bio: 'Увлекаюсь фотографией и кулинарией.',
//     images: ['woman1.jpg', 'woman2.jpg', 'woman3.jpg'],
//     location: 'Казань, Россия',
//     verified: true,
//     interests: ['Фотография', 'Кулинария', 'Книги', 'Йога', 'Природа'],
//     lifestyle: {
//       kids: 'Нет детей',
//       smoking: 'Не курю',
//       alcohol: 'Не пью',
//       education: 'Высшее образование',
//     },
//     passion: ['Фотография', 'Кулинария', 'Книги', 'Йога', 'Природа'],
//   },
// ];

// const getDefaultFilter = (userId: string) => {
//   return {
//     userId: userId,
//     genderPriority: GenderPriority.male,
//     minAge: 18,
//     maxAge: 100,
//     minHeight: 100,
//     maxHeight: 250,
//     distance: 100,
//     goal: Goal.unspecified,
//     zodiac: Zodiac.unspecified,
//     education: Education.unspecified,
//     children: Children.unspecified,
//     alcohol: Alcohol.unspecified,
//     smoking: Smoking.unspecified,
//     interest: {
//       sport: [Sport.unspecified],
//       selfDevelopment: [Selfdevelopment.unspecified],
//       art: [Art.unspecified],
//       social: [Social.unspecified],
//       hobby: [Hobby.unspecified],
//       gastronomy: [Gastronomy.unspecified],
//     },
//     onlyVerified: false,
//     onlyPremium: false,
//   };
// };

// export const ProfileSlider = () => {
//   // const [searchUsers] = useSearchUsersMutation();
//   // const [createFilter] = useCreateFilterMutation();

//   // const userId = useSelector(getUserId);
//   // const [profiles, setProfiles] = useState([]);

//   // useEffect(() => {
//   //   if (!userId) return; // Защита от пустого userId

//   //   const fetchData = async () => {
//   //     const defaultFilter = getDefaultFilter(userId);
//   //     await createFilter(defaultFilter).unwrap(); // (опционально) если нужно дождаться

//   //     const response = await searchUsers({ userId, limit: 10, offset: 0 }).unwrap();
//   //     setProfiles(response.users); // response должен быть массивом
//   //   };

//   //   fetchData();
//   // }, [userId]);

//   // console.log(profiles);

//   // const [currentIndex, setCurrentIndex] = useState(0);
//   // const [currentPhotoIndex, setCurrentPhotoIndex] = useState(0);
//   // const [selectedProfile, setSelectedProfile] = useState<(typeof profiles)[0] | null>(null);

//   // const handleNextUser = () => {
//   //   setCurrentIndex(prevIndex => (prevIndex + 1) % profiles.length);
//   //   setCurrentPhotoIndex(0);
//   // };

//   const [searchUsers] = useSearchUsersMutation();
//   const [createFilter] = useCreateFilterMutation();
//   const userId = useSelector(getUserId);
//   const [profiles, setProfiles] = useState<(typeof profiles)[0][]>([]);
//   const [loading, setLoading] = useState(true);
//   const [error, setError] = useState<string | null>(null);

//   useEffect(() => {
//     if (!userId) return;

//     const fetchData = async () => {
//       try {
//         setLoading(true);
//         const defaultFilter = getDefaultFilter(userId);
//         await createFilter(defaultFilter).unwrap();
//         const response = await searchUsers({ userId, limit: 10, offset: 0 }).unwrap();
//         setProfiles(response.users || []);
//       } catch (err) {
//         setError('Ошибка при загрузке профилей');
//         console.error(err);
//       } finally {
//         setLoading(false);
//       }
//     };

//     fetchData();
//   }, [userId]);

//   console.log(profiles);

//   const [currentIndex, setCurrentIndex] = useState(0);
//   const [currentPhotoIndex, setCurrentPhotoIndex] = useState(0); // Новое состояние для текущей фотографии
//   const [selectedProfile, setSelectedProfile] = useState<(typeof profiles)[0] | null>(null);

//   const handleNextUser = () => {
//     setCurrentIndex(prevIndex => (prevIndex + 1) % profiles.length);
//     setCurrentPhotoIndex(0); // Сбрасываем индекс фото при переходе к новому пользователю
//   };

//   const handlePrevUser = () => {
//     setCurrentIndex(prevIndex => (prevIndex - 1 + profiles.length) % profiles.length);
//     setCurrentPhotoIndex(0); // Сбрасываем индекс фото при переходе к новому пользователю
//   };

//   const handleNextPhoto = () => {
//     const currentUser = profiles[currentIndex];
//     setCurrentPhotoIndex(prevIndex => (prevIndex + 1) % currentUser.images.length);
//   };

//   const handlePrevPhoto = () => {
//     const currentUser = profiles[currentIndex];
//     setCurrentPhotoIndex(
//       prevIndex => (prevIndex - 1 + currentUser.images.length) % currentUser.images.length,
//     );
//   };

//   const handlers = useSwipeable({
//     onSwipedLeft: handleNextUser,
//     onSwipedRight: handlePrevUser,
//     trackMouse: true,
//   });

//   const currentProfile = profiles[currentIndex];

//   console.log(currentProfile);

//   if (loading) {
//     return <div className={styles.loading}>Загрузка...</div>;
//   }

//   if (error) {
//     return <div className={styles.error}>{error}</div>;
//   }

//   if (!profiles.length) {
//     return <div className={styles.empty}>Нет профилей для отображения</div>;
//   }

//   const handleProfileClick = (e: React.MouseEvent<HTMLDivElement>) => {
//     const rect = e.currentTarget.getBoundingClientRect();
//     const clickPosition = e.clientX - rect.left;
//     const width = rect.width;

//     // Если клик в правой четверти изображения - следующее фото
//     if (clickPosition > width * 0.75) {
//       handleNextPhoto();
//     }
//     // Если клик в левой четверти изображения - предыдущее фото
//     else if (clickPosition < width * 0.25) {
//       handlePrevPhoto();
//     }
//     // Если клик по центру - открываем профиль
//     else {
//       setSelectedProfile(currentProfile);
//     }
//   };

//   const handleCloseProfile = () => {
//     setSelectedProfile(null);
//   };

//   const handleLike = () => {
//     handleCloseProfile();
//   };

//   let bioLines = 0;

//   if (currentProfile.bio.length > 0 && currentProfile.bio.length <= 50) {
//     bioLines = 1;
//   } else if (currentProfile.bio.length > 50) {
//     bioLines = 2;
//   }

//   const nameClass = [
//     styles.name,
//     bioLines === 1 && styles.nameWithBioOne,
//     bioLines === 2 && styles.nameWithBioTwo,
//   ]
//     .filter(Boolean)
//     .join(' ');

//   return (
//     <div className={styles.slider}>
//       <div className={styles.profileСard} {...handlers} onClick={handleProfileClick}>
//         <img
//           src="man1.jpg" // Используем текущую фотографию
//           alt={currentProfile.name}
//           className={styles.profileImage}
//           draggable="false" // Запрещаем перетаскивание изображения
//         />
//         <div className={nameClass}>
//           {currentProfile.name}, {currentProfile.age}
//           {currentProfile.verified && (
//             <div className={styles.verifiedBadge}>
//               <VerifiedIcon />
//             </div>
//           )}
//         </div>
//         <div className={styles.bio}>{currentProfile.bio}</div>
//         <div className={styles.photoCounter}>
//           {/* {currentPhotoIndex + 1}/{currentProfile.images.length} */}
//         </div>
//       </div>
//       <div className={styles.controls}>
//         <DislikeButton onClick={handleNextUser} className={styles.dislikeButton} />
//         <LikeButton onClick={handlePrevUser} className={styles.likeButton} />
//       </div>

//       {selectedProfile && (
//         <ProfileView
//           profile={{
//             ...selectedProfile,
//             imageUrl: selectedProfile.images[currentPhotoIndex],
//           }}
//           onClose={handleCloseProfile}
//           onLike={handleLike}
//         />
//       )}
//     </div>
//   );
// };

// export default ProfileSlider;
