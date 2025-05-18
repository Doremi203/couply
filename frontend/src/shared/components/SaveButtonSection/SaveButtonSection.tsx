import React from 'react';

import { CustomButton } from '../CustomButton';

import styles from './saveButtonSection.module.css';

interface SaveButtonSectionProps {
  onSave: () => void;
  text?: string;
}

export const SaveButtonSection: React.FC<SaveButtonSectionProps> = ({
  onSave,
  text = 'Cохранить',
}) => {
  return (
    <div className={styles.container}>
      <CustomButton text={text} onClick={onSave} className={styles.button} />
    </div>
  );
};
