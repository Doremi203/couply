export enum Goal {
  unspecified = 'GOAL_UNSPECIFIED',
  dating = 'GOAL_DATING',
  relationship = 'GOAL_RELATIONSHIP',
  friendship = 'GOAL_FRIENDSHIP',
  justChatting = 'GOAL_JUST_CHATTING',
}

// enum Sport {
//   SPORT_UNSPECIFIED = 0;
//   SPORT_GYM = 1;             // Тренажерный зал
//   SPORT_RUNNING = 2;          // Бег
//   SPORT_YOGA = 3;             // Йога
//   SPORT_SWIMMING = 4;         // Плавание
//   SPORT_CYCLING = 5;          // Велоспорт
//   SPORT_TENNIS = 6;           // Теннис
//   SPORT_BASKETBALL = 7;       // Баскетбол
//   SPORT_HIKING = 8;           // Походы
//   SPORT_DANCING = 9;          // Танцы
//   SPORT_MARTIAL_ARTS = 10;    // Боевые искусства
//   SPORT_FOOTBALL = 11;        // Футбол
//   SPORT_SKIING = 12;          // Лыжи/Сноуборд
//   SPORT_CLIMBING = 13;        // Скалолазание
// }

export enum Sport {
  unspecified = 'SPORT_UNSPECIFIED',
  gym = 'SPORT_GYM',
  running = 'SPORT_RUNNING',
  swimming = 'SPORT_SWIMMING',
  cycling = 'SPORT_CYCLING',
  tennis = 'SPORT_TENNIS',
  basketball = 'SPORT_BASKETBALL',
  hiking = 'SPORT_HIKING',
  dancing = 'SPORT_DANCING',
  martialArts = 'SPORT_MARTIAL_ARTS',
  football = 'SPORT_FOOTBALL',
  skiing = 'SPORT_SKIING',
  climbing = 'SPORT_CLIMBING',
}

// enum SelfDevelopment {
//   SELFDEVELOPMENT_UNSPECIFIED = 0;
//   SELFDEVELOPMENT_LANGUAGES = 1;    // Изучение языков
//   SELFDEVELOPMENT_LECTURES = 2;     // Лекции
//   SELFDEVELOPMENT_ONLINE_COURSES = 3; // Онлайн-курсы
//   SELFDEVELOPMENT_SELF_EDUCATION = 4; // Самообразование
//   SELFDEVELOPMENT_MEDITATION = 5;   // Медитация
//   SELFDEVELOPMENT_PSYCHOLOGY = 6;   // Психология
//   SELFDEVELOPMENT_PHILOSOPHY = 7;   // Философия
//   SELFDEVELOPMENT_HISTORY = 8;      // История
//   SELFDEVELOPMENT_TECHNOLOGY = 9;   // Технологии
//   SELFDEVELOPMENT_READING = 10;     // Чтение
// }

export enum Selfdevelopment {
  unspecified = 'SELFDEVELOPMENT_UNSPECIFIED',
  languages = 'SELFDEVELOPMENT_LANGUAGES',
  lectures = 'SELFDEVELOPMENT_LECTURES',
  onlineCourses = 'SELFDEVELOPMENT_ONLINE_COURSES',
  selfEducation = 'SELFDEVELOPMENT_SELF_EDUCATION',
  meditation = 'SELFDEVELOPMENT_MEDITATION',
  psychology = 'SELFDEVELOPMENT_PSYCHOLOGY',
  philosophy = 'SELFDEVELOPMENT_PHILOSOPHY',
  history = 'SELFDEVELOPMENT_HISTORY',
  technology = 'SELFDEVELOPMENT_TECHNOLOGY',
  reading = 'SELFDEVELOPMENT_READING',
}

// enum Hobby {
//   HOBBY_UNSPECIFIED = 0;
//   HOBBY_PHOTOGRAPHY = 1;      // Фотография
//   HOBBY_PAINTING = 2;         // Рисование
//   HOBBY_BOARDGAMES = 3;       // Настольные игры
//   HOBBY_READING = 4;          // Чтение
//   HOBBY_COOKING = 5;          // Готовка
//   HOBBY_GARDENING = 6;        // Садоводство
//   HOBBY_TRAVEL = 7;           // Путешествия
//   HOBBY_WRITING = 8;          // Писательство
//   HOBBY_CHESS = 9;            // Шахматы
//   HOBBY_CRAFTS = 10;          // Рукоделие/DIY
//   HOBBY_ANIMALS = 11;         // Уход за животными
//   HOBBY_ASTROLOGY = 12;       // Астрология
// }

