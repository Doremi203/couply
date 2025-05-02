import React from 'react';

import ToggleButtons from '../../../../../../shared/components/ToggleButtons/ToggleButtons';
import styles from '../../filtersDrawer.module.css';

type GenderOption = {
  label: string;
  value: string;
};

type GenderFilterProps = {
  value: string;
  options: GenderOption[];
  onChange: (value: string) => void;
};

const GenderFilter: React.FC<GenderFilterProps> = ({ value, options, onChange }) => {
  return (
    <div className={styles.section}>
      <h3 className={styles.sectionTitle}>Заинтересован в</h3>
      <ToggleButtons options={options} onSelect={onChange} value={value} />
    </div>
  );
};

export default GenderFilter;
