import React from 'react';

import ToggleButtons from '../../../../../../shared/components/ToggleButtons/ToggleButtons';
import styles from '../../filtersDrawer.module.css';

type GoalOption = {
  label: string;
  value: string;
};

type GoalFilterProps = {
  value: string;
  options: GoalOption[];
  onChange: (value: string) => void;
};

const GoalFilter: React.FC<GoalFilterProps> = ({ value, options, onChange }) => {
  return (
    <div className={styles.section}>
      <h3 className={styles.sectionTitle}>Ищу</h3>
      <ToggleButtons options={options} onSelect={onChange} value={value} />
    </div>
  );
};

export default GoalFilter;
