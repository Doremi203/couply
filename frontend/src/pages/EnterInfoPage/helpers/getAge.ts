export function getAge(birthdate: string) {
  if (!/^\d{4}-\d{2}-\d{2}$/.test(birthdate)) {
    return 'Неверный формат даты';
  }

  const [inputYear, inputMonth, inputDay] = birthdate.split('-').map(Number);
  const birthDate = new Date(inputYear, inputMonth - 1, inputDay);
  const today = new Date();

  if (birthDate > today) {
    return 'Дата рождения в будущем';
  }

  if (
    birthDate.getFullYear() !== inputYear ||
    birthDate.getMonth() + 1 !== inputMonth ||
    birthDate.getDate() !== inputDay
  ) {
    return 'Некорректная дата';
  }

  let age = today.getFullYear() - birthDate.getFullYear();
  const monthDiff = today.getMonth() - birthDate.getMonth();
  const dayDiff = today.getDate() - birthDate.getDate();

  if (monthDiff < 0 || (monthDiff === 0 && dayDiff < 0)) {
    age--;
  }

  return age;
}

export default getAge;
