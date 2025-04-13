import React from 'react';

import styles from '../EditProfile/editProfile.module.css';

interface AboutMeSectionProps {
  about: string;
  onInputChange: (field: string, value: string) => void;
}

export const AboutMeSection: React.FC<AboutMeSectionProps> = ({ about, onInputChange }) => {
  return (
    <div className={styles.editSection}>
      <h3>About Me</h3>
      <div className={styles.textareaContainer}>
        <textarea
          className={styles.textareaInput}
          placeholder="Tell something about yourself"
          value={about}
          onChange={e => onInputChange('about', e.target.value)}
          maxLength={500}
        />
        <div className={styles.characterCount}>{about.length}/500</div>
      </div>
    </div>
  );
};
