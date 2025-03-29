import React, { useState, useEffect } from "react";
import styles from "./toggleButtons.module.css";

interface ToggleButtonOption {
  label: string;
  value: string;
}

interface ToggleButtonsProps {
  options: ToggleButtonOption[];
  onSelect: (value: string) => void;
  value?: string;
}

export const ToggleButtons: React.FC<ToggleButtonsProps> = ({
  options,
  onSelect,
  value,
}) => {
  const [selected, setSelected] = useState<string>(value || "");

  useEffect(() => {
    if (value) {
      setSelected(value);
    }
  }, [value]);

  const handleClick = (value: string) => {
    setSelected(value);
    onSelect(value);
  };

  return (
    <div className={styles.toggleButtons}>
      {options.map((option) => (
        <button
          key={option.value}
          className={selected === option.value ? styles.active : styles.button}
          onClick={() => handleClick(option.value)}
        >
          {option.label}
        </button>
      ))}
    </div>
  );
};

export default ToggleButtons;
