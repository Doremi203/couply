import styles from './customButton.module.css';

interface CustomButtonProps {
  text: string;
  onClick: () => void;
  className?: string;
  disabled?: boolean;
}

export const CustomButton = ({ text, onClick, className, disabled = false }: CustomButtonProps) => {
  return (
    <div className={styles.customButtonWrapper}>
      <button
        className={`${styles.customButton} ${className ? className : ''} ${
          disabled ? styles.disabled : ''
        }`}
        onClick={onClick}
        disabled={disabled}
      >
        {text}
      </button>
    </div>
  );
};

export default CustomButton;
