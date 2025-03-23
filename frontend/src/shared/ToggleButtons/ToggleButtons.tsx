import React, { useState } from "react";
import styles from "./toggleButtons.module.css";

interface ToggleButtonOption {
  label: string;
  value: string;
}

interface ToggleButtonsProps {
  options: ToggleButtonOption[];
  onSelect: (value: string) => void;
}

export const ToggleButtons: React.FC<ToggleButtonsProps> = ({
  options,
  onSelect,
}) => {
  const [selected, setSelected] = useState<string>(options[0]?.value);

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
