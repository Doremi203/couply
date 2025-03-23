import React from "react";
import styles from "./customInput.module.css";

interface CustomInputProps {
  placeholder: string;
  type: string;
  className?: string;
}

export const CustomInput = ({
  placeholder,
  type,
  className,
}: CustomInputProps) => {
  return (
    <div className={styles.customInputWrapper}>
      <input
        type={type}
        placeholder={placeholder}
        className={`${styles.customInput} ${className ? className : ""}`}
      />
    </div>
  );
};

export default CustomInput;
