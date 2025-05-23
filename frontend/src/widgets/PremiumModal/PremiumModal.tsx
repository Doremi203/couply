import BoltIcon from '@mui/icons-material/Bolt';
import CloseIcon from '@mui/icons-material/Close';
import React from 'react';
import { createPortal } from 'react-dom';
import { useNavigate } from 'react-router-dom';

import styles from './premiumModal.module.css';

interface PremiumModalProps {
  isOpen: boolean;
  onClose: () => void;
}

export const PremiumModal: React.FC<PremiumModalProps> = ({ isOpen, onClose }) => {
  const navigate = useNavigate();

  const handleSubscribe = () => {
    onClose();
    navigate('/premium');
  };

  const handleViewAllPlans = () => {
    onClose();
    navigate('/premium');
  };

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

        <h2 className={styles.title}>Couply Premium</h2>

        <div className={styles.benefits}>
          <div className={styles.benefitItem}>
            <BoltIcon className={styles.benefitIcon} />
            <span>Никакой рекламы</span>
          </div>
          <div className={styles.benefitItem}>
            <BoltIcon className={styles.benefitIcon} />
            <span>Возможность писать сообщение вместе с лайком</span>
          </div>
          <div className={styles.benefitItem}>
            <BoltIcon className={styles.benefitIcon} />
            <span>Безграничное возвращение назад при поиске</span>
          </div>
        </div>

        <div className={styles.planCard} onClick={() => handleSubscribe()}>
          <div className={styles.planInfo}>
            <div className={styles.planDuration}>1 месяц</div>
          </div>
          <div className={styles.planPrice}>
            <div className={styles.priceAmount}>199₽</div>
          </div>
        </div>

        <button className={styles.viewAllButton} onClick={handleViewAllPlans}>
          Смотреть все планы
        </button>
      </div>
    </div>,
    document.body,
  );
};

export default PremiumModal;
