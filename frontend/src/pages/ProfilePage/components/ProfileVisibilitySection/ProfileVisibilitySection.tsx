import React from 'react';

import { ToggleButtons } from '../../../../shared/components/ToggleButtons';
import styles from '../EditProfile/editProfile.module.css';

interface ProfileVisibilitySectionProps {
  isHidden: boolean;
  onInputChange: (field: string, value: string) => void;
}

export const ProfileVisibilitySection: React.FC<ProfileVisibilitySectionProps> = ({
  isHidden,
  onInputChange,
}) => {
  return (
    <div className={styles.editSection}>
      <h3>Profile Visibility</h3>
      <div className={styles.toggleOption}>
        <span>Hide my profile</span>
        <ToggleButtons
          options={[
            { label: 'No', value: 'visible' },
            { label: 'Yes', value: 'hidden' },
          ]}
          value={isHidden ? 'hidden' : 'visible'}
          onSelect={value => onInputChange('isHidden', value === 'hidden' ? 'true' : 'false')}
        />
      </div>
    </div>
  );
};
