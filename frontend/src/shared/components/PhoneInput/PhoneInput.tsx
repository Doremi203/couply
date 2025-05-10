import { useRef, KeyboardEvent } from 'react';

// @ts-ignore
import styles from './PhoneInput.module.css';

interface PhoneInputProps {
  value: string;
  onChange: (value: string) => void;
}

export const PhoneInput = ({ value, onChange }: PhoneInputProps) => {
  const inputRef = useRef<HTMLInputElement>(null);

  const formatPhone = (input: string): string => {
    // Сохраняем только цифры и плюс
    const numbers = input.replace(/[^\d+]/g, '');

    // Всегда начинаем с +7
    let cleanValue = numbers.startsWith('+7') ? numbers : `+7${numbers.replace(/\D/g, '')}`;

    // Ограничиваем длину (11 цифр после +7)
    cleanValue = cleanValue.slice(0, 12);

    // Форматируем оставшиеся цифры
    const rest = cleanValue.slice(2).replace(/\D/g, '');

    let formatted = '+7';
    if (rest.length > 0) formatted += ` (${rest.slice(0, 3)}`;
    if (rest.length >= 4) formatted += `) ${rest.slice(3, 6)}`;
    if (rest.length >= 7) formatted += ` ${rest.slice(6, 8)}`;
    if (rest.length >= 9) formatted += ` ${rest.slice(8, 10)}`;

    return formatted;
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const formatted = formatPhone(e.target.value);
    onChange(formatted);
  };

  const handleKeyDown = (e: KeyboardEvent<HTMLInputElement>) => {
    // Запрещаем удаление базовой маски "+7"

    //@ts-ignore
    if (e.key === 'Backspace' && inputRef.current?.selectionStart <= 3) {
      e.preventDefault();
    }
  };

  const handleFocus = () => {
    if (value === '') {
      onChange('+7 ');
    }
    // setShowPlaceholder(false);
    setTimeout(() => {
      if (inputRef.current) {
        inputRef.current.setSelectionRange(4, 4); // Устанавливаем курсор после +7
      }
    }, 0);
  };

  return (
    <input
      type="tel"
      inputMode="tel"
      className={styles.input}
      value={value}
      onChange={handleChange}
      onKeyDown={handleKeyDown}
      ref={inputRef}
      placeholder="+7 (999) 999 99 99"
      onFocus={handleFocus}
    />
  );
};
