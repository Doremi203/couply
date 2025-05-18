import React, { useEffect, useRef } from 'react';

import styles from './matchModal.module.css';

interface MatchModalProps {
  userImage: string;
  matchImage: string;
  matchName: string;
  onKeepSwiping: () => void;
}

export const MatchModal: React.FC<MatchModalProps> = ({
  userImage,
  matchImage,
  matchName,
  onKeepSwiping,
}) => {
  const modalRef = useRef<HTMLDivElement>(null);

  const handleClickOutside = (event: MouseEvent) => {
    if (modalRef.current && !modalRef.current.contains(event.target as Node)) {
      onKeepSwiping();
    }
  };

  useEffect(() => {
    document.body.style.overflow = 'hidden';

    document.addEventListener('mousedown', handleClickOutside);

    if (modalRef.current) {
      modalRef.current.focus();
    }

    return () => {
      document.removeEventListener('mousedown', handleClickOutside);
      document.body.style.overflow = '';
    };
  }, []);

  const handleModalClick = (e: React.MouseEvent) => {
    e.stopPropagation();
  };

  return (
    <div className={styles.darkOverlay} onClick={handleModalClick}>
      <div className={styles.modalContent} ref={modalRef} tabIndex={-1}>
        <div className={styles.profileImages}>
          <div className={styles.profileImageWrapper}>
            <img src={userImage} alt="Your profile" className={styles.profileImage} />
          </div>
          <div className={styles.profileImageWrapper}>
            <img src={matchImage} alt={matchName} className={styles.profileImage} />
          </div>
        </div>
        <h1 className={styles.title}>КАК ДВЕ КАПЛИ</h1>

        <div className={styles.dropContainer}>
          <div className={styles.drop} />
        </div>
      </div>
    </div>
  );
};

export default MatchModal;
