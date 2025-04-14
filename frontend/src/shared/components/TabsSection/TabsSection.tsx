import React from 'react';

import styles from './tabsSection.module.css';

interface TabsSectionProps<T extends string> {
  tabs: T[];
  activeTab: T;
  onTabChange: (tab: T) => void;
  tabLabels?: Record<T, string>;
}

export const TabsSection = <T extends string>({
  tabs,
  activeTab,
  onTabChange,
  tabLabels,
}: TabsSectionProps<T>) => {
  return (
    <div className={styles.tabs}>
      {tabs.map(tab => (
        <div
          key={tab}
          className={`${styles.tab} ${activeTab === tab ? styles.activeTab : ''}`}
          onClick={() => onTabChange(tab)}
        >
          {tabLabels?.[tab] || tab}
        </div>
      ))}
    </div>
  );
};

export default TabsSection;
