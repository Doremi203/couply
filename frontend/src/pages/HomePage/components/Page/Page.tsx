import styles from "./homePage.module.css";
import { FiltersIcon } from "../../../../shared/components/FiltersIcon";
import { NavBar } from "../../../../shared/components/NavBar";
import { ProfileSlider } from "../../../../features/ProfileSlider";
import FiltersDrawer from "../FiltersDrawer/FiltersDrawer";
import { useState } from "react";

export const HomePage = () => {
  const [isFiltersOpen, setIsFiltersOpen] = useState(false);

  const handleFiltersOpen = () => {
    setIsFiltersOpen(true);
  };
  const handleFiltersClose = () => {
    setIsFiltersOpen(false);
  };

  return (
    <div>
      <div className={styles.header}>
        <div className={styles.appName}> couply</div>
        <div className={styles.filtersIcon}>
          <div onClick={handleFiltersOpen}>
            {" "}
            <FiltersIcon />{" "}
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
