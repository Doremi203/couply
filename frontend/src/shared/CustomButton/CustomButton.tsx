import React from "react";
import styles from "./customButton.module.css";

interface CustomButtonProps {
  text: string;
  onClick: () => void;
  className?: string;
}

export const CustomButton = ({
  text,
  onClick,
  className,
}: CustomButtonProps) => {
  return (
    <div className={styles.customButtonWrapper}>
      <button
        className={`${styles.customButton} ${className ? className : ""}`}
        onClick={onClick}
      >
        {text}
      </button>
    </div>
  );
};

export default CustomButton;
