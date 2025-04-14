import React from 'react';

import styles from './emptyState.module.css';

interface EmptyStateProps {
  title: string;
  subtitle: string;
}

export const EmptyState: React.FC<EmptyStateProps> = ({ title, subtitle }) => {
  return (
    <div className={styles.emptyState}>
      <div className={styles.emptyStateIcon}>
        <svg
          width="64"
          height="64"
          viewBox="0 0 24 24"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            d="M12 21.35L10.55 20.03C5.4 15.36 2 12.28 2 8.5C2 5.42 4.42 3 7.5 3C9.24 3 10.91 3.81 12 5.09C13.09 3.81 14.76 3 16.5 3C19.58 3 22 5.42 22 8.5C22 12.28 18.6 15.36 13.45 20.04L12 21.35Z"
            fill="#E0E0E0"
          />
        </svg>
      </div>
      <p className={styles.emptyStateText}>{title}</p>
      <p className={styles.emptyStateSubtext}>{subtitle}</p>
    </div>
  );
};

export default EmptyState;
