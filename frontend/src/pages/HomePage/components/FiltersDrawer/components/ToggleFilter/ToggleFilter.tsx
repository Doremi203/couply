import React from "react";
import styles from "../../filtersDrawer.module.css";

type ToggleFilterProps = {
  title: string;
  description: string;
  value: boolean;
  onChange: () => void;
};

const ToggleFilter: React.FC<ToggleFilterProps> = ({
  title,
  description,
  value,
  onChange
}) => {
  return (
    <div className={styles.section}>
      <div className={styles.toggleContainer}>
        <h3 className={styles.sectionTitle}>{title}</h3>
        <label className={styles.switch}>
          <input 
            type="checkbox" 
            checked={value}
            onChange={onChange}
          />
          <span className={styles.slider}></span>
        </label>
      </div>
      <p className={styles.verificationText}>
        {description}
      </p>
    </div>
  );
};

export default ToggleFilter;