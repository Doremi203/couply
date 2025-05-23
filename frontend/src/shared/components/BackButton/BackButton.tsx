import styles from './backButton.module.css';

interface BackButtonProps {
  onClose: (e?: React.MouseEvent) => void;
}

export const BackButton = ({ onClose }: BackButtonProps) => {
  return (
    <button
      className={styles.backButton}
      onClick={e => {
        e.stopPropagation();
        onClose(e);
      }}
    >
      <svg
        width="24"
        height="24"
        viewBox="0 0 24 24"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          d="M19 12H5"
          stroke="white"
          stroke-width="1.5"
          stroke-linecap="round"
          stroke-linejoin="round"
        />
        <path
          d="M12 19L5 12L12 5"
          stroke="white"
          stroke-width="1.5"
          stroke-linecap="round"
          stroke-linejoin="round"
        />
      </svg>
    </button>
  );
};
