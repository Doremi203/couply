export const createToggleHandler = <T>(
  setter: React.Dispatch<React.SetStateAction<T[]>>,
  currentValues: T[],
) => {
  return (value: T) => {
    setter(
      currentValues.includes(value)
        ? currentValues.filter(item => item !== value)
        : [...currentValues, value],
    );
  };
};