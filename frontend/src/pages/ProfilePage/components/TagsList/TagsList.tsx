import React from 'react';

import styles from './tagsList.module.css';

interface TagsListProps {
  items: string[];
  commonItems?: string[];
}

export const TagsList: React.FC<TagsListProps> = ({ items, commonItems = [] }) => {
  return (
    <div className={styles.tagsList}>
      {items.map((item, index) => {
        const isCommon = commonItems.includes(item);
        return (
          <div key={index} className={`${styles.tag} ${isCommon ? styles.commonInterest : ''}`}>
            {item}
            {isCommon && <span className={styles.commonBadge}>Common</span>}
          </div>
        );
      })}
    </div>
  );
};

export default TagsList;
