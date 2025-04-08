import { ReactNode } from "react";
import styles from "./buttonWithIcon.module.css";

interface ButtonWithIconProps {
  icon: ReactNode;
  text: string;
  onClick: () => void;
  className?: string;
  disabled?: boolean;
}

export const ButtonWithIcon = ({
  icon,
  text,
  onClick,
  className,
  disabled = false,
}: ButtonWithIconProps) => {
  return (
    <div className={styles.buttonWrapper}>
      <button
        className={`${styles.buttonWithIcon} ${className ? className : ""} ${
          disabled ? styles.disabled : ""
        }`}
        onClick={onClick}
        disabled={disabled}
      >
        <div className={styles.iconContainer}>{icon}</div>
        <span className={styles.buttonText}>{text}</span>
      </button>
    </div>
  );
};

export default ButtonWithIcon;
