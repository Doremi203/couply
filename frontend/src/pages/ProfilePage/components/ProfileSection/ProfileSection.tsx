import React, { ReactNode } from "react";
import styles from "./profileSection.module.css";

interface ProfileSectionProps {
  title: string;
  children: ReactNode;
  onEdit?: () => void;
  showEditLink?: boolean;
}

export const ProfileSection: React.FC<ProfileSectionProps> = ({
  title,
  children,
  onEdit,
  showEditLink = false
}) => {
  return (
    <div className={styles.profileSection}>
      <div className={styles.sectionHeader}>
        <h3>{title}</h3>
        {showEditLink && onEdit && (
          <span className={styles.editLink} onClick={onEdit}>
            Edit
          </span>
        )}
      </div>
      {children}
    </div>
  );
};

export default ProfileSection;