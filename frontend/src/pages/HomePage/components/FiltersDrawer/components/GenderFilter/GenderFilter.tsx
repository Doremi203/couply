import React from "react";
import styles from "../../filtersDrawer.module.css";
import ToggleButtons from "../../../../../../shared/components/ToggleButtons/ToggleButtons";

type GenderOption = {
  label: string;
  value: string;
};

type GenderFilterProps = {
  value: string;
  options: GenderOption[];
  onChange: (value: string) => void;
};

const GenderFilter: React.FC<GenderFilterProps> = ({ value, options, onChange }) => {
  return (
    <div className={styles.section}>
      <h3 className={styles.sectionTitle}>Interested in</h3>
      <ToggleButtons
        options={options}
        onSelect={onChange}
        value={value}
      />
    </div>
  );
};

export default GenderFilter;