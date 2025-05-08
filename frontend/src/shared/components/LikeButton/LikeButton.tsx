import Like from '../Like/Like';

import styles from './likeButton.module.css';

interface LikeButtonProps {
  onClick: () => void;
  className?: string;
  likeClassName?: string;
}

export const LikeButton = ({ onClick, className, likeClassName }: LikeButtonProps) => {
  return (
    <div className={`${styles.likeCircle} ${className || ''}`} onClick={onClick}>
      <Like className={`${styles.like} ${likeClassName || ''}`} />
    </div>
  );
};

export default LikeButton;
