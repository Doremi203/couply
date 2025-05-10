import FavoriteBorderOutlinedIcon from '@mui/icons-material/FavoriteBorderOutlined';
import VerifiedIcon from '@mui/icons-material/Verified';
import React, { useEffect, useState, useRef } from 'react';
import { useSwipeable } from 'react-swipeable';

import { GenderPriority } from '../../../../entities/search';
import {
  useCreateFilterMutation,
  useSearchUsersMutation,
} from '../../../../entities/search/api/searchApi';
import {
  Alcohol,
  Art,
  Education,
  Gastronomy,
  Goal,
  Hobby,
  Selfdevelopment,
  Smoking,
  Social,
  Sport,
  Zodiac,
  Children,
} from '../../../../entities/user';
import { MessageModal } from '../../../../pages/HomePage/components/MessageModal/MessageModal';
import { NoUsersLeft } from '../../../../pages/HomePage/components/NoUsersLeft/NoUsersLeft';
import { DislikeButton } from '../../../../shared/components/DislikeButton';
import { LikeButton } from '../../../../shared/components/LikeButton';
//@ts-ignore
import MessageButton from '../../../../shared/components/MessageButton/messageButton';
import UndoButton from '../../../../shared/components/UndoButton/UndoButton';
import { PremiumModal } from '../../../../widgets/PremiumModal';
import { ProfileView } from '../../../../widgets/ProfileView';
import { ComplaintModal } from '../ComplaintModal/CompliantModal';

import styles from './profileSlider.module.css';

