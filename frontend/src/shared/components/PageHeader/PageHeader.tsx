import KeyboardBackspaceIcon from '@mui/icons-material/KeyboardBackspace';
import React from 'react';

import styles from './pageHeader.module.css';

interface PageHeaderProps {
  onBack: () => void;
  title?: string;
}

export const PageHeader: React.FC<PageHeaderProps> = ({ onBack, title = 'edit profile' }) => {
  return (
    <div className={styles.profileHeader}>
      <div className={styles.backButton} onClick={onBack}>
        <KeyboardBackspaceIcon />
      </div>
      <div className={styles.header}>{title}</div>
    </div>
  );
};

export default PageHeader;
