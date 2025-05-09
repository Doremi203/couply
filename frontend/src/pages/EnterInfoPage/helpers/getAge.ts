export function getAge(birthdate: string) {
  // Проверка корректности даты
  if (!/^\d{4}-\d{2}-\d{2}$/.test(birthdate)) {
    return 'Неверный формат даты';
  }

  const [inputYear, inputMonth, inputDay] = birthdate.split('-').map(Number);
  const birthDate = new Date(inputYear, inputMonth - 1, inputDay); // Месяцы в JS: 0-11
  const today = new Date();

  // Проверка на будущую дату
  if (birthDate > today) {
    return 'Дата рождения в будущем';
  }

  // Проверка на валидность даты (например, 31 февраля)
  if (
    birthDate.getFullYear() !== inputYear ||
    birthDate.getMonth() + 1 !== inputMonth ||
    birthDate.getDate() !== inputDay
  ) {
    return 'Некорректная дата';
  }

  // Расчет возраста
  let age = today.getFullYear() - birthDate.getFullYear();
  const monthDiff = today.getMonth() - birthDate.getMonth();
  const dayDiff = today.getDate() - birthDate.getDate();

  // Корректировка, если день рождения еще не наступил
  if (monthDiff < 0 || (monthDiff === 0 && dayDiff < 0)) {
    age--;
  }

  return age;
}

export default getAge;
