export const mapFiltersToApi = <T extends string>(
  selectedFilters: string[],
  filterToApi: Record<string, T>,
  defaultUnspecified?: T,
): T[] => {
  // Маппим и фильтруем невалидные значения
  const mapped = selectedFilters.map(filter => filterToApi[filter]).filter(Boolean);

  // Добавляем значение по умолчанию если нужно
  if (mapped.length === 0 && defaultUnspecified) {
    return [defaultUnspecified];
  }

  return mapped;
};