export enum Hobby {
  unspecified = 'HOBBY_UNSPECIFIED',
  photography = 'HOBBY_PHOTOGRAPHY',
  painting = 'HOBBY_PAINTING',
  boardGames = 'HOBBY_BOARD_GAMES',
  reading = 'HOBBY_READING',
  cooking = 'HOBBY_COOKING',
  gardening = 'HOBBY_GARDENING',
  travel = 'HOBBY_TRAVEL',
  writing = 'HOBBY_WRITING',
  chess = 'HOBBY_CHESS',
  crafts = 'HOBBY_CRAFTS',
  animals = 'HOBBY_ANIMALS',
  astrology = 'HOBBY_ASTROLOGY',
}

export enum Zodiac {
  unspecified = 'ZODIAC_UNSPECIFIED',
  aries = 'ZODIAC_ARIES',
  taurus = 'ZODIAC_TAURUS',
  gemini = 'ZODIAC_GEMINI',
  cancer = 'ZODIAC_CANCER',
  leo = 'ZODIAC_LEO',
  virgo = 'ZODIAC_VIRGO',
  libra = 'ZODIAC_LIBRA',
  scorpio = 'ZODIAC_SCORPIO',
  sagittarius = 'ZODIAC_SAGITTARIUS',
  capricorn = 'ZODIAC_CAPRICORN',
  aquarius = 'ZODIAC_AQUARIUS',
  pisces = 'ZODIAC_PISCES',
}

export enum Education {
  unspecified = 'EDUCATION_UNSPECIFIED',
  secondary = 'EDUCATION_SECONDARY',
  higher = 'EDUCATION_HIGHER',
  phd = 'EDUCATION_PHD',
}

export enum Children {
  unspecified = 'CHILDREN_UNSPECIFIED',
  no = 'CHILDREN_NO',
  notYet = 'CHILDREN_NOT_YET',
  yes = 'CHILDREN_YES',
}

export enum Alcohol {
  unspecified = 'ALCOHOL_UNSPECIFIED',
  negatively = 'ALCOHOL_NEGATIVELY',
  neutrally = 'ALCOHOL_NEUTRALLY',
  positively = 'ALCOHOL_POSITIVELY',
}

export enum Smoking {
  unspecified = 'SMOKING_UNSPECIFIED',
  negatively = 'SMOKING_NEGATIVELY',
  neutrally = 'SMOKING_NEUTRALLY',
  positively = 'SMOKING_POSITIVELY',
}

export enum Gender {
  unspecified = 'GENDER_UNSPECIFIED',
  male = 'GENDER_MALE',
  female = 'GENDER_FEMALE',
}

// enum MoviesTV {
//   MOVIESTV_UNSPECIFIED = 0;
//   MOVIESTV_ACTION = 1;         // Боевики
//   MOVIESTV_COMEDY = 2;         // Комедии
//   MOVIESTV_DRAMA = 3;          // Драмы
//   MOVIESTV_SCIFI = 4;          // Научная фантастика
//   MOVIESTV_ANIME = 5;          // Аниме
//   MOVIESTV_DOCUMENTARIES = 6;  // Документалки
//   MOVIESTV_HORROR = 7;         // Ужасы
//   MOVIESTV_FANTASY = 8;        // Фэнтези
//   MOVIESTV_THRILLER = 9;       // Триллеры
//   MOVIESTV_ROMANCE = 10;       // Мелодрамы
//   MOVIESTV_HISTORICAL = 11;    // Исторические
// }

export enum MoviesTV {
  unspecified = 'MOVIESTV_UNSPECIFIED',
  action = 'MOVIESTV_ACTION',
  comedy = 'MOVIESTV_COMEDY',
  drama = 'MOVIESTV_DRAMA',
  scifi = 'MOVIESTV_SCIFI',
  anime = 'MOVIESTV_ANIME',
  documentaries = 'MOVIESTV_DOCUMENTARIES',
  horror = 'MOVIESTV_HORROR',
  fantasy = 'MOVIESTV_FANTASY',
  thriller = 'MOVIESTV_THRILLER',
  romance = 'MOVIESTV_ROMANCE',
  historical = 'MOVIESTV_HISTORICAL',
}

// enum Music {
//   MUSIC_UNSPECIFIED = 0;
//   MUSIC_POP = 1;              // Поп-музыка
//   MUSIC_ROCK = 2;             // Рок
//   MUSIC_HIPHOP = 3;           // Хип-хоп
//   MUSIC_RAP = 4;              // Рэп
//   MUSIC_ELECTRONIC = 5;       // Электронная
//   MUSIC_JAZZ = 6;             // Джаз
//   MUSIC_CLASSICAL = 7;        // Классическая
//   MUSIC_INDIE = 8;            // Инди
//   MUSIC_RNB = 9;              // R&B
//   MUSIC_METAL = 10;            // Метал
//   MUSIC_FOLK = 11;            // Фолк
//   MUSIC_COUNTRY = 12;         // Кантри
//   MUSIC_ALTERNATIVE = 13;     // Альтернатива
// }

