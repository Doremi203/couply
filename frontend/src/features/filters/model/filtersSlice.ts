import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { RootState } from '../../../app/store';

// Интерфейс для фильтров
export interface FiltersState {
  ageRange: [number, number];
  distance: number;
  interests: string[];
  showOnlyWithPhoto: boolean;
  // Добавьте другие фильтры по необходимости
}

// Начальное состояние
const initialState: FiltersState = {
  ageRange: [18, 50],
  distance: 50,
  interests: [],
  showOnlyWithPhoto: true,
};

// Создание слайса
export const filtersSlice = createSlice({
  name: 'filters',
  initialState,
  reducers: {
    // Установка возрастного диапазона
    setAgeRange: (state, action: PayloadAction<[number, number]>) => {
      state.ageRange = action.payload;
    },
    
    // Установка дистанции
    setDistance: (state, action: PayloadAction<number>) => {
      state.distance = action.payload;
    },
    
    // Добавление интереса
    addInterest: (state, action: PayloadAction<string>) => {
      if (!state.interests.includes(action.payload)) {
        state.interests.push(action.payload);
      }
    },
    
    // Удаление интереса
    removeInterest: (state, action: PayloadAction<string>) => {
      state.interests = state.interests.filter(interest => interest !== action.payload);
    },
    
    // Переключение опции "только с фото"
    toggleShowOnlyWithPhoto: (state) => {
      state.showOnlyWithPhoto = !state.showOnlyWithPhoto;
    },
    
    // Сброс всех фильтров
    resetFilters: () => initialState,
  },
});

// Экспорт actions
export const {
  setAgeRange,
  setDistance,
  addInterest,
  removeInterest,
  toggleShowOnlyWithPhoto,
  resetFilters,
} = filtersSlice.actions;

// Селекторы
export const selectFilters = (state: RootState) => state.filters;
export const selectAgeRange = (state: RootState) => state.filters.ageRange;
export const selectDistance = (state: RootState) => state.filters.distance;
export const selectInterests = (state: RootState) => state.filters.interests;
export const selectShowOnlyWithPhoto = (state: RootState) => state.filters.showOnlyWithPhoto;

// Экспорт reducer
export default filtersSlice.reducer;