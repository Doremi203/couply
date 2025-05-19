export const genderOptions = [
  { label: 'Женщины', value: 'Girls' },
  { label: 'Мужчины', value: 'Boys' },
  { label: 'Оба', value: 'Both' },
];

export const educationOptions = {
  secondary: 'Среднее',
  higher: 'Высшее',
  phd: 'Phd',
};

export const childrenOptions = {
  no: 'Нет детей',
  notYet: 'Нет, но хочу',
  yes: 'Есть дети',
};

export const alcoholOptions = {
  negatively: 'Негативно',
  neutrally: 'Нейтрально',
  positively: 'Положительно',
};

export const smokingOptions = {
  negatively: 'Негативно',
  neutrally: 'Нейтрально',
  positively: 'Положительно',
};

export const goalOptions = {
  dating: 'Знакомства',
  relationship: 'Отношения',
  friendship: 'Дружба',
  justChatting: 'Общение',
};

export const sportOptions = {
  gym: 'Тренажерный зал',
  running: 'Бег',
  swimming: 'Плавание',
  cycling: 'Велоспорт',
  tennis: 'Теннис',
  basketball: 'Баскетбол',
  hiking: 'Походы',
  dancing: 'Танцы',
  martialArts: 'Боевые искусства',
  football: 'Футбол',
  skiing: 'Лыжи/Сноуборд',
  climbing: 'Скалолазание',
};

export const selfdevelopmentOptions = {
  languages: 'Изучение языков',
  lectures: 'Лекции',
  onlineCourses: 'Онлайн-курсы',
  selfEducation: 'Самообразование',
  meditation: 'Медитация',
  psychology: 'Психология',
  philosophy: 'Философия',
  history: 'История',
  technology: 'Технологии',
  reading: 'Чтение',
};

export const hobbyOptions = {
  photography: 'Фотография',
  painting: 'Рисование',
  boardGames: 'Настольные игры',
  reading: 'Чтение',
  cooking: 'Готовка',
  gardening: 'Садоводство',
  travel: 'Путешествия',
  writing: 'Писательство',
  chess: 'Шахматы',
  crafts: 'Рукоделие/DIY',
  animals: 'Уход за животными',
  astrology: 'Астрология',
};

export const zodiacOptions = {
  aries: 'Овен',
  taurus: 'Телец',
  gemini: 'Близнецы',
  cancer: 'Рак',
  leo: 'Лев',
  virgo: 'Дева',
  libra: 'Весы',
  scorpio: 'Скорпион',
  sagittarius: 'Стрелец',
  capricorn: 'Козерог',
  aquarius: 'Водолей',
  pisces: 'Рыбы',
};

export const goalToApi = {
  Знакомства: 'GOAL_DATING',
  Отношения: 'GOAL_RELATIONSHIP',
  Дружба: 'GOAL_FRIENDSHIP',
  Общение: 'GOAL_JUST_CHATTING',
};

export const moviesTVOptions = {
  action: 'Боевики',
  comedy: 'Комедии',
  drama: 'Драмы',
  scifi: 'Научная фантастика',
  anime: 'Аниме',
  documentaries: 'Документалки',
  horror: 'Ужасы',
  fantasy: 'Фэнтези',
  thriller: 'Триллеры',
  romance: 'Мелодрамы',
  historical: 'Исторические',
};

export const musicOptions = {
  pop: 'Поп',
  rock: 'Рок',
  hiphop: 'Хип-хоп',
  rap: 'Рэп',
  electronic: 'Электронная',
  jazz: 'Джаз',
  classical: 'Классическая',
  indie: 'Инди',
  rnb: 'R&B',
  metal: 'Метал',
  folk: 'Фолк',
  country: 'Кантри',
  alternative: 'Альтернатива',
};

export const foodDrinkOptions = {
  coffee: 'Кофе',
  wine: 'Вино',
  cocktails: 'Коктейли',
  vegan: 'Веганство',
  baking: 'Выпечка',
  fine_dining: 'Рестораны',
  street_food: 'Уличная еда',
  tea: 'Чайные церемонии',
  barbecue: 'Шашлыки/Гриль',
  craft_beer: 'Крафтовое пиво',
};

export const personalityTraitsOptions = {
  introvert: 'Интроверт',
  extrovert: 'Экстраверт',
  adventurous: 'Авантюрист',
  homebody: 'Домосед',
  optimist: 'Оптимист',
  ambitious: 'Амбициозный',
  creative: 'Творческий',
  empathetic: 'Эмпат',
  analytical: 'Аналитик',
  sarcastic: 'Саркастичность',
};

export const petsOptions = {
  dogs: 'Собаки',
  cats: 'Кошки',
  other: 'Другие питомцы',
  none: 'Без питомцев',
};

