import {
  Sport,
  Selfdevelopment,
  Hobby,
  Music,
  MoviesTV,
  FoodDrink,
  PersonalityTraits,
  Pets,
} from '../../../entities/user';
import {
  sportOptions,
  sportToApi,
  selfdevelopmentOptions,
  selfdevelopmentToApi,
  hobbyOptions,
  hobbyToApi,
  musicOptions,
  musicToApi,
  moviesTVOptions,
  moviesTVToApi,
  foodDrinkOptions,
  foodDrinkToApi,
  personalityTraitsOptions,
  personalityTraitsToApi,
  petsOptions,
  petsToApi,
} from '../components/constants';

const interestCategoriesMap = {
  sport: {
    options: sportOptions,
    apiMap: sportToApi,
    default: Sport.unspecified,
  },
  selfDevelopment: {
    options: selfdevelopmentOptions,
    apiMap: selfdevelopmentToApi,
    default: Selfdevelopment.unspecified,
  },
  hobby: {
    options: hobbyOptions,
    apiMap: hobbyToApi,
    default: Hobby.unspecified,
  },
  music: {
    options: musicOptions,
    apiMap: musicToApi,
    default: Music.unspecified,
  },
  moviesTv: {
    options: moviesTVOptions,
    apiMap: moviesTVToApi,
    default: MoviesTV.unspecified,
  },
  foodDrink: {
    options: foodDrinkOptions,
    apiMap: foodDrinkToApi,
    default: FoodDrink.unspecified,
  },
  personalityTraits: {
    options: personalityTraitsOptions,
    apiMap: personalityTraitsToApi,
    default: PersonalityTraits.unspecified,
  },
  pets: {
    options: petsOptions,
    apiMap: petsToApi,
    default: Pets.unspecified,
  },
};

export const mapInterestsFromBackendFormat = (apiInterests: Record<string, string[]>) => {
  const result: string[] = [];

  (Object.entries(interestCategoriesMap) as [keyof typeof interestCategoriesMap, any][]).forEach(
    ([category, config]) => {
      const apiValues = apiInterests[category] || [config.default];
      
      // Create reverse mapping from API values to frontend values
      const reverseApiMap = Object.entries(config.apiMap).reduce(
        (acc, [frontendValue, apiValue]) => {
          acc[apiValue] = frontendValue;
          return acc;
        },
        {} as Record<string, string>,
      );

      // Map API values back to frontend values
      apiValues.forEach((apiValue: string) => {
        const frontendValue = reverseApiMap[apiValue];
        if (frontendValue && frontendValue !== config.default) {
          result.push(frontendValue);
        }
      });
    },
  );

  return result;
}; 