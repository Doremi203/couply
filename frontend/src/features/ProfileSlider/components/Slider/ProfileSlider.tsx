import { useState } from "react";
import { useSwipeable } from "react-swipeable";
import styles from "./profileSlider.module.css";
import { Dislike } from "../../../../shared/Dislike";
import { Like } from "../../../../shared/Like";


const profiles = [
  {
    id: 1,
    name: "Анна",
    age: 25,
    bio: "Люблю путешествовать и заниматься спортом.",
    imageUrl: "man1.jpg",
  },
  {
    id: 2,
    name: "Иван",
    age: 30,
    bio: "Пишу музыку и люблю кататься на велосипеде.",
    imageUrl: "photo1.png",
  },
  {
    id: 3,
    name: "Ольга",
    age: 28,
    bio: "Увлекаюсь фотографией и кулинарией.",
    imageUrl: "woman1.jpg",
  },
  // Добавьте больше профилей по потребности
];

export const ProfileSlider = () => {
  const [currentIndex, setCurrentIndex] = useState(0);

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

  return (
    <div className={styles.slider}>
      <div className={styles.profileCard} {...handlers}>
        <img
          src={currentProfile.imageUrl}
          alt={currentProfile.name}
          className={styles.profileImage}
        />
        <h2>
          {currentProfile.name}, {currentProfile.age}
        </h2>
        {/* <p>{currentProfile.bio}</p> */}
      </div>
      <div className={styles.controls}>
        <div className={styles.likeCircle} onClick={handlePrev}>
          <Like />
        </div>

        <div className={styles.dislikeCircle} onClick={handleNext}>
          <Dislike />
        </div>
      </div>
    </div>
  );
};

export default ProfileSlider;
