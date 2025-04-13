import React from 'react';

import styles from '../../activityHistory.module.css';
import { ActivityItem as ActivityItemType } from '../../types';
import { ActivityItem } from '../ActivityItem';

interface ActivityListProps {
  activities: ActivityItemType[];
  formatDate: (dateString: string) => string;
}

export const ActivityList: React.FC<ActivityListProps> = ({ activities, formatDate }) => {
  return (
    <div className={styles.list}>
      {activities.map((activity, index) => (
        <ActivityItem key={index} activity={activity} formatDate={formatDate} />
      ))}
    </div>
  );
};
