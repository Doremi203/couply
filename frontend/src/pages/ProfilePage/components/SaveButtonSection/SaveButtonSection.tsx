import React from 'react';

import { CustomButton } from '../../../../shared/components/CustomButton';
import styles from '../EditProfile/editProfile.module.css';

interface SaveButtonSectionProps {
  onSave: () => void;
}

export const SaveButtonSection: React.FC<SaveButtonSectionProps> = ({ onSave }) => {
  return (
    <div className={styles.saveButtonContainer}>
      <CustomButton text="Save Changes" onClick={onSave} className={styles.saveButton} />
    </div>
  );
};
