import React from "react";

import styles from "./buttonWithIcon.module.css";

export const ButtonWithIcon = () => {
  return (
    <button className={styles.buttonWithIcon}>
      <img src="logo.png" alt="Google Icon" className={styles.buttonIcon} />
      LOGIN WITH PHONE
    </button>
  );
};

export default ButtonWithIcon;