export const genderToApi = {
  Женщины: 'GENDER_FEMALE',
  Мужчины: 'GENDER_MALE',
  Оба: 'GENDER_UNSPECIFIED',
};

export const educationToApi = {
  Среднее: 'EDUCATION_SECONDARY',
  Высшее: 'EDUCATION_HIGHER',
  Phd: 'EDUCATION_PHD',
};

export const childrenToApi = {
  'Нет детей': 'CHILDREN_NO',
  'Нет, но хочу': 'CHILDREN_NOT_YET',
  'Есть дети': 'CHILDREN_YES',
};

export const alcoholToApi = {
  Негативно: 'ALCOHOL_NEGATIVELY',
  Нейтрально: 'ALCOHOL_NEUTRALLY',
  Положительно: 'ALCOHOL_POSITIVELY',
};

export const smokingToApi = {
  Негативно: 'SMOKING_NEGATIVELY',
  Нейтрально: 'SMOKING_NEUTRALLY',
  Положительно: 'SMOKING_POSITIVELY',
};

export const sportToApi = {
  'Тренажерный зал': 'SPORT_GYM',
  Бег: 'SPORT_RUNNING',
  Плавание: 'SPORT_SWIMMING',
  Велоспорт: 'SPORT_CYCLING',
  Теннис: 'SPORT_TENNIS',
  Баскетбол: 'SPORT_BASKETBALL',
  Походы: 'SPORT_HIKING',
  Танцы: 'SPORT_DANCING',
  'Боевые искусства': 'SPORT_MARTIAL_ARTS',
  Футбол: 'SPORT_FOOTBALL',
  'Лыжи/Сноуборд': 'SPORT_SKIING',
  Скалолазание: 'SPORT_CLIMBING',
};

export const selfdevelopmentToApi = {
  'Изучение языков': 'SELFDEVELOPMENT_LANGUAGES',
  Лекции: 'SELFDEVELOPMENT_LECTURES',
  'Онлайн-курсы': 'SELFDEVELOPMENT_ONLINE_COURSES',
  Самообразование: 'SELFDEVELOPMENT_SELF_EDUCATION',
  Медитация: 'SELFDEVELOPMENT_MEDITATION',
  Психология: 'SELFDEVELOPMENT_PSYCHOLOGY',
  Философия: 'SELFDEVELOPMENT_PHILOSOPHY',
  История: 'SELFDEVELOPMENT_HISTORY',
  Технологии: 'SELFDEVELOPMENT_TECHNOLOGY',
  Чтение: 'SELFDEVELOPMENT_READING',
};

export const hobbyToApi = {
  Фотография: 'HOBBY_PHOTOGRAPHY',
  Рисование: 'HOBBY_PAINTING',
  'Настольные игры': 'HOBBY_BOARD_GAMES',
  Чтение: 'HOBBY_READING',
  Готовка: 'HOBBY_COOKING',
  Садоводство: 'HOBBY_GARDENING',
  Путешествия: 'HOBBY_TRAVEL',
  Писательство: 'HOBBY_WRITING',
  Шахматы: 'HOBBY_CHESS',
  'Рукоделие/DIY': 'HOBBY_CRAFTS',
  'Уход за животными': 'HOBBY_ANIMALS',
  Астрология: 'HOBBY_ASTROLOGY',
};

export const zodiacToApi = {
  Овен: 'ZODIAC_ARIES',
  Телец: 'ZODIAC_TAURUS',
  Близнецы: 'ZODIAC_GEMINI',
  Рак: 'ZODIAC_CANCER',
  Лев: 'ZODIAC_LEO',
  Дева: 'ZODIAC_VIRGO',
  Весы: 'ZODIAC_LIBRA',
  Скорпион: 'ZODIAC_SCORPIO',
  Стрелец: 'ZODIAC_SAGITTARIUS',
  Козерог: 'ZODIAC_CAPRICORN',
  Водолей: 'ZODIAC_AQUARIUS',
  Рыбы: 'ZODIAC_PISCES',
};

export const moviesTVToApi = {
  Боевики: 'MOVIESTV_ACTION',
  Комедии: 'MOVIESTV_COMEDY',
  Драмы: 'MOVIESTV_DRAMA',
  'Научная фантастика': 'MOVIESTV_SCIFI',
  Аниме: 'MOVIESTV_ANIME',
  Документалки: 'MOVIESTV_DOCUMENTARIES',
  Ужасы: 'MOVIESTV_HORROR',
  Фэнтези: 'MOVIESTV_FANTASY',
  Триллеры: 'MOVIESTV_THRILLER',
  Мелодрамы: 'MOVIESTV_ROMANCE',
  Исторические: 'MOVIESTV_HISTORICAL',
};

