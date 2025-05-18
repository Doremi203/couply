import React from 'react';

import { ToggleButtons } from '../../../../shared/components/ToggleButtons';

import styles from './profileVisibilitySection.module.css';

interface ProfileVisibilitySectionProps {
  isHidden: boolean;
  onInputChange: (field: string, value: string) => void;
  title?: string;
}

export const ProfileVisibilitySection: React.FC<ProfileVisibilitySectionProps> = ({
  isHidden,
  onInputChange,
  title = 'Видимость профиля',
}) => {
  return (
    <div className={styles.section}>
      <h3>{title}</h3>
      <div className={styles.toggleOption}>
        <span>Скрыть профиль</span>
        <ToggleButtons
          options={[
            { label: 'Нет', value: 'visible' },
            { label: 'Да', value: 'hidden' },
          ]}
          value={isHidden ? 'hidden' : 'visible'}
          onSelect={value => onInputChange('isHidden', value === 'hidden' ? 'true' : 'false')}
        />
      </div>
    </div>
  );
};
