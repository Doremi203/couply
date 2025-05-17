import { useState, useRef, useEffect, KeyboardEvent } from 'react';

import styles from './codeInput.module.css';

interface CodeInputProps {
  length?: number;
  onCodeChange: (code: string) => void;
}

export const CodeInput = ({ length = 6, onCodeChange }: CodeInputProps) => {
  const [code, setCode] = useState<string[]>(Array(length).fill(''));
  const inputsRef = useRef<HTMLInputElement[]>([]);

  useEffect(() => {
    inputsRef.current[0]?.focus();
  }, []);

  const handleChange = (index: number, value: string) => {
    if (!/^\d*$/.test(value)) return;

    const newCode = [...code];
    newCode[index] = value.slice(-1);
    setCode(newCode);

    if (value && index < length - 1) {
      inputsRef.current[index + 1]?.focus();
    }

    onCodeChange(newCode.join(''));
  };

  const handleKeyDown = (index: number, e: KeyboardEvent<HTMLInputElement>) => {
    if (e.key === 'Backspace' && !code[index] && index > 0) {
      inputsRef.current[index - 1]?.focus();
    }
  };

  const handlePaste = (e: React.ClipboardEvent) => {
    e.preventDefault();
    const pastedData = e.clipboardData.getData('text').slice(0, length);
    const newCode = [...code];

    pastedData.split('').forEach((char, i) => {
      if (i < length) {
        newCode[i] = char;
      }
    });

    setCode(newCode);
    onCodeChange(newCode.join(''));
    inputsRef.current[Math.min(pastedData.length, length - 1)]?.focus();
  };

  return (
    <div className={styles.container}>
      {code.map((digit, index) => (
        <input
          key={index}
          type="text"
          inputMode="numeric"
          pattern="[0-9]*"
          maxLength={1}
          value={digit}
          onChange={e => handleChange(index, e.target.value)}
          onKeyDown={e => handleKeyDown(index, e)}
          onPaste={handlePaste}
          ref={el => el && (inputsRef.current[index] = el)}
          className={styles.input}
        />
      ))}
    </div>
  );
};

export default CodeInput;
