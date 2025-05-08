import React from 'react';

import { ActivityIcon } from '../../../../shared/components/ActivityIcon';
import styles from '../../activityHistory.module.css';
import { ActivityItem as ActivityItemType } from '../../types';

interface ActivityItemProps {
  activity: ActivityItemType;
  formatDate: (dateString: string) => string;
}

export const ActivityItem: React.FC<ActivityItemProps> = ({ activity, formatDate }) => {
  const getActivityText = (type: string) => {
    switch (type) {
      case 'view':
        return 'viewed your profile';
      case 'like':
        return 'liked your profile';
      case 'message':
        return 'sent you a message';
      default:
        return '';
    }
  };

  return (
    <div className={styles.item}>
      <ActivityIcon type={activity.type} />
      <div className={styles.details}>
        <span className={styles.user}>{activity.user}</span>
        <span className={styles.type}>{getActivityText(activity.type)}</span>
        <span className={styles.date}>{formatDate(activity.date)}</span>
      </div>
    </div>
  );
};
