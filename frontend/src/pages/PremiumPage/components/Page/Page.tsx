import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import BoltIcon from '@mui/icons-material/Bolt';
import CloseIcon from '@mui/icons-material/Close';
import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

import { NavBar } from '../../../../shared/components/NavBar';

import styles from './premiumPage.module.css';

export const PremiumPage: React.FC = () => {
  const navigate = useNavigate();

  const handleBack = () => {
    navigate(-1);
  };

  const handleClose = () => {
    navigate(-1);
  };

  const [selectedPlan, setSelectedPlan] = useState<string | null>('6-month');
  const [selectedPrice, setSelectedPrice] = useState<string | null>('999₽');

  const plans = [
    { id: 'monthly', duration: '1 месяц', price: '199₽' },
    { id: '6-month', duration: '6 месяцев', price: '999₽' },
    { id: 'yearly', duration: '12 месяцев', price: '1799₽' },
  ];

  const handleSubscribe = (plan: string, price: string) => {
    setSelectedPlan(plan);
    setSelectedPrice(price);
  };

  return (
    <div className={styles.pageContainer}>
      <div className={styles.content}>
        <div className={styles.header}>
          <button className={styles.backButton} onClick={handleBack}>
            <ArrowBackIcon />
          </button>
          <button className={styles.closeButton} onClick={handleClose}>
            <CloseIcon />
          </button>
        </div>
        <div className={styles.decorations}>
          <div className={styles.star1} />
          <div className={styles.star2} />
          <div className={styles.circle} />
          <div className={styles.wave1} />
          <div className={styles.wave2} />
          <div className={styles.square} />
        </div>
        <h1 className={styles.title}>Couply Premium</h1>
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
        <div className={styles.subscriptionPlans}>
          {plans.map(plan => (
            <div
              key={plan.id}
              className={`${styles.planCard} ${selectedPlan === plan.id ? styles.selectedPlan : ''}`}
              onClick={() => handleSubscribe(plan.id, plan.price)}
            >
              <div className={styles.planInfo}>
                <div className={styles.planDuration}>{plan.duration}</div>
              </div>
              <div className={styles.planPrice}>
                <div className={styles.priceAmount}>{plan.price}</div>
              </div>
            </div>
          ))}
        </div>

        <button className={styles.restoreButton}>Подписаться за {selectedPrice}</button>
      </div>

      <div className={styles.navBarWrapper}>
        <NavBar />
      </div>
    </div>
  );
};

export default PremiumPage;
