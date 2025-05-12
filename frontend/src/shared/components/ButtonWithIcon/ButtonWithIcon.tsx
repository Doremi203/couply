import { ReactNode } from 'react';

import styles from './buttonWithIcon.module.css';

interface ButtonWithIconProps {
  icon: ReactNode;
  text: string;
  onClick: () => void;
  className?: string;
  disabled?: boolean;
  iconClassName?: string;
}

export const ButtonWithIcon = ({
  icon,
  text,
  onClick,
  className,
  disabled = false,
  iconClassName,
}: ButtonWithIconProps) => {
  return (
    <div className={styles.buttonWrapper}>
      <button
        className={`${styles.buttonWithIcon} ${className ? className : ''} ${
          disabled ? styles.disabled : ''
        }`}
        onClick={onClick}
        disabled={disabled}
      >
        <div className={`${styles.iconContainer} ${iconClassName ? iconClassName : ''}`}>
          {icon}
        </div>
        <span className={styles.buttonText}>{text}</span>
      </button>
    </div>
  );
};

export default ButtonWithIcon;
