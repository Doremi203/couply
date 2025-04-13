import React from 'react';

import styles from '../ActivityHistory/activityHistory.module.css';
import { ActivityItem as ActivityItemType } from '../ActivityHistory/types';
import { ActivityIcon } from '../ActivityIcon';

interface ActivityItemProps {
  activity: ActivityItemType;
  formatDate: (dateString: string) => string;
}

export const ActivityItemComponent: React.FC<ActivityItemProps> = ({ activity, formatDate }) => {
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
    <div className={styles.activityItem}>
      <ActivityIcon type={activity.type} />
      <div className={styles.activityDetails}>
        <span className={styles.activityUser}>{activity.user}</span>
        <span className={styles.activityType}>{getActivityText(activity.type)}</span>
        <span className={styles.activityDate}>{formatDate(activity.date)}</span>
      </div>
    </div>
  );
};
