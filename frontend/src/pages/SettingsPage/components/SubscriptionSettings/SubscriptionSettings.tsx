import WorkspacePremiumIcon from '@mui/icons-material/WorkspacePremium';
import { CircularProgress } from '@mui/material';
import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';

import {
  useGetSubscriptionMutation,
  useCancelSubscriptionMutation,
} from '../../../../entities/subscription/api/subscriptionApi';
import { Plan, Status } from '../../../../entities/subscription/types';
import ConfirmModal from '../../../../shared/components/ConfirmModal';

import styles from './subscriptionSettings.module.css';

interface SubscriptionSettingsProps {
  className?: string;
}

export const SubscriptionSettings: React.FC<SubscriptionSettingsProps> = ({ className }) => {
  const navigate = useNavigate();
  const [isLoading, setIsLoading] = useState(false);
  const [getSubscription] = useGetSubscriptionMutation();
  const [cancelSubscription] = useCancelSubscriptionMutation();
  const [isPremium, setIsPremium] = useState(false);
  const [subscriptionId, setSubscriptionId] = useState<string>('');
  const [subscriptionPlan, setSubscriptionPlan] = useState<Plan>(Plan.unspecified);
  const [subscriptionEndDate, setSubscriptionEndDate] = useState<string>('');
  const [showUnsubscribeModal, setShowUnsubscribeModal] = useState(false);

  useEffect(() => {
    const fetchSubscription = async () => {
      try {
        setIsLoading(true);
        const subscription = await getSubscription({}).unwrap();
        setIsPremium(subscription.status === Status.active);
        setSubscriptionId(subscription.subscriptionId);
        setSubscriptionPlan(subscription.plan);
        setSubscriptionEndDate(subscription.endDate);
      } catch (error) {
        console.error(error);
      } finally {
        setIsLoading(false);
      }
    };

    fetchSubscription();
  }, [getSubscription]);

  const handleCancelSubscription = () => {
    setShowUnsubscribeModal(true);
  };

  const handleUnsubscribe = async () => {
    try {
      setIsLoading(true);
      await cancelSubscription({ subscriptionId }).unwrap();
      setIsPremium(false);
      setShowUnsubscribeModal(false);
    } catch (error) {
      console.error('Error canceling subscription:', error);
      alert('Произошла ошибка при отмене подписки');
    } finally {
      setIsLoading(false);
    }
  };

  const handleSubscribe = () => {
    navigate('/premium');
  };

  const formatDate = (dateString: string) => {
    if (!dateString) return '';

    const date = new Date(dateString);
    return date.toLocaleDateString('ru-RU', {
      day: 'numeric',
      month: 'long',
      year: 'numeric',
    });
  };

  const getPlanName = (plan: Plan) => {
    switch (plan) {
      case Plan.monthly:
        return 'Ежемесячный';
      case Plan.annual:
        return 'Годовой';
      case Plan.semiAnnual:
        return 'Полугодовой';
      default:
        return 'Стандартный';
    }
  };

  return (
    <div className={`${styles.container} ${className || ''}`}>
      <div className={styles.header}>
        <div className={styles.icon}>
          <WorkspacePremiumIcon />
        </div>
        <h3 className={styles.title}>Couply Premium</h3>
      </div>

      {isLoading ? (
        <div style={{ display: 'flex', justifyContent: 'center', padding: '20px' }}>
          <CircularProgress size={24} />
        </div>
      ) : isPremium ? (
        <>
          <div className={styles.planInfo}>
            <div className={styles.planName}>{getPlanName(subscriptionPlan)} план</div>
            <div className={styles.planDetails}>Активен до {formatDate(subscriptionEndDate)}</div>
          </div>
          <button
            className={styles.cancelButton}
            onClick={handleCancelSubscription}
            disabled={isLoading}
          >
            {isLoading ? (
              <CircularProgress size={16} style={{ color: 'white' }} />
            ) : (
              'Отключить подписку'
            )}
          </button>
        </>
      ) : (
        <>
          <div className={styles.noSubscription}>У вас нет активной подписки Premium</div>
          <button className={styles.subscribeButton} onClick={handleSubscribe} disabled={isLoading}>
            Оформить подписку
          </button>
        </>
      )}

      <ConfirmModal
        isOpen={showUnsubscribeModal}
        onClose={() => setShowUnsubscribeModal(false)}
        onConfirm={handleUnsubscribe}
        title="Отмена подписки"
        message="Вы уверены, что хотите отменить подписку? После отмены вы потеряете доступ к премиум-функциям."
        confirmText="Отменить подписку"
        cancelText="Вернуться"
      />
    </div>
  );
};

export default SubscriptionSettings;
