import KeyboardBackspaceIcon from '@mui/icons-material/KeyboardBackspace';
import React from 'react';

import styles from '../../activityHistory.module.css';

interface ActivityHeaderProps {
  onBack: () => void;
  title?: string;
}

export const ActivityHeader: React.FC<ActivityHeaderProps> = ({
  onBack,
  title = 'activity history',
}) => {
  return (
    <div className={styles.header}>
      <div className={styles.backButton} onClick={onBack}>
        <KeyboardBackspaceIcon />
      </div>
      <div className={styles.headerTitle}>{title}</div>
    </div>
  );
};
