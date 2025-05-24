import CloseIcon from '@mui/icons-material/Close';
import React from 'react';
import { createPortal } from 'react-dom';

import styles from './messageModal.module.css';

interface MessageModalProps {
  isOpen: boolean;
  onClose: () => void;
  message?: string;
  senderName?: string;
}

export const MessageModal: React.FC<MessageModalProps> = ({
  isOpen,
  onClose,
  message = 'No message available',
  senderName = 'Someone',
}) => {
  const handleOverlayClick = (e: React.MouseEvent<HTMLDivElement>) => {
    if (e.target === e.currentTarget) {
      onClose();
    }
  };

  if (!isOpen) return null;

  return createPortal(
    <div className={styles.modalOverlay} onClick={handleOverlayClick}>
      <div className={styles.modalContent}>
        <button className={styles.closeButton} onClick={onClose}>
          <CloseIcon />
        </button>

        <div className={styles.decorations}>
          <div className={styles.star1} />
          <div className={styles.star2} />
          <div className={styles.circle} />
          <div className={styles.wave1} />
        </div>

        <h2 className={styles.title}>Сообщение от {senderName}</h2>

        <div className={styles.messageContent}>
          <p className={styles.messageText}>{message}</p>
        </div>
      </div>
    </div>,
    document.body,
  );
};

export default MessageModal;
