export enum Goal {
  unspecified = 'GOAL_UNSPECIFIED',
  dating = 'GOAL_DATING',
  relationship = 'GOAL_RELATIONSHIP',
  friendship = 'GOAL_FRIENDSHIP',
  justChatting = 'GOAL_JUST_CHATTING',
}

export enum Sport {
  unspecified = 'SPORT_UNSPECIFIED',
  running = 'SPORT_RUNNING',
  swimming = 'SPORT_SWIMMING',
  yoga = 'SPORT_YOGA',
  bicycle = 'SPORT_BICYCLE',
  gym = 'SPORT_GYM',
  skiing = 'SPORT_SKIING',
  snowboarding = 'SPORT_SNOWBOARDING',
  dancing = 'SPORT_DANCING',
  martialArts = 'SPORT_MARTIAL_ARTS',
  surfing = 'SPORT_SURFING',
  hiking = 'SPORT_HIKING',
  tennis = 'SPORT_TENNIS',
  climbing = 'SPORT_CLIMBING',
}

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

export enum Art {
  unspecified = 'ART_UNSPECIFIED',
  painting = 'ART_PAINTING',
  photograph = 'ART_PHOTOGRAPH',
  music = 'ART_MUSIC',
  singing = 'ART_SINGING',
  writing = 'ART_WRITING',
  sculpture = 'ART_SCULPTURE',
  theater = 'ART_THEATER',
  cinema = 'ART_CINEMA',
  needlework = 'ART_NEEDLEWORK',
}

export enum Social {
  unspecified = 'SOCIAL_UNSPECIFIED',
  volunteering = 'SOCIAL_VOLUNTEERING',
  charity = 'SOCIAL_CHARITY',
  ecoActivism = 'SOCIAL_ECO_ACTIVISM',
  elderlyCare = 'SOCIAL_ELDERLY_CARE',
  childcare = 'SOCIAL_CHILDCARE',
  animalWelfare = 'SOCIAL_ANIMAL_WELFARE',
}

export enum Hobby {
  unspecified = 'HOBBY_UNSPECIFIED',
  literature = 'HOBBY_LITERATURE',
  videoGames = 'HOBBY_VIDEO_GAMES',
  boardGames = 'HOBBY_BOARD_GAMES',
  travels = 'HOBBY_TRAVELS',
  plantCultivation = 'HOBBY_PLANT_CULTIVATION',
  fishing = 'HOBBY_FISHING',
  dogWalks = 'HOBBY_DOG_WALKS',
  catsLover = 'HOBBY_CATS_LOVER',
  carsAndMotorcycles = 'HOBBY_CARS_AND_MOTORCYCLES',
  concerts = 'HOBBY_CONCERTS',
}

export enum Gastronomy {
  unspecified = 'GASTRONOMY_UNSPECIFIED',
  cooking = 'GASTRONOMY_COOKING',
  wineDegustation = 'GASTRONOMY_WINE_DEGUSTATION',
  bars = 'GASTRONOMY_BARS',
  coffee = 'GASTRONOMY_COFFEE',
  tea = 'GASTRONOMY_TEA',
  vegan = 'GASTRONOMY_VEGAN',
  foodCritic = 'GASTRONOMY_FOOD_CRITIC',
  sugarLover = 'GASTRONOMY_SUGAR_LOVER',
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

export interface UserRequest {
  id?: string;
  name: string;
  age: number;
  gender: string;
  location: string;
  bio: string;
  goal: Goal;
  interest: null; //
  zodiac: Zodiac;
  height: number;
  education: Education;
  children: Children;
  alcohol: Alcohol;
  smoking: Smoking;
  hidden: boolean;
  verified: boolean;
  photos: null; //
}

export interface UserResponse {
  user: {
    id: string;
    name: string;
    age: number;
    gender: string;
    location: string;
    bio: string;
    goal: Goal;
    interest: null; //
    zodiac: Zodiac;
    height: number;
    education: Education;
    children: Children;
    alcohol: Alcohol;
    smoking: Smoking;
    hidden: boolean;
    verified: boolean;
    photos: null; //
  };
}

export interface UpdateUserParams {
  id: string;
  data: Partial<UserRequest>;
}

export interface GetUserRequest {
  id: string;
}
