import Dislike from '../Dislike/Dislike';

import styles from './dislikeButton.module.css';

interface DislikeButtonProps {
  onClick: () => void
  className?: string
}

export const DislikeButton = ({onClick, className}: DislikeButtonProps) => {
  return (
    <div className={`${styles.dislikeCircle} ${className || ''}`} onClick={onClick}>
      <Dislike />
    </div>
  );
};

export default DislikeButton;