import React from 'react';

import styles from '../../filtersDrawer.module.css';

type ChipFilterProps = {
  title: string;
  options: string[];
  selectedOptions: string[];
  onToggle: (option: string) => void;
};

const GenderFilter: React.FC<ChipFilterProps> = ({ title, options, selectedOptions, onToggle }) => {
  return (
    <div className={styles.section}>
      <h3 className={styles.sectionTitle}>{title}</h3>
      <div className={styles.chipContainer}>
        {options.map((option, index) => (
          <div
            key={index}
            className={`${styles.chipBig} ${selectedOptions.includes(option) ? styles.chipBigSelected : ''}`}
            onClick={() => onToggle(option)}
          >
            {option}
          </div>
        ))}
      </div>
    </div>
  );
};

export default GenderFilter;
