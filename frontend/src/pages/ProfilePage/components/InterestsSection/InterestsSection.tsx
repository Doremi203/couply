import React from 'react';

import { CustomInput } from '../../../../shared/components/CustomInput';
import styles from '../EditProfile/editProfile.module.css';

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
    <div className={styles.editSection}>
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
