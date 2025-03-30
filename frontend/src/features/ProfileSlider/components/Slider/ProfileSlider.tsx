import { useState } from "react";
import { useSwipeable } from "react-swipeable";
import styles from "./profileSlider.module.css";
import { DislikeButton } from "../../../../shared/components/DislikeButton";
import { LikeButton } from "../../../../shared/components/LikeButton";
import { ProfileView } from "../../../../pages/LikesPage/components/ProfileView";


const profiles = [
  {
    id: 1,
    name: "Анна",
    age: 25,
    bio: "Люблю путешествовать и заниматься спортом.",
    imageUrl: "man1.jpg",
    location: "Москва, Россия",
    interests: ["Музыка", "Путешествия", "Фотография", "Спорт", "Искусство"],
    lifestyle: {
      kids: "Нет детей",
      smoking: "Не курю",
      alcohol: "Иногда",
      education: "Высшее образование"
    },
    passion: ["Музыка", "Путешествия", "Фотография", "Спорт", "Искусство"]
  },
  {
    id: 2,
    name: "Иван",
    age: 30,
    bio: "Пишу музыку и люблю кататься на велосипеде.",
    imageUrl: "photo1.png",
    location: "Санкт-Петербург, Россия",
    interests: ["Музыка", "Велоспорт", "Технологии", "Кино", "Путешествия"],
    lifestyle: {
      kids: "Нет детей",
      smoking: "Не курю",
      alcohol: "Редко",
      education: "Высшее образование"
    },
    passion: ["Музыка", "Велоспорт", "Технологии", "Кино", "Путешествия"]
  },
  {
    id: 3,
    name: "Ольга",
    age: 28,
    bio: "Увлекаюсь фотографией и кулинарией.",
    imageUrl: "woman1.jpg",
    location: "Казань, Россия",
    interests: ["Фотография", "Кулинария", "Книги", "Йога", "Природа"],
    lifestyle: {
      kids: "Нет детей",
      smoking: "Не курю",
      alcohol: "Не пью",
      education: "Высшее образование"
    },
    passion: ["Фотография", "Кулинария", "Книги", "Йога", "Природа"]
  },
  // Добавьте больше профилей по потребности
];

export const ProfileSlider = () => {
  const [currentIndex, setCurrentIndex] = useState(0);
  const [selectedProfile, setSelectedProfile] = useState<typeof profiles[0] | null>(null);

  const handleNext = () => {
    setCurrentIndex((prevIndex) => (prevIndex + 1) % profiles.length);
  };

  const handlePrev = () => {
    setCurrentIndex(
      (prevIndex) => (prevIndex - 1 + profiles.length) % profiles.length
    );
  };

  const handlers = useSwipeable({
    onSwipedLeft: handleNext,
    onSwipedRight: handlePrev,
    trackMouse: true,
  });

  const currentProfile = profiles[currentIndex];

  const handleProfileClick = () => {
    setSelectedProfile(currentProfile);
  };

  const handleCloseProfile = () => {
    setSelectedProfile(null);
  };

  const handleLike = (id: number) => {
    // Handle like functionality if needed
    console.log(`Liked profile with id: ${id}`);
    handleCloseProfile();
  };

  return (
    <div className={styles.slider}>
      <div className={styles.profileСard} {...handlers} onClick={handleProfileClick}>
        <img
          src={currentProfile.imageUrl}
          alt={currentProfile.name}
          className={styles.profileImage}
        />
        <h2 className={styles.name}>
          {currentProfile.name}, {currentProfile.age}
        </h2>
        {/* <p>{currentProfile.bio}</p> */}
      </div>
      <div className={styles.controls}>
        {/* <div className={styles.likeCircle} onClick={handlePrev}>
          <Like />
        </div>

        <div className={styles.dislikeCircle} onClick={handleNext}>
          <Dislike />
        </div> */}

        <LikeButton onClick={handlePrev} className={styles.likeButton} />
        <DislikeButton onClick={handleNext} className={styles.dislikeButton} />
      </div>

      {selectedProfile && (
        <ProfileView
          profile={{
            id: selectedProfile.id,
            name: selectedProfile.name,
            age: selectedProfile.age,
            imageUrl: selectedProfile.imageUrl,
            bio: selectedProfile.bio,
            location: selectedProfile.location,
            interests: selectedProfile.interests,
            lifestyle: selectedProfile.lifestyle,
            passion: selectedProfile.passion
          }}
          onClose={handleCloseProfile}
          onLike={handleLike}
        />
      )}
    </div>
  );
};

export default ProfileSlider;
