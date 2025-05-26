import CloseIcon from '@mui/icons-material/Close';
import React, { useState } from 'react';
import { createPortal } from 'react-dom';
import { useDispatch } from 'react-redux';
import { useNavigate } from 'react-router-dom';

// import { useCreatePaymentMutation } from '../../entities/payments/api/paymentsApi';
import { useCreatePaymentMutation } from '../../entities/payments/api/paymentsApi';
import { useCreateSubscriptionMutation } from '../../entities/subscription/api/subscriptionApi';
import { setSubscriptionData } from '../../entities/subscription/model/subscriptionSlice';
import { Plan, Status } from '../../entities/subscription/types';

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

  const [cardNumber, setCardNumber] = useState('');
  const [expiryDate, setExpiryDate] = useState('');
  const [cvv, setCvv] = useState('');
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
          <div className={styles.formGroup}>
            <label htmlFor="cardNumber">Номер карты</label>
            <input
              type="text"
              id="cardNumber"
              value={cardNumber}
              onChange={e => setCardNumber(e.target.value)}
              placeholder="1234 5678 9012 3456"
              maxLength={19}
              required
            />
          </div>

          <div className={styles.formRow}>
            <div className={styles.formGroup}>
              <label htmlFor="expiryDate">Срок действия</label>
              <input
                type="text"
                id="expiryDate"
                value={expiryDate}
                onChange={e => setExpiryDate(e.target.value)}
                placeholder="MM/YY"
                maxLength={5}
                required
              />
            </div>

            <div className={styles.formGroup}>
              <label htmlFor="cvv">CVV</label>
              <input
                type="text"
                id="cvv"
                value={cvv}
                onChange={e => setCvv(e.target.value)}
                placeholder="123"
                maxLength={3}
                required
              />
            </div>
          </div>

          <button type="submit" className={styles.submitButton} disabled={isProcessing}>
            {isProcessing ? 'Обработка...' : 'Оплатить'}
          </button>
        </form>

        <div className={styles.securityNote}>
          <p>Ваши данные защищены и не будут переданы третьим лицам</p>
        </div>
      </div>
    </div>,
    document.body,
  );
};

export default PaymentModal;
