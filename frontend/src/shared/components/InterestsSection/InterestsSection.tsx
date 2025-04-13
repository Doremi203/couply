import React from 'react';

import { CustomInput } from '../CustomInput';

import styles from './interestsSection.module.css';

interface InterestsSectionProps {
  title: string;
  placeholder: string;
  values: string[];
  fieldName: string;
  onArrayInputChange: (field: string, value: string) => void;
}

export const InterestsSection: React.FC<InterestsSectionProps> = ({
  title,
  placeholder,
  values,
  fieldName,
  onArrayInputChange,
}) => {
  return (
    <div className={styles.section}>
      <h3>{title}</h3>
      <CustomInput
        type="text"
        placeholder={placeholder}
        value={values.join(', ')}
        onChange={e => onArrayInputChange(fieldName, e.target.value)}
      />
    </div>
  );
};
