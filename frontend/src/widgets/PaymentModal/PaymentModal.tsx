import CloseIcon from '@mui/icons-material/Close';
import React, { useState } from 'react';
import { createPortal } from 'react-dom';
import { useNavigate } from 'react-router-dom';

import { useCreatePaymentMutation } from '../../entities/payments/api/paymentsApi';
import { useCreateSubscriptionMutation } from '../../entities/subscription/api/subscriptionApi';
import { Plan } from '../../entities/subscription/types';

import styles from './paymentModal.module.css';

interface PaymentModalProps {
  isOpen: boolean;
  onClose: () => void;
  selectedPlan: Plan;
  price: string;
}

export const PaymentModal: React.FC<PaymentModalProps> = ({
  isOpen,
  onClose,
  selectedPlan,
  price,
}) => {
  const navigate = useNavigate();
  const [createPayment] = useCreatePaymentMutation();
  const [createSubscription] = useCreateSubscriptionMutation();

  const [isProcessing, setIsProcessing] = useState(false);

  const handleOverlayClick = (e: React.MouseEvent<HTMLDivElement>) => {
    if (e.target === e.currentTarget) {
      onClose();
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsProcessing(true);
    try {
      const subscriptionResponse = await createSubscription({
        autoRenew: true,
        plan: selectedPlan,
      }).unwrap();

      await createPayment({
        subscriptionId: subscriptionResponse.subscriptionId,
        amount: price.replace('₽', ''),
        currency: 'RUB',
      });

      navigate('/profile');
      onClose();
    } catch (error) {
      console.error('Payment error:', error);
    } finally {
      setIsProcessing(false);
    }
  };

  if (!isOpen) return null;

  return createPortal(
    <div className={styles.modalOverlay} onClick={handleOverlayClick}>
      <div className={styles.modalContent}>
        <button className={styles.closeButton} onClick={onClose}>
          <CloseIcon />
        </button>

        <h2 className={styles.title}>Оплата подписки</h2>
        <p className={styles.subtitle}>Сумма к оплате: {price}</p>

        <form onSubmit={handleSubmit} className={styles.paymentForm}>
          <button type="submit" className={styles.submitButton} disabled={isProcessing}>
            {isProcessing ? 'Обработка...' : 'Оплатить'}
          </button>
        </form>
      </div>
    </div>,
    document.body,
  );
};

export default PaymentModal;
