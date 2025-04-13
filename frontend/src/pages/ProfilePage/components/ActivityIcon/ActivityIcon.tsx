import VisibilityIcon from '@mui/icons-material/Visibility';
import React from 'react';

import styles from '../ActivityHistory/activityHistory.module.css';

interface ActivityIconProps {
  type: string;
}

export const ActivityIcon: React.FC<ActivityIconProps> = ({ type }) => {
  return (
    <div className={styles.activityIcon}>
      {type === 'view' && <VisibilityIcon />}
      {type === 'like' && <span>❤️</span>}
      {type === 'message' && <span>💬</span>}
    </div>
  );
};
