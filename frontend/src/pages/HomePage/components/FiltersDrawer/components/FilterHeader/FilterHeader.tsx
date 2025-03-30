import React from "react";
import styles from "../../filtersDrawer.module.css";
import KeyboardBackspaceIcon from '@mui/icons-material/KeyboardBackspace';

type FilterHeaderProps = {
  onBack: () => void;
  onClear: () => void;
};

const FilterHeader: React.FC<FilterHeaderProps> = ({ onBack, onClear }) => {
  return (
    <div className={styles.header}>
      <button className={styles.backButton} onClick={onBack}>
        <KeyboardBackspaceIcon />
      </button>
      <h2 className={styles.title}>Filters</h2>
      <button className={styles.clearButton} onClick={onClear}>
        Clear
      </button>
    </div>
  );
};

export default FilterHeader;