import VerifiedIcon from '@mui/icons-material/Verified';
import { useState } from 'react';
import { useSwipeable } from 'react-swipeable';

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
    images: ['man1.jpg', 'man2.jpg', 'man3.jpg'], // Изменено на массив фотографий
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
    images: ['woman1.jpg', 'woman2.jpg', 'woman3.jpg'],
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
  const [currentIndex, setCurrentIndex] = useState(0);
  const [currentPhotoIndex, setCurrentPhotoIndex] = useState(0); // Новое состояние для текущей фотографии
  const [selectedProfile, setSelectedProfile] = useState<(typeof profiles)[0] | null>(null);

  const handleNextUser = () => {
    setCurrentIndex(prevIndex => (prevIndex + 1) % profiles.length);
    setCurrentPhotoIndex(0); // Сбрасываем индекс фото при переходе к новому пользователю
  };

  const handlePrevUser = () => {
    setCurrentIndex(prevIndex => (prevIndex - 1 + profiles.length) % profiles.length);
    setCurrentPhotoIndex(0); // Сбрасываем индекс фото при переходе к новому пользователю
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

  const handlers = useSwipeable({
    onSwipedLeft: handleNextUser,
    onSwipedRight: handlePrevUser,
    trackMouse: true,
  });

  const currentProfile = profiles[currentIndex];

  const handleProfileClick = (e: React.MouseEvent<HTMLDivElement>) => {
    const rect = e.currentTarget.getBoundingClientRect();
    const clickPosition = e.clientX - rect.left;
    const width = rect.width;

    // Если клик в правой четверти изображения - следующее фото
    if (clickPosition > width * 0.75) {
      handleNextPhoto();
    }
    // Если клик в левой четверти изображения - предыдущее фото
    else if (clickPosition < width * 0.25) {
      handlePrevPhoto();
    }
    // Если клик по центру - открываем профиль
    else {
      setSelectedProfile(currentProfile);
    }
  };

  const handleCloseProfile = () => {
    setSelectedProfile(null);
  };

  const handleLike = () => {
    handleCloseProfile();
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
          src={currentProfile.images[currentPhotoIndex]} // Используем текущую фотографию
          alt={currentProfile.name}
          className={styles.profileImage}
          draggable="false" // Запрещаем перетаскивание изображения
        />
        <div className={nameClass}>
          {currentProfile.name}, {currentProfile.age}
          {currentProfile.verified && (
            <div className={styles.verifiedBadge}>
              <VerifiedIcon />
            </div>
          )}
        </div>
        <div className={styles.bio}>{currentProfile.bio}</div>
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