export const musicToApi = {
  Поп: 'MUSIC_POP',
  Рок: 'MUSIC_ROCK',
  'Хип-хоп': 'MUSIC_HIPHOP',
  Рэп: 'MUSIC_RAP',
  Электронная: 'MUSIC_ELECTRONIC',
  Джаз: 'MUSIC_JAZZ',
  Классическая: 'MUSIC_CLASSICAL',
  Инди: 'MUSIC_INDIE',
  'R&B': 'MUSIC_RNB',
  Метал: 'MUSIC_METAL',
  Фолк: 'MUSIC_FOLK',
  Кантри: 'MUSIC_COUNTRY',
  Альтернатива: 'MUSIC_ALTERNATIVE',
};

export const foodDrinkToApi = {
  Кофе: 'FOODDRINK_COFFEE',
  Вино: 'FOODDRINK_WINE',
  Коктейли: 'FOODDRINK_COCKTAILS',
  Веганство: 'FOODDRINK_VEGAN',
  Выпечка: 'FOODDRINK_BAKING',
  Рестораны: 'FOODDRINK_FINE_DINING',
  'Уличная еда': 'FOODDRINK_STREET_FOOD',
  'Чайные церемонии': 'FOODDRINK_TEA',
  'Шашлыки/Гриль': 'FOODDRINK_BARBECUE',
  'Крафтовое пиво': 'FOODDRINK_CRAFT_BEER',
};

export const personalityTraitsToApi = {
  Интроверт: 'TRAIT_INTROVERT',
  Экстраверт: 'TRAIT_EXTROVERT',
  Авантюрист: 'TRAIT_ADVENTUROUS',
  Домосед: 'TRAIT_HOMEBODY',
  Оптимист: 'TRAIT_OPTIMIST',
  Амбициозный: 'TRAIT_AMBITIOUS',
  Творческий: 'TRAIT_CREATIVE',
  Эмпат: 'TRAIT_EMPATHIC',
  Аналитик: 'TRAIT_ANALYTICAL',
  Саркастичность: 'TRAIT_SARCASM',
};

export const petsToApi = {
  Собаки: 'PETS_DOGS',
  Кошки: 'PETS_CATS',
  'Другие питомцы': 'PETS_OTHER',
  'Без питомцев': 'PETS_NONE',
};

export const goalFromApi = {
  GOAL_DATING: 'Знакомства',
  GOAL_RELATIONSHIP: 'Отношения',
  GOAL_FRIENDSHIP: 'Дружба',
  GOAL_JUST_CHATTING: 'Общение',
};

export const genderFromApi = {
  GENDER_FEMALE: 'Женщины',
  GENDER_MALE: 'Мужчины',
  GENDER_UNSPECIFIED: 'Оба',
};

export const educationFromApi = {
  EDUCATION_SECONDARY: 'Среднее',
  EDUCATION_HIGHER: 'Высшее',
  EDUCATION_PHD: 'Phd',
};

export const childrenFromApi = {
  CHILDREN_NO: 'Нет детей',
  CHILDREN_NOT_YET: 'Нет, но хочу',
  CHILDREN_YES: 'Есть дети',
};

export const alcoholFromApi = {
  ALCOHOL_NEGATIVELY: 'Негативно',
  ALCOHOL_NEUTRALLY: 'Нейтрально',
  ALCOHOL_POSITIVELY: 'Положительно',
};

export const smokingFromApi = {
  SMOKING_NEGATIVELY: 'Негативно',
  SMOKING_NEUTRALLY: 'Нейтрально',
  SMOKING_POSITIVELY: 'Положительно',
};

export const sportFromApi = {
  SPORT_GYM: 'Тренажерный зал',
  SPORT_RUNNING: 'Бег',
  SPORT_SWIMMING: 'Плавание',
  SPORT_CYCLING: 'Велоспорт',
  SPORT_TENNIS: 'Теннис',
  SPORT_BASKETBALL: 'Баскетбол',
  SPORT_HIKING: 'Походы',
  SPORT_DANCING: 'Танцы',
  SPORT_MARTIAL_ARTS: 'Боевые искусства',
  SPORT_FOOTBALL: 'Футбол',
  SPORT_SKIING: 'Лыжи/Сноуборд',
  SPORT_CLIMBING: 'Скалолазание',
};

export const selfdevelopmentFromApi = {
  SELFDEVELOPMENT_LANGUAGES: 'Изучение языков',
  SELFDEVELOPMENT_LECTURES: 'Лекции',
  SELFDEVELOPMENT_ONLINE_COURSES: 'Онлайн-курсы',
  SELFDEVELOPMENT_SELF_EDUCATION: 'Самообразование',
  SELFDEVELOPMENT_MEDITATION: 'Медитация',
  SELFDEVELOPMENT_PSYCHOLOGY: 'Психология',
  SELFDEVELOPMENT_PHILOSOPHY: 'Философия',
  SELFDEVELOPMENT_HISTORY: 'История',
  SELFDEVELOPMENT_TECHNOLOGY: 'Технологии',
  SELFDEVELOPMENT_READING: 'Чтение',
};

