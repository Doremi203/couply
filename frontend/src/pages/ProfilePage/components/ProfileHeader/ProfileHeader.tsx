import React from "react";
import styles from "./profileHeader.module.css";
import { IconButton } from "../../../../shared/components/IconButton";
import { CustomButton } from "../../../../shared/components/CustomButton";
import EditIcon from "@mui/icons-material/Edit";
import VisibilityIcon from "@mui/icons-material/Visibility";
import VisibilityOffIcon from "@mui/icons-material/VisibilityOff";
import HistoryIcon from "@mui/icons-material/History";

interface ProfileHeaderProps {
  isProfileHidden: boolean;
  onEditToggle: () => void;
  onVisibilityToggle: () => void;
  onActivityClick: () => void;
  onPreviewClick: () => void;
}

export const ProfileHeader: React.FC<ProfileHeaderProps> = ({
  isProfileHidden,
  onEditToggle,
  onVisibilityToggle,
  onActivityClick,
  onPreviewClick
}) => {
  return (
    <div className={styles.profileHeader}>
      <div className={styles.header}>profile</div>
      <div className={styles.profileActions}>
        <IconButton onClick={onEditToggle} touchFriendly={true}>
          <EditIcon />
        </IconButton>
        <IconButton onClick={onVisibilityToggle}>
          {isProfileHidden ? <VisibilityOffIcon /> : <VisibilityIcon />}
        </IconButton>
        <IconButton onClick={onActivityClick} touchFriendly={true}>
          <HistoryIcon />
        </IconButton>
      </div>
      <div className={styles.previewButtonContainer}>
        <CustomButton
          text="preview"
          onClick={onPreviewClick}
          className={styles.previewButton}
        />
      </div>
    </div>
  );
};

export default ProfileHeader;