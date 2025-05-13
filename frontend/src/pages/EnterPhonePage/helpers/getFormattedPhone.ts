export const getFormattedPhone = (phone: string) => {
  const cleanedPhone = phone.replace(/\D/g, '');
  const formattedPhone = cleanedPhone.startsWith('7') ? `+${cleanedPhone}` : `+7${cleanedPhone}`;
  return formattedPhone;
};