export const hobbyFromApi = {
  HOBBY_PHOTOGRAPHY: 'Фотография',
  HOBBY_PAINTING: 'Рисование',
  HOBBY_BOARD_GAMES: 'Настольные игры',
  HOBBY_READING: 'Чтение',
  HOBBY_COOKING: 'Готовка',
  HOBBY_GARDENING: 'Садоводство',
  HOBBY_TRAVEL: 'Путешествия',
  HOBBY_WRITING: 'Писательство',
  HOBBY_CHESS: 'Шахматы',
  HOBBY_CRAFTS: 'Рукоделие/DIY',
  HOBBY_ANIMALS: 'Уход за животными',
  HOBBY_ASTROLOGY: 'Астрология',
};

export const zodiacFromApi = {
  ZODIAC_ARIES: 'Овен',
  ZODIAC_TAURUS: 'Телец',
  ZODIAC_GEMINI: 'Близнецы',
  ZODIAC_CANCER: 'Рак',
  ZODIAC_LEO: 'Лев',
  ZODIAC_VIRGO: 'Дева',
  ZODIAC_LIBRA: 'Весы',
  ZODIAC_SCORPIO: 'Скорпион',
  ZODIAC_SAGITTARIUS: 'Стрелец',
  ZODIAC_CAPRICORN: 'Козерог',
  ZODIAC_AQUARIUS: 'Водолей',
  ZODIAC_PISCES: 'Рыбы',
};

export const moviesTVFromApi = {
  MOVIESTV_ACTION: 'Боевики',
  MOVIESTV_COMEDY: 'Комедии',
  MOVIESTV_DRAMA: 'Драмы',
  MOVIESTV_SCIFI: 'Научная фантастика',
  MOVIESTV_ANIME: 'Аниме',
  MOVIESTV_DOCUMENTARIES: 'Документалки',
  MOVIESTV_HORROR: 'Ужасы',
  MOVIESTV_FANTASY: 'Фэнтези',
  MOVIESTV_THRILLER: 'Триллеры',
  MOVIESTV_ROMANCE: 'Мелодрамы',
  MOVIESTV_HISTORICAL: 'Исторические',
};

export const musicFromApi = {
  MUSIC_POP: 'Поп',
  MUSIC_ROCK: 'Рок',
  MUSIC_HIPHOP: 'Хип-хоп',
  MUSIC_RAP: 'Рэп',
  MUSIC_ELECTRONIC: 'Электронная',
  MUSIC_JAZZ: 'Джаз',
  MUSIC_CLASSICAL: 'Классическая',
  MUSIC_INDIE: 'Инди',
  MUSIC_RNB: 'R&B',
  MUSIC_METAL: 'Метал',
  MUSIC_FOLK: 'Фолк',
  MUSIC_COUNTRY: 'Кантри',
  MUSIC_ALTERNATIVE: 'Альтернатива',
};

export const foodDrinkFromApi = {
  FOODDRINK_COFFEE: 'Кофе',
  FOODDRINK_WINE: 'Вино',
  FOODDRINK_COCKTAILS: 'Коктейли',
  FOODDRINK_VEGAN: 'Веганство',
  FOODDRINK_BAKING: 'Выпечка',
  FOODDRINK_FINE_DINING: 'Рестораны',
  FOODDRINK_STREET_FOOD: 'Уличная еда',
  FOODDRINK_TEA: 'Чайные церемонии',
  FOODDRINK_BARBECUE: 'Шашлыки/Гриль',
  FOODDRINK_CRAFT_BEER: 'Крафтовое пиво',
};

export const personalityTraitsFromApi = {
  TRAIT_INTROVERT: 'Интроверт',
  TRAIT_EXTROVERT: 'Экстраверт',
  TRAIT_ADVENTUROUS: 'Авантюрист',
  TRAIT_HOMEBODY: 'Домосед',
  TRAIT_OPTIMIST: 'Оптимист',
  TRAIT_AMBITIOUS: 'Амбициозный',
  TRAIT_CREATIVE: 'Творческий',
  TRAIT_EMPATHIC: 'Эмпат',
  TRAIT_ANALYTICAL: 'Аналитик',
  TRAIT_SARCASM: 'Саркастичность',
};

export const petsFromApi = {
  PETS_DOGS: 'Собаки',
  PETS_CATS: 'Кошки',
  PETS_OTHER: 'Другие питомцы',
  PETS_NONE: 'Без питомцев',
};
