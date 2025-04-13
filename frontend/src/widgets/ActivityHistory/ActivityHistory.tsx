import React from 'react';

import styles from './activityHistory.module.css';
import { ActivityHeader } from './components/ActivityHeader';
import { ActivityList } from './components/ActivityList';
import { ActivityHistoryProps } from './types';

export const ActivityHistory: React.FC<ActivityHistoryProps> = ({
  activityHistory,
  onBack,
  formatDate,
}) => {
  return (
    <div className={styles.activityContent}>
      <ActivityHeader onBack={onBack} />
      <ActivityList activities={activityHistory} formatDate={formatDate} />
    </div>
  );
};

export default ActivityHistory;
