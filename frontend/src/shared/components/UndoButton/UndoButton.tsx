import UndoIcon from '@mui/icons-material/Undo';

import styles from './undoButton.module.css';

interface BackCircleButtonProps {
  onClick: () => void;
  className?: string;
}

export const UndoButton = ({ onClick, className }: BackCircleButtonProps) => {
  return (
    <div className={`${styles.circle} ${className || ''}`} onClick={onClick}>
      <UndoIcon className={styles.icon} />
    </div>
  );
};

export default UndoButton;
