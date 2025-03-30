import React from "react";
import styles from "./infoGrid.module.css";

interface InfoItem {
  label: string;
  value: string | number;
}

interface InfoGridProps {
  infoItems: InfoItem[];
}

export const InfoGrid: React.FC<InfoGridProps> = ({ infoItems }) => {
  return (
    <div className={styles.infoGrid}>
      {infoItems.map((item, index) => (
        <div key={index} className={styles.infoItem}>
          <span className={styles.infoLabel}>{item.label}:</span>
          <span>{item.value}</span>
        </div>
      ))}
    </div>
  );
};

export default InfoGrid;