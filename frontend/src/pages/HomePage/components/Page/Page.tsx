import { useState } from 'react';

import { ProfileSlider } from '../../../../features/ProfileSlider';
import { FiltersIcon } from '../../../../shared/components/FiltersIcon';
import { NavBar } from '../../../../shared/components/NavBar';
import FiltersDrawer from '../FiltersDrawer/FiltersDrawer';

import styles from './homePage.module.css';

export const HomePage = () => {
  const [isFiltersOpen, setIsFiltersOpen] = useState(false);

  const handleFiltersOpen = () => {
    setIsFiltersOpen(true);
  };
  const handleFiltersClose = () => {
    setIsFiltersOpen(false);
  };

  return (
    <div className={styles.pageContainer}>
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

      <NavBar />
    </div>
  );
};

export default HomePage;
