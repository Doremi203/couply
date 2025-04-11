import KeyboardBackspaceIcon from '@mui/icons-material/KeyboardBackspace';
import VisibilityIcon from '@mui/icons-material/Visibility';
import React from 'react';

import styles from './activityHistory.module.css';

interface ActivityItem {
  type: string;
  user: string;
  date: string;
}

interface ActivityHistoryProps {
  activityHistory: ActivityItem[];
  onBack: () => void;
  formatDate: (dateString: string) => string;
}

export const ActivityHistory: React.FC<ActivityHistoryProps> = ({
  activityHistory,
  onBack,
  formatDate,
}) => {
  return (
    <div className={styles.activityContent}>
      <div className={styles.profileHeader}>
        <div className={styles.backButton} onClick={onBack}>
          <KeyboardBackspaceIcon />
        </div>
        <div className={styles.header}>activity history</div>
      </div>

      <div className={styles.activityList}>
        {activityHistory.map((activity, index) => (
          <div key={index} className={styles.activityItem}>
            <div className={styles.activityIcon}>
              {activity.type === 'view' && <VisibilityIcon />}
              {activity.type === 'like' && <span>❤️</span>}
              {activity.type === 'message' && <span>💬</span>}
            </div>
            <div className={styles.activityDetails}>
              <span className={styles.activityUser}>{activity.user}</span>
              <span className={styles.activityType}>
                {activity.type === 'view' && 'viewed your profile'}
                {activity.type === 'like' && 'liked your profile'}
                {activity.type === 'message' && 'sent you a message'}
              </span>
              <span className={styles.activityDate}>{formatDate(activity.date)}</span>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default ActivityHistory;