import React from 'react';

import InterestFilter from '../../../../features/filters/components/FiltersDrawer/components/InterestFilter/InterestFilter';

import styles from './interestsSection.module.css';

interface InterestsSectionProps {
  selectedOptions: string[];
  onSelect: (value: string[]) => void;
}

export const InterestSection: React.FC<InterestsSectionProps> = ({ selectedOptions, onSelect }) => {
  return (
    <div className={styles.section}>
      <div className={styles.inner}>
        <InterestFilter title="Интересы" selectedOptions={selectedOptions} onSelect={onSelect} />
      </div>
    </div>
  );
};
