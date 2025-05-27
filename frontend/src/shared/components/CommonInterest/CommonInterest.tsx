import React from 'react';

import styles from './commonInterest.module.css';

interface CommonInterestProps {
  text: string;
  isCommon?: boolean;
  className?: string;
}

export const CommonInterest: React.FC<CommonInterestProps> = ({
  text,
  isCommon = false,
  className = '',
}) => {
  return (
    <div className={`${styles.tag} ${isCommon ? styles.commonInterest : ''} ${className}`}>
      {text}
      {isCommon && <span className={styles.commonBadge}>Common</span>}
    </div>
  );
};

export default CommonInterest;
