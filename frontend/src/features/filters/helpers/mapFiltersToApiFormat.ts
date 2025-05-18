export const mapFiltersToApi = (
  selectedFilters: string[],
  //@ts-ignore
  filterToApi,
  //@ts-ignore
  defaultUnspecified?,
) => {
  // Маппим и фильтруем невалидные значения
  const mapped = selectedFilters.map(filter => filterToApi[filter]).filter(Boolean);

  // Добавляем значение по умолчанию если нужно
  if (mapped.length === 0 && defaultUnspecified) {
    return [defaultUnspecified];
  }

  return mapped;
};
