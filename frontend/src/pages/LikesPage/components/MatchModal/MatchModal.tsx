import React from "react";
import styles from "./matchModal.module.css";
import { CustomButton } from "../../../../shared/CustomButton";

interface MatchModalProps {
  userImage: string;
  matchImage: string;
  matchName: string;
  onSendMessage: () => void;
  onKeepSwiping: () => void;
}

export const MatchModal: React.FC<MatchModalProps> = ({
  userImage,
  matchImage,
  matchName,
  onSendMessage,
  onKeepSwiping,
}) => {
  return (
    <div className={styles.overlay}>
      <div className={styles.modal}>
        <div className={styles.confetti}>
          {/* Confetti elements */}
          <div className={`${styles.confettiItem} ${styles.confetti1}`}></div>
          <div className={`${styles.confettiItem} ${styles.confetti2}`}></div>
          <div className={`${styles.confettiItem} ${styles.confetti3}`}></div>
          <div className={`${styles.confettiItem} ${styles.confetti4}`}></div>
          <div className={`${styles.confettiItem} ${styles.confetti5}`}></div>
          <div className={`${styles.confettiItem} ${styles.confetti6}`}></div>
          <div className={`${styles.confettiItem} ${styles.confetti7}`}></div>
          <div className={`${styles.confettiItem} ${styles.confetti8}`}></div>
        </div>

        <div className={styles.profileImages}>
          <div className={styles.profileImageWrapper}>
            <img src={userImage} alt="Your profile" className={styles.profileImage} />
          </div>
          <div className={styles.profileImageWrapper}>
            <img src={matchImage} alt={matchName} className={styles.profileImage} />
          </div>
        </div>

        <h2 className={styles.title}>Love is in the air!</h2>
        <p className={styles.subtitle}>
          You and <span className={styles.matchName}>{matchName}</span> like each other
        </p>

        <div className={styles.buttons}>
          <CustomButton 
            text="Send Message" 
            onClick={onSendMessage} 
            className={styles.sendMessageButton} 
          />
          <button className={styles.keepSwipingButton} onClick={onKeepSwiping}>
            Keep Swiping
          </button>
        </div>
      </div>
    </div>
  );
};

export default MatchModal;