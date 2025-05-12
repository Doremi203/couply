import React from 'react';

import styles from './aboutMeSection.module.css';

interface AboutMeSectionProps {
  about: string;
  onInputChange: (field: string, value: string) => void;
  title?: string;
  maxLength?: number;
  placeholder?: string;
}

export const AboutMeSection: React.FC<AboutMeSectionProps> = ({
  about,
  onInputChange,
  title = 'Обо мне',
  maxLength = 500,
  placeholder = 'Напишите что-нибудь о себе',
}) => {
  return (
    <div className={styles.editSection}>
      <h3>{title}</h3>
      <div className={styles.textareaContainer}>
        <textarea
          className={styles.textareaInput}
          placeholder={placeholder}
          value={about}
          onChange={e => onInputChange('about', e.target.value)}
          maxLength={maxLength}
        />
        <div className={styles.characterCount}>
          {about.length}/{maxLength}
        </div>
      </div>
    </div>
  );
};

export default AboutMeSection;
