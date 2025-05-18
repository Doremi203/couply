import React from 'react';

import Like from '../Like/Like';

import styles from './likeButton.module.css';

interface LikeButtonProps {
  onClick: (e: React.MouseEvent) => void;
  className?: string;
  likeClassName?: string;
}

export const LikeButton = ({ onClick, className, likeClassName }: LikeButtonProps) => {
  const handleClick = (e: React.MouseEvent) => {
    e.stopPropagation();
    e.preventDefault();
    onClick(e);
  };

  return (
    <div
      className={`${styles.likeCircle} ${className || ''}`}
      onClick={handleClick}
      style={{ pointerEvents: 'auto', cursor: 'pointer' }}
    >
      <Like className={`${styles.like} ${likeClassName || ''}`} />
    </div>
  );
};

export default LikeButton;
