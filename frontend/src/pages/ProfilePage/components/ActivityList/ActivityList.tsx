import React from 'react';

// import styles from '../../activityHistory.module.css';
import styles from '../ActivityHistory/activityHistory.module.css';
import { ActivityItem } from '../ActivityHistory/types';
import { ActivityItem as ActivityItemComponent } from '../ActivityItem';

interface ActivityListProps {
  activities: ActivityItem[];
  formatDate: (dateString: string) => string;
}

export const ActivityList: React.FC<ActivityListProps> = ({ activities, formatDate }) => {
  return (
    <div className={styles.activityList}>
      {activities.map((activity, index) => (
        <ActivityItemComponent key={index} activity={activity} formatDate={formatDate} />
      ))}
    </div>
  );
};
