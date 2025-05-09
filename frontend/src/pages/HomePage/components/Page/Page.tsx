import { useState } from 'react';

import FiltersDrawer from '../../../../features/filters/components/FiltersDrawer';
import { ProfileSlider } from '../../../../features/ProfileSlider';
import { FiltersIcon } from '../../../../shared/components/FiltersIcon';
import { NavBar } from '../../../../shared/components/NavBar';
import { HiddenAcc } from '../HiddenAcc/HiddenAcc';

import styles from './homePage.module.css';

export const HomePage = () => {
  const [isFiltersOpen, setIsFiltersOpen] = useState(false);

  const handleFiltersOpen = () => {
    setIsFiltersOpen(true);
  };
  const handleFiltersClose = () => {
    setIsFiltersOpen(false);
  };

  //TODO
  const hidden = false;

  if (hidden) {
    return (
      <body className={styles.pageContainer}>
        <div className={styles.headerHidden}>
          <div className={styles.spacer} />
          <div className={styles.appName}> couply</div>
        </div>

        <HiddenAcc />

        <div style={{ position: 'relative', zIndex: 1010 }}>
          <NavBar />
        </div>
      </body>
    );
  }

  return (
    <body className={styles.pageContainer}>
      <div className={styles.header}>
        <div className={styles.spacer} />
        <div className={styles.appName}> couply</div>
        <div className={styles.filtersIcon}>
          <div onClick={handleFiltersOpen}>
            <FiltersIcon />
          </div>
          <FiltersDrawer open={isFiltersOpen} onClose={handleFiltersClose} />
        </div>
      </div>

      <ProfileSlider />

      <div style={{ position: 'relative', zIndex: 1010 }}>
        <NavBar />
      </div>
    </body>
  );
};

export default HomePage;
