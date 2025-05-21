
import { GenderPriority } from '../../../entities/search';
import { Goal, Zodiac, Education, Alcohol, Smoking, Sport, Selfdevelopment, Hobby, Music, MoviesTV, FoodDrink, PersonalityTraits, Pets, Children } from '../../../entities/user';

export const getDefaultFilter = () => {
  return {
    genderPriority: GenderPriority.any,
    minAge: 18,
    maxAge: 100,
    minHeight: 100,
    maxHeight: 250,
    minDistanceKm: 0,
    maxDistanceKm: 100,
    goal: Goal.unspecified,
    zodiac: Zodiac.unspecified,
    education: Education.unspecified,
    children: Children.unspecified,
    alcohol: Alcohol.unspecified,
    smoking: Smoking.unspecified,
    interest: {
      sport: [Sport.unspecified],
      selfDevelopment: [Selfdevelopment.unspecified],
      hobby: [Hobby.unspecified],
      music: [Music.unspecified],
      moviesTv: [MoviesTV.unspecified],
      foodDrink: [FoodDrink.unspecified],
      personalityTraits: [PersonalityTraits.unspecified],
      pets: [Pets.unspecified],
    },
    onlyVerified: false,
    onlyPremium: false,
  };
};