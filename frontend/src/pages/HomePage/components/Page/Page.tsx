import { useEffect, useState } from 'react';
import { useSelector } from 'react-redux';

import { useGetFilterQuery } from '../../../../entities/search';
import { getUserId } from '../../../../entities/user/model/userSlice';
import FiltersDrawer from '../../../../features/filters/components/FiltersDrawer';
import { ProfileSlider } from '../../../../features/ProfileSlider';
import { FiltersIcon } from '../../../../shared/components/FiltersIcon';
import { NavBar } from '../../../../shared/components/NavBar';
import { useGeolocation } from '../../../../shared/lib/hooks/useGeolocation';
import { HiddenAcc } from '../HiddenAcc/HiddenAcc';

import styles from './homePage.module.css';

export const HomePage = () => {
  useEffect(() => {
    document.documentElement.classList.add(styles.noScroll);
    return () => {
      document.documentElement.classList.remove(styles.noScroll);
    };
  }, []);

  const userId = useSelector(getUserId);
  const { updateUserLocation } = useGeolocation();
  const [isFiltersOpen, setIsFiltersOpen] = useState(false);
  const { data: filterData } = useGetFilterQuery({});

  useEffect(() => {
    if (userId) {
      updateUserLocation()
        .then(success => {
          if (success) {
            console.log('Location updated successfully on HomePage');
          }
        })
        .catch(err => {
          console.error('Failed to update location on HomePage:', err);
        });
    }
  }, [userId, updateUserLocation]);

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
      <div className={styles.pageContainer}>
        <div className={styles.headerHidden}>
          <div className={styles.spacer} />
          <div className={styles.appName}> couply</div>
        </div>

        <HiddenAcc />

        <div style={{ position: 'relative', zIndex: 1010 }}>
          <NavBar />
        </div>
      </div>
    );
  }

  return (
    <div className={styles.pageContainer}>
      <div className={styles.header}>
        <div className={styles.spacer} />
        <div className={styles.appName}> couply</div>
        <div className={styles.filtersIcon}>
          <div onClick={handleFiltersOpen}>
            <FiltersIcon />
          </div>
          <FiltersDrawer
            open={isFiltersOpen}
            onClose={handleFiltersClose}
            initialFilterData={filterData}
          />
        </div>
      </div>

      <ProfileSlider />

      <div style={{ position: 'relative', zIndex: 1010 }}>
        <NavBar />
      </div>
    </div>
  );
};

export default HomePage;
