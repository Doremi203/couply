import KeyboardBackspaceIcon from '@mui/icons-material/KeyboardBackspace';
import React from 'react';

// import styles from '../../activityHistory.module.css';
import styles from '../ActivityHistory/activityHistory.module.css';

interface ActivityHeaderProps {
  onBack: () => void;
}

export const ActivityHeader: React.FC<ActivityHeaderProps> = ({ onBack }) => {
  return (
    <div className={styles.profileHeader}>
      <div className={styles.backButton} onClick={onBack}>
        <KeyboardBackspaceIcon />
      </div>
      <div className={styles.header}>activity history</div>
    </div>
  );
};
