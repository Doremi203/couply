
import { Sport, Selfdevelopment, Hobby, Music, MoviesTV, FoodDrink, PersonalityTraits, Pets } from '../../../entities/user';
import { sportOptions, sportToApi, selfdevelopmentOptions, selfdevelopmentToApi, hobbyOptions, hobbyToApi, musicOptions, musicToApi, moviesTVOptions, moviesTVToApi, foodDrinkOptions, foodDrinkToApi, personalityTraitsOptions, personalityTraitsToApi, petsOptions, petsToApi } from '../components/constants';

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
  
  export const mapInterestsToBackendFormat = (selectedInterests: string[]) => {
    const result: Record<string, string[]> = {};
  
    const reverseCategoryMap: Record<string, keyof typeof interestCategoriesMap> = {};
  
    (Object.entries(interestCategoriesMap) as [keyof typeof interestCategoriesMap, any][])
      .forEach(([category, config]) => {
        Object.values(config.options).forEach(option => {
            //@ts-ignore
          reverseCategoryMap[option] = category;
        });
      });
  
    const grouped = selectedInterests.reduce((acc, interest) => {
      const category = reverseCategoryMap[interest];
      if (category) {
        acc[category] = acc[category] || [];
        acc[category].push(interest);
      }
      return acc;
    }, {} as Record<string, string[]>);

    (Object.entries(interestCategoriesMap) as [keyof typeof interestCategoriesMap, any][])
      .forEach(([category, config]) => {
        const selected = grouped[category] || [];
        const mapped = selected
          .map(interest => config.apiMap[interest])
          .filter(Boolean);
  
        result[category] = mapped.length > 0 ? mapped : [config.default];
      });
  
    return result;
  };
  