export enum Music {
  unspecified = 'MUSIC_UNSPECIFIED',
  pop = 'MUSIC_POP',
  rock = 'MUSIC_ROCK',
  hiphop = 'MUSIC_HIPHOP',
  rap = 'MUSIC_RAP',
  electronic = 'MUSIC_ELECTRONIC',
  jazz = 'MUSIC_JAZZ',
  classical = 'MUSIC_CLASSICAL',
  indie = 'MUSIC_INDIE',
  rnb = 'MUSIC_RNB',
  metal = 'MUSIC_METAL',
  folk = 'MUSIC_FOLK',
  country = 'MUSIC_COUNTRY',
  alternative = 'MUSIC_ALTERNATIVE',
}

// enum FoodDrink {
//   FOODDRINK_UNSPECIFIED = 0;
//   FOODDRINK_COFFEE = 1;        // Кофе
//   FOODDRINK_WINE = 2;          // Вино
//   FOODDRINK_COCKTAILS = 3;     // Коктейли
//   FOODDRINK_VEGAN = 4;         // Веганство
//   FOODDRINK_BAKING = 5;        // Выпечка
//   FOODDRINK_FINE_DINING = 6;   // Рестораны
//   FOODDRINK_STREET_FOOD = 7;   // Уличная еда
//   FOODDRINK_TEA = 8;           // Чайные церемонии
//   FOODDRINK_BARBECUE = 9;      // Шашлыки/Гриль
//   FOODDRINK_CRAFT_BEER = 10;   // Крафтовое пиво
// }

export enum FoodDrink {
  unspecified = 'FOODDRINK_UNSPECIFIED',
  coffee = 'FOODDRINK_COFFEE',
  wine = 'FOODDRINK_WINE',
  cocktails = 'FOODDRINK_COCKTAILS',
  vegan = 'FOODDRINK_VEGAN',
  baking = 'FOODDRINK_BAKING',
  fine_dining = 'FOODDRINK_FINE_DINING',
  street_food = 'FOODDRINK_STREET_FOOD',
  tea = 'FOODDRINK_TEA',
  barbecue = 'FOODDRINK_BARBECUE',
  craft_beer = 'FOODDRINK_CRAFT_BEER',
}

// enum PersonalityTraits {
//   TRAIT_UNSPECIFIED = 0;
//   TRAIT_INTROVERT = 1;         // Интроверт
//   TRAIT_EXTROVERT = 2;         // Экстраверт
//   TRAIT_ADVENTUROUS = 3;       // Авантюрист
//   TRAIT_HOMEBODY = 4;          // Домосед
//   TRAIT_OPTIMIST = 5;          // Оптимист
//   TRAIT_AMBITIOUS = 6;         // Амбициозный
//   TRAIT_CREATIVE = 7;          // Творческий
//   TRAIT_EMPATHIC = 8;          // Эмпат
//   TRAIT_ANALYTICAL = 9;        // Аналитик
//   TRAIT_SARCASM = 10;          // Саркастичность
// }

export enum PersonalityTraits {
  unspecified = 'TRAIT_UNSPECIFIED',
  introvert = 'TRAIT_INTROVERT',
  extrovert = 'TRAIT_EXTROVERT',
  adventurous = 'TRAIT_ADVENTUROUS',
  homebody = 'TRAIT_HOMEBODY',
  optimist = 'TRAIT_OPTIMIST',
  ambitious = 'TRAIT_AMBITIOUS',
  creative = 'TRAIT_CREATIVE',
  empathetic = 'TRAIT_EMPATHIC',
  analytical = 'TRAIT_ANALYTICAL',
  sarcastic = 'TRAIT_SARCASM',
}

// enum Pets {
//   PETS_UNSPECIFIED = 0;
//   PETS_DOGS = 1;               // Собаки
//   PETS_CATS = 2;               // Кошки
//   PETS_OTHER = 3;              // Другие питомцы
//   PETS_NONE = 4;               // Без питомцев
// }

export enum Pets {
  unspecified = 'PETS_UNSPECIFIED',
  dogs = 'PETS_DOGS',
  cats = 'PETS_CATS',
  other = 'PETS_OTHER',
  none = 'PETS_NONE',
}
