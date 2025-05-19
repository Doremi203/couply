import React from 'react';

import ChipFilter from '../../../../features/filters/components/FiltersDrawer/components/ChipFilter';

import styles from './editSection.module.css';

interface InterestsSectionProps {
  title: string;
  placeholder?: string;
  values?: string[];
  fieldName?: string;
  onArrayInputChange?: (field: string, value: string) => void;
  options?: string[];
  selectedOptions?: string[];
  onToggle?: (value: string) => void;
}

export const EditSection: React.FC<InterestsSectionProps> = props => {
  const {
    title,
    options,
    selectedOptions,
    onToggle,
    placeholder,
    values,
    fieldName,
    onArrayInputChange,
  } = props;

  // If options are provided, render a ChipFilter
  if (options && selectedOptions && onToggle) {
    return (
      <div className={styles.section}>
        <ChipFilter
          title={title}
          options={options}
          selectedOptions={selectedOptions}
          onToggle={onToggle}
        />
      </div>
    );
  }

  if (values && fieldName && onArrayInputChange) {
    return (
      <div className={styles.section}>
        <h3>{title}</h3>
        <input
          type="text"
          placeholder={placeholder}
          value={values.join(', ')}
          onChange={e => onArrayInputChange(fieldName, e.target.value)}
        />
      </div>
    );
  }

  return null;
};
