import React from "react";
import styles from "./customInput.module.css";

interface CustomInputProps {
  placeholder: string;
  type: string;
  className?: string;
  value?: string;
  onChange?: (e: React.ChangeEvent<HTMLInputElement>) => void;
}

export const CustomInput = ({
  placeholder,
  type,
  className,
  value,
  onChange,
}: CustomInputProps) => {
  return (
    <div className={styles.customInputWrapper}>
      <input
        type={type}
        placeholder={placeholder}
        className={`${styles.customInput} ${className ? className : ""}`}
        value={value}
        onChange={onChange}
      />
    </div>
  );
};

export default CustomInput;
