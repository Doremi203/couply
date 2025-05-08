import KeyboardBackspaceIcon from '@mui/icons-material/KeyboardBackspace';
import React from 'react';

import styles from '../../filtersDrawer.module.css';

type FilterHeaderProps = {
  onBack: () => void;
  onClear: () => void;
};

const FilterHeader: React.FC<FilterHeaderProps> = ({ onBack, onClear }) => {
  return (
    <div className={styles.header}>
      <button className={styles.backButton} onClick={onBack}>
        <KeyboardBackspaceIcon />
      </button>
      <h2 className={styles.title}>Фильтры</h2>
      <button className={styles.clearButton} onClick={onClear}>
        Очистить
      </button>
    </div>
  );
};

export default FilterHeader;