const profiles = [
  {
    id: 1,
    name: 'Анна',
    age: 25,
    bio: '',
    photos: ['man1.jpg', 'man1.jpg', 'man1.jpg'],
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
    photos: ['photo1.png', 'photo2.png', 'photo3.png'],
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
    photos: ['woman1.jpg', 'woman1.jpg', 'woman1.jpg'],
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
  {
    id: 3,
    name: 'Превиак',
    age: 20,
    bio: 'пидорское млекопитающее',
    photos: ['previak.jpg'],
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

// Advertisement profiles
const adProfiles = [
  {
    id: 'ad1',
    isAd: true,
    adText: 'Ебать какая классная игра всем советую!!',
    adLink: 'https://t.me/cactus_carnage',
    photos: ['cactus3.jpg'],
    name: 'Cactus Carnage',
  },
  // {
  //   id: 'ad2',
  //   isAd: true,
  //   adText: 'Получите премиум-подписку со скидкой 50%!',
  //   adLink: 'https://example.com/premium',
  //   photos: ['photo1.png'],
  // },
];

const getDefaultFilter = () => {
  return {
    genderPriority: GenderPriority.male,
    minAge: 18,
    maxAge: 100,
    minHeight: 100,
    maxHeight: 250,
    distance: 100,
    goal: Goal.unspecified,
    zodiac: Zodiac.unspecified,
    education: Education.unspecified,
    children: Children.unspecified,
    alcohol: Alcohol.unspecified,
    smoking: Smoking.unspecified,
    interest: {
      sport: [Sport.unspecified],
      selfDevelopment: [Selfdevelopment.unspecified],
      art: [Art.unspecified],
      social: [Social.unspecified],
      hobby: [Hobby.unspecified],
      gastronomy: [Gastronomy.unspecified],
    },
    onlyVerified: false,
    onlyPremium: false,
  };
};

export const ProfileSlider = () => {
  const [searchUsers] = useSearchUsersMutation();
  const [createFilter] = useCreateFilterMutation();

  // const [profiles1, setProfiles] = useState<(typeof profiles)[0][]>([]);
  // const [loading, setLoading] = useState(true);
  // const [error, setError] = useState<string | null>(null);
  const [currentIndex, setCurrentIndex] = useState(0);
  const [currentPhotoIndex, setCurrentPhotoIndex] = useState(0);
  const [selectedProfile, setSelectedProfile] = useState<(typeof profiles)[0] | null>(null);
  const [swipeCount, setSwipeCount] = useState(0);
  const [showingAd, setShowingAd] = useState(false);
  const [adIndex, setAdIndex] = useState(0);
  const [timer, setTimer] = useState(5);
  const [timerActive, setTimerActive] = useState(false);
  //@ts-ignore
  const timerRef = useRef<NodeJS.Timeout | null>(null);

  const [complaintOpen, setComplaintOpen] = useState(false);
  const [messageOpen, setMessageOpen] = useState(false);
  const [premiumOpen, setPremiumOpen] = useState(false);

  useEffect(() => {
    const fetchData = async () => {
      try {
        //@ts-ignore
        setLoading(true);
        const defaultFilter = getDefaultFilter();
        //@ts-ignore
        await createFilter(defaultFilter).unwrap();
        const response = await searchUsers({ limit: 10, offset: 0 }).unwrap();
        console.log(response);
        //@ts-ignore
        setProfiles(response.users || []);
      } catch (err) {
        //@ts-ignore
        setError('Ошибка при загрузке профилей');
        console.error(err);
      } finally {
        //@ts-ignore
        setLoading(false);
      }
    };

    fetchData();
  }, [createFilter, searchUsers]);

  const handleNextUser = () => {
    // If showing an ad and timer is active, don't allow swiping
    if (showingAd && timerActive) {
      return;
    }

    // Increment swipe count if not showing an ad
    if (!showingAd) {
      const newSwipeCount = swipeCount + 1;
      setSwipeCount(newSwipeCount);

      // Show ad after every 3 swipes
      if (newSwipeCount % 3 === 0) {
        setShowingAd(true);
        setAdIndex((adIndex + 1) % adProfiles.length);
        setTimerActive(true);
        setTimer(5);

        // Start the countdown timer
        if (timerRef.current) {
          clearInterval(timerRef.current);
        }

        timerRef.current = setInterval(() => {
          setTimer(prevTimer => {
            if (prevTimer <= 1) {
              if (timerRef.current) {
                clearInterval(timerRef.current);
                timerRef.current = null;
              }
              setTimerActive(false);
              return 0;
            }
            return prevTimer - 1;
          });
        }, 1000);

        return;
      }
    } else {
      // If showing an ad, reset to show regular profiles
      setShowingAd(false);
      if (timerRef.current) {
        clearInterval(timerRef.current);
        timerRef.current = null;
      }
      setTimerActive(false);
    }

    setCurrentIndex(prev => prev + 1);
    setCurrentPhotoIndex(0);
  };

  const handlePrevUser = () => {
    if (currentIndex === 0) return;
    setCurrentIndex(prev => prev - 1);
    setCurrentPhotoIndex(0);
  };

  const handleNextPhoto = () => {
    const currentUser = profiles[currentIndex];
    setCurrentPhotoIndex(prevIndex => (prevIndex + 1) % currentUser.photos.length);
  };

  const handlePrevPhoto = () => {
    const currentUser = profiles[currentIndex];
    setCurrentPhotoIndex(
      prevIndex => (prevIndex - 1 + currentUser.photos.length) % currentUser.photos.length,
    );
  };

  const handlers = useSwipeable({
    onSwipedLeft: () => handleNextUser(),
    onSwipedRight: () => handleNextUser(),
    trackMouse: true,
    // @ts-ignore
    preventDefaultTouchmoveEvent: !timerActive, // Prevent swiping when timer is active
  });

  // Clean up timer on unmount
  useEffect(() => {
    return () => {
      if (timerRef.current) {
        clearInterval(timerRef.current);
      }
    };
  }, []);

  if (currentIndex >= profiles.length && !showingAd) {
    return <NoUsersLeft />;
  }

  const currentProfile = showingAd
    ? adProfiles[adIndex % adProfiles.length]
    : profiles[currentIndex];

  // if (loading) {
  //   return <div className={styles.loading}>Загрузка...</div>;
  // }

  // if (error) {
  //   return <div className={styles.error}>{error}</div>;
  // }

  // if (!profiles.length) {
  //   return <div className={styles.empty}>Нет профилей для отображения</div>;
  // }

  const handleProfileClick = (e: React.MouseEvent<HTMLDivElement>) => {
    // If it's an ad, navigate to the ad link
    if (showingAd && 'adLink' in currentProfile) {
      window.open(currentProfile.adLink, '_blank');
      return;
    }

    const rect = e.currentTarget.getBoundingClientRect();
    const clickPosition = e.clientX - rect.left;
    const width = rect.width;

    if (clickPosition > width * 0.75) {
      handleNextPhoto();
    } else if (clickPosition < width * 0.25) {
      handlePrevPhoto();
    } else {
      // @ts-ignore
      setSelectedProfile(currentProfile);
    }
  };

  const handleCloseProfile = () => {
    setSelectedProfile(null);
  };

  const handleLike = () => {
    handleCloseProfile();
  };

  const renderName = (nameClass: string) => {
    return (
      <div className={nameClass}>
        {/** @ts-ignore */}
        {currentProfile.name}, {currentProfile.age}
        {/* @ts-ignore */}
        {currentProfile.verified && (
          <div className={styles.verifiedBadge}>
            <VerifiedIcon />
          </div>
        )}
      </div>
    );
  };

  const renderProfileInfo = () => {
    switch (currentPhotoIndex) {
      case 0: {
        let bioLines = 0;

        //@ts-ignore
        if (currentProfile.bio.length > 0 && currentProfile.bio.length <= 50) {
          bioLines = 1;
          //@ts-ignore
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
          <>
            {renderName(nameClass)}
            {/** @ts-ignore */}
            <div className={styles.bio}>{currentProfile.bio}</div>
          </>
        );
      }
      case 1: {
        const nameClass = styles.nameWithGoal;
        return (
          <>
            {renderName(nameClass)}
            <div className={styles.goal}>
              <FavoriteBorderOutlinedIcon className={styles.goalIcon} fontSize="small" />
              <span>Отношения</span>
            </div>
          </>
        );
      }
      case 2: {
        const nameClass = styles.nameWithInterests;
        return (
          <>
            {renderName(nameClass)}
            <div className={styles.interests}>
              <div className={styles.interestsList}>
                {/** @ts-ignore */}
                {currentProfile.interests.slice(0, 3).map((interest, index) => (
                  <span key={index} className={styles.interestTag}>
                    {interest}
                  </span>
                ))}
              </div>
            </div>
          </>
        );
      }
      default:
        return null;
    }
  };

  const onBlock = () => {
    setComplaintOpen(true);
  };

  const isAd = showingAd;

  return (
    <div className={styles.slider}>
      <div className={styles.profileCard} {...handlers} onClick={handleProfileClick}>
        <img
          src={currentProfile.photos[currentPhotoIndex]}
          // src="man1.jpg"
          alt={currentProfile.name}
          className={styles.profileImage}
          draggable="false"
        />

        {!isAd && (
          <>
            <div className={styles.distance}>0 км</div>
            <img
              src="data:image/png;base64,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"
              alt="error--v1"
              className={styles.block}
              width="20px"
              height="20px"
              onClick={onBlock}
            />
          </>
        )}

        {!isAd && renderProfileInfo()}

        {isAd && (
          <div>
            <div className={styles.nameWithBioOne}>{currentProfile.name}</div>
            <div className={styles.bio}>
              {'adText' in currentProfile ? currentProfile.adText : ''}
            </div>
            {timerActive && <div className={styles.adTimer}>{timer}</div>}
            <div className={styles.adText}>Реклама</div>
          </div>
        )}

        {!isAd && (
          <div className={styles.photoCounter}>
            {currentPhotoIndex + 1}/{currentProfile.photos.length}
          </div>
        )}
      </div>
      <div className={styles.controls}>
        <UndoButton onClick={handlePrevUser} />
        <DislikeButton onClick={handleNextUser} className={styles.dislikeButton} />
        <LikeButton
          onClick={handleNextUser}
          className={styles.likeButton}
          likeClassName={styles.like}
        />
        <MessageButton onClick={() => setPremiumOpen(true)} />
      </div>

      {selectedProfile && (
        <ProfileView
          profile={{
            ...selectedProfile,
            imageUrl: selectedProfile.photos[currentPhotoIndex],
          }}
          onClose={handleCloseProfile}
          onLike={handleLike}
        />
      )}

      <ComplaintModal isOpen={complaintOpen} onClose={() => setComplaintOpen(false)} />
      <MessageModal isOpen={messageOpen} onClose={() => setMessageOpen(false)} />
      <PremiumModal isOpen={premiumOpen} onClose={() => setPremiumOpen(false)} />
    </div>
  );
};

export default ProfileSlider;
