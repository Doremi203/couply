import React from "react";
import styles from "./homePage.module.css";
import ProfileSlider from "../../features/ProfileSlider/components/ProfileSlider";
import FiltersIcon from "../../shared/FiltersIcon/FiltersIcon";

export default function HomePage() {
  return (
    <div>
      <div className={styles.header}>
        <div className={styles.appName}> couply</div>
        <div className={styles.filtersIcon}>
          <FiltersIcon />
        </div>
      </div>

      <ProfileSlider />
    </div>
  );
}
