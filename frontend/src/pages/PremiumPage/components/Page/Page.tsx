import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import BoltIcon from '@mui/icons-material/Bolt';
import CloseIcon from '@mui/icons-material/Close';
import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

import { Plan } from '../../../../entities/subscription/types';
import { NavBar } from '../../../../shared/components/NavBar';
import { PaymentModal } from '../../../../widgets/PaymentModal';

import styles from './premiumPage.module.css';

export const PremiumPage: React.FC = () => {
  const navigate = useNavigate();
  const [isPaymentModalOpen, setIsPaymentModalOpen] = useState(false);

  const handleBack = () => {
    navigate(-1);
  };

  const handleClose = () => {
    navigate(-1);
  };

  const [selectedPlan, setSelectedPlan] = useState<string | null>('6-month');
  const [selectedPrice, setSelectedPrice] = useState<string | null>('999₽');
  const [selectedToApi, setSelectedToApi] = useState<Plan>(Plan.semiAnnual);

  const plans = [
    { id: 'monthly', duration: '1 месяц', price: '199₽', toApi: Plan.monthly },
    { id: '6-month', duration: '6 месяцев', price: '999₽', toApi: Plan.semiAnnual },
    { id: 'yearly', duration: '12 месяцев', price: '1799₽', toApi: Plan.annual },
  ];

  const handleSubscribe = async (plan: string, price: string, toApi: Plan) => {
    setSelectedPlan(plan);
    setSelectedPrice(price);
    setSelectedToApi(toApi);
    setIsPaymentModalOpen(true);
  };

  const handleCreateSub = async () => {
    setIsPaymentModalOpen(true);
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
              onClick={() => handleSubscribe(plan.id, plan.price, plan.toApi)}
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

        <button className={styles.restoreButton} onClick={handleCreateSub}>
          Подписаться за {selectedPrice}
        </button>
      </div>

      <div className={styles.navBarWrapper}>
        <NavBar />
      </div>

      <PaymentModal
        isOpen={isPaymentModalOpen}
        onClose={() => setIsPaymentModalOpen(false)}
        selectedPlan={selectedToApi}
        price={selectedPrice || '0₽'}
      />
    </div>
  );
};

export default PremiumPage;
