import { LikeProfile, MatchProfile } from './types';

// Mock data for likes
export const likesData: LikeProfile[] = [
  {
    id: 1,
    name: 'Анна',
    age: 25,
    imageUrl: 'woman1.jpg',
    liked: false,
    hasLikedYou: true, // This profile has already liked the user
    location: 'Москва, Россия',
    interests: ['Музыка', 'Путешествия', 'Фотография', 'Мода', 'Искусство'],
  },
  {
    id: 2,
    name: 'Иван',
    age: 30,
    imageUrl: 'man1.jpg',
    liked: false,
    hasLikedYou: true, // This profile has already liked the user
    location: 'Санкт-Петербург, Россия',
    interests: ['Спорт', 'Кино', 'Технологии', 'Путешествия'],
  },
  {
    id: 3,
    name: 'Ольга',
    age: 28,
    imageUrl: 'photo1.png',
    liked: false,
    hasLikedYou: false,
    location: 'Казань, Россия',
    interests: ['Книги', 'Йога', 'Кулинария', 'Природа'],
  },
  {
    id: 4,
    name: 'Алексей',
    age: 32,
    imageUrl: 'man1.jpg',
    liked: false,
    hasLikedYou: false,
    location: 'Екатеринбург, Россия',
    interests: ['Музыка', 'Горы', 'Фотография', 'Путешествия'],
  },
];

// Mock data for matches
export const matchesData: MatchProfile[] = [
  {
    id: 101, // Using different ID range for matches
    name: 'Мария',
    age: 27,
    imageUrl: 'woman1.jpg',
    telegram: '@maria_27',
    instagram: '@maria_insta',
  },
  {
    id: 102, // Using different ID range for matches
    name: 'Дмитрий',
    age: 31,
    imageUrl: 'man1.jpg',
    telegram: '@dmitry_31',
    instagram: '@dmitry_insta',
  },
];

// Helper function to generate a unique ID for new matches
export const generateMatchId = (existingMatches: MatchProfile[]): number => {
  const maxId = existingMatches.reduce(
    (max: number, match: MatchProfile) => Math.max(max, match.id),
    100,
  );
  return maxId + 1;
};
