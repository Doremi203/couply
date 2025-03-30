import React from "react";
import styles from "./profileInfo.module.css";
import { CustomButton } from "../../../../shared/components/CustomButton";
import VerifiedIcon from "@mui/icons-material/Verified";
import { ProfileData } from "../../types";

interface ProfileInfoProps {
  profileData: ProfileData;
  isVerified: boolean;
  onVerificationRequest: () => void;
}

export const ProfileInfo: React.FC<ProfileInfoProps> = ({
  profileData,
  isVerified,
  onVerificationRequest
}) => {
  return (
    <div className={styles.profileInfo}>
      <div className={styles.profileImageContainer}>
        <img
          src={profileData.photos[0] || "/photo1.png"}
          alt="Profile"
          className={styles.profilePic}
        />
        {isVerified && (
          <div className={styles.verificationBadge}>
            <VerifiedIcon />
          </div>
        )}
      </div>
      <h2>{profileData.name}, {profileData.age}</h2>
      {!isVerified && (
        <CustomButton
          text="Verify Profile"
          onClick={onVerificationRequest}
          className={styles.verifyButton}
        />
      )}
    </div>
  );
};

export default ProfileInfo;