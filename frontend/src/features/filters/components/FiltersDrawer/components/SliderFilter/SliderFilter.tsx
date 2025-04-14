import React from 'react';

import styles from '../../filtersDrawer.module.css';
import { CustomSlider } from '../../styled/CustomSlider';

type SliderFilterProps = {
  title: string;
  value: number | number[];
  min: number;
  max: number;
  onChange: (event: Event, value: number | number[]) => void;
  unit?: string;
  valueLabelDisplay?: 'auto' | 'on' | 'off';
};

const SliderFilter: React.FC<SliderFilterProps> = ({
  title,
  value,
  min,
  max,
  onChange,
  unit = '',
  valueLabelDisplay = 'off',
}) => {
  // Format the display value based on whether it's a single value or range
  const displayValue = Array.isArray(value) 
    ? `${value[0]}-${value[1]}${unit}` 
    : `${value}${unit}`;

  return (
    <div className={styles.section}>
      <div className={styles.sliderHeader}>
        <h3 className={styles.sectionTitle}>{title}</h3>
        <span className={styles.sliderValue}>{displayValue}</span>
      </div>
      <CustomSlider
        value={value}
        onChange={onChange}
        valueLabelDisplay={valueLabelDisplay}
        aria-labelledby={`${title.toLowerCase()}-slider`}
        min={min}
        max={max}
      />
    </div>
  );
};

export default SliderFilter;